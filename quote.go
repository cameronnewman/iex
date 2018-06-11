package iex

import (
	"errors"
	"net/url"
)

// Quote struct
type Quote struct {
	Symbol           string  `json:"symbol"`
	CompanyName      string  `json:"companyName"`
	PrimaryExchange  string  `json:"primaryExchange"`
	Sector           string  `json:"sector"`
	CalculationPrice string  `json:"calculationPrice"`
	Open             float64 `json:"open"`
	OpenTime         int64   `json:"openTime"`
	Close            float64 `json:"close"`
	CloseTime        int64   `json:"closeTime"`
	High             float64 `json:"high"`
	Low              float64 `json:"low"`
	LatestPrice      float64 `json:"latestPrice"`
	LatestSource     string  `json:"latestSource"`
	LatestTime       string  `json:"latestTime"`
	LatestUpdate     int64   `json:"latestUpdate"`
	LatestVolume     int     `json:"latestVolume"`
	IexRealtimePrice float64 `json:"iexRealtimePrice"`
	IexRealtimeSize  int     `json:"iexRealtimeSize"`
	IexLastUpdated   int64   `json:"iexLastUpdated"`
	DelayedPrice     float64 `json:"delayedPrice"`
	DelayedPriceTime int64   `json:"delayedPriceTime"`
	PreviousClose    float64 `json:"previousClose"`
	Change           float64 `json:"change"`
	ChangePercent    float64 `json:"changePercent"`
	IexMarketPercent float64 `json:"iexMarketPercent"`
	IexVolume        int     `json:"iexVolume"`
	AvgTotalVolume   int     `json:"avgTotalVolume"`
	IexBidPrice      int     `json:"iexBidPrice"`
	IexBidSize       int     `json:"iexBidSize"`
	IexAskPrice      int     `json:"iexAskPrice"`
	IexAskSize       int     `json:"iexAskSize"`
	MarketCap        int64   `json:"marketCap"`
	PeRatio          float64 `json:"peRatio"`
	Week52High       float64 `json:"week52High"`
	Week52Low        float64 `json:"week52Low"`
	YtdChange        float64 `json:"ytdChange"`
}

// GetQuote is used to receive real-time depth of book quotations direct from IEX.
// The depth of book quotations received via DEEP provide an aggregated size
// of resting displayed orders at a price and side, and do not indicate the
// size or number of individual orders at any price level. Non-displayed
// orders and non-displayed portions of reserve orders are not represented
// in DEEP.
//
// DEEP also provides last trade price and size information.
// Trades resulting from either displayed or non-displayed orders
// matching on IEX will be reported. Routed executions will not be reported.
func (i *IEX) GetQuote(symbol string) (*Quote, error) {
	result := Quote{}

	if symbol == "" {
		return &result, errors.New("Symbol needs to be defined")
	}

	err := i.getJSON(i.endpoint("/stock/"+url.QueryEscape(symbol)+"/quote"), &result)
	if err != nil {
		return &result, err
	}
	return &result, nil
}
