// Code generated by "callbackgen -type Stream -interface"; DO NOT EDIT.

package okex

import (
	"github.com/puper/bbgo/pkg/exchange/okex/okexapi"
)

func (s *Stream) OnKLineEvent(cb func(candle KLineEvent)) {
	s.kLineEventCallbacks = append(s.kLineEventCallbacks, cb)
}

func (s *Stream) EmitKLineEvent(candle KLineEvent) {
	for _, cb := range s.kLineEventCallbacks {
		cb(candle)
	}
}

func (s *Stream) OnBookEvent(cb func(book BookEvent)) {
	s.bookEventCallbacks = append(s.bookEventCallbacks, cb)
}

func (s *Stream) EmitBookEvent(book BookEvent) {
	for _, cb := range s.bookEventCallbacks {
		cb(book)
	}
}

func (s *Stream) OnAccountEvent(cb func(account okexapi.Account)) {
	s.accountEventCallbacks = append(s.accountEventCallbacks, cb)
}

func (s *Stream) EmitAccountEvent(account okexapi.Account) {
	for _, cb := range s.accountEventCallbacks {
		cb(account)
	}
}

func (s *Stream) OnOrderTradesEvent(cb func(orderTrades []OrderTradeEvent)) {
	s.orderTradesEventCallbacks = append(s.orderTradesEventCallbacks, cb)
}

func (s *Stream) EmitOrderTradesEvent(orderTrades []OrderTradeEvent) {
	for _, cb := range s.orderTradesEventCallbacks {
		cb(orderTrades)
	}
}

func (s *Stream) OnMarketTradeEvent(cb func(tradeDetail []MarketTradeEvent)) {
	s.marketTradeEventCallbacks = append(s.marketTradeEventCallbacks, cb)
}

func (s *Stream) EmitMarketTradeEvent(tradeDetail []MarketTradeEvent) {
	for _, cb := range s.marketTradeEventCallbacks {
		cb(tradeDetail)
	}
}

type StreamEventHub interface {
	OnKLineEvent(cb func(candle KLineEvent))

	OnBookEvent(cb func(book BookEvent))

	OnAccountEvent(cb func(account okexapi.Account))

	OnOrderTradesEvent(cb func(orderTrades []OrderTradeEvent))

	OnMarketTradeEvent(cb func(tradeDetail []MarketTradeEvent))
}
