package plugin

import "plugin_temp/proto"

type IPlugin interface {
	Id() string
	GetAuthType() *proto.AuthType
	CheckAuth(*proto.AuthType) *proto.Status
	GetAuthData() []byte
	InitAuth([]byte) *proto.Status
	AuthId() string
	GetDirEntry(dir_path string, page, page_size uint64) *proto.DirEntry
	GetFileResource(file_path string) *proto.FileResource
	Close()
}

func NewPlugin() IPlugin {
	return &PluginImpl{}
}
