package items

import (
	"context"

	"github.com/mercadolibre/go-meli-toolkit/goutils/apierrors"
)

type ClientMock struct {
	HandleGetItem func(ctx context.Context, itemID string) (*ItemDTO, apierrors.ApiError)
}

func NewClientMock() *ClientMock {
	return &ClientMock{}
}

func (m ClientMock) GetItem(ctx context.Context, itemID string) (*ItemDTO, apierrors.ApiError) {
	if m.HandleGetItem != nil {
		return m.HandleGetItem(ctx, itemID)
	}
	return nil, nil
}

func (m ClientMock) GetRateUSD(ctx context.Context, currencyID string) (float64, apierrors.ApiError) {
	return 0.0, nil
}
