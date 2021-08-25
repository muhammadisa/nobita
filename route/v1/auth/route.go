package auth

import (
	authrouteinterface "github.com/muhammadisa/nobita/route/v1/auth/interface"
	authusecaseinterface "github.com/muhammadisa/nobita/usecase/v1/auth/interface"
)

type route struct {
	UseCase authusecaseinterface.UseCase
}

func NewAuthRouteV1(usecase authusecaseinterface.UseCase) authrouteinterface.Route {
	return &route{UseCase: usecase}
}
