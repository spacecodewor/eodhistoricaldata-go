package eodhistoricaldata

import (
	"testing"
)

func TestProfile(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	p, err := APIClient.Fundamental.Profile("DJI.INDX")
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(p)
}
