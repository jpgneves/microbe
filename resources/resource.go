package resources

import (
	"github.com/jpgneves/microbe/requests"
)

type Resource interface {}

type GetMethodHandler interface {
	Get(*requests.Request) *requests.Response
}

type PostMethodHandler interface {
	Post(*requests.Request) *requests.Response
}

type PutMethodHandler interface {
	Put(*requests.Request) *requests.Response
}

type DeleteMethodHandler interface {
	Delete(*requests.Request) *requests.Response
}