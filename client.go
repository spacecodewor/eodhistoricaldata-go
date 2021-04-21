package fmpcloud

import (
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Config for create new API client
type Config struct {
	HTTPClient    *resty.Client
	Logger        *zap.Logger
	APIToken      string
	FMTFromat     string
	Debug         bool
	RetryCount    *int
	RetryWaitTime *time.Duration
	Timeout       int
}

// APIClient ...
type APIClient struct {
	Stock    *Stock
	Exchange *Exchange
	Logger   *zap.Logger
	Debug    bool
}

// Core params
const (
	APIURL            = "https://eodhistoricaldata.com/api"
	apiDefaultToken   = "demo"
	fmtFromat         = "json"
	apiDefaultTimeout = 25
)

// NewAPIClient creates a new API client
func NewAPIClient(cfg Config) (*APIClient, error) {
	APIClient := &APIClient{Logger: cfg.Logger, Debug: cfg.Debug}
	if APIClient.Logger == nil {
		logger, err := createNewLogger()
		if err != nil {
			return nil, errors.Wrap(err, "Error create new zap logger")
		}

		APIClient.Logger = logger
	}

	// Check set timeout param
	if cfg.Timeout == 0 {
		cfg.Timeout = apiDefaultTimeout
	}

	if len(cfg.FMTFromat) == 0 {
		cfg.FMTFromat = fmtFromat
	}

	if len(cfg.APIToken) == 0 {
		cfg.APIToken = apiDefaultToken
	}

	// Init rest client (resty)
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = resty.New()
	}

	cfg.HTTPClient.SetDebug(APIClient.Debug)
	cfg.HTTPClient.SetTimeout(time.Duration(cfg.Timeout) * time.Second)
	cfg.HTTPClient.SetHostURL(APIURL)

	HTTPClient := &HTTPClient{
		client:    cfg.HTTPClient,
		apiToken:  cfg.APIToken,
		fmtFromat: cfg.FMTFromat,
		logger:    APIClient.Logger,
	}

	if cfg.RetryCount != nil {
		HTTPClient.retryCount = cfg.RetryCount
	}

	if cfg.RetryWaitTime != nil {
		HTTPClient.retryWaitTime = cfg.RetryWaitTime
	}

	if HTTPClient.retryCount == nil || *HTTPClient.retryCount == 0 {
		retryCount := 1
		HTTPClient.retryCount = &retryCount
	}

	if HTTPClient.retryWaitTime == nil || *HTTPClient.retryWaitTime == 0 {
		retryWaitTime := 1 * time.Second
		HTTPClient.retryWaitTime = &retryWaitTime
	}

	APIClient.Stock = &Stock{Client: HTTPClient}
	APIClient.Exchange = &Exchange{Client: HTTPClient}

	return APIClient, nil
}

// Create new logger
func createNewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	logger, err := cfg.Build()
	if err != nil {
		return nil, errors.Wrap(err, "Logger Error: init")
	}

	return logger, nil
}
