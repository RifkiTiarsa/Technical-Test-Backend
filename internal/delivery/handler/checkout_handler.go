package handler

import (
	"errors"
	"technical-test/internal/config"
	"technical-test/internal/entity"
	"technical-test/internal/middleware"
	"technical-test/internal/shared/common"
	"technical-test/internal/usecase"

	"github.com/gin-gonic/gin"
)

type CheckoutHandler struct {
	checkoutUC     usecase.CheckoutUsecase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (c *CheckoutHandler) createCheckoutHandler(ctx *gin.Context) {
	var checkout entity.Checkout

	if err := ctx.ShouldBindJSON(&checkout); err != nil {
		common.SendErrorResponse(ctx, 400, err.Error())
		return
	}

	newCheckout, err := c.checkoutUC.CreateCheckout(checkout)
	if err != nil {
		if errors.Is(err, common.ErrInvalidInput) {
			common.SendErrorResponse(ctx, 400, err.Error())
			return
		} else {
			common.SendErrorResponse(ctx, 500, err.Error())
			return
		}
	}

	common.SendSuccessResponse(ctx, newCheckout, "Checkout created successfully")
}

func (c *CheckoutHandler) Route() {
	c.rg.POST(config.CreateCheckout, c.authMiddleware.RequireToken(), c.createCheckoutHandler)
}

func NewCheckoutHandler(checkoutUC usecase.CheckoutUsecase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *CheckoutHandler {
	return &CheckoutHandler{checkoutUC: checkoutUC, rg: rg, authMiddleware: authMiddleware}
}
