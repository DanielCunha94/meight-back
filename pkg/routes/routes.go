package routes

import "github.com/DanielCunha94/Meight-backend/internal/app/domain"

type RouteFinder interface {
	SortByFastestRoute(orders []*domain.Order) ([]*domain.Order, error)
}
