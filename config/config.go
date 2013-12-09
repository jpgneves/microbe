package config

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"strconv"
)

type HostPortPair struct {
	Address *string
	Port    int
}

type Configuration struct {
	ListenAddr *HostPortPair
	WwwRoot    *string
	RouterType string
}

func ReadConfig(filename string) *Configuration {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer f.Close()
	decoder := json.NewDecoder(f)
	config := &Configuration{}
	err = decoder.Decode(config)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return config
}

func (c *Configuration) HostPortString() string {
	return net.JoinHostPort(*c.ListenAddr.Address, strconv.Itoa(c.ListenAddr.Port))
}
