package endpoint

import (
	"github.com/gin-gonic/gin"
	endpointinterface "github.com/muhammadisa/nobita/endpoint/interface"
	routeversion "github.com/muhammadisa/nobita/route"
)

type endpoint struct {
	Router        *gin.Engine
	RouteVersions routeversion.Versions
}

func (e endpoint) MiddlewareEndpoint() {
}

func NewEndpoint(router *gin.Engine, routes routeversion.Versions) endpointinterface.Endpoints {
	ep := &endpoint{
		Router:        router,
		RouteVersions: routes,
	}
	ep.EndpointsPut()
	ep.EndpointsPost()
	ep.EndpointsDelete()
	ep.EndpointsGet()
	return ep
}
