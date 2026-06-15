package product

import "context"

type Service interface {
	ListProducts(context.Context) error
}

type svc struct {
}

func (s *svc) ListProducts(ctx context.Context) error {
	return nil
}

func NewService() Service {
	return &svc{}
}
