package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"go-web/setting"
	"time"
)

func CreateToken(id int64, account string, cl map[string]interface{}) (token string, err error) {
	jwtConfig := setting.Application.Jwt
	t := jwt.New(jwt.GetSigningMethod(jwtConfig.Alg))
	claims := t.Claims.(jwt.MapClaims)
	claims["account"] = account
	claims["id"] = id
	claims["exp"] = time.Now().Add(jwtConfig.Time).Unix()
	for key, value := range cl {
		claims[key] = value
	}
	token, err = t.SignedString([]byte(jwtConfig.Key))
	return
}

func ParseToken(tokenString string) (token *jwt.Token, err error) {
	jwtConfig := setting.Application.Jwt
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtConfig.Key), nil
	})
	return
}
