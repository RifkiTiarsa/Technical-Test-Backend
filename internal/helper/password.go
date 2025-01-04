package helper

import (
	"regexp"
	"technical-test/internal/shared/common"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", common.ErrInternalError
	}

	return string(hashPassword), nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return common.ErrInvalidPassword
	}

	if matched, _ := regexp.MatchString(`[A-Z]`, password); !matched {
		return common.ErrInvalidPassword
	}

	if matched, _ := regexp.MatchString(`[a-z]`, password); !matched {
		return common.ErrInvalidPassword
	}

	if matched, _ := regexp.MatchString(`[0-9]`, password); !matched {
		return common.ErrInvalidPassword
	}

	if matched, _ := regexp.MatchString(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};':",\.<>\/\?\\|~]`, password); !matched {
		return common.ErrInvalidPassword
	}

	return nil
}

func CheckPassword(hashedPassword, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return common.ErrInternalError
	}

	return nil
}
