package main

import (
	"github.com/labulakalia/plugin_api"
)

// Id implements IPlugin.
func (p *PluginImpl) PluginId() string {

	panic("unimplemented")
}

// GetAuthType implements IPlugin.
func (p *PluginImpl) GetAuthType() *plugin_api.AuthType {
	panic("unimplemented")
}

// CheckAuth implements IPlugin.
func (p *PluginImpl) CheckAuth(*plugin_api.AuthType) *plugin_api.Status {
	panic("unimplemented")
}

// GetAuthData implements IPlugin.
func (p *PluginImpl) GetAuthData() []byte {
	panic("unimplemented")
}

// InitAuth implements IPlugin.
func (p *PluginImpl) InitAuth([]byte) *plugin_api.Status {
	panic("unimplemented")
}

// AuthId implements IPlugin.
func (p *PluginImpl) PluginAuthId() string {
	panic("unimplemented")
}

// GetDirEntry implements IPlugin.
func (p *PluginImpl) GetDirEntry(dir_path string, page uint64, page_size uint64) *plugin_api.DirEntry {
	panic("unimplemented")
}

// GetFileResource implements IPlugin.
func (p *PluginImpl) GetFileResource(file_path string) *plugin_api.FileResource {
	panic("unimplemented")
}

// Close implements IPlugin.
func (p *PluginImpl) Close() {
	panic("unimplemented")
}
