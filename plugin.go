// DO NOT EDIT.
//go:build wasip1

package main

import (
	"github.com/labulakalia/plugin_api"
)

func init() {
	plugin_api.RegistryPlugin(NewPluginImpl())
}

func main() {}
