package jwt

import (
	"errors"
	jwtinterface "github.com/muhammadisa/nobita/util/jwt/interface"
	"time"
)

const (
	blankString = ``
)

var (
	errorNoTokenGiven          = errors.New("no bearer token given")
	errorUnexpectSigningMethod = errors.New("unexpected signing method")
	errorClaimNotOK            = errors.New("claim not ok or token invalid")
	errorClaimKeyNotFound      = errors.New("fail claim key or key not recognized")
	errorClaimCastingFailed    = errors.New("claim key found but cast error")
)

type jwt struct {
	Bearer, Secret string
	Exp            time.Duration
}

func NewJWT(bearer, secret string) jwtinterface.JWT {
	return &jwt{Bearer: bearer, Secret: secret}
}

func NewJWTGenerateMode(exp time.Duration, secret string) jwtinterface.JWT {
	return &jwt{Exp: exp, Secret: secret}
}
