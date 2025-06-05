package tokenservice

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Config struct {
	SignKey               string        `koanf:"sign_key"`
	AccessExpirationTime  time.Duration `koanf:"access_expiration_time"`
	RefreshExpirationTime time.Duration `koanf:"refresh_expiration_time"`
	AccessSubject         string        `koanf:"access_subject"`
	RefreshSubject        string        `koanf:"refresh_subject"`
}

type Claims struct {
	jwt.RegisteredClaims
	AdminID uint `json:"admin_id"`
}

type Service struct {
	config Config
}

func New(cfg Config) Service {
	return Service{config: cfg}
}

func (s Service) CreateAccessToken(adminID uint) (string, error) {
	return s.createToken(adminID, s.config.AccessSubject, s.config.AccessExpirationTime)
}

func (s Service) CreateRefreshToken(adminID uint) (string, error) {
	return s.createToken(adminID, s.config.RefreshSubject, s.config.RefreshExpirationTime)
}

func (s Service) createToken(userID uint, subject string, expireDuration time.Duration) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
		},
		AdminID: userID,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString([]byte(s.config.SignKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
