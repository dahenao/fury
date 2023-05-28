package items

import (
	"context"

	"github.com/mercadolibre/go-meli-toolkit/goutils/apierrors"
)

type Item struct {
	ItemDTO
}

type ItemDTO struct {
	ItemID     string  `json:"id"`
	Price      float64 `json:"price"`
	CurrencyID string  `json:"currency_id"`
}

type ItemUSD struct {
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
	GetItem(ctx context.Context, itemID string) (*Item, apierrors.ApiError)
}

type APIClient interface {
	GetItem(ctx context.Context, itemID string) (*ItemDTO, apierrors.ApiError)
}
