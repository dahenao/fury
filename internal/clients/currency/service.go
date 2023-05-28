package currency

import (
	"context"

	"github.com/mercadolibre/go-meli-toolkit/goutils/apierrors"
)

type Service struct {
	currencyConvertClient APIClient
}

func NewService(currencyConvert APIClient) *Service {
	return &Service{
		currencyConvertClient: currencyConvert,
	}
}

func (s *Service) GetRate(ctx context.Context, currencyID string) (float64, apierrors.ApiError) {
	rate, err := s.currencyConvertClient.GetpriceInfo(ctx, currencyID)
	if err != nil || rate == 0.0 {
		return 0.0, err
	}
	return rate, err
}
