package strategy

import (
	"fmt"
	"io"
	"reflect"
	"sort"
	"strings"

	"github.com/c9s/bbgo/pkg/bbgo"
	"github.com/c9s/bbgo/pkg/util"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type JsonStruct struct {
	Key   string
	Json  string
	Type  string
	Value interface{}
}

type JsonArr []JsonStruct

func (a JsonArr) Len() int           { return len(a) }
func (a JsonArr) Less(i, j int) bool { return a[i].Key < a[j].Key }
func (a JsonArr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func DefaultStyle() *table.Style {
	style := table.Style{
		Name:    "StyleRounded",
		Box:     table.StyleBoxRounded,
		Format:  table.FormatOptionsDefault,
		HTML:    table.DefaultHTMLOptions,
		Options: table.OptionsDefault,
		Title:   table.TitleOptionsDefault,
		Color:   table.ColorOptionsYellowWhiteOnBlack,
	}
	style.Color.Row = text.Colors{text.FgHiYellow, text.BgHiBlack}
	style.Color.RowAlternate = text.Colors{text.FgYellow, text.BgBlack}
	return &style
}

func DefaultWhiteList() []string {
	return []string{"Window", "Interval", "Symbol"}
}

// @param s: strategy object
// @param f: io.Writer used for writing the config dump
// @param style: pretty print table style. Use DefaultStyle() to get default one.
// @param withColor: whether to print with color
// @param whiteLists: fields to be printed out from embedded struct (1st layer only)
func PrintConfig(s interface{}, f io.Writer, style *table.Style, withColor bool, whiteLists ...string) {
	t := table.NewWriter()
	var write func(io.Writer, string, ...interface{})

	if withColor {
		write = color.New(color.FgHiYellow).FprintfFunc()
	} else {
		write = func(a io.Writer, format string, args ...interface{}) {
			fmt.Fprintf(a, format, args...)
		}
	}
	if style != nil {
		t.SetOutputMirror(f)
		t.SetStyle(*style)
		t.SetColumnConfigs([]table.ColumnConfig{
			{Number: 4, WidthMax: 50, WidthMaxEnforcer: text.WrapText},
		})
		t.AppendHeader(table.Row{"json", "struct field name", "type", "value"})
	}
	write(f, "---- %s Settings ---\n", bbgo.CallID(s))

	embeddedWhiteSet := map[string]struct{}{}
	for _, whiteList := range whiteLists {
		embeddedWhiteSet[whiteList] = struct{}{}
	}

	redundantSet := map[string]struct{}{}

	var rows []table.Row

	val := reflect.ValueOf(s)

	if val.Type().Kind() == util.Pointer {
		val = val.Elem()
	}
	var values JsonArr
	for i := 0; i < val.Type().NumField(); i++ {
		t := val.Type().Field(i)
		if !t.IsExported() {
			continue
		}
		fieldName := t.Name
		switch jsonTag := t.Tag.Get("json"); jsonTag {
		case "-":
		case "":
			// we only fetch fields from the first layer of the embedded struct
			if t.Anonymous {
				var target reflect.Type
				var field reflect.Value
				if t.Type.Kind() == util.Pointer {
					target = t.Type.Elem()
					field = val.Field(i).Elem()
				} else {
					target = t.Type
					field = val.Field(i)
				}
				for j := 0; j < target.NumField(); j++ {
					tt := target.Field(j)
					if !tt.IsExported() {
						continue
					}
					fieldName := tt.Name
					if _, ok := embeddedWhiteSet[fieldName]; !ok {
						continue
					}
					if jtag := tt.Tag.Get("json"); jtag != "" && jtag != "-" {
						name := strings.Split(jtag, ",")[0]
						if _, ok := redundantSet[name]; ok {
							continue
						}
						redundantSet[name] = struct{}{}
						value := field.Field(j).Interface()
						values = append(values, JsonStruct{Key: fieldName, Json: name, Type: tt.Type.String(), Value: value})
					}
				}
			}
		default:
			name := strings.Split(jsonTag, ",")[0]
			if _, ok := redundantSet[name]; ok {
				continue
			}
			redundantSet[name] = struct{}{}
			values = append(values, JsonStruct{Key: fieldName, Json: name, Type: t.Type.String(), Value: val.Field(i).Interface()})
		}
	}
	sort.Sort(values)
	for _, value := range values {
		if style != nil {
			rows = append(rows, table.Row{value.Json, value.Key, value.Type, value.Value})
		} else {
			write(f, "%s: %v\n", value.Json, value.Value)
		}
	}
	if style != nil {
		t.AppendRows(rows)
		t.Render()
	}
}