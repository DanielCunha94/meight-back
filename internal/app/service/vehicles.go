package service

import "github.com/DanielCunha94/Meight-backend/internal/app/domain"

func (a *Aggregator) CreateVehicle(vehicle *domain.Vehicle) (uint, error) {
	err := vehicle.Validate()
	if err != nil {
		return 0, err
	}

	return a.repo.CreateVehicle(vehicle)
}

func (a *Aggregator) GetVehicles() ([]domain.Vehicle, error) {
	return a.repo.GetVehicles()
}
