package controller

import (
	"net/http"

	"bonus/service"

	"github.com/lnq99/rsoi-2022-lab3-fault-tolerance-lnq99/src/pkg/model"

	"github.com/gin-gonic/gin"
)

const UsernameHeader = model.UsernameHeader

type GinController struct {
	service service.Service
}

func NewGinController(service service.Service) *GinController {
	return &GinController{service}
}

func (c *GinController) ListPrivilegeHistories(ctx *gin.Context) {
	username := ctx.GetHeader(UsernameHeader)
	r := c.service.GetPrivilege(ctx, username)
	ctx.JSON(http.StatusOK, r)
}

func (c *GinController) UpdateBalanceAndHistory(ctx *gin.Context) {
	username := ctx.GetHeader(UsernameHeader)

	history := model.BalanceHistory{}
	if err := ctx.ShouldBindJSON(&history); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err := c.service.UpdateBalanceAndHistory(ctx, username, history)
	if err == nil {
		ctx.Status(http.StatusOK)
	} else {
		ctx.Status(http.StatusInternalServerError)
	}
}

func (c *GinController) RevertBalanceAndHistory(ctx *gin.Context) {
	ticketUid := ctx.Param("ticketUid")

	err := c.service.RevertBalanceAndHistory(ctx, ticketUid)
	if err == nil {
		ctx.Status(http.StatusNoContent)
	} else {
		ctx.Status(http.StatusInternalServerError)
	}
}
