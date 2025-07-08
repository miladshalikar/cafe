package usertokenauthservice

import (
	"github.com/golang-jwt/jwt/v5"
	"strings"
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
	UserID uint `json:"user_id"`
}

type Service struct {
	config Config
}

func New(cfg Config) Service {
	return Service{config: cfg}
}

func (s Service) CreateAccessToken(userID uint) (string, error) {
	return s.createToken(userID, s.config.AccessSubject, s.config.AccessExpirationTime)
}

func (s Service) CreateRefreshToken(userID uint) (string, error) {
	return s.createToken(userID, s.config.RefreshSubject, s.config.RefreshExpirationTime)
}

func (s Service) createToken(userID uint, subject string, expireDuration time.Duration) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
		},
		UserID: userID,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString([]byte(s.config.SignKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s Service) ParseToken(bearerToken string) (*Claims, error) {
	tokenStr := strings.Replace(bearerToken, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.SignKey), nil
	},
		jwt.WithLeeway(5*time.Second),
		//jwt.WithSubject(s.config.AccessSubject),
		//jwt.WithSubject(s.config.RefreshSubject),
	)

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
