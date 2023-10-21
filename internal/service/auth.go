package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/Kosodaka/todo-service/internal/models"
	"github.com/Kosodaka/todo-service/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"userId"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(c *gin.Context, user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(c, user)
}

func (s *AuthService) CreateToken(c *gin.Context, username, password string) (string, error) {
	user, err := s.repo.GetUser(c, username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	if err = godotenv.Load(); err != nil {
		log.Logger.Fatal().Msgf("error to load env variables: %s", err.Error())
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
		},
		UserId: user.Id,
	})
	return token.SignedString([]byte(os.Getenv(os.Getenv("JWT_SALT"))))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	if err := godotenv.Load(); err != nil {
		log.Logger.Fatal().Msgf("error to load env variables: %s", err.Error())
	}
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(os.Getenv("JWT_SALT")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, nil

}

func generatePasswordHash(password string) string {
	if err := godotenv.Load(); err != nil {
		log.Logger.Fatal().Msgf("error to load env variables: %s", err.Error())
	}
	salt := os.Getenv("SALT")
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
