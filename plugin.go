package main

import (
	"github.com/labulakalia/plugin_api"
	"github.com/labulakalia/wazero_net/wasi/http"
	"github.com/labulakalia/wazero_net/wasi/net"
)

func main() {
	_ = net.Conn{}
	_ = http.Do
}


var _ plugin_api.IPluginApi = &PluginImpl{}

type PluginImpl struct {
}
