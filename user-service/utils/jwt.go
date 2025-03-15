package utils

import (
	"errors"
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
		"id":   user.Id,
		"role": user.RoleId,
		"iat":  time.Now().Unix(),
		"eat":  time.Now().Add(time.Second * time.Duration(totalTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}

func validateToken(ctx *gin.Context) error {
	token, err := getToken(ctx)

	if err != nil {
		log.Warn().Msg("There is an issue in utils/jwt.go/validateToken")
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return errors.New("there is an error in validateToken method")
	}
	return nil

}

func validateAdminRole(ctx *gin.Context) error {
	token, err := getToken(ctx)

	if err != nil {
		log.Warn().Msg("There is an issue in utils/jwt.go/validateAdminRole")
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	roleId := claims["role"].(float64)

	if ok && token.Valid && roleId == 1 {
		return nil
	}
	return errors.New("there is an issue in validateToken part2")
}

func validateUserRole(ctx *gin.Context) error {
	token, err := getToken(ctx)

	if err != nil {
		log.Warn().Msg("There is an issue in utils/jwt.go/validateUserRole")
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	roleId := claims["role"].(float64)

	if ok && token.Valid && roleId == 2 {
		return nil
	}
	return errors.New("there is an issue in validateToken part3")
}

func validateStoreKeeperRole(ctx *gin.Context) error {
	token, err := getToken(ctx)

	if err != nil {
		log.Warn().Msg("There is an issue in utils/jwt.go/validateStoreKeeperRole")
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	roleId := claims["role"].(float64)

	if ok && token.Valid && roleId == 3 {
		return nil
	}

	return errors.New("there is an issue in validateToken part4")
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
