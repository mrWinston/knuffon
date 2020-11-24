package manager

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	JWT_SECRET string = "KLASJFnealjf83hafs"
)

func CreateToken(userID string) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(time.Hour * 6).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidateToken(tokenString string) (string, error) {
	if tokenString == "" {
		return "", errors.New("Token was Empty")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, ok := claims["user_id"].(string)
		if ok {
			return userID, nil
		}
		return "", errors.New("Error while parsing JWT Metadata")
	}
	return "", errors.New("Error while Casting JWT Metadata or Token is Invalid")
}
