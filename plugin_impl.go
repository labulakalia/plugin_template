//go:build wasip1

package main

import (
	"github.com/labulakalia/plugin_api/plugin"
)

/*
NOTE: net and http use package
"github.com/labulakalia/wazero_net/wasi/http"
"github.com/labulakalia/wazero_net/wasi/net"
*/

type PluginImpl struct {
}

func NewPluginImpl() *PluginImpl {
	return &PluginImpl{}
}

// Id implements IPlugin.
func (p *PluginImpl) PluginId() (string, error) {
	panic("unimplemented")
}

// GetAuthType implements IPlugin.
func (p *PluginImpl) GetAuthType() (*plugin.AuthType, error) {
	panic("unimplemented")
}

// CheckAuth implements IPlugin.
func (p *PluginImpl) CheckAuthType(*plugin.AuthType) (authData []byte, err error) {
	panic("unimplemented")
}

// InitAuth implements IPlugin.
func (p *PluginImpl) CheckAuthData([]byte) error {
	panic("unimplemented")
}

// AuthId implements IPlugin.
func (p *PluginImpl) PluginAuthId() (string, error) {
	panic("unimplemented")
}

// GetDirEntry implements IPlugin.
func (p *PluginImpl) GetDirEntry(dir_path string, page uint64, page_size uint64) (*plugin.DirEntry, error) {
	panic("unimplemented")
}

// GetFileResource implements IPlugin.
func (p *PluginImpl) GetFileResource(file_path string) (*plugin.FileResource, error) {
	panic("unimplemented")
}
