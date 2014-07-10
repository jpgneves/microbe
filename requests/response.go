package requests

import "net/http"

type Response struct {
	StatusCode int
	Header     http.Header
	Data       *string
}
