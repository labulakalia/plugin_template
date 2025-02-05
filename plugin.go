//go:build wasip1

package main

import (
	"github.com/labulakalia/plugin_api"
)

func main() {
	plugin_api.RegistryPlugin(&PluginImpl{})
}
