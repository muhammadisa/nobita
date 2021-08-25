package _interface

import jwtlib "github.com/golang-jwt/jwt"

type Data struct {
	RefreshToken string `json:"refresh_token"`
	Token        string `json:"token"`
}

type JWT interface {
	Parser() (*jwtlib.Token, error)
	Claim(*jwtlib.Token, string) (string, error)
	ExtractKey(string) (string, error)
	ExtractKeys([]string) (map[string]string, error)
	Generate(map[string]string) (*Data, error)
}
