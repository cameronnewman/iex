package iex

import (
	"errors"
)

// DEEP struct
type DEEP struct {
	Symbol        string        `json:"symbol"`
	MarketPercent int           `json:"marketPercent"`
	Volume        int           `json:"volume"`
	LastSalePrice int           `json:"lastSalePrice"`
	LastSaleSize  int           `json:"lastSaleSize"`
	LastSaleTime  int           `json:"lastSaleTime"`
	LastUpdated   int           `json:"lastUpdated"`
	Bids          []interface{} `json:"bids"`
	Asks          []interface{} `json:"asks"`
	SystemEvent   struct {
	} `json:"systemEvent"`
	TradingStatus struct {
		Status    string `json:"status"`
		Reason    string `json:"reason"`
		Timestamp int64  `json:"timestamp"`
	} `json:"tradingStatus"`
	OpHaltStatus struct {
		IsHalted  bool  `json:"isHalted"`
		Timestamp int64 `json:"timestamp"`
	} `json:"opHaltStatus"`
	SsrStatus struct {
		IsSSR     bool   `json:"isSSR"`
		Detail    string `json:"detail"`
		Timestamp int64  `json:"timestamp"`
	} `json:"ssrStatus"`
	SecurityEvent struct {
	} `json:"securityEvent"`
	Trades      []interface{} `json:"trades"`
	TradeBreaks []interface{} `json:"tradeBreaks"`
	Auction     struct {
	} `json:"auction"`
	OfficialPrice struct {
	} `json:"officialPrice"`
}

// GetDEEP is used to receive real-time depth of book quotations direct from IEX.
// The depth of book quotations received via DEEP provide an aggregated size
// of resting displayed orders at a price and side, and do not indicate the
// size or number of individual orders at any price level. Non-displayed
// orders and non-displayed portions of reserve orders are not represented
// in DEEP.
//
// DEEP also provides last trade price and size information.
// Trades resulting from either displayed or non-displayed orders
// matching on IEX will be reported. Routed executions will not be reported.
func (i *IEX) GetDEEP(symbol string) (*DEEP, error) {
	result := DEEP{}

	if symbol == "" {
		return &result, errors.New("Symbol needs to be defined")
	}

	err := i.getJSON(i.endpoint("/deep?symbols="+symbol), &result)
	if err != nil {
		return &result, err
	}
	return &result, nil
}
