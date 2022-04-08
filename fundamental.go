package eodhistoricaldata

import (
	"encoding/json"
	"fmt"

	"github.com/spacecodewor/eodhistoricaldata-go/objects"
)

// Url const for request
const (
	urlAPIFundamentalProfile = "/fundamentals/%s"
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
