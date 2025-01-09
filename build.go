package main

import (
	"archive/zip"
	"compress/flate"
	"fmt"
	"io"
	"plugin_temp/proto"

	"log/slog"
	"os"
	"os/exec"

	"github.com/BurntSushi/toml"
)

func main() {
	pluginTomlFile := "plugin.toml"
	driverPluginConfig := &proto.DriverPluginConfig{}
	_, err := toml.DecodeFile(pluginTomlFile, driverPluginConfig)
	if err != nil {
		slog.Error("decode plugin.toml failed", "err", err)
		os.Exit(1)
	}
	if len(driverPluginConfig.Icon) > 0 {
		_, err = os.Stat(driverPluginConfig.Icon)
		if err != nil {
			slog.Error("decode plugin.toml failed", "err", err)
			os.Exit(1)
		}
	}
	slog.Info("driver plugin start building...")
	buildWasmFile := fmt.Sprintf("%s.wasm", driverPluginConfig.Id)
	defer os.Remove(buildWasmFile)
	buildCmd := fmt.Sprintf("GOOS=wasip1 GOARCH=wasm go1.24rc1 build -buildmode=c-shared -o %s export.go", buildWasmFile)
	cmd := exec.Command("bash", "-c", buildCmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		slog.Error("exec command failed", "err", err)
		os.Exit(1)
	}
	outCompressFile := fmt.Sprintf("%s.zip", driverPluginConfig.Id)
	outFile, err := os.Create(fmt.Sprintf("%s.zip", driverPluginConfig.Id))
	if err != nil {
		slog.Error("create file failed", "file", outCompressFile, "err", err)
		os.Exit(1)
	}
	defer outFile.Close()
	zw := zip.NewWriter(outFile)
	zw.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(out, flate.BestCompression)
	})
	defer zw.Close()

	fileList := []string{
		pluginTomlFile,
		fmt.Sprintf("%s.wasm", driverPluginConfig.Id),
		driverPluginConfig.Icon,
	}

	for _, file := range fileList {
		localFile, err := os.Open(file)
		if err != nil {
			slog.Error("zip create file failed", "file", file, "err", err)
			os.Exit(1)
		}

		zwfile, err := zw.Create(file)
		if err != nil {
			slog.Error("zip create file failed", "file", file, "err", err)
			os.Exit(1)
		}
		io.Copy(zwfile, localFile)
		localFile.Close()
	}
	slog.Info("driver plugin build success", "file", outCompressFile)
}
