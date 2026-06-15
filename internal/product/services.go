package product

import (
	"context"

	repo "github.com/leosanner/ecommerce-go/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListProducts(context.Context) ([]repo.Product, error)
	GetProductById(context.Context, int64) (repo.Product, error)
}

type svc struct {
	repo repo.Querier
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	return s.repo.ListProducts(ctx)
}

func (s *svc) GetProductById(ctx context.Context, id int64) (repo.Product, error) {
	return s.repo.FindProductById(ctx, id)
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}
