package indicator

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/puper/bbgo/pkg/fixedpoint"
	"github.com/puper/bbgo/pkg/types"
)

// generated from:
// 2020/12/05 10:25
// curl -s 'https://www.binance.com/api/v3/klines?symbol=ETHUSDT&interval=5m&endTime=1607135400000&limit=1000' | jq '. | map({ closePrice: (.[4] | tonumber), openTime: .[0] })'
// curl -s 'https://www.binance.com/api/v3/klines?symbol=ETHUSDT&interval=5m&endTime=1607135400000&limit=1000' | jq '. | map(.[4] | tonumber)'
var ethusdt5m = []byte(`[
	614.36,
	613.62,
	611.68,
	614.7,
	607.69,
	606.82,
	602.91,
	606.35,
	606.04,
	603.94,
	607.37,
	599.72,
	597.77,
	594.36,
	595.18,
	592.65,
	595.44,
	595.08,
	595.82,
	596.31,
	592.47,
	594.19,
	590.71,
	586.77,
	589.74,
	592.33,
	593.44,
	594.34,
	593.89,
	594,
	590.81,
	595.51,
	594.08,
	596.58,
	592.97,
	593.3,
	590.89,
	591.9,
	589.67,
	590.25,
	592.53,
	594.01,
	591.46,
	588.5,
	591.34,
	589.24,
	592,
	594.57,
	593.98,
	593.93,
	593.46,
	594.22,
	595.57,
	595.81,
	595.44,
	599.65,
	600.3,
	599.5,
	601,
	598.24,
	596.56,
	596.55,
	595.61,
	596.91,
	593.49,
	592.03,
	592.56,
	592.95,
	593.56,
	593.78,
	594.69,
	594.17,
	595.62,
	596.06,
	595.83,
	598.51,
	598.29,
	596.94,
	597.87,
	598.02,
	595.13,
	594.87,
	597,
	599.89,
	599.49,
	599.32,
	598.82,
	597.26,
	597.46,
	593.32,
	592.69,
	589.42,
	590.39,
	589.91,
	586.83,
	586.53,
	588.85,
	589.44,
	584.95,
	584.06,
	588.99,
	588.75,
	588.62,
	589.61,
	589.27,
	585.26,
	582.06,
	577.06,
	583.13,
	588.01,
	589.25,
	589.42,
	591.03,
	590.99,
	592.33,
	593.86,
	592.12,
	593.21,
	592.4,
	590.86,
	590.1,
	592.32,
	590.79,
	592.01,
	592.07,
	596.66,
	595.63,
	595.76,
	594.8,
	594.21,
	592.58,
	592.19,
	592.6,
	591.29,
	591.89,
	590.03,
	590.33,
	588.53,
	589.26,
	587.18,
	589.03,
	587.49,
	590.83,
	590.63,
	587.13,
	584.86,
	584.86,
	584,
	587.95,
	589.48,
	589.31,
	586.96,
	584.29,
	585.68,
	587.57,
	582.19,
	580.7,
	580.01,
	580.83,
	578.84,
	578.97,
	582.23,
	581.45,
	582,
	582.69,
	582.68,
	581.54,
	581.32,
	584.95,
	582.57,
	580.75,
	579.58,
	580.13,
	582.42,
	582.85,
	585.5,
	584.87,
	586.1,
	585.39,
	587.32,
	587.83,
	590.11,
	589.84,
	589.78,
	592.13,
	590.5,
	589.54,
	589.01,
	590.7,
	592.73,
	593.54,
	591.56,
	591.32,
	590.49,
	590.64,
	591.43,
	591.94,
	593.12,
	593.97,
	595.47,
	596.23,
	597.63,
	595.06,
	595,
	594.88,
	597.61,
	596.67,
	600.23,
	599.22,
	596.36,
	597.75,
	596.34,
	596.05,
	595.9,
	595.8,
	596.33,
	597.31,
	598.82,
	603.56,
	602.33,
	599.58,
	598.7,
	598.41,
	596.14,
	596.21,
	595.13,
	596.47,
	595.81,
	594.36,
	595.78,
	594.66,
	596,
	597.95,
	597.77,
	597.57,
	598.43,
	597.99,
	598.4,
	600.05,
	597.81,
	596.52,
	597.04,
	598.86,
	600.16,
	601.05,
	600.3,
	599.6,
	597.42,
	598.49,
	598.62,
	589.69,
	591.62,
	592.73,
	593.15,
	593.62,
	593.77,
	592.34,
	592.54,
	588.71,
	591.31,
	593.5,
	593.36,
	593.49,
	593.63,
	593.62,
	593.37,
	596.52,
	597.04,
	597.59,
	596.94,
	596.09,
	595.51,
	595.13,
	593.1,
	594.48,
	594.61,
	596.8,
	595.36,
	593.32,
	593.58,
	593,
	591.84,
	590.61,
	592.27,
	591.81,
	590.27,
	587.59,
	587.96,
	591.07,
	591.86,
	589.61,
	588.16,
	588.08,
	585.46,
	586.69,
	588.21,
	588.7,
	589.51,
	585.28,
	586.27,
	589.88,
	590.97,
	590.92,
	590.7,
	590.09,
	589.38,
	588.51,
	586.69,
	585.33,
	584.32,
	585.42,
	586.3,
	586.08,
	587.32,
	587.55,
	587.31,
	587.07,
	587.57,
	587.44,
	588.56,
	588.91,
	589.13,
	591.17,
	594.2,
	593.36,
	595.75,
	596.29,
	596.78,
	596.82,
	596.39,
	596.31,
	595.77,
	595.95,
	594.09,
	596.86,
	595.64,
	595.41,
	592.67,
	594,
	594.42,
	594.92,
	595,
	596.46,
	596.5,
	596.98,
	596.57,
	599.27,
	596.93,
	596.47,
	596.59,
	595.8,
	596.01,
	596.38,
	598.19,
	598.72,
	598.13,
	596.83,
	596.36,
	597.12,
	597.95,
	596.02,
	596,
	596.33,
	594.97,
	595.82,
	596.58,
	596.37,
	597.88,
	598.5,
	598.06,
	599.71,
	599.97,
	600.7,
	601.44,
	599.72,
	597.62,
	597.24,
	597.7,
	597.67,
	599.85,
	599.43,
	596.28,
	596.21,
	597.68,
	597.75,
	597.35,
	598.18,
	598.39,
	598.54,
	599.73,
	600.1,
	598.7,
	597.96,
	596.89,
	598.8,
	599.57,
	598.4,
	598.12,
	597.39,
	597.64,
	594.69,
	595.61,
	595.82,
	595.97,
	596.71,
	596,
	592.87,
	592.94,
	591.5,
	591.81,
	592.02,
	593.01,
	591.34,
	590.34,
	590.27,
	591.48,
	591.92,
	591.87,
	592.07,
	592.43,
	592.79,
	592.25,
	591.28,
	591.18,
	592,
	592.73,
	593.61,
	592.93,
	592.61,
	592.68,
	593.84,
	594.28,
	595.41,
	595.17,
	594.8,
	594.67,
	595.99,
	596.16,
	597.58,
	596.63,
	596.06,
	595.24,
	595.17,
	595.69,
	596.51,
	596.5,
	595.56,
	594.47,
	593.49,
	593.47,
	593.81,
	593.93,
	594,
	592.85,
	593.04,
	593.88,
	593.78,
	593.16,
	594.9,
	594.99,
	593.91,
	591.28,
	588.23,
	589.01,
	589.12,
	588.28,
	589.36,
	589.94,
	589.6,
	589.77,
	590.36,
	590.85,
	590.18,
	590.01,
	590.77,
	589.97,
	590.78,
	590.99,
	588.95,
	588.9,
	588.02,
	590.51,
	592.01,
	591.8,
	596.14,
	595.68,
	595.72,
	598.47,
	598.1,
	598.47,
	597.27,
	595.94,
	596.43,
	602.8,
	607.08,
	606.03,
	605.98,
	603.91,
	604.5,
	604.19,
	604.98,
	605.25,
	606.41,
	599.67,
	601.87,
	603.85,
	603.18,
	602.31,
	602.5,
	603.02,
	603.84,
	605.19,
	605.29,
	606.56,
	605.54,
	605.36,
	605.11,
	609.64,
	608.82,
	610.79,
	609.74,
	612.51,
	614.67,
	614.8,
	613.42,
	612.79,
	614.4,
	610.96,
	610.57,
	612.19,
	613.22,
	614.24,
	612.67,
	612.4,
	611.6,
	610.69,
	611.5,
	613.6,
	610.99,
	611.22,
	610.54,
	607.53,
	608.7,
	608.98,
	608.4,
	608.72,
	609.35,
	609.84,
	608.5,
	609.27,
	611.71,
	610.67,
	611.18,
	610.87,
	612.38,
	611.53,
	610.7,
	610.08,
	609,
	609.51,
	608.16,
	609.17,
	611.02,
	612.63,
	612.25,
	612.71,
	613.27,
	612.47,
	611.12,
	611.34,
	611.14,
	610.58,
	610.01,
	610.64,
	610.07,
	609.73,
	610.85,
	611.75,
	612.76,
	612.51,
	613.12,
	613.03,
	614.49,
	615.24,
	617.37,
	618.92,
	621.61,
	618.66,
	608.37,
	609.37,
	612.48,
	614.14,
	616.23,
	615.57,
	615.53,
	613.25,
	614.73,
	614.2,
	615.19,
	612.99,
	609.91,
	610.28,
	611.5,
	612.55,
	613.21,
	612.74,
	611.8,
	612.18,
	612.2,
	612.56,
	614.05,
	614.67,
	614.45,
	614.46,
	614.16,
	614.22,
	615.62,
	614.73,
	614.56,
	613.89,
	613.38,
	614.12,
	612.24,
	612.39,
	611.83,
	611.83,
	611.24,
	611.62,
	611.44,
	612.21,
	612.12,
	613.14,
	613.53,
	613.76,
	613.38,
	612.76,
	613.31,
	613.19,
	613.71,
	613.89,
	613.52,
	612.32,
	611.39,
	610.82,
	611.03,
	611.33,
	611.52,
	612.69,
	612.63,
	612.9,
	613.25,
	612.92,
	613.35,
	614.23,
	614.66,
	615.71,
	615.39,
	614.95,
	614.74,
	614.84,
	615.27,
	615.41,
	615.74,
	615.99,
	615.22,
	615.22,
	614.99,
	614.43,
	614.7,
	616.13,
	616.78,
	618.41,
	618.35,
	618.02,
	615.89,
	616.52,
	617.12,
	617.46,
	616.88,
	616.67,
	616.9,
	615.85,
	614.86,
	614.16,
	614.95,
	615.19,
	615.17,
	615.61,
	617.79,
	619.39,
	618.27,
	618.25,
	617.81,
	617.36,
	617.22,
	614.3,
	613.67,
	613.7,
	613.79,
	614.07,
	612.17,
	613.29,
	612.94,
	612.8,
	613.1,
	611.32,
	612.93,
	612.42,
	611.86,
	611.16,
	609.93,
	610.95,
	610.16,
	603.83,
	606.02,
	606.72,
	604.39,
	604.26,
	607,
	606.84,
	607.84,
	607.8,
	605.64,
	604.26,
	605.22,
	605.08,
	606.01,
	605.48,
	604.73,
	602.64,
	604.64,
	605.7,
	605.58,
	605.68,
	606.05,
	606.2,
	606.62,
	606.16,
	607.36,
	606.38,
	605.48,
	606.12,
	607.16,
	607.51,
	607.4,
	607.27,
	606.66,
	607.21,
	606.1,
	606.66,
	605.73,
	605.29,
	603.52,
	603.46,
	604.73,
	605.57,
	606.04,
	606.05,
	607.02,
	608.31,
	607.69,
	607.68,
	607.45,
	607.36,
	607.18,
	606.89,
	605.9,
	605.03,
	604.85,
	604.16,
	604.49,
	604.43,
	604.94,
	606.74,
	606.51,
	606.44,
	605.03,
	604.28,
	606.71,
	608.23,
	609.64,
	611.5,
	611.6,
	612.57,
	611.81,
	612.12,
	611.87,
	610.47,
	611.22,
	610.86,
	609.75,
	609.76,
	609.06,
	609.52,
	607.76,
	608.81,
	609.2,
	608.17,
	607.88,
	608.37,
	609.27,
	608.69,
	608.33,
	603.37,
	600.33,
	592.31,
	592.43,
	592.5,
	591.56,
	592.21,
	594.3,
	594.29,
	593.4,
	593.25,
	594.53,
	594.85,
	597.23,
	596.2,
	595.5,
	593.74,
	591.08,
	594.12,
	595,
	591.83,
	590.93,
	585.57,
	588.02,
	587.75,
	584.22,
	583.38,
	586.86,
	587.51,
	589.39,
	589.98,
	591.01,
	590.7,
	591.73,
	592.09,
	590.98,
	591.22,
	591.04,
	590.08,
	588.41,
	590.28,
	590.95,
	589.63,
	591.09,
	591.04,
	590.03,
	590.92,
	589.92,
	589.77,
	587.7,
	589.1,
	591.42,
	595.3,
	594.44,
	593.91,
	593.85,
	593.08,
	592.17,
	591.94,
	592.46,
	591.24,
	590.66,
	590.42,
	590.24,
	589.58,
	589.38,
	590.48,
	590,
	589.55,
	589.3,
	586.89,
	584.95,
	588.32,
	583.45,
	586.51,
	586.6,
	586.03,
	584.14,
	583.92,
	585.94,
	584.76,
	587.6,
	588.47,
	589.39,
	588.17,
	588.99,
	588.52,
	589.23,
	589.61,
	590.62,
	589.68,
	589.88,
	590.22,
	589.78,
	589.34,
	588.11,
	588.63,
	588.27,
	588.8,
	587.96,
	589.45,
	589.99,
	591,
	590.82,
	591.62,
	591.17,
	591.14,
	590.68,
	590.78,
	591.87,
	591.02,
	590.61,
	589.85,
	589.66,
	589.83,
	591.09,
	589.6,
	590.89,
	590.83,
	589.29,
	588.84,
	588.51,
	588.56,
	588.26,
	588.83,
	587.95,
	587.32,
	588.3,
	588.96,
	588.54,
	588.87,
	589.12,
	589.01,
	587.54,
	585.49,
	584.81,
	584.11,
	584.77,
	584.68,
	585.07,
	582.69,
	580.38,
	580.83,
	580.53,
	577.67,
	582.09,
	579.3,
	578.17,
	575.88,
	576.2,
	577.53,
	572.94,
	572.63,
	570.55,
	569.16,
	573.56,
	573.78,
	574.3,
	574.69,
	575.85,
	575.59,
	574.36,
	571.4,
	571.9,
	569.9,
	571.62,
	567.24,
	563.9,
	564.1,
	565.2,
	567.49,
	567.81,
	568.34,
	570.68,
	571.31,
	570.79,
	571.36,
	572.66,
	571.79,
	570.58,
	570.63,
	572.96,
	572.04,
	571.64,
	570.91,
	570.17,
	570.94,
	572.12,
	572.18,
	569.6,
	567.31,
	568.21,
	570.32,
	572.85,
	572.21,
	572.63,
	572.74
]`)

