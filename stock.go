package eodhistoricaldata

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/spacecodewor/eodhistoricaldata-go/objects"
)

// Url const for request
const (
	urlAPIQuote = "/real-time/%s"
	urlAPIEOD   = "/eod/%s"
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

// EOD - get historical stock price data
func (s *Stock) EOD(symbol string, period string, from, to *time.Time) (qList []objects.Candle, err error) {
	reqParam := map[string]string{"period": period}
	if from != nil {
		reqParam["from"] = from.Format("2006-01-02")
	}

	if to != nil {
		reqParam["to"] = to.Format("2006-01-02")
	}

	data, err := s.Client.Get(fmt.Sprintf(urlAPIEOD, symbol), reqParam)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &qList)
	if err != nil {
		return nil, err
	}

	return qList, nil
}
