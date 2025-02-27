
## How to develop a Driver plugin
1. Fork this repo
2. Install `go1.24rc1`
3. impl this `interface` in `plugin_impl.go`
```
type IPlugin interface {
	// plugin id
	PluginId() (string, error)
	// get auth type like form edit,qrcode,oauth2
	GetAuthType() (*plugin.AuthType, error)
	// check auth data status
	CheckAuth(*plugin.AuthType) (*plugin.Status, error)
	// get auth data when check auth status is success
	GetAuthData() ([]byte, error)
	// use auth data init auth
	CheckAuthData([]byte) (*plugin.Status, error)
	// plugin auth id,it need unqiue for same driver
	PluginAuthId() (string, error)
	// get dir entry from driver plugin
	GetDirEntry(dir_path string, page, page_size uint64) (*plugin.DirEntry, error)
	// get file entry resource from driver plugin
	GetFileResource(file_path string) (*plugin.FileResource, error)
	// close driver plugin
	Close() error
}
```
4. Modify Your Driver Plugin info and plugin icon in file `plugin.toml`
5. Build For Driver Plugin
```
make
```
6. get you driver plugin `dist/{name}.zip`
