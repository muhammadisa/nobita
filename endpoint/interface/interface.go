package _interface

import "github.com/gin-gonic/gin"

type Endpoints interface {
	MiddlewareEndpoint()
	GroupEndpoint() map[string]*gin.RouterGroup
	EndpointsGet()
	EndpointsPost()
	EndpointsDelete()
	EndpointsPut()
}
