package domain

import (
	"fmt"
	"github.com/DanielCunha94/Meight-backend/pkg/errors"
	"github.com/go-playground/validator/v10"
	"time"
)

type Assignment struct {
	ID                    uint `gorm:"primarykey"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	Date                  string   `validate:"required" gorm:"uniqueIndex:idx_plate_date"`
	Plate                 string   `validate:"required" gorm:"uniqueIndex:idx_plate_date"`
	Orders                []*Order `validate:"gte=1,dive,required"  gorm:"foreignkey:AssignmentID"`
	MaxWeightCapacity     float64  `validate:"required"`
	CurrentWeightCapacity float64  `validate:"required"`
}

func (a *Assignment) Validate() error {
	validate := validator.New()
	err := validate.Struct(a)
	if err != nil {
		return errors.NewBadRequest(fmt.Sprintf("invalid assignment :%s", err.Error()))
	}

	_, err = time.Parse("2006-01-02", a.Date)
	if err != nil {
		return errors.NewBadRequest("invalid date")
	}

	if a.CurrentWeightCapacity > a.MaxWeightCapacity {
		return errors.NewBadRequest("currentCapacity exceeds maxWeightCapacity")
	}

	return nil
}

func (a *Assignment) SubtractFromCurrentCapacity(capacity float64) {
	a.CurrentWeightCapacity = a.CurrentWeightCapacity - capacity
}
