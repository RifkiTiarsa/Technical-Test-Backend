package usecase

import (
	"technical-test/internal/entity"
	"technical-test/internal/entity/dto"
	"technical-test/internal/shared/common"
	"technical-test/internal/shared/service"
)

type authUsecase struct {
	userUsecase UserUsecase
	jwt         service.JwtService
}

// Register implements AuthUsecase.
func (a *authUsecase) Register(user dto.AuthRegisterDto) (entity.User, error) {
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return entity.User{}, common.ErrInvalidInput
	}

	return a.userUsecase.RegisterUser(entity.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	})
}

// Login implements AuthUsecase.
func (a *authUsecase) Login(user dto.AuthLoginDto) (dto.AuthLoginResponseDto, error) {
	if user.Email == "" || user.Password == "" {
		return dto.AuthLoginResponseDto{}, common.ErrUnauthorized
	}

	userExist, err := a.userUsecase.FindUserByEmailPassword(user.Email, user.Password)
	if err != nil {
		return dto.AuthLoginResponseDto{}, common.ErrWrongPassword
	}

	token, err := a.jwt.GenerateToken(userExist)
	if err != nil {
		return dto.AuthLoginResponseDto{}, common.ErrInternalError
	}

	return token, nil
}

type AuthUsecase interface {
	Register(user dto.AuthRegisterDto) (entity.User, error)
	Login(user dto.AuthLoginDto) (dto.AuthLoginResponseDto, error)
}

func NewAuthUsecase(userUsecase UserUsecase, jwt service.JwtService) AuthUsecase {
	return &authUsecase{userUsecase: userUsecase, jwt: jwt}
}
