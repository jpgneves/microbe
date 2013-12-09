package microbe

import (
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
	InitResource(resource *resources.Resource) *resources.Resource
	doStart()
}

type MicrobeInstance struct {
	config         *config.Configuration
	routinghandler *routers.RoutingHandler
}

func Init(config_filename string) *MicrobeInstance {
	config := config.ReadConfig(config_filename)

	routingHandler := routers.MakeRoutingHandler(config)

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

func (m *MicrobeInstance) InitResource(resource resources.Resource) resources.Resource {
	return resource.Init(m.config)
}