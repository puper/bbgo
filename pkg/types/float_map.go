package types

import "github.com/puper/bbgo/pkg/datatype/floats"

var _ Series = floats.Slice([]float64{}).Addr()
