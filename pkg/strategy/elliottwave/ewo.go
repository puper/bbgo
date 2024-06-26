package elliottwave

import "github.com/puper/bbgo/pkg/indicator"

type ElliottWave struct {
	maSlow  *indicator.SMA
	maQuick *indicator.SMA
}

func (s *ElliottWave) Index(i int) float64 {
	return s.Last(i)
}

func (s *ElliottWave) Last(i int) float64 {
	return s.maQuick.Index(i)/s.maSlow.Index(i) - 1.0
}

func (s *ElliottWave) Length() int {
	return s.maSlow.Length()
}

func (s *ElliottWave) Update(v float64) {
	s.maSlow.Update(v)
	s.maQuick.Update(v)
}
