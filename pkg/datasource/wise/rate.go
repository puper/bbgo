package wise

import "github.com/puper/bbgo/pkg/fixedpoint"

type Rate struct {
	Value  fixedpoint.Value `json:"rate"`
	Target string           `json:"target"`
	Source string           `json:"source"`
	Time   Time             `json:"time"`
}
