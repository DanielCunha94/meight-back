package repository

import (
	"github.com/DanielCunha94/Meight-backend/internal/app/domain"
	appErrors "github.com/DanielCunha94/Meight-backend/pkg/errors"
	"gorm.io/gorm"
)

type Repository interface {
	Transaction(fn func(repo Repository) error) error

	CreateVehicle(vehicle *domain.Vehicle) (uint, error)
	GetVehicles() ([]domain.Vehicle, error)
	GetVehicleByPlate(plate string) (*domain.Vehicle, error)

	CreateOrder(order *domain.Order) (uint, error)
	GetOrders(filters map[string]interface{}) ([]domain.Order, error)
	UpdateOrder(order *domain.Order) error
	GetOrderByID(orderID string) (*domain.Order, error)

	CreateAssigment(assigment *domain.Assignment) (uint, error)
	GetAssignmentByPlateAndDate(plate string, date string) (*domain.Assignment, error)
	GetAssignmentByID(id uint) (*domain.Assignment, error)
	UpdateAssignment(assignment *domain.Assignment) error
}

type SQLRepository struct {
	db *gorm.DB
}

func NewSQLRepository(db *gorm.DB) (*SQLRepository, error) {
	err := db.AutoMigrate(&domain.Vehicle{})
	if err != nil {
		return nil, appErrors.NewInternalServer(err.Error())
	}

	err = db.AutoMigrate(&domain.Order{})
	if err != nil {
		return nil, appErrors.NewInternalServer(err.Error())
	}

	err = db.AutoMigrate(&domain.Assignment{})
	if err != nil {
		return nil, appErrors.NewInternalServer(err.Error())
	}

	return &SQLRepository{db: db}, nil
}

func (s *SQLRepository) withTx(tx *gorm.DB) *SQLRepository {
	return &SQLRepository{
		db: tx,
	}
}

func (s *SQLRepository) Transaction(fn func(repo Repository) error) error {
	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	repo := s.withTx(tx)
	err := fn(repo)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
