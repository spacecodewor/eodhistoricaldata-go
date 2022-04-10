package eodhistoricaldata

import (
	"testing"
)

func TestProfile(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Fundamental.Profile("DJI.INDX")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestEconomicCalendar(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Fundamental.EconomicCalendar(nil, nil, nil, nil, 0, 100)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestInsiderTransactionList(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Fundamental.InsiderTransactionList(nil, nil, nil, 100)
	if err != nil {
		t.Fatal(err.Error())
	}
}
