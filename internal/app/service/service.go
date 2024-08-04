package service

import (
	"github.com/DanielCunha94/Meight-backend/internal/app/domain"
	"github.com/DanielCunha94/Meight-backend/internal/app/repository"
	"github.com/DanielCunha94/Meight-backend/pkg/routes"
	"github.com/DanielCunha94/Meight-backend/pkg/sse"
)

type Service interface {
	CreateVehicle(vehicle *domain.Vehicle) (uint, error)
	GetVehicles() ([]domain.Vehicle, error)

	CreateOrder(order *domain.Order) (uint, error)
	GetOrders(filters map[string]interface{}) ([]domain.Order, error)
	CompleteOrder(id string) error
	UpdateOrderObservations(id string, observations string) error
	CreateAssigment(assigment *domain.Assignment) (uint, error)
	GetAssignmentByPlateAndDate(plate string, date string) (*domain.Assignment, error)
}

type Aggregator struct {
	repo   repository.Repository
	routes routes.RouteFinder
	sse    sse.Events
}

func NewService(repo repository.Repository, finder routes.RouteFinder, sse sse.Events) *Aggregator {
	return &Aggregator{
		repo:   repo,
		routes: finder,
		sse:    sse,
	}
}
