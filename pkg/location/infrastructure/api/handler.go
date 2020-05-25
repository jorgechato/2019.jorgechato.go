package api

import (
	"github.com/gin-gonic/gin"

	"jorgechato.com/pkg/location/application"
)

type LocationHandler interface {
	GetCurrent(ctx *gin.Context)
	GetNext(ctx *gin.Context)
}

type handler struct {
	locationUseCase application.LocationUseCase
}

func NewHandler(locationUseCase application.LocationUseCase) LocationHandler {
	return &handler{locationUseCase: locationUseCase}
}
