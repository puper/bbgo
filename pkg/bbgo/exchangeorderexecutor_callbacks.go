// Code generated by "callbackgen -type ExchangeOrderExecutor"; DO NOT EDIT.

package bbgo

import (
	"github.com/puper/bbgo/pkg/types"
)

func (e *ExchangeOrderExecutor) OnTradeUpdate(cb func(trade types.Trade)) {
	e.tradeUpdateCallbacks = append(e.tradeUpdateCallbacks, cb)
}

func (e *ExchangeOrderExecutor) EmitTradeUpdate(trade types.Trade) {
	for _, cb := range e.tradeUpdateCallbacks {
		cb(trade)
	}
}

func (e *ExchangeOrderExecutor) OnOrderUpdate(cb func(order types.Order)) {
	e.orderUpdateCallbacks = append(e.orderUpdateCallbacks, cb)
}

func (e *ExchangeOrderExecutor) EmitOrderUpdate(order types.Order) {
	for _, cb := range e.orderUpdateCallbacks {
		cb(order)
	}
}
