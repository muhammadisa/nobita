package endpoint

import "github.com/muhammadisa/nobita/middleware"

func (e endpoint) EndpointsPut() {
	v1 := e.GroupEndpoint()[versionPathV1]
	v1.Use(middleware.JWTAuthMiddleware())
	v1.PUT(e.RouteVersions.RouteAuthV1.UpdateProfile())
}
