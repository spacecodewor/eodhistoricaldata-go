package eodhistoricaldata

import (
	"testing"
)

func TestExchangeList(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Exchange.ExchangeList()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestExchangeDetail(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Exchange.ExchangeDetail("MCX")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestExchangeSymbolList(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Exchange.ExchangeSymbolList("INDX")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestNews(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Exchange.News(nil, nil, nil, nil, 100, 0)
	if err != nil {
		t.Fatal(err.Error())
	}
}
