package currency

import (
	"context"

	"github.com/mercadolibre/go-meli-toolkit/goutils/apierrors"
)

type currencyConvert struct {
	ConvertUSD
}

type ConvertUSD struct {
	CurrencyBase  string  `json:"currency_base"`
	CurrencyQuote string  `json:"currency_quote"`
	Ratio         float64 `json:"ratio"`
	Rate          float64 `json:"rate"`
	InvRate       float64 `json:"inv_rate"`
	CreationDate  string  `json:"creation_date"`
	ValidUntil    string  `json:"valid_until"`
}

type ItemResponse struct {
	ItemID    string  `json:"id"`
	Price_usd float64 `json:"price_usd"`
}

type IService interface {
	GetRate(ctx context.Context, currencyID string) (float64, apierrors.ApiError)
}

type APIClient interface {
	GetpriceInfo(ctx context.Context, currencyID string) (float64, apierrors.ApiError)
}
