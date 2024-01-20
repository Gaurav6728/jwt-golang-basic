package helpers

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJwtToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		log.Fatal("error creating token", err)
		return "", err
	}

	return signedToken, nil
}

func VerifyJwtToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) { return []byte(os.Getenv("JWT_SECRET_KEY")), nil })
	if err != nil {

		return err
	}

	// see if token is valid
	if !token.Valid {

		return errors.New("invalid token")
	}

	return nil
}
