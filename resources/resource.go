package resources

import (
	"github.com/jpgneves/microbe/requests"
)

type Resource interface {
	Get(*requests.Request) *requests.Response
	Post(*requests.Request) *requests.Response
}
