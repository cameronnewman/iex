package iex

// Symbol struct
type Symbol struct {
	Symbol    string `json:"symbol"`
	Name      string `json:"name"`
	Date      string `json:"date"`
	IsEnabled bool   `json:"isEnabled"`
	Type      string `json:"type"`
	IexID     string `json:"iexId"`
}

// GetSymbols returns an array of symbols IEX supports for trading.
// This list is updated daily as of 7:45 a.m. ET. Symbols may be added
// or removed by IEX after the list was produced.
func (i *IEX) GetSymbols() (*[]Symbol, error) {
	result := []Symbol{}

	err := i.getJSON(i.endpoint("/ref-data/symbols"), &result)
	if err != nil {
		return &result, err
	}

	return &result, nil
}
