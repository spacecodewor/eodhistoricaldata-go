package eodhistoricaldata

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/spacecodewor/eodhistoricaldata-go/objects"
)

// Url const for request
const (
	urlAPIExchangesList      = "/exchanges-list"
	urlAPIExchangeSymbolList = "/exchange-symbol-list/%s"
	urlAPIExchangeDetails    = "/exchange-details/%s"
	urlAPIExchangeNews       = "/news"
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

// News - financial news
func (e *Exchange) News(from, to *time.Time, s, t *string, limit, offset int) (nList []objects.News, err error) {
	reqParam := map[string]string{
		"limit":  fmt.Sprint(limit),
		"offset": fmt.Sprint(offset),
	}

	if s != nil {
		reqParam["s"] = *s
	}

	if t != nil {
		reqParam["t"] = *t
	}

	if from != nil {
		reqParam["from"] = from.Format("2006-01-02")
	}

	if to != nil {
		reqParam["to"] = to.Format("2006-01-02")
	}

	data, err := e.Client.Get(urlAPIExchangeNews, reqParam)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &nList)
	if err != nil {
		return nil, err
	}

	return nList, nil
}
