package models

import "github.com/DanielCunha94/Meight-backend/internal/app/domain"

type Assignment struct {
	ID                    uint     `json:"id"`
	Date                  string   `json:"date"`
	Plate                 string   `json:"plate"`
	Orders                []*Order `json:"orders"`
	MaxWeightCapacity     float64  `json:"maxWeightCapacity"`
	CurrentWeightCapacity float64  `json:"currentWeightCapacity"`
}

func ToDomainAssignment(assignment *Assignment) *domain.Assignment {
	orders := make([]*domain.Order, len(assignment.Orders))
	for i, order := range assignment.Orders {
		orders[i] = ToDomainOrder(order)
	}

	return &domain.Assignment{
		ID:                    assignment.ID,
		Date:                  assignment.Date,
		Plate:                 assignment.Plate,
		Orders:                orders,
		MaxWeightCapacity:     assignment.MaxWeightCapacity,
		CurrentWeightCapacity: assignment.CurrentWeightCapacity,
	}
}

func FromDomainAssignment(assignment *domain.Assignment) *Assignment {
	orders := make([]*Order, len(assignment.Orders))
	for i, order := range assignment.Orders {
		orders[i] = FromDomainOrder(order)
	}

	return &Assignment{
		ID:                    assignment.ID,
		Date:                  assignment.Date,
		Plate:                 assignment.Plate,
		Orders:                orders,
		MaxWeightCapacity:     assignment.MaxWeightCapacity,
		CurrentWeightCapacity: assignment.CurrentWeightCapacity,
	}
}
