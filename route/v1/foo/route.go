package foo

import (
	foorouteinterface "github.com/muhammadisa/nobita/route/v1/foo/interface"
)

type route struct{}

func NewFooRouteV1() foorouteinterface.Route {
	return &route{}
}
