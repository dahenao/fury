package clients

import (
	"github.com/mercadolibre/go-meli-toolkit/goutils/apierrors"
	"net/url"
	"path"
)

type Config struct {
	BaseURL string
}

func Query() url.Values {
	return make(url.Values)
}

func BuildURL(params []string, query url.Values) (string, apierrors.ApiError) {
	if len(params) == 0 {
		return "", apierrors.NewInternalServerApiError("invalid params URL", nil)
	}

	uri := &url.URL{Path: path.Join(params...)}
	if _, err := url.Parse(uri.Path); err != nil {
		return "", apierrors.NewInternalServerApiError("error parsing URL", err)
	}
	uri.RawQuery = query.Encode()

	return uri.String(), nil
}
