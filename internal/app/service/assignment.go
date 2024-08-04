package service

import (
	"github.com/DanielCunha94/Meight-backend/internal/app/domain"
	"github.com/DanielCunha94/Meight-backend/pkg/errors"
)

func (a *Aggregator) CreateAssigment(assignment *domain.Assignment) (uint, error) {
	vehicle, err := a.repo.GetVehicleByPlate(assignment.Plate)
	if err != nil {
		return 0, err
	}

	assignment.MaxWeightCapacity = vehicle.MaxWeightCapacity
	var currentCapacity float64 = 0
	for _, order := range assignment.Orders {
		currentCapacity += order.Weight
	}

	assignment.CurrentWeightCapacity = currentCapacity

	err = assignment.Validate()
	if err != nil {
		return 0, err
	}

	assignment.Orders, err = a.routes.SortByFastestRoute(assignment.Orders)
	if err != nil {
		return 0, errors.NewInternalServer(err.Error())
	}

	return a.repo.CreateAssigment(assignment)
}

func (a *Aggregator) GetAssignmentByPlateAndDate(plate string, date string) (*domain.Assignment, error) {
	return a.repo.GetAssignmentByPlateAndDate(plate, date)
}
