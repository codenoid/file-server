package cli

import (
	"flag"

	"github.com/codenoid/file-server/core"
)

func NewServeCmd() {

	var bind string
	var configFile string

	flag.StringVar(&bind, "bind", ":8080", "bind addr")
	flag.StringVar(&configFile, "config", "config.yml", "path to config file")
	flag.Parse()

	core.StartServer(bind, configFile)

}
