package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"go-web/setting"
	"log"
	"time"
)

func CreateToken(id int64, account string, cl map[string]interface{}) (token string, err error) {
	jwtConfig := setting.Application.Jwt
	duration, err := time.ParseDuration(jwtConfig.Time)
	if err != nil {
		log.Fatal("parseDuration fail !", err)
	}
	t := jwt.New(jwt.SigningMethodHS256)
	claims := t.Claims.(jwt.MapClaims)
	claims["account"] = account
	claims["id"] = id
	claims["exp"] = time.Now().Add(duration).Unix()
	token, err = t.SignedString([]byte(jwtConfig.Key))
	return token, err
}
