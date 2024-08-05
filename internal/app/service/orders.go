package service

import (
	"fmt"
	"github.com/DanielCunha94/Meight-backend/internal/app/domain"
	"github.com/DanielCunha94/Meight-backend/internal/app/repository"
)

func (a *Aggregator) CreateOrder(order *domain.Order) (uint, error) {
	return a.repo.CreateOrder(order)
}

func (a *Aggregator) GetOrders(filters map[string]interface{}) ([]domain.Order, error) {
	return a.repo.GetOrders(filters)
}

func (a *Aggregator) CompleteOrder(orderID string) error {
	return a.repo.Transaction(func(repo repository.Repository) error {
		order, err := a.repo.GetOrderByID(orderID)
		if err != nil {
			return err
		}
		order.UpdateIsCompleted()
		err = a.repo.UpdateOrder(order)
		if err != nil {
			return err
		}

		assignment, err := a.repo.GetAssignmentByID(order.AssignmentID)
		if err != nil {
			return err
		}

		assignment.SubtractFromCurrentCapacity(order.Weight)

		err = a.repo.UpdateAssignment(assignment)
		if err != nil {
			return err
		}

		return nil
	})
}

func (a *Aggregator) UpdateOrderObservations(orderID string, observations string) error {
	order, err := a.repo.GetOrderByID(orderID)
	if err != nil {
		return err
	}

	order.UpdateObservations(observations)

	err = a.repo.UpdateOrder(order)
	if err != nil {
		return err
	}

	a.sse.Publish("orders", []byte(fmt.Sprintf("order %s updated", orderID)))
	return nil
}
