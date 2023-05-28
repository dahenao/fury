package items

import (
	"context"

	"github.com/mercadolibre/go-meli-toolkit/goutils/apierrors"
)

type ServiceMock struct {
	HandleGetItem func(ctx context.Context, itemID string) (*Item, apierrors.ApiError)
}

func NewServiceMock() *ServiceMock {
	return &ServiceMock{}
}

func (sm ServiceMock) GetItem(ctx context.Context, itemID string) (*Item, apierrors.ApiError) {
	if sm.HandleGetItem != nil {
		return sm.HandleGetItem(ctx, itemID)
	}
	return &Item{}, nil
}

func (sm ServiceMock) GetRateUSD(ctx context.Context, currencyID string) (float64, apierrors.ApiError) {
	return 0.0, nil
}
