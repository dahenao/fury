package handlers

import (
	"net/http"

	"github.com/mercadolibre/fury_bootcamp-go-demo/cmd/api/dependencies"
	"github.com/mercadolibre/fury_bootcamp-go-demo/internal/clients/currency"
	"github.com/mercadolibre/fury_bootcamp-go-demo/internal/clients/items"
	"github.com/mercadolibre/fury_go-core/pkg/web"
)

type ItemHandler struct {
	itemService     items.IService
	currencyService currency.IService
}

func NewItemHandler(depend dependencies.Dependencies) *ItemHandler {
	return &ItemHandler{
		itemService:     depend.ItemsService,
		currencyService: depend.CurrencyService,
	}
}

func (i *ItemHandler) GetItemPrice(w http.ResponseWriter, r *http.Request) error {
	params := web.Params(r)

	itemID, err := params.String("id")
	if err != nil {
		return web.EncodeJSON(w, "error fetching item id from url", http.StatusInternalServerError)
	}

	item, apiErr := i.itemService.GetItem(r.Context(), itemID)
	if apiErr != nil {
		return web.EncodeJSON(w, apiErr, apiErr.Status())
	}

	return web.EncodeJSON(w, item.ItemDTO, http.StatusOK)
}

func (i *ItemHandler) GetItemPriceUSD(w http.ResponseWriter, r *http.Request) error {
	params := web.Params(r)

	itemID, err := params.String("id")
	if err != nil {
		return web.EncodeJSON(w, "error fetching item id from url", http.StatusInternalServerError)
	}

	item, apiErr := i.itemService.GetItem(r.Context(), itemID)
	if apiErr != nil {
		return web.EncodeJSON(w, apiErr, apiErr.Status())
	}

	rate, apiErr := i.currencyService.GetRate(r.Context(), item.ItemDTO.CurrencyID)
	if apiErr != nil {
		println(apiErr.Error() + "********************************************")
		return web.EncodeJSON(w, apiErr, apiErr.Status())
	}

	convertPrice := getConvertedPrice(item.Price, rate)

	var response items.ItemResponse
	response.ItemID = item.ItemID
	response.Price_usd = convertPrice

	return web.EncodeJSON(w, response, http.StatusOK)
}

func getConvertedPrice(currentPrice float64, rate float64) float64 {
	return currentPrice * rate
}
