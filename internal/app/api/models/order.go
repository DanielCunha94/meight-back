package models

import "github.com/DanielCunha94/Meight-backend/internal/app/domain"

type Order struct {
	ID           uint    `json:"id"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Weight       float64 `json:"weight"`
	Observations string  `json:"observations"`
	IsCompleted  bool    `json:"isCompleted"`
	AssignmentID uint    `json:"assignmentId"`
}

func ToDomainOrder(order *Order) *domain.Order {
	return &domain.Order{
		ID:           order.ID,
		Latitude:     order.Latitude,
		Longitude:    order.Longitude,
		Weight:       order.Weight,
		Observations: order.Observations,
		IsCompleted:  order.IsCompleted,
	}
}

func FromDomainOrder(order *domain.Order) *Order {
	return &Order{
		ID:           order.ID,
		Latitude:     order.Latitude,
		Longitude:    order.Longitude,
		Weight:       order.Weight,
		Observations: order.Observations,
		IsCompleted:  order.IsCompleted,
		AssignmentID: order.AssignmentID,
	}
}
