package eodhistoricaldata

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/spacecodewor/eodhistoricaldata-go/objects"
)

// Url const for request
const (
	urlAPIFundamentalProfile             = "/fundamentals/%s"
	urlAPIFundamentalEconomicEvents      = "/economic-events"
	urlAPIFundamentalInsiderTransactions = "/insider-transactions"
)

// Fundamental client
type Fundamental struct {
	Client *HTTPClient
}

// Profile - get profile by symbol
func (e *Fundamental) Profile(symbol string) (p *objects.Profile, err error) {
	data, err := e.Client.Get(fmt.Sprintf(urlAPIFundamentalProfile, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// EconomicCalendar - output for Economic Events Calendar
func (e *Fundamental) EconomicCalendar(from, to *time.Time, country, comparison *string, offset, limit int) (eList []objects.EconomicEvent, err error) {
	reqParam := map[string]string{
		"limit":  fmt.Sprint(limit),
		"offset": fmt.Sprint(offset),
	}

	if comparison != nil {
		reqParam["comparison"] = *comparison
	}

	if country != nil {
		reqParam["country"] = *country
	}

	if from != nil {
		reqParam["from"] = from.Format("2006-01-02")
	}

	if to != nil {
		reqParam["to"] = to.Format("2006-01-02")
	}

	data, err := e.Client.Get(urlAPIFundamentalEconomicEvents, reqParam)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &eList)
	if err != nil {
		return nil, err
	}

	return eList, nil
}

// InsiderTransactionList - output for Insider Transaction list
func (e *Fundamental) InsiderTransactionList(from, to *time.Time, code *string, limit int) (iList []objects.InsiderTransaction, err error) {
	reqParam := map[string]string{"limit": fmt.Sprint(limit)}
	if code != nil {
		reqParam["code"] = *code
	}

	if from != nil {
		reqParam["from"] = from.Format("2006-01-02")
	}

	if to != nil {
		reqParam["to"] = to.Format("2006-01-02")
	}

	data, err := e.Client.Get(urlAPIFundamentalInsiderTransactions, reqParam)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &iList)
	if err != nil {
		return nil, err
	}

	return iList, nil
}