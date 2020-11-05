package model

import "time"

// CallType enum
type CallType string

const (
	// BUY action
	BUY CallType = "BUY"

	// SELL action
	SELL CallType = "SELL"
)

// TradeCall data model
type TradeCall struct {
	Action CallType
	Stock  string
	Units  int
	Price  float32
	Date   time.Time
}

// NewBuy returns instance of Buy TradeCall
func NewBuy(stock string, units int, price float32) TradeCall {
	return TradeCall{
		Action: BUY,
		Stock:  stock,
		Units:  units,
		Price:  price,
		Date:   time.Now(),
	}
}

// NewSell returns instance of Buy TradeCall
func NewSell(stock string, units int, price float32) *TradeCall {
	return &TradeCall{
		Action: SELL,
		Stock:  stock,
		Units:  units,
		Price:  price,
		Date:   time.Now(),
	}
}
