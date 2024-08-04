package models

import "github.com/DanielCunha94/Meight-backend/internal/app/domain"

type Vehicle struct {
	ID                uint    `json:"id"`
	Plate             string  `json:"plate"`
	MaxWeightCapacity float64 `json:"maxWeightCapacity"`
}

func ToDomainVehicle(vehicle *Vehicle) *domain.Vehicle {
	return &domain.Vehicle{ID: vehicle.ID, Plate: vehicle.Plate, MaxWeightCapacity: vehicle.MaxWeightCapacity}
}

func FromDomainVehicle(vehicle *domain.Vehicle) *Vehicle {
	return &Vehicle{ID: vehicle.ID, Plate: vehicle.Plate, MaxWeightCapacity: vehicle.MaxWeightCapacity}
}
