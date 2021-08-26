package route

import (
	routeauthinterface "github.com/muhammadisa/nobita/route/v1/auth/interface"
	routefoointerface "github.com/muhammadisa/nobita/route/v1/foo/interface"
)

type Versions struct {
	RouteAuthV1 routeauthinterface.Route
	RouteFooV1  routefoointerface.Route
}
