package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "superSecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		//Need to check the token I am Verifying was signed  with that sigining method you chose for creating  it. In this case I am using SigningMethodHS256
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("could nor parse token")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, errors.New("token is expired")
	}

	claims, _ := parsedToken.Claims.(jwt.MapClaims)

	//FOr these to check I need to add  expectedEmail and expectedUserID as argument in method
	// // Verify email claim
	// email, ok := claims["email"].(string)
	// if !ok || email != expectedEmail {
	// 	return errors.New("email claim is invalid")
	// }

	// // Verify userID claim
	// userID, ok := claims["userID"].(int64)
	// if !ok || int64(userID) != expectedUserID {
	// 	return errors.New("userID claim is invalid")
	// }

	userID := int64(claims["userID"].(float64))

	//fetch the expiry time
	exp, _ := claims["exp"].(float64)
	iExp := int64(exp)

	//compare the expiry time with the current time
	// Check if the token has expired
	if time.Unix(iExp, 0).Before(time.Now()) {
		return userID, errors.New("auth token has expired")
	}

	return userID, nil
}
