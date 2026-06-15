package product

import (
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
	sample := []string{"ola"}

	json.Write(w, http.StatusOK, sample)
}
