package usecase

import (
	"technical-test/internal/entity"
	"technical-test/internal/helper"
	"technical-test/internal/repository"
	"technical-test/internal/shared/common"
	"time"
)

type userUsecase struct {
	userRepo repository.UserRepository
}

// RegisterUser implements UserUsecase.
func (u *userUsecase) RegisterUser(user entity.User) (entity.User, error) {
	existUser, _ := u.userRepo.GetUserByEmail(user.Email)

	if existUser.Email == user.Email {
		return entity.User{}, common.ErrCoflict
	}

	if err := helper.ValidatePassword(user.Password); err != nil {
		return entity.User{}, common.ErrInvalidPassword
	}

	hashPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return entity.User{}, common.ErrInternalError
	}

	user.Password = hashPassword

	return u.userRepo.CreateUser(user)
}

// FindUserByEmailPassword implements UserUsecase.
func (u *userUsecase) FindUserByEmailPassword(email string, password string) (entity.User, error) {
	userExist, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return entity.User{}, common.ErrCoflict
	}

	if err := helper.CheckPassword(userExist.Password, password); err != nil {
		return entity.User{}, common.ErrWrongPassword
	}

	return userExist, nil
}

// GetUserByUsername implements UserUsecase.
func (u *userUsecase) GetUserByUsername(username string) (entity.User, error) {
	return u.userRepo.GetUserByUsername(username)

}

// GetUserByEmail implements UserUsecase.
func (u *userUsecase) GetUserByEmail(email string) (entity.User, error) {
	return u.userRepo.GetUserByEmail(email)
}

// GetUserByDate implements UserUsecase.
func (u *userUsecase) GetUserByDate(date time.Time) (entity.User, error) {
	return u.userRepo.GetUserByDate(date)
}

type UserUsecase interface {
	RegisterUser(user entity.User) (entity.User, error)
	FindUserByEmailPassword(email, password string) (entity.User, error)
	GetUserByUsername(username string) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	GetUserByDate(date time.Time) (entity.User, error)
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}
