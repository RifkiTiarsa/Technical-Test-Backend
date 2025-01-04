package handler

import (
	"errors"
	"technical-test/internal/config"
	"technical-test/internal/shared/common"
	"technical-test/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUsecase usecase.ProductUsecase
	rg             *gin.RouterGroup
}

func (p *ProductHandler) getProductHandler(c *gin.Context) {
	name := c.Query("name")

	if name != "" {
		product, err := p.productUsecase.GetProductByName(name)
		if err != nil {
			if errors.Is(err, common.ErrNotFound) {
				common.SendErrorResponse(c, 404, err.Error())
				return
			} else {
				common.SendErrorResponse(c, 500, err.Error())
				return
			}
		}

		common.SendSuccessResponse(c, product, "Get product by name successfully")
		return
	} else {
		products, err := p.productUsecase.ListProduct()
		if err != nil {
			common.SendErrorResponse(c, 500, err.Error())
			return
		}
		common.SendSuccessResponse(c, products, "List product successfully")
	}

}

func (p *ProductHandler) getProductByIDHandler(c *gin.Context) {
	id := c.Param("id")

	product, err := p.productUsecase.GetProductByID(id)
	if err != nil {
		if errors.Is(err, common.ErrNotFound) {
			common.SendErrorResponse(c, 404, err.Error())
			return
		} else {
			common.SendErrorResponse(c, 500, err.Error())
			return
		}
	}

	common.SendSuccessResponse(c, product, "Get product by ID successfully")
}

func (p *ProductHandler) Route() {
	p.rg.GET(config.GetProduct, p.getProductHandler)
	p.rg.GET(config.GetProductByID, p.getProductByIDHandler)
}

func NewProductHandler(productUsecase usecase.ProductUsecase, rg *gin.RouterGroup) *ProductHandler {
	return &ProductHandler{productUsecase: productUsecase, rg: rg}
}
