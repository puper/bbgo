package indicatorv2

import (
	"github.com/puper/bbgo/pkg/datatype/floats"
	"github.com/puper/bbgo/pkg/types"
)

type PivotLowStream struct {
	*types.Float64Series
	rawValues           floats.Slice
	window, rightWindow int
}

func PivotLow(source types.Float64Source, window int, args ...int) *PivotLowStream {
	rightWindow := window
	if len(args) > 0 {
		rightWindow = args[0]
	}

	s := &PivotLowStream{
		Float64Series: types.NewFloat64Series(),
		window:        window,
		rightWindow:   rightWindow,
	}

	s.Subscribe(source, func(x float64) {
		s.rawValues.Push(x)
		if low, ok := s.calculatePivotLow(s.rawValues, s.window, s.rightWindow); ok {
			s.PushAndEmit(low)
		}
	})
	return s
}

func (s *PivotLowStream) calculatePivotLow(lows floats.Slice, left, right int) (float64, bool) {
	return floats.FindPivot(lows, left, right, func(a, pivot float64) bool {
		return a > pivot
	})
}
