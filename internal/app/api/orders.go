package api

import (
	"github.com/DanielCunha94/Meight-backend/internal/app/api/models"
	"github.com/DanielCunha94/Meight-backend/pkg/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (api API) initOrdersRouter(router *gin.RouterGroup) {
	orders := router.Group("orders")

	orders.POST("", api.handleCreateOrder)
	orders.GET("", api.handleGetOrders)
	orders.PATCH("/:id/observations", api.handleUpdateOrderObservations)
	orders.POST("/:id/complete", api.handleCompleteOrder)

}

func (api API) handleCreateOrder(ctx *gin.Context) {
	var order models.Order

	err := ctx.ShouldBind(&order)
	if err != nil {
		_ = ctx.Error(errors.NewBadRequest(err.Error()))
		return
	}

	id, err := api.service.CreateOrder(models.ToDomainOrder(&order))
	if err != nil {
		_ = ctx.Error(err)
		return

	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (api API) handleGetOrders(ctx *gin.Context) {

	filters := make(map[string]interface{})

	if assignmentID := ctx.Query("assignmentid"); assignmentID != "" {
		uintAssignmentID, err := strconv.ParseUint(assignmentID, 10, 0)
		if err != nil {
			_ = ctx.Error(errors.NewBadRequest(err.Error()))
			return
		}
		filters["assignment_id"] = uintAssignmentID
	}

	orders, err := api.service.GetOrders(filters)
	if err != nil {
		_ = ctx.Error(err)
		return

	}

	var apiOrders []models.Order
	for _, order := range orders {
		apiOrders = append(apiOrders, *models.FromDomainOrder(&order))
	}

	ctx.JSON(http.StatusOK, apiOrders)
}

func (api API) handleUpdateOrderObservations(ctx *gin.Context) {
	type Body struct {
		Observations string `json:"observations"`
	}

	id := ctx.Param("id")

	var body Body
	err := ctx.ShouldBind(&body)
	if err != nil {
		_ = ctx.Error(errors.NewBadRequest(err.Error()))
		return
	}

	err = api.service.UpdateOrderObservations(id, body.Observations)
	if err != nil {
		_ = ctx.Error(err)
		return

	}

}

func (api API) handleCompleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	err := api.service.CompleteOrder(id)
	if err != nil {
		_ = ctx.Error(err)
		return

	}
}
