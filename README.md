
## How to develop a Driver plugin
1. Fork this repo
2. Install `go1.24rc1`
```
go install golang.org/dl/go1.24rc1@latest
```
3. impl this `interface` in `plugin/plugin_impl.go`
```
type IPlugin interface {
	// plugin id
	Id() string
	// get auth type like form edit,qrcode,oauth2
	GetAuthType() *proto.AuthType
	// check auth data status
	CheckAuth(*proto.AuthType) *proto.Status
	// get auth data when check auth status is success
	GetAuthData() []byte
	// use auth data init auth
	InitAuth([]byte) *proto.Status
	// plugin auth id,it need unqiue for same driver
	AuthId() string
	// get dir entry from driver plugin
	GetDirEntry(dir_path string, page, page_size uint64) *proto.DirEntry
	// get file entry resource from driver plugin
	GetFileResource(file_path string) *proto.FileResource
	// close driver plugin
	Close()
}
```
4. Modify Your Driver Plugin info and plugin icon in file `plugin.toml`
5. Build For Driver Plugin
```
make
```
6. get you driver plugin `{name}.zip`
