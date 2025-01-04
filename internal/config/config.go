package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Driver   string
}

type ApiConfig struct {
	ApiPort string
}

type TokenConfig struct {
	IssuerName           string `json:"IssuerName"`
	JwtSignatureKy       []byte `json:"JwtSignatureKy"`
	JwtSigningMethod     *jwt.SigningMethodHMAC
	AccessJwtExpiresTime time.Duration
}

type Config struct {
	DBConfig
	ApiConfig
	TokenConfig
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func (c *Config) readConfig() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("missing env file %v", err.Error())
	}
	c.DBConfig = DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "rahasia"),
		Name:     getEnv("DB_NAME", "e_commerce"),
		Driver:   getEnv("DB_DRIVER", "postgres"),
	}

	c.ApiConfig = ApiConfig{ApiPort: getEnv("API_PORT", "8080")}

	accessTokenExpire, _ := strconv.Atoi(getEnv("ACCESS_TOKEN_EXPIRE", "60"))
	c.TokenConfig = TokenConfig{
		IssuerName:           getEnv("TOKEN_ISSUE", "rifkiTiarsa"),
		JwtSignatureKy:       []byte(getEnv("TOKEN_SECRET", "sangatAmatRahasia")),
		JwtSigningMethod:     jwt.SigningMethodHS256,
		AccessJwtExpiresTime: time.Duration(accessTokenExpire) * time.Minute,
	}

	if c.Host == "" || c.Port == "" || c.User == "" || c.Name == "" || c.Driver == "" || c.ApiPort == "" ||
		c.IssuerName == "" || c.AccessJwtExpiresTime < 0 || len(c.JwtSignatureKy) == 0 {
		return fmt.Errorf("missing required environment")
	}

	return nil

}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.readConfig(); err != nil {
		return nil, err
	}
	return cfg, nil
}
