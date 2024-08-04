package api

import (
	"github.com/DanielCunha94/Meight-backend/internal/app/api/models"
	"github.com/DanielCunha94/Meight-backend/pkg/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (api API) initAssignmentsRouter(router *gin.RouterGroup) {
	assignments := router.Group("assignments")

	assignments.POST("", api.handleCreateAssignment)
	assignments.GET("/plate/:plate/date/:date", api.handleGetAssignmentByPlateAndDate)
}

func (api API) handleCreateAssignment(ctx *gin.Context) {
	var assignment models.Assignment
	var err error
	err = ctx.ShouldBind(&assignment)
	if err != nil {
		_ = ctx.Error(errors.NewBadRequest(err.Error()))
		return
	}

	var id uint
	id, err = api.service.CreateAssigment(models.ToDomainAssignment(&assignment))
	if err != nil {
		_ = ctx.Error(err)
		return

	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (api API) handleGetAssignmentByPlateAndDate(ctx *gin.Context) {
	plate := ctx.Param("plate")
	date := ctx.Param("date")

	assignment, err := api.service.GetAssignmentByPlateAndDate(plate, date)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, models.FromDomainAssignment(assignment))
}
