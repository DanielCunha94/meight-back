package domain

import (
	"fmt"
	"github.com/DanielCunha94/Meight-backend/pkg/errors"
	"github.com/go-playground/validator/v10"
	"time"
)

type Vehicle struct {
	ID                uint `gorm:"primarykey"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Plate             string  `gorm:"uniqueIndex" validate:"required"`
	MaxWeightCapacity float64 `validate:"required"`
}

func (v *Vehicle) Validate() error {
	validate := validator.New()
	err := validate.Struct(v)
	if err != nil {

		return errors.NewBadRequest(fmt.Sprintf("invalid vehicle :%s", err.Error()))
	}
	return nil
}
