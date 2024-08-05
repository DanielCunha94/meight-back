package service

import (
	"github.com/DanielCunha94/Meight-backend/internal/app/domain"
	"github.com/DanielCunha94/Meight-backend/internal/app/repository"
	"github.com/DanielCunha94/Meight-backend/pkg/routes"
	"github.com/DanielCunha94/Meight-backend/pkg/sse"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"
)

var service Service

func TestMain(m *testing.M) {

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = db.AutoMigrate(&domain.Vehicle{}, &domain.Order{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	repo, err := repository.NewSQLRepository(db)
	if err != nil {
		log.Fatalf("failed to create repo: %v", err)
	}

	service = NewService(repo, routes.NewRouteMock(), sse.NewSSEMock())

	code := m.Run()
	os.Exit(code)

}

func TestCreateVehicleIntegration(t *testing.T) {
	vehicle := &domain.Vehicle{
		Plate:             "test",
		MaxWeightCapacity: 1000,
	}

	id, err := service.CreateVehicle(vehicle)

	assert.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetVehiclesIntegration(t *testing.T) {
	vehicles, err := service.GetVehicles()
	assert.NoError(t, err)
	assert.Len(t, vehicles, 1)
}

func TestCreateOrderIntegration(t *testing.T) {
	order := &domain.Order{
		Latitude:     45,
		Longitude:    45,
		Weight:       2000,
		Observations: "test",
	}

	id, err := service.CreateOrder(order)

	assert.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetOrdersIntegration(t *testing.T) {
	filters := map[string]interface{}{}
	orders, err := service.GetOrders(filters)
	assert.NoError(t, err)
	assert.Len(t, orders, 1)
}

func TestCreateAssignmentIntegration(t *testing.T) {
	order := &domain.Order{
		Latitude:     45,
		Longitude:    45,
		Weight:       500,
		Observations: "test",
	}

	assignment := &domain.Assignment{
		Date:                  "2024-08-05",
		Plate:                 "test",
		MaxWeightCapacity:     1000,
		CurrentWeightCapacity: 500,
		Orders:                []*domain.Order{order},
	}

	id, err := service.CreateAssigment(assignment)

	assert.NoError(t, err)
	assert.NotZero(t, id)

	id, err = service.CreateAssigment(assignment)

	assert.Error(t, err)
}

func TestGetAssigmentIntegration(t *testing.T) {
	_, err := service.GetAssignmentByPlateAndDate("test", "2024-08-05")
	assert.NoError(t, err)

	_, err = service.GetAssignmentByPlateAndDate("test2", "2024-08-05")
	assert.Error(t, err)
}

//TODO: test other services
