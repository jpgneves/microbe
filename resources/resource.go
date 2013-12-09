package resources

import (
	"github.com/jpgneves/microbe/config"
	"github.com/jpgneves/microbe/requests"
)

type Resource interface {
	Init(*config.Configuration) Resource
	Get(*requests.Request) *requests.Response
	Post(*requests.Request) *requests.Response
}
