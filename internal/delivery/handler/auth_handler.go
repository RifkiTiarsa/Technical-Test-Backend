package handler

import (
	"errors"
	"technical-test/internal/config"
	"technical-test/internal/entity/dto"
	"technical-test/internal/shared/common"
	"technical-test/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
	rg          *gin.RouterGroup
}

func (a *AuthHandler) registerHandler(c *gin.Context) {
	var payload dto.AuthRegisterDto

	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, 400, err.Error())
		return
	}

	user, err := a.authUsecase.Register(payload)
	if err != nil {
		if errors.Is(err, common.ErrInvalidInput) {
			common.SendErrorResponse(c, 400, err.Error())
			return
		} else if errors.Is(err, common.ErrInvalidPassword) {
			common.SendErrorResponse(c, 400, err.Error())
			return
		} else if errors.Is(err, common.ErrCoflict) {
			common.SendErrorResponse(c, 409, err.Error())
			return
		} else {
			common.SendErrorResponse(c, 500, err.Error())
			return
		}
	}

	common.SendCreateResponse(c, user, "Register user successfully")
}

func (a *AuthHandler) loginHandler(c *gin.Context) {
	var payload dto.AuthLoginDto

	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, 400, err.Error())
		return
	}

	token, err := a.authUsecase.Login(payload)
	if err != nil {
		if errors.Is(err, common.ErrInvalidInput) {
			common.SendErrorResponse(c, 400, err.Error())
			return
		} else if errors.Is(err, common.ErrUnauthorized) {
			common.SendErrorResponse(c, 401, err.Error())
			return
		} else if errors.Is(err, common.ErrNotFound) {
			common.SendErrorResponse(c, 404, err.Error())
			return
		} else if errors.Is(err, common.ErrWrongPassword) {
			common.SendErrorResponse(c, 401, err.Error())
			return
		} else {
			common.SendErrorResponse(c, 500, err.Error())
			return
		}
	}

	common.SendSuccessResponse(c, token, "Login successfully")
}

func (a *AuthHandler) Route() {
	a.rg.POST(config.Register, a.registerHandler)
	a.rg.POST(config.Login, a.loginHandler)
}

func NewAuthHandler(authUsecase usecase.AuthUsecase, rg *gin.RouterGroup) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase, rg: rg}
}