func buildKLines(prices []fixedpoint.Value) (klines []types.KLine) {
	for _, p := range prices {
		klines = append(klines, types.KLine{Close: p})
	}

	return klines
}

func Test_calculateEWMA(t *testing.T) {
	type args struct {
		allKLines []types.KLine
		priceF    types.KLineValueMapper
		window    int
	}
	var input []fixedpoint.Value
	if err := json.Unmarshal(ethusdt5m, &input); err != nil {
		panic(err)
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "ETHUSDT EMA 7",
			args: args{
				allKLines: buildKLines(input),
				priceF:    types.KLineClosePriceMapper,
				window:    7,
			},
			want: 571.72, // with open price, binance desktop returns 571.45, trading view returns 570.8957, for close price, binance mobile returns 571.72
		},
		{
			name: "ETHUSDT EMA 25",
			args: args{
				allKLines: buildKLines(input),
				priceF:    types.KLineClosePriceMapper,
				window:    25,
			},
			want: 571.30,
		},
		{
			name: "ETHUSDT EMA 99",
			args: args{
				allKLines: buildKLines(input),
				priceF:    types.KLineClosePriceMapper,
				window:    99,
			},
			want: 577.62, // binance mobile uses 577.58
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateKLinesEMA(tt.args.allKLines, tt.args.priceF, tt.args.window)
			got = math.Trunc(got*100.0) / 100.0
			if got != tt.want {
				t.Errorf("CalculateKLinesEMA() = %v, want %v", got, tt.want)
			}
		})
	}
}
