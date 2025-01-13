package utils

import (
	"gin-template/internal/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(user *models.User, secret string, exp time.Time) (string, error) {
	claims := &models.Claims{
		StandardClaims: jwt.StandardClaims{
			Subject: strconv.Itoa(user.ID),
			ExpiresAt: exp.Unix(),
		},
		Email: user.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
