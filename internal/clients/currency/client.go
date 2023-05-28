package currency

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mercadolibre/fury_bootcamp-go-demo/internal/clients"
	"github.com/mercadolibre/go-meli-toolkit/gorelic"
	"github.com/mercadolibre/go-meli-toolkit/goutils/apierrors"
	"github.com/mercadolibre/go-meli-toolkit/goutils/logger"
	"github.com/mercadolibre/go-meli-toolkit/restful/rest"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Client struct {
	restClient *rest.RequestBuilder
}

func NewClient(config clients.Config) APIClient {
	customPool := &rest.CustomPool{
		MaxIdleConnsPerHost: 100,
	}
	headers := make(http.Header)

	restClient := &rest.RequestBuilder{
		BaseURL:        config.BaseURL,
		Headers:        headers,
		Timeout:        3 * time.Second,
		ContentType:    rest.JSON,
		EnableCache:    false,
		DisableTimeout: false,
		CustomPool:     customPool,
		MetricsConfig:  rest.MetricsReportConfig{TargetId: "items-api"},
	}
	return Client{restClient: restClient}
}

func (c Client) GetpriceInfo(ctx context.Context, currencyID string) (float64, apierrors.ApiError) {
	var response *rest.Response
	txn := newrelic.FromContext(ctx)
	query := clients.Query()
	query.Add("from", currencyID)
	query.Add("to", "USD")

	uri, err := clients.BuildURL([]string{"/currency_conversions/search", ""}, query)
	if err != nil {
		return 0.0, apierrors.NewInternalServerApiError("error parsing items URL", err)
	}
	gorelic.WrapExternalSegment(txn, c.restClient.BaseURL+".currency_conversions.get", func() {
		response = c.restClient.Get(uri, rest.Context(ctx))
	})

	if response.Err != nil || response.Response == nil {
		txn.AddAttribute("url", uri)
		errMsg := "Unexpected error getting rate"
		logger.Errorf(fmt.Sprintf("[currencyID:%s] %s, url: %s", currencyID, errMsg, uri), response.Err)
		return 0.0, apierrors.NewInternalServerApiError(errMsg, response.Err)
	}

	if response.StatusCode == http.StatusNotFound {
		return 0.0, apierrors.NewNotFoundApiError("item not found")
	}

	if response.StatusCode != http.StatusOK {
		txn.AddAttribute("url", uri)
		errMsg := "Unexpected error getting item response"
		logger.Errorf(fmt.Sprintf("[currencyID:%s] %s, url: %s, status code: %d", currencyID, errMsg, uri, response.StatusCode), response.Err)
		return 0.0, apierrors.NewApiError(errMsg, "Unexpected error getting item", response.StatusCode, apierrors.CauseList{})
	}
	var currencyConvert currencyConvert
	rawItem := response.Bytes()
	if unmarshalError := json.Unmarshal(rawItem, &currencyConvert); unmarshalError != nil {
		txn.AddAttribute("config_value", string(rawItem))
		errMsg := "Unexpected error unmarshalling item"
		logger.Errorf(fmt.Sprintf("[currencyID:%s] %s, value: %s", currencyID, errMsg, string(rawItem)), unmarshalError)
		return 0.0, apierrors.NewInternalServerApiError(errMsg, unmarshalError)
	}

	return currencyConvert.Rate, nil
}
