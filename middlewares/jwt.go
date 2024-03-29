package middlewares

import (
	"batch6/constants"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type jwtCustomClaims struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
	jwt.RegisteredClaims
}

func GenerateTokenJWT(userId int, userName string) string {
	fmt.Println(userId)
	var userClaims = jwtCustomClaims{
		userId, userName,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	resultJWT, _ := token.SignedString([]byte(constants.PRIVATE_KEY_JWT))
	return resultJWT
}
