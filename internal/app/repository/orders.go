package repository

import (
	"errors"
	"github.com/DanielCunha94/Meight-backend/internal/app/domain"
	appErrors "github.com/DanielCunha94/Meight-backend/pkg/errors"
	"gorm.io/gorm"
)

func (s *SQLRepository) CreateOrder(order *domain.Order) (uint, error) {
	result := s.db.Create(order)
	if result.Error != nil {
		return order.ID, appErrors.NewInternalServer(result.Error.Error())
	}

	return order.ID, nil

}

func (s *SQLRepository) GetOrders(filters map[string]interface{}) ([]domain.Order, error) {
	var orders []domain.Order

	query := s.db
	for key, value := range filters {
		query = query.Where(key+" = ?", value)
	}

	result := query.Find(&orders)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, appErrors.NewNotFound(result.Error.Error())
		}
		return nil, appErrors.NewInternalServer(result.Error.Error())
	}

	return orders, nil
}

func (s *SQLRepository) GetOrderByID(orderID string) (*domain.Order, error) {
	order := &domain.Order{}
	result := s.db.First(&order, orderID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, appErrors.NewNotFound(result.Error.Error())
		} else {
			return nil, appErrors.NewInternalServer(result.Error.Error())
		}
	}

	return order, nil
}

func (s *SQLRepository) UpdateOrder(order *domain.Order) error {
	result := s.db.Save(order)

	if result.Error != nil {
		return appErrors.NewInternalServer(result.Error.Error())
	}
	return nil
}
