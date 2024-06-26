package indicator

import (
	"math"

	"github.com/puper/bbgo/pkg/datatype/floats"
	"github.com/puper/bbgo/pkg/types"
)

// Refer: https://tradingview.com/script/aDymGrFx-Drift-Study-Inspired-by-Monte-Carlo-Simulations-with-BM-KL/
// Brownian Motion's drift factor
// could be used in Monte Carlo Simulations
//
// In the context of Brownian motion, drift can be measured by calculating the simple moving average (SMA) of the logarithm
// of the price changes of a security over a specified period of time. This SMA can be used to identify the long-term trend
// or bias in the random movement of the security's price. A security with a positive drift is said to be trending upwards,
// while a security with a negative drift is said to be trending downwards. Drift can be used by traders to identify potential
// entry and exit points for trades, or to confirm other technical analysis signals.
// It is typically used in conjunction with other indicators to provide a more comprehensive view of the security's price.

//go:generate callbackgen -type Drift
type Drift struct {
	types.SeriesBase
	types.IntervalWindow
	chng      *types.Queue
	Values    floats.Slice
	MA        types.UpdatableSeriesExtend
	LastValue float64

	UpdateCallbacks []func(value float64)
}

func (inc *Drift) Update(value float64) {
	if inc.chng == nil {
		inc.SeriesBase.Series = inc
		if inc.MA == nil {
			inc.MA = &SMA{IntervalWindow: types.IntervalWindow{Interval: inc.Interval, Window: inc.Window}}
		}
		inc.chng = types.NewQueue(inc.Window)
		inc.LastValue = value
		return
	}
	var chng float64
	if value == 0 {
		chng = 0
	} else {
		chng = math.Log(value / inc.LastValue)
		inc.LastValue = value
	}
	inc.MA.Update(chng)
	inc.chng.Update(chng)
	if inc.chng.Length() >= inc.Window {
		stdev := types.Stdev(inc.chng, inc.Window)
		drift := inc.MA.Last(0) - stdev*stdev*0.5
		inc.Values.Push(drift)
	}
}

// Assume that MA is SMA
func (inc *Drift) ZeroPoint() float64 {
	window := float64(inc.Window)
	stdev := types.Stdev(inc.chng, inc.Window)
	chng := inc.chng.Index(inc.Window - 1)
	/*b := -2 * inc.MA.Last() - 2
	c := window * stdev * stdev - chng * chng + 2 * chng * (inc.MA.Last() + 1) - 2 * inc.MA.Last() * window

	root := math.Sqrt(b*b - 4*c)
	K1 := (-b + root)/2
	K2 := (-b - root)/2
	N1 := math.Exp(K1) * inc.LastValue
	N2 := math.Exp(K2) * inc.LastValue
	if math.Abs(inc.LastValue-N1) < math.Abs(inc.LastValue-N2) {
		return N1
	} else {
		return N2
	}*/
	return inc.LastValue * math.Exp(window*(0.5*stdev*stdev)+chng-inc.MA.Last(0)*window)
}

func (inc *Drift) Clone() (out *Drift) {
	out = &Drift{
		IntervalWindow: inc.IntervalWindow,
		chng:           inc.chng.Clone(),
		Values:         inc.Values[:],
		MA:             types.Clone(inc.MA),
		LastValue:      inc.LastValue,
	}
	out.SeriesBase.Series = out
	return out
}

func (inc *Drift) TestUpdate(value float64) *Drift {
	out := inc.Clone()
	out.Update(value)
	return out
}

func (inc *Drift) Index(i int) float64 {
	return inc.Last(i)
}

func (inc *Drift) Last(i int) float64 {
	return inc.Values.Last(i)
}

func (inc *Drift) Length() int {
	if inc.Values == nil {
		return 0
	}
	return inc.Values.Length()
}

var _ types.SeriesExtend = &Drift{}

func (inc *Drift) PushK(k types.KLine) {
	inc.Update(k.Close.Float64())
}

func (inc *Drift) CalculateAndUpdate(allKLines []types.KLine) {
	if inc.chng == nil {
		for _, k := range allKLines {
			inc.PushK(k)
			inc.EmitUpdate(inc.Last(0))
		}
	} else {
		k := allKLines[len(allKLines)-1]
		inc.PushK(k)
		inc.EmitUpdate(inc.Last(0))
	}
}

func (inc *Drift) handleKLineWindowUpdate(interval types.Interval, window types.KLineWindow) {
	if inc.Interval != interval {
		return
	}

	inc.CalculateAndUpdate(window)
}

func (inc *Drift) Bind(updater KLineWindowUpdater) {
	updater.OnKLineWindowUpdate(inc.handleKLineWindowUpdate)
}
