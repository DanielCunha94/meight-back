package repository

import (
	"errors"
	"github.com/DanielCunha94/Meight-backend/internal/app/domain"
	appErrors "github.com/DanielCunha94/Meight-backend/pkg/errors"
	"gorm.io/gorm"
)

func (s *SQLRepository) CreateAssigment(assignment *domain.Assignment) (uint, error) {
	// TODO: move data consistency logic to service

	tx := s.db.Begin()
	result := tx.Create(assignment)
	if result.Error != nil {
		tx.Rollback()
		return 0, appErrors.NewInternalServer(result.Error.Error())
	}

	var orderIDs []uint
	for _, order := range assignment.Orders {
		orderIDs = append(orderIDs, order.ID)
	}

	if err := tx.Model(&domain.Order{}).Where("id IN ?", orderIDs).Update("assignment_id", assignment.ID).Error; err != nil {
		tx.Rollback()
		return 0, appErrors.NewInternalServer(err.Error())
	}

	tx.Commit()
	return assignment.ID, nil
}

func (s *SQLRepository) GetAssignmentByPlateAndDate(plate string, date string) (*domain.Assignment, error) {
	assignment := &domain.Assignment{}
	result := s.db.Preload("Orders").Where("plate = ? AND date = ?", plate, date).First(&assignment)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, appErrors.NewNotFound(result.Error.Error())
		} else {
			return nil, appErrors.NewInternalServer(result.Error.Error())
		}
	}

	return assignment, nil
}

func (s *SQLRepository) GetAssignmentByID(assignmentID uint) (*domain.Assignment, error) {
	assignment := &domain.Assignment{}
	result := s.db.First(&assignment, assignmentID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, appErrors.NewNotFound(result.Error.Error())
		} else {
			return nil, appErrors.NewInternalServer(result.Error.Error())
		}
	}

	return assignment, nil
}

func (s *SQLRepository) UpdateAssignment(assigment *domain.Assignment) error {
	result := s.db.Save(assigment)

	if result.Error != nil {
		return appErrors.NewInternalServer(result.Error.Error())
	}
	return nil
}
