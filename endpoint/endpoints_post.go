package endpoint

func (e endpoint) EndpointsPost() {
	v1 := e.GroupEndpoint()[versionPathV1]
	v1.POST(e.RouteVersions.RouteAuthV1.PostAuthAccount())
	v1.POST(e.RouteVersions.RouteAuthV1.PostVerifySecret())
}
