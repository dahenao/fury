package items

import (
	"context"

	"github.com/mercadolibre/go-meli-toolkit/goutils/apierrors"
)

type Service struct {
	itemsClient APIClient
}

func NewService(itemsClient APIClient) *Service {
	return &Service{
		itemsClient: itemsClient,
	}
}

func (s *Service) GetItem(ctx context.Context, itemID string) (*Item, apierrors.ApiError) {
	itemDTO, err := s.itemsClient.GetItem(ctx, itemID)
	if err != nil || itemDTO == nil {
		return nil, err
	}
	return &Item{*itemDTO}, err
}
