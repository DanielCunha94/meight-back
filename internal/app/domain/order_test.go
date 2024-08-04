package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestOrder_Validate(t *testing.T) {
	tests := []struct {
		name    string
		order   Order
		wantErr bool
	}{
		{
			name: "valid order",
			order: Order{
				Latitude:    45.0,
				Longitude:   -93.0,
				Weight:      10.5,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				IsCompleted: false,
			},
			wantErr: false,
		},
		{
			name: "missing latitude",
			order: Order{
				Longitude:   -93.0,
				Weight:      10.5,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				IsCompleted: false,
			},
			wantErr: true,
		},
		{
			name: "missing longitude",
			order: Order{
				Latitude:    45.0,
				Weight:      10.5,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				IsCompleted: false,
			},
			wantErr: true,
		},
		{
			name: "missing weight",
			order: Order{
				Latitude:    45.0,
				Longitude:   -93.0,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				IsCompleted: false,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.order.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestOrder_UpdateObservations(t *testing.T) {
	order := Order{
		Latitude:    45.0,
		Longitude:   -93.0,
		Weight:      10.5,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsCompleted: false,
	}

	newObservations := "New observations"
	order.UpdateObservations(newObservations)

	assert.Equal(t, newObservations, order.Observations)
}

func TestOrder_UpdateIsCompleted(t *testing.T) {
	order := Order{
		Latitude:    45.0,
		Longitude:   -93.0,
		Weight:      10.5,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsCompleted: false,
	}

	order.UpdateIsCompleted()

	assert.True(t, order.IsCompleted)
}
