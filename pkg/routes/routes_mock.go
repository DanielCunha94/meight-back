package routes

import "github.com/DanielCunha94/Meight-backend/internal/app/domain"

type RouteMock struct {
}

func NewRouteMock() *RouteMock {
	return &RouteMock{}
}
func (rm *RouteMock) SortByFastestRoute(orders []*domain.Order) ([]*domain.Order, error) {
	return orders, nil
}
