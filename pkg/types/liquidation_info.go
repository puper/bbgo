package types

import "github.com/puper/bbgo/pkg/fixedpoint"

type LiquidationInfo struct {
	Symbol       string
	Side         SideType
	OrderType    OrderType
	TimeInForce  TimeInForce
	Quantity     fixedpoint.Value
	Price        fixedpoint.Value
	AveragePrice fixedpoint.Value
	OrderStatus  OrderStatus
	TradeTime    Time
}
