package service

import (
	"fmt"
	"technical-test/internal/config"
	"technical-test/internal/entity"
	"technical-test/internal/entity/dto"
	"technical-test/internal/shared/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtService struct {
	cfgToken config.TokenConfig
}

// GenerateToken implements JwtService.
func (j *jwtService) GenerateToken(user entity.User) (dto.AuthLoginResponseDto, error) {
	// Access Token
	accessClaims := model.Claim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.cfgToken.IssuerName,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.cfgToken.AccessJwtExpiresTime)),
		},
		UserID:   user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
	}

	accessToken := jwt.NewWithClaims(j.cfgToken.JwtSigningMethod, accessClaims)
	accessTokenString, err := accessToken.SignedString(j.cfgToken.JwtSignatureKy)
	if err != nil {
		return dto.AuthLoginResponseDto{}, fmt.Errorf("failed to create token: %v", err)
	}

	// Refresh Token
	refreshClaims := model.Claim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.cfgToken.IssuerName,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.cfgToken.RefreshJwtExpiresTime)),
		},
		UserID:   user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
	}

	refreshToken := jwt.NewWithClaims(j.cfgToken.JwtSigningMethod, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(j.cfgToken.JwtSignatureKy)
	if err != nil {
		return dto.AuthLoginResponseDto{}, fmt.Errorf("failed to create token: %v", err)
	}

	return dto.AuthLoginResponseDto{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

// ValidateToken implements JwtService.
func (j *jwtService) ValidateToken(tokenString string) (*model.Claim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return j.cfgToken.JwtSignatureKy, nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	claim, ok := token.Claims.(*model.Claim)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("unauthorized : %v", err)
	}

	return claim, nil
}

type JwtService interface {
	GenerateToken(user entity.User) (dto.AuthLoginResponseDto, error)
	ValidateToken(tokenString string) (*model.Claim, error)
}

func NewJwtService(cfg config.TokenConfig) JwtService {
	return &jwtService{cfgToken: cfg}
}
