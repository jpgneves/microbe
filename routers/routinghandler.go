package routers

import (
	"errors"
	"fmt"
	"github.com/jpgneves/microbe/config"
	"github.com/jpgneves/microbe/requests"
	"github.com/jpgneves/microbe/resources"
	"log"
	"net/http"
)

type RoutingHandler struct {
	router *Router
}

func MakeRoutingHandler(config *config.Configuration) *RoutingHandler {
	routergen, err := selectRouter(config.RouterType)
	if err != nil {
		log.Fatal(err)
	}

	if routergen != nil {
		router := routergen()
		return &RoutingHandler{&router}
	}

	return &RoutingHandler{nil}
}

func (rh *RoutingHandler) Router() *Router {
	return rh.router
}

func (rh *RoutingHandler) SetRouter(r *Router) {
	rh.router = r
}

func (rh *RoutingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var response *requests.Response
	if rh.router != nil {
		path := r.URL.Path
		router := *rh.router
		match := router.Route(path)
		req := &requests.Request{r, match.matches}
		if _, ok := match.value.(resources.Resource); ok {
			switch r.Method {
			case "GET":
				if handler, ok := match.value.(resources.GetMethodHandler); ok {
					response = handler.Get(req)
				} else {
					handleError(w, r, http.StatusMethodNotAllowed)
				}
			case "POST":
				if handler, ok := match.value.(resources.PostMethodHandler); ok {
					response = handler.Post(req)
				} else {
					handleError(w, r, http.StatusMethodNotAllowed)
				}
			case "PUT":
				if handler, ok := match.value.(resources.PutMethodHandler); ok {
					response = handler.Put(req)
				} else {
					handleError(w, r, http.StatusMethodNotAllowed)
				}
			case "DELETE":
				if handler, ok := match.value.(resources.DeleteMethodHandler); ok {
					response = handler.Delete(req)
				} else {
					handleError(w, r, http.StatusMethodNotAllowed)
				}
			default:
				handleError(w, r, http.StatusMethodNotAllowed)
				return
			}
			if response != nil {
				handleResponse(w, r, response)
			}
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
	for k, vs := range response.Header {
		for _, v := range vs {
			w.Header().Add(k, v)
		}
	}
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

func selectRouter(routertype string) (router func() Router, err error) {
	switch routertype {
	case "static":
		return NewStaticRouter, nil
	case "matching":
		return NewMatchingRouter, nil
	case "custom": // Use *only* if you want to manually specify a custom router
		return nil, nil
	}
	return nil, errors.New("Unknown router type")
}
