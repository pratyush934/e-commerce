package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pratyush934/e-commerce/user-service/models"
	"github.com/rs/zerolog/log"
	"strings"
	"time"
)

var privateKey = []byte("57bdcc478aec3c27e91838f1247bc6244b12e4b63bb2a3cc62bb2b2fa15683e0")

func GenerateJWT(user models.User) (string, error) {

	totalTTL := 1800
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     user.Id,
		"roleId": user.RoleId,
		"iat":    time.Now().Unix(),
		"eat":    time.Now().Add(time.Second * time.Duration(totalTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}

func getToken(ctx *gin.Context) (*jwt.Token, error) {
	header, err := getTokenFromHeader(ctx)

	if err != nil {
		log.Warn().Msg("There is an issue coming in utils/jwt.go/getToken, Part1")
		return nil, err
	}

	parsedToken, err := jwt.Parse(header, func(token *jwt.Token) (interface{}, error) {
		if _, ok := header.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected siging method: %v", token.Header["alg"])

		}
		return privateKey, nil
	})

	if err != nil {
		return nil, err
	}
	return parsedToken, nil
}

func getTokenFromHeader(context *gin.Context) (string, error) {
	bearerToken := context.Request.Header.Get("Authorization")

	splitToken := strings.Split(bearerToken, " ")

	if len(splitToken) != 2 {
		log.Warn().Msg("There is an Error in utils/jwt.go/getTokenFromHeader, it seems it is not able to extract Token from Header")
		return "", nil
	}

	return splitToken[1], nil

}
