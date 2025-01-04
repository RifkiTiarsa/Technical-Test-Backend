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

type CartHandler struct {
	cartUC         usecase.CartUsecase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (c *CartHandler) addCartHandler(ctx *gin.Context) {
	var cart entity.Cart

	if err := ctx.ShouldBindJSON(&cart); err != nil {
		common.SendErrorResponse(ctx, 400, err.Error())
		return
	}

	newCart, err := c.cartUC.AddCart(cart)
	if err != nil {
		if errors.Is(err, common.ErrInvalidInput) {
			common.SendErrorResponse(ctx, 400, err.Error())
			return
		} else if errors.Is(err, common.ErrNotFound) {
			common.SendErrorResponse(ctx, 404, err.Error())
			return
		} else if errors.Is(err, common.ErrStockAvailability) {
			common.SendErrorResponse(ctx, 400, "Insufficient stock available")
			return
		} else if errors.Is(err, common.ErrUnauthorized) {
			common.SendErrorResponse(ctx, 401, err.Error())
			return
		} else {
			common.SendErrorResponse(ctx, 500, err.Error())
			return
		}
	}

	common.SendCreateResponse(ctx, newCart, "The product has been successfully added to the cart.")
}

func (c *CartHandler) Route() {
	c.rg.POST(config.AddCart, c.authMiddleware.RequireToken(), c.addCartHandler)
}

func NewCartHandler(cartUC usecase.CartUsecase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *CartHandler {
	return &CartHandler{cartUC: cartUC, rg: rg, authMiddleware: authMiddleware}
}
