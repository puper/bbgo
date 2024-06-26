// Code generated by "callbackgen -type ActiveOrderBook"; DO NOT EDIT.

package bbgo

import (
	"github.com/puper/bbgo/pkg/types"
)

func (b *ActiveOrderBook) OnNew(cb func(o types.Order)) {
	b.newCallbacks = append(b.newCallbacks, cb)
}

func (b *ActiveOrderBook) EmitNew(o types.Order) {
	for _, cb := range b.newCallbacks {
		cb(o)
	}
}

func (b *ActiveOrderBook) OnFilled(cb func(o types.Order)) {
	b.filledCallbacks = append(b.filledCallbacks, cb)
}

func (b *ActiveOrderBook) EmitFilled(o types.Order) {
	for _, cb := range b.filledCallbacks {
		cb(o)
	}
}

func (b *ActiveOrderBook) OnCanceled(cb func(o types.Order)) {
	b.canceledCallbacks = append(b.canceledCallbacks, cb)
}

func (b *ActiveOrderBook) EmitCanceled(o types.Order) {
	for _, cb := range b.canceledCallbacks {
		cb(o)
	}
}
