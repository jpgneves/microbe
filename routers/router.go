package routers

import (
	"github.com/jpgneves/microbe/resources"
)

type Router interface {
	AddRoute(route string, resource resources.Resource)
	RemoveRoute(route string)
	Route(path string) *RouteMatch
}
