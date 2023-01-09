package controller

import (
	"flight/service"
	"net/http"

	errors "github.com/lnq99/rsoi-2022-lab3-fault-tolerance-lnq99/src/pkg/error"

	"github.com/lnq99/rsoi-2022-lab3-fault-tolerance-lnq99/src/pkg/util"

	"github.com/gin-gonic/gin"
)

type GinController struct {
	service service.Service
}

func NewGinController(service service.Service) *GinController {
	return &GinController{service}
}

func (c *GinController) ListFlights(ctx *gin.Context) {
	page := util.ToInt(ctx.Query("page"))
	size := util.ToInt(ctx.Query("size"))

	if page <= 0 || size <= 0 {
		ctx.JSON(http.StatusBadRequest, errors.ErrorResponse{"Invalid params"})
	}

	r := c.service.ListFlights(ctx, int32(page), int32(size))
	ctx.JSON(http.StatusOK, r)
}

func (c *GinController) GetFlight(ctx *gin.Context) {
	flightNumber := ctx.Param("flightNumber")

	r := c.service.GetFlight(ctx, flightNumber)
	ctx.JSON(http.StatusOK, r)
}
