package fixer

import (
	"testing"

	"github.com/thrasher-corp/gocryptotrader/currency/forexprovider/base"
)

// Please set API key and apikey subscription level for correct due diligence
// testing - NOTE please be aware tests will diminish your monthly API calls

const (
	apikey    = ""
	apiKeyLvl = 3
)

var f Fixer

var isSetup bool

func setup(t *testing.T) {
	if !isSetup {
		err := f.Setup(base.Settings{})
		if err != nil {
			t.Fatal("Test Failed - Setup error", err)
		}
		isSetup = true
	}
}

func TestGetRates(t *testing.T) {
	setup(t)
	_, err := f.GetRates("EUR", "AUD")
	if err == nil {
		t.Error("test failed - fixer GetRates() error", err)
	}
}

func TestGetLatestRates(t *testing.T) {
	setup(t)
	_, err := f.GetLatestRates("EUR", "AUD")
	if err == nil {
		t.Error("test failed - fixer GetLatestRates() error", err)
	}
}

func TestGetHistoricalRates(t *testing.T) {
	setup(t)
	_, err := f.GetHistoricalRates("2013-12-24", "EUR", []string{"AUD,KRW"})
	if err == nil {
		t.Error("test failed - fixer GetHistoricalRates() error", err)
	}
}

func TestConvertCurrency(t *testing.T) {
	setup(t)
	_, err := f.ConvertCurrency("AUD", "EUR", "", 1337)
	if err == nil {
		t.Error("test failed - fixer ConvertCurrency() error", err)
	}
}

func TestGetTimeSeriesData(t *testing.T) {
	setup(t)
	_, err := f.GetTimeSeriesData("2013-12-24", "2013-12-25", "EUR", []string{"AUD,KRW"})
	if err == nil {
		t.Error("test failed - fixer GetTimeSeriesData() error", err)
	}
}

func TestGetFluctuationData(t *testing.T) {
	setup(t)
	_, err := f.GetFluctuationData("2013-12-24", "2013-12-25", "EUR", []string{"AUD,KRW"})
	if err == nil {
		t.Error("test failed - fixer GetFluctuationData() error", err)
	}
}
