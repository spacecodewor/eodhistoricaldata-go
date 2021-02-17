package fmpcloud

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spacecodewor/eodhistoricaldata-go/objects"
)

// Url const for request
const (
	urlAPIQuote = "real-time/%s"
)

// Stock client
type Stock struct {
	Client *HTTPClient
}

// Quote - Live (Delayed) Stock Prices API
func (s *Stock) Quote(symbol string) (q *objects.Quote, err error) {
	data, err := s.Client.Get(fmt.Sprintf(urlAPIQuote, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &q)
	if err != nil {
		return nil, err
	}

	return q, nil
}

// BatchQuote - Live (Delayed) Stock Prices API (Batch)
func (s *Stock) BatchQuote(symbolList []string) (qList []objects.Quote, err error) {
	data, err := s.Client.Get(
		fmt.Sprintf(urlAPIQuote, symbolList[0]),
		map[string]string{"s": strings.Join(symbolList[1:], ",")},
	)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &qList)
	if err != nil {
		return nil, err
	}

	return qList, nil
}
