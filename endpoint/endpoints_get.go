package endpoint

import "github.com/muhammadisa/nobita/middleware"

func (e endpoint) EndpointsGet() {
	v1 := e.GroupEndpoint()[versionPathV1]

	v1.Use(middleware.JWTAuthMiddleware())
	v1.Use(middleware.RoleAccessibleMiddleware(e.FeatureRepo))

	v1.GET(e.RouteVersions.RouteFooV1.GetFoo())
	v1.GET(e.RouteVersions.RouteAuthV1.GetProfile())
}
