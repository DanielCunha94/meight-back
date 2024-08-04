package api

import (
	"github.com/DanielCunha94/Meight-backend/internal/app/service"
	"github.com/DanielCunha94/Meight-backend/pkg/sse"
	"github.com/gin-gonic/gin"
)

type API struct {
	service service.Service
	sse     sse.Handler
}

func NewAPI(service service.Service, handler sse.Handler) *API {
	return &API{service: service, sse: handler}
}

func (api API) InitRouter(router *gin.RouterGroup) {
	api.initVehiclesRouter(router)
	api.initOrdersRouter(router)
	api.initAssignmentsRouter(router)
	api.initEventsRouter(router)
}

func (api API) initEventsRouter(router *gin.RouterGroup) {
	assignments := router.Group("events")

	assignments.GET("", api.sse.Serve)
}
