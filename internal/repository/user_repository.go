package repository

import (
	"technical-test/internal/entity"
	"time"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// CreateUser implements UserRepository.
func (u *userRepository) CreateUser(user entity.User) (entity.User, error) {
	if err := u.db.Create(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

// GetUserByDate implements UserRepository.
func (u *userRepository) GetUserByDate(date time.Time) (entity.User, error) {
	var user entity.User

	if err := u.db.Where("created_at > ?", date).First(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

// GetUserByEmail implements UserRepository.
func (u *userRepository) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User

	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

// GetUserByUsername implements UserRepository.
func (u *userRepository) GetUserByUsername(username string) (entity.User, error) {
	var user entity.User

	if err := u.db.Where("username = ?", username).First(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

// GetUserById implements UserRepository.
func (u *userRepository) GetUserById(id string) (entity.User, error) {
	var user entity.User
	if err := u.db.Where("id =?", id).First(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

type UserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUserByUsername(username string) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	GetUserByDate(date time.Time) (entity.User, error)
	GetUserById(id string) (entity.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
