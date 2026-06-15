package product

import (
	"log"
	"net/http"

	"github.com/leosanner/ecommerce-go/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(s Service) *handler {
	return &handler{
		service: s,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.ListProducts(r.Context())

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, products)
}

func (h *handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	val, match := ctx.Value("id").(int64)

	if !match {
		json.Write(w, http.StatusBadRequest, "invalid id type")
		return
	}

	if product, err := h.service.GetProductById(ctx, val); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	} else {
		json.Write(w, http.StatusOK, product)
	}
}
