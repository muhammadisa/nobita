package jwt

import (
	jwtlib "github.com/golang-jwt/jwt"
	"strings"
)

func (j jwt) Parser() (*jwtlib.Token, error) {
	strs := strings.Split(j.Bearer, " ")
	if !(len(strs) > 1) {
		return nil, errorNoTokenGiven
	}
	token, err := jwtlib.Parse(strs[1], func(token *jwtlib.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtlib.SigningMethodHMAC); !ok {
			return nil, errorUnexpectSigningMethod
		}
		return []byte(j.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
