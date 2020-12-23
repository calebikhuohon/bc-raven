package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

var (
	secret = []byte("secretsiyutyrerethjknbmvgdghjj")
)

func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"]=username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenstr, err := token.SignedString(secret)

	if err != nil {
		log.Fatal("Error in Generating key: ", err)
		return "", err
	}
	return tokenstr, nil
}

func ParseToken(tokenStr string) (string, error)  {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", err
	}
}
