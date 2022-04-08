package eodhistoricaldata

import (
	"encoding/json"
	"fmt"

	"github.com/spacecodewor/eodhistoricaldata-go/objects"
)

// Url const for request
const (
	urlAPIExchangesList      = "/exchanges-list"
	urlAPIExchangeSymbolList = "/exchange-symbol-list/%s"
	urlAPIExchangeDetails    = "/exchange-details/%s"
)

// Exchange client
type Exchange struct {
	Client *HTTPClient
}

// ExchangeSymbolList - Get List of Tickers (Exchange Symbols)
func (e *Exchange) ExchangeSymbolList(exchange string) (sList []objects.Symbol, err error) {
	data, err := e.Client.Get(fmt.Sprintf(urlAPIExchangeSymbolList, exchange), nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// ExchangeList - Get List of Exchanges
func (e *Exchange) ExchangeList() (eList []objects.Exchange, err error) {
	data, err := e.Client.Get(urlAPIExchangesList, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &eList)
	if err != nil {
		return nil, err
	}

	return eList, nil
}

// ExchangeDetail - Get Exchange Details and Trading Hours
func (e *Exchange) ExchangeDetail(exchange string) (ex *objects.ExchangeDetail, err error) {
	data, err := e.Client.Get(fmt.Sprintf(urlAPIExchangeDetails, exchange), nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &ex)
	if err != nil {
		return nil, err
	}

	return ex, nil
}
