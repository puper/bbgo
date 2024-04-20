// Code generated by "callbackgen -type Stream"; DO NOT EDIT.

package max

import (
	"github.com/puper/bbgo/pkg/exchange/max/maxapi"
)

func (s *Stream) OnAuthEvent(cb func(e max.AuthEvent)) {
	s.authEventCallbacks = append(s.authEventCallbacks, cb)
}

func (s *Stream) EmitAuthEvent(e max.AuthEvent) {
	for _, cb := range s.authEventCallbacks {
		cb(e)
	}
}

func (s *Stream) OnBookEvent(cb func(e max.BookEvent)) {
	s.bookEventCallbacks = append(s.bookEventCallbacks, cb)
}

func (s *Stream) EmitBookEvent(e max.BookEvent) {
	for _, cb := range s.bookEventCallbacks {
		cb(e)
	}
}

func (s *Stream) OnTradeEvent(cb func(e max.PublicTradeEvent)) {
	s.tradeEventCallbacks = append(s.tradeEventCallbacks, cb)
}

func (s *Stream) EmitTradeEvent(e max.PublicTradeEvent) {
	for _, cb := range s.tradeEventCallbacks {
		cb(e)
	}
}

func (s *Stream) OnKLineEvent(cb func(e max.KLineEvent)) {
	s.kLineEventCallbacks = append(s.kLineEventCallbacks, cb)
}

func (s *Stream) EmitKLineEvent(e max.KLineEvent) {
	for _, cb := range s.kLineEventCallbacks {
		cb(e)
	}
}

func (s *Stream) OnErrorEvent(cb func(e max.ErrorEvent)) {
	s.errorEventCallbacks = append(s.errorEventCallbacks, cb)
}

func (s *Stream) EmitErrorEvent(e max.ErrorEvent) {
	for _, cb := range s.errorEventCallbacks {
		cb(e)
	}
}

func (s *Stream) OnSubscriptionEvent(cb func(e max.SubscriptionEvent)) {
	s.subscriptionEventCallbacks = append(s.subscriptionEventCallbacks, cb)
}

func (s *Stream) EmitSubscriptionEvent(e max.SubscriptionEvent) {
	for _, cb := range s.subscriptionEventCallbacks {
		cb(e)
	}
}

func (s *Stream) OnTradeUpdateEvent(cb func(e max.TradeUpdateEvent)) {
	s.tradeUpdateEventCallbacks = append(s.tradeUpdateEventCallbacks, cb)
}

func (s *Stream) EmitTradeUpdateEvent(e max.TradeUpdateEvent) {
	for _, cb := range s.tradeUpdateEventCallbacks {
		cb(e)
	}
}

func (s *Stream) OnTradeSnapshotEvent(cb func(e max.TradeSnapshotEvent)) {
	s.tradeSnapshotEventCallbacks = append(s.tradeSnapshotEventCallbacks, cb)
}

func (s *Stream) EmitTradeSnapshotEvent(e max.TradeSnapshotEvent) {
	for _, cb := range s.tradeSnapshotEventCallbacks {
		cb(e)
	}
}

func (s *Stream) OnOrderUpdateEvent(cb func(e max.OrderUpdateEvent)) {
	s.orderUpdateEventCallbacks = append(s.orderUpdateEventCallbacks, cb)
}

func (s *Stream) EmitOrderUpdateEvent(e max.OrderUpdateEvent) {
	for _, cb := range s.orderUpdateEventCallbacks {
		cb(e)
	}
}

func (s *Stream) OnOrderSnapshotEvent(cb func(e max.OrderSnapshotEvent)) {
	s.orderSnapshotEventCallbacks = append(s.orderSnapshotEventCallbacks, cb)
}

func (s *Stream) EmitOrderSnapshotEvent(e max.OrderSnapshotEvent) {
	for _, cb := range s.orderSnapshotEventCallbacks {
		cb(e)
	}
}

func (s *Stream) OnAdRatioEvent(cb func(e max.ADRatioEvent)) {
	s.adRatioEventCallbacks = append(s.adRatioEventCallbacks, cb)
}

func (s *Stream) EmitAdRatioEvent(e max.ADRatioEvent) {
	for _, cb := range s.adRatioEventCallbacks {
		cb(e)
	}
}

func (s *Stream) OnDebtEvent(cb func(e max.DebtEvent)) {
	s.debtEventCallbacks = append(s.debtEventCallbacks, cb)
}

func (s *Stream) EmitDebtEvent(e max.DebtEvent) {
	for _, cb := range s.debtEventCallbacks {
		cb(e)
	}
}

func (s *Stream) OnAccountSnapshotEvent(cb func(e max.AccountSnapshotEvent)) {
	s.accountSnapshotEventCallbacks = append(s.accountSnapshotEventCallbacks, cb)
}

func (s *Stream) EmitAccountSnapshotEvent(e max.AccountSnapshotEvent) {
	for _, cb := range s.accountSnapshotEventCallbacks {
		cb(e)
	}
}

func (s *Stream) OnAccountUpdateEvent(cb func(e max.AccountUpdateEvent)) {
	s.accountUpdateEventCallbacks = append(s.accountUpdateEventCallbacks, cb)
}

func (s *Stream) EmitAccountUpdateEvent(e max.AccountUpdateEvent) {
	for _, cb := range s.accountUpdateEventCallbacks {
		cb(e)
	}
}
