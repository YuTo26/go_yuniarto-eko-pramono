package utils

import (
	"project_1/constants"
	"project_1/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    user.ID,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(constants.JWT_SECRET))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
