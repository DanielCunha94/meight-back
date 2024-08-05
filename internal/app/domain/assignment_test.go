package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssignment_Validate(t *testing.T) {
	tests := []struct {
		name            string
		assignment      Assignment
		shouldHaveError bool
	}{
		{
			name: "Valid Assignment",
			assignment: Assignment{
				Date:                  "2024-08-03",
				Plate:                 "ABC123",
				MaxWeightCapacity:     1000.0,
				CurrentWeightCapacity: 500.0,
				Orders: []*Order{
					&Order{ID: 1, Latitude: 45, Longitude: 70, Weight: 100},
				},
			},
			shouldHaveError: false,
		},
		{
			name: "Missing Required Fields",
			assignment: Assignment{
				Date:                  "",
				Plate:                 "",
				MaxWeightCapacity:     0,
				CurrentWeightCapacity: 0,
				Orders:                nil,
			},
			shouldHaveError: true,
		},
		{
			name: "Invalid Date Format",
			assignment: Assignment{
				Date:                  "03-08-2024",
				Plate:                 "ABC123",
				MaxWeightCapacity:     1000.0,
				CurrentWeightCapacity: 500.0,
				Orders: []*Order{
					{ID: 1, Latitude: 45, Longitude: 70, Weight: 100},
				},
			},
			shouldHaveError: true,
		},
		{
			name: "Current Weight Exceeds Max Weight",
			assignment: Assignment{
				Date:                  "2024-08-03",
				Plate:                 "ABC123",
				MaxWeightCapacity:     1000.0,
				CurrentWeightCapacity: 1500.0,
				Orders: []*Order{
					{ID: 1, Latitude: 45, Longitude: 70, Weight: 100},
				},
			},
			shouldHaveError: true,
		},
		{
			name: "Invalid Orders",
			assignment: Assignment{
				Date:                  "2024-08-03",
				Plate:                 "ABC123",
				MaxWeightCapacity:     1000.0,
				CurrentWeightCapacity: 500.0,
				Orders:                []*Order{},
			},
			shouldHaveError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.assignment.Validate()
			if tt.shouldHaveError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestAssignment_SubtractFromCurrentCapacity(t *testing.T) {
	assignment := &Assignment{
		Date:                  "2024-08-05",
		Plate:                 "XYZ123",
		Orders:                []*Order{{}, {}},
		MaxWeightCapacity:     100.0,
		CurrentWeightCapacity: 50.0,
	}

	assignment.SubtractFromCurrentCapacity(10.0)
	assert.Equal(t, 40.0, assignment.CurrentWeightCapacity)

	assignment.SubtractFromCurrentCapacity(5.0)
	assert.Equal(t, 35.0, assignment.CurrentWeightCapacity)
}
