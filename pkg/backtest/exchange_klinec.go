package backtest

import (
	"github.com/puper/bbgo/pkg/bbgo"
	"github.com/puper/bbgo/pkg/types"
)

type ExchangeDataSource struct {
	C         chan types.KLine
	Exchange  *Exchange
	Session   *bbgo.ExchangeSession
	Callbacks []func(types.KLine, *ExchangeDataSource)
}
