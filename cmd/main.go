package main

import (
	"flag"
	"mynginx/logic"
	"mynginx/pkg/ttviper"
)

var (
	filename = flag.String("f", "config.yaml", "set config filename")
	dir      = flag.String("d", "./etc", "set config dir")
)

func main() {
	flag.Parse()

	config := ttviper.ReadConfig(*dir, *filename)
	nginx := logic.ParseNginx(config)
	nginx.Run()
}
