package domain

import (
	"fmt"
	"github.com/DanielCunha94/Meight-backend/pkg/errors"
	"github.com/go-playground/validator/v10"
	"time"
)

type Order struct {
	ID           uint `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Latitude     float64 `validate:"required"`
	Longitude    float64 `validate:"required"`
	Weight       float64 `validate:"required"`
	Observations string
	IsCompleted  bool
	AssignmentID uint
}

func (o *Order) Validate() error {
	validate := validator.New()
	err := validate.Struct(o)
	if err != nil {
		return errors.NewBadRequest(fmt.Sprintf("invalid order :%s", err.Error()))
	}
	return nil
}

func (o *Order) UpdateObservations(observations string) {
	o.Observations = observations
}

func (o *Order) UpdateIsCompleted() {
	o.IsCompleted = true
}
