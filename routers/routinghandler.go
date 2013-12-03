package routers

import (
	"fmt"
	"github.com/jpgneves/microbe/requests"
	"github.com/jpgneves/microbe/resources"
	"log"
	"net/http"
)

type RoutingHandler struct {
	router Router
}

func MakeRoutingHandler(router Router) *RoutingHandler {
	return &RoutingHandler{router}
}

func (rh *RoutingHandler) SetRouter(r Router) {
	rh.router = r
}

func (rh *RoutingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var response *requests.Response
	if rh.router != nil {
		path := r.URL.Path
		match := rh.router.Route(path)
		if resource, ok := match.value.(resources.Resource); ok {
			req := &requests.Request{r, match.matches}
			switch r.Method {
			case "GET":
				response = resource.Get(req)
			case "POST":
				response = resource.Post(req)
			default:
				handleError(w, r, http.StatusMethodNotAllowed)
				return
			}
			handleResponse(w, r, response)
		} else {
			handleError(w, r, http.StatusNotFound)
		}
	} else {
		handleError(w, r, http.StatusNotFound)
	}
	return
}

func handleResponse(w http.ResponseWriter, r *http.Request, response *requests.Response) {
	log.Printf("%s %s - %v", r.Method, r.URL.Path, response.StatusCode)
	switch response.StatusCode {
	case http.StatusTemporaryRedirect:
		fallthrough
	case http.StatusMovedPermanently:
		http.Redirect(w, r, *response.Data, response.StatusCode)
	default:
		w.WriteHeader(response.StatusCode)
		var response_txt string
		if response.Data == nil {
			response_txt = http.StatusText(response.StatusCode)
		} else {
			response_txt = *response.Data
		}
		fmt.Fprintf(w, response_txt)
	}
}

func handleError(w http.ResponseWriter, r *http.Request, code int) {
	log.Printf("%s %s - %v", r.Method, r.URL.Path, code)
	w.WriteHeader(code)
	fmt.Fprintf(w, http.StatusText(code))
}