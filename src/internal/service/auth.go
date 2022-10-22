package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/weeabook/src/university-platform-backend/internal/model"
	"github.com/weeabook/src/university-platform-backend/internal/repository"
)

const (
	// я знаю, что это надо в env
	salt       = "weeabookpro"
	signingKey = "weeabook"
	tokenSSL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type Auth struct {
	repo repository.AuthRepository
}

func NewAuth(repo repository.AuthRepository) *Auth {
	return &Auth{repo: repo}
}

func (s *Auth) CreateUser(user model.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *Auth) GenerateToken(email string, password string) (string, error) {
	user, err := s.repo.GetUser(model.InputUser{
		Email:    email,
		Password: generatePasswordHash(password),
	})
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenSSL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *Auth) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, nil
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (s *Auth) GetUserInfo(userId int) (model.User, error) {
	return s.repo.GetUserInfo(userId)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
