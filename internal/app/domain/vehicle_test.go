package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestVehicleValidate(t *testing.T) {
	tests := []struct {
		name    string
		vehicle Vehicle
		wantErr bool
	}{
		{
			name: "Valid Vehicle",
			vehicle: Vehicle{
				ID:                1,
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
				Plate:             "ABC1234",
				MaxWeightCapacity: 1000.0,
			},
			wantErr: false,
		},
		{
			name: "Missing Plate",
			vehicle: Vehicle{
				ID:                2,
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
				MaxWeightCapacity: 1000.0,
			},
			wantErr: true,
		},
		{
			name: "Missing MaxWeightCapacity",
			vehicle: Vehicle{
				ID:        3,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Plate:     "XYZ5678",
			},
			wantErr: true,
		},
		{
			name: "Empty Vehicle",
			vehicle: Vehicle{
				ID:        4,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.vehicle.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
