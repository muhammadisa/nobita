package endpoint

import (
	"github.com/gin-gonic/gin"
)

func (e endpoint) GroupEndpoint() map[string]*gin.RouterGroup {
	routerGroupMap := make(map[string]*gin.RouterGroup)
	routerGroupMap[versionPathV1] = e.Router.Group(versionPathV1)
	return routerGroupMap
}
