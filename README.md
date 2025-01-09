
## How to develop a Driver plugin
1. Fork this repo
2. impl this `interface` in `plugin/plugin_impl.go`
```
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
```
3. Modify Your Driver Plugin info and plugin icon in file `plugin.toml`
4. Build For Driver Plugin
```
make
```
5. get you driver plugin `{name}.zip`
