package repository

import (
	"errors"
	"github.com/DanielCunha94/Meight-backend/internal/app/domain"
	appErrors "github.com/DanielCunha94/Meight-backend/pkg/errors"
	"gorm.io/gorm"
)

func (s *SQLRepository) CreateVehicle(vehicle *domain.Vehicle) (uint, error) {
	result := s.db.Create(vehicle)

	if result.Error != nil {
		return vehicle.ID, appErrors.NewInternalServer(result.Error.Error())
	}

	return vehicle.ID, nil
}

func (s *SQLRepository) GetVehicles() ([]domain.Vehicle, error) {
	var vehicles []domain.Vehicle

	result := s.db.Find(&vehicles)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, appErrors.NewNotFound(result.Error.Error())
		}
		return nil, appErrors.NewInternalServer(result.Error.Error())
	}

	return vehicles, nil
}

func (s *SQLRepository) GetVehicleByPlate(plate string) (*domain.Vehicle, error) {
	vehicle := &domain.Vehicle{}

	result := s.db.Where("plate = ?", plate).First(vehicle)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, appErrors.NewNotFound(result.Error.Error())
		} else {
			return nil, appErrors.NewInternalServer(result.Error.Error())
		}
	}

	return vehicle, nil
}
