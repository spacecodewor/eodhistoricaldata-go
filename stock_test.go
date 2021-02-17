package fmpcloud

import (
	"testing"
	"time"
)

// Varibles for test check cases
var (
	retryCount         = 5
	retryWaitTime      = 1 * time.Second
	testCaseSymbolList = []string{"AAPL", "ES.INDX", "NQ.INDX", "VTI"}
	testCaseAPIConfig  = Config{Debug: true, Timeout: 60, RetryCount: &retryCount, RetryWaitTime: &retryWaitTime}
)

func TestQuote(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.Quote("AAPL")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestBatchQuote(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.BatchQuote(testCaseSymbolList)
	if err != nil {
		t.Fatal(err.Error())
	}
}
