package utility

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type authCustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(email string) string {
	claims := &authCustomClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		panic(err)
	}
	return t
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, errors.New("Invalid Token")

		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

}
