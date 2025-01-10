package plugin

import "github.com/labulakalia/plugin_temp/proto"

type PluginImpl struct {
}

// AuthId implements IPlugin.
func (p *PluginImpl) AuthId() string {
	panic("unimplemented")
}

// CheckAuth implements IPlugin.
func (p *PluginImpl) CheckAuth(*proto.AuthType) *proto.Status {
	panic("unimplemented")
}

// Close implements IPlugin.
func (p *PluginImpl) Close() {
	panic("unimplemented")
}

// GetAuthData implements IPlugin.
func (p *PluginImpl) GetAuthData() []byte {
	panic("unimplemented")
}

// GetAuthType implements IPlugin.
func (p *PluginImpl) GetAuthType() *proto.AuthType {
	panic("unimplemented")
}

// GetDirEntry implements IPlugin.
func (p *PluginImpl) GetDirEntry(dir_path string, page uint64, page_size uint64) *proto.DirEntry {
	panic("unimplemented")
}

// GetFileResource implements IPlugin.
func (p *PluginImpl) GetFileResource(file_path string) *proto.FileResource {
	panic("unimplemented")
}

// Id implements IPlugin.
func (p *PluginImpl) Id() string {
	panic("unimplemented")
}

// InitAuth implements IPlugin.
func (p *PluginImpl) InitAuth([]byte) *proto.Status {
	panic("unimplemented")
}
