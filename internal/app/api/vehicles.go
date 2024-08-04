package api

import (
	"github.com/DanielCunha94/Meight-backend/internal/app/api/models"
	"github.com/DanielCunha94/Meight-backend/pkg/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (api API) initVehiclesRouter(router *gin.RouterGroup) {
	vehicles := router.Group("vehicles")

	vehicles.POST("", api.handleCreateVehicle)
	vehicles.GET("", api.handleGetVehicles)

}

func (api API) handleCreateVehicle(ctx *gin.Context) {
	var vehicle models.Vehicle

	err := ctx.ShouldBind(&vehicle)
	if err != nil {
		_ = ctx.Error(errors.NewBadRequest(err.Error()))
		return
	}

	id, err := api.service.CreateVehicle(models.ToDomainVehicle(&vehicle))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (api API) handleGetVehicles(ctx *gin.Context) {
	vehicles, err := api.service.GetVehicles()
	if err != nil {
		_ = ctx.Error(err)
		return

	}

	var apiVehicles []models.Vehicle
	for _, vehicle := range vehicles {
		apiVehicles = append(apiVehicles, *models.FromDomainVehicle(&vehicle))
	}

	ctx.JSON(http.StatusOK, apiVehicles)
}
