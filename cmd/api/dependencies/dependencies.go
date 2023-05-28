package dependencies

import (
	"github.com/mercadolibre/fury_bootcamp-go-demo/internal/clients"
	"github.com/mercadolibre/fury_bootcamp-go-demo/internal/clients/currency"
	"github.com/mercadolibre/fury_bootcamp-go-demo/internal/clients/items"
	"github.com/mercadolibre/fury_bootcamp-go-demo/internal/platform/environment"
	"github.com/mercadolibre/go-meli-toolkit/goutils/apierrors"
)

type Dependencies struct {
	ItemsService    items.IService
	CurrencyService currency.IService
}

func BuildDependencies(env environment.Environment) (Dependencies, apierrors.ApiError) {
	mlBaseURL := "https://api.mercadolibre.com"
	switch env {
	case environment.Production:
		mlBaseURL = "http://internal.mercadolibre.com"
	}
	itemsClient := items.NewClient(clients.Config{BaseURL: mlBaseURL})
	itemsService := items.NewService(itemsClient)

	currencyClient := currency.NewClient(clients.Config{BaseURL: mlBaseURL})
	currencyService := currency.NewService(currencyClient)

	return Dependencies{
		ItemsService:    itemsService,
		CurrencyService: currencyService,
	}, nil
}
