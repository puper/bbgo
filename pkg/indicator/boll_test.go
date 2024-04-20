package indicator

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/puper/bbgo/pkg/fixedpoint"
	"github.com/puper/bbgo/pkg/types"
)

/*
python:

import numpy as np
import pandas as pd

np.random.seed(1)

window = 14
n = 100

s = pd.Series(10 + np.sin(2 * np.pi * np.arange(n) / n) + np.random.rand(n))
print(s.tolist())

std = s.rolling(window).std()
ma = s.rolling(window).mean()

boll_up = ma + std
boll_down = ma - std
print(boll_up)
print(boll_down)
*/
func TestBOLL(t *testing.T) {
	var Delta = 4e-2
	var randomPrices = []byte(`[10.417022004702574, 10.783115012971471, 10.12544760838165, 10.489713887217565, 10.395445777981967, 10.401355589143744, 10.55438476406235, 10.77134001860812, 10.878521148332386, 11.074643528982353, 11.006979766695768, 11.322643490145449, 10.888999355660205, 11.607086063812357, 10.797900835973715, 11.47948450455335, 11.261632727869141, 11.434996508489617, 11.045213991061253, 11.12787797497313, 11.75180108497069, 11.936844736848029, 11.295711428887932, 11.684437316983791, 11.87441588072431, 11.894606663503847, 11.08307093979805, 11.03116948454736, 11.152117670293258, 11.846725664558043, 11.049403350128204, 11.350884110893302, 11.862716582616521, 11.40947196501688, 11.536205039452488, 11.12453262538101, 11.457014170457374, 11.563594299318783, 10.70283538327288, 11.387568304693657, 11.576646341198968, 11.283992449358836, 10.76219766616612, 11.215058620016562, 10.471350559262321, 10.756910520550854, 11.157285390257952, 10.480995462959404, 10.413108572150653, 10.192819091647591, 10.019366957870297, 10.616045013410577, 10.086294882435753, 10.078165344786502, 10.242883272115485, 9.744345550742134, 10.205993052807335, 9.720949283340737, 10.107551862801568, 10.163931565041935, 9.514549176535352, 9.776631998070878, 10.009853051799057, 9.685210642105492, 9.279440216170297, 9.726879411540565, 9.819466719717774, 9.638582432014445, 10.039767703524793, 9.656778554613743, 9.95234539899273, 9.168891543017606, 9.15698909652207, 9.815276587395047, 9.399650108557262, 9.165354197116933, 9.929481851967761, 9.355651158431028, 9.768524852407467, 9.75741482422182, 9.932249574910657, 9.693895721167358, 9.846115381561317, 9.47259166193398, 9.425599966263011, 10.086869223821118, 9.657577947095504, 10.235871419726973, 9.97889439188976, 9.984271730460431, 9.526960720660902, 10.413662463728075, 9.968158459378225, 10.152610322822058, 10.040012250076602, 9.92800998586808, 10.654689633397398, 10.386298172086562, 9.877537093466854, 10.55435439409141]`)
	var input []fixedpoint.Value
	if err := json.Unmarshal(randomPrices, &input); err != nil {
		panic(err)
	}
	tests := []struct {
		name   string
		kLines []types.KLine
		window int
		k      float64
		up     float64
		down   float64
	}{
		{
			name:   "random_case",
			kLines: buildKLines(input),
			window: 14,
			k:      1,
			up:     10.421434,
			down:   9.772696,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			boll := BOLL{IntervalWindow: types.IntervalWindow{Window: tt.window}, K: tt.k}
			boll.CalculateAndUpdate(tt.kLines)
			assert.InDelta(t, tt.up, boll.UpBand.Last(0), Delta)
			assert.InDelta(t, tt.down, boll.DownBand.Last(0), Delta)
		})
	}

}
