package main

import (
	"flag"
)

func main() {

	var cmd = flag.String("d", "error", "Description: -d start/stop")
	flag.Parse()
	switch *cmd {
	case "start":
		println("开始")
	case "stop":
		println("停止")
	default:
		println("参数不对")
	}

}
