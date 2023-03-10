package controller

import (
	"net/http"

	"ticket/service"

	"github.com/lnq99/rsoi-2022-lab3-fault-tolerance-lnq99/src/pkg/model"

	errors "github.com/lnq99/rsoi-2022-lab3-fault-tolerance-lnq99/src/pkg/error"

	"github.com/gin-gonic/gin"
)

const (
	UsernameHeader = model.UsernameHeader
	AuthHeader     = "Authorization"
)

type GinController struct {
	service service.Service
}

func NewGinController(service service.Service) *GinController {
	return &GinController{service}
}

func (c *GinController) ListTickets(ctx *gin.Context) {
	authHeader := ctx.GetHeader(AuthHeader)
	username := ctx.GetHeader(UsernameHeader)

	ctx.Set(AuthHeader, authHeader)

	r := c.service.ListTickets(ctx, username)
	ctx.JSON(http.StatusOK, r)
}
func (c *GinController) CreateTicket(ctx *gin.Context) {
	var ticketReq model.TicketPurchaseRequest

	if err := ctx.ShouldBindJSON(&ticketReq); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	authHeader := ctx.GetHeader(AuthHeader)
	username := ctx.GetHeader(UsernameHeader)

	ctx.Set(AuthHeader, authHeader)

	r, err := c.service.CreateTicket(ctx, username, &ticketReq)
	if err != nil {
		// TODO: check StatusBadRequest
		ctx.JSON(http.StatusServiceUnavailable, errors.ErrorResponse{err.Error()})
	} else {
		ctx.JSON(http.StatusOK, r)
	}
}
func (c *GinController) GetTicket(ctx *gin.Context) {
	authHeader := ctx.GetHeader(AuthHeader)
	username := ctx.GetHeader(UsernameHeader)
	ticketUid := ctx.Param("ticketUid")

	ctx.Set(AuthHeader, authHeader)

	r := c.service.GetTicket(ctx, username, ticketUid)
	if r != nil {
		ctx.JSON(http.StatusOK, r)
	} else {
		ctx.JSON(http.StatusNotFound, errors.ErrorResponse{"Not found"})
	}

}
func (c *GinController) DeleteTicket(ctx *gin.Context) {
	authHeader := ctx.GetHeader(AuthHeader)
	username := ctx.GetHeader(UsernameHeader)
	ticketUid := ctx.Param("ticketUid")

	ctx.Set(AuthHeader, authHeader)

	err := c.service.DeleteTicket(ctx, username, ticketUid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errors.ErrorResponse{err.Error()})
	} else {
		ctx.Status(http.StatusNoContent)
	}
}
