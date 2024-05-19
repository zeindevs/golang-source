package api

import (
	"encoding/json"
	"net/http"

	"github.com/zeindevs/gocommerce/store"
	"github.com/zeindevs/gocommerce/types"
)

type ProductCreateRequest struct {
	SKU  string `json:"sku"`
	Name string `json:"name"`
}

type ProductHandler struct {
	store store.ProductStorer
}

func NewProductHandler(pStore store.ProductStorer) *ProductHandler {
	return &ProductHandler{
		store: pStore,
	}
}

func (h *ProductHandler) HandlePostProduct() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var pcr ProductCreateRequest
		if err := json.NewDecoder(r.Body).Decode(&pcr); err != nil {
			WriteJSON(w, http.StatusNotFound, map[string]any{"message": err.Error()})
			return
		}

		p, err := types.NewProductFormRequest(&types.ProductCreateRequest{
			SKU:  pcr.SKU,
			Name: pcr.Name,
		})
		if err != nil {
			WriteJSON(w, http.StatusNotFound, map[string]any{"message": err.Error()})
			return
		}

		if err := h.store.Insert(r.Context(), p); err != nil {
			WriteJSON(w, http.StatusNotFound, map[string]any{"message": err.Error()})
			return
		}

		WriteJSON(w, http.StatusOK, &p)
	}
}

func (h *ProductHandler) HandleGetProductByID() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		p, err := h.store.GetByID(r.Context(), id)
		if err != nil {
			WriteJSON(w, http.StatusNotFound, map[string]any{"message": err.Error()})
			return
		}

		WriteJSON(w, http.StatusOK, p)
	}
}

func (h *ProductHandler) HandleGetAllProduct() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		p, err := h.store.GetAll(r.Context())
		if err != nil {
			WriteJSON(w, http.StatusNotFound, map[string]any{"message": err.Error()})
			return
		}

		WriteJSON(w, http.StatusOK, p)
	}
}
