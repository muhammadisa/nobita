package endpoint

import (
	"github.com/gin-gonic/gin"
	endpointinterface "github.com/muhammadisa/nobita/endpoint/interface"
	featurerepositoryv1 "github.com/muhammadisa/nobita/repository/v1/feature/interface"
	routeversion "github.com/muhammadisa/nobita/route"
)

type endpoint struct {
	Router        *gin.Engine
	FeatureRepo   featurerepositoryv1.RW
	RouteVersions routeversion.Versions
}

func (e endpoint) MiddlewareEndpoint() {
}

func NewEndpoint(router *gin.Engine, featureRepo featurerepositoryv1.RW, routes routeversion.Versions) endpointinterface.Endpoints {
	ep := &endpoint{
		Router:        router,
		FeatureRepo:   featureRepo,
		RouteVersions: routes,
	}
	ep.EndpointsPut()
	ep.EndpointsPost()
	ep.EndpointsDelete()
	ep.EndpointsGet()
	return ep
}
