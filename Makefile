CHECK_PROGRAM := $(shell which plugin_build1 2>/dev/null)
ifeq ($(CHECK_PROGRAM),)
   $(error "plugin_build not found,install cmd: go install github.com/labulakalia/plugin_api/cmd/plugin_build@latest")
endif
build:
	plugin_build
