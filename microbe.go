package microbe

import (
	"errors"
	"github.com/jpgneves/microbe/config"
	"github.com/jpgneves/microbe/resources"
	"github.com/jpgneves/microbe/routers"
	"log"
	"net/http"
)

type Microbe interface {
	Start(daemon bool)
	AddRoute(route string, resource *resources.Resource)
	RemoveRoute(route string)
	SetRouter(router *routers.Router)
	doStart()
}

type MicrobeInstance struct {
	config         *config.Configuration
	routinghandler *routers.RoutingHandler
}

func Init(config_filename string) *MicrobeInstance {
	config := config.ReadConfig(config_filename)
	routergen, err := selectRouter(config.RouterType)
	if err != nil {
		log.Fatal(err)
	}
	routingHandler := routers.MakeRoutingHandler(routergen())

	return &MicrobeInstance{config, routingHandler}
}

func (m *MicrobeInstance) Start(daemon bool) {
	if daemon {
		go m.doStart()
	} else {
		m.doStart()
	}
}

func (m *MicrobeInstance) doStart() {
	addr := m.config.HostPortString()
	log.Printf("Starting server on %s", addr)
	http.ListenAndServe(addr, m.routinghandler)
}

func (m *MicrobeInstance) AddRoute(route string, resource resources.Resource) {
	router := *(m.routinghandler.Router())
	router.AddRoute(route, resource)
}

func (m *MicrobeInstance) RemoveRoute(route string) {
	router := *(m.routinghandler.Router())
	router.RemoveRoute(route)
}

func (m *MicrobeInstance) SetRouter(router *routers.Router) {
	m.routinghandler.SetRouter(router)
}

func selectRouter(routertype string) (router func() routers.Router, err error) {
	switch routertype {
	case "static":
		return routers.NewStaticRouter, nil
	case "matching":
		return routers.NewMatchingRouter, nil
	case "custom": // Use *only* if you want to manually specify a custom router
		return nil, nil
	}
	return nil, errors.New("Unknown router type")
}
