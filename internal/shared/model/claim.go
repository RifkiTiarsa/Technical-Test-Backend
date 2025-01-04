package model

import "github.com/golang-jwt/jwt/v5"

type Claim struct {
	jwt.RegisteredClaims
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
