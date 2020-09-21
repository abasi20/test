package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type serverInfo struct {
	// tcp port
	Tcpport int16 `yaml:"tcpport"`
	// http
	Ssl      bool  `yaml:"ssl"`
	Httpport int16 `yaml:"httpport"`
	// mysql conf
	Mysqlpath string `yaml:"mysqlpath"`
	Mysqlport int16  `yaml:"mysqlport"`
	Mysqluser string `yaml:"mysqluser"`
	Mysqlpwd  string `yaml:"mysqlpwd"`
	Table     string `yaml:"table"`
}

func (c *serverInfo) getConf() *serverInfo {
	yamlFile, err := ioutil.ReadFile("conf/conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

var conf *serverInfo

// init
func init() {
	var c serverInfo
	conf = c.getConf()

}

// main
func main() {
	// test
	var cmd = flag.String("c", "error", "Description")
	var path = flag.String("d", "error", "Description")
	flag.Parse()
	println(conf.Mysqlport, conf.Mysqlpath)

	println(conf.Mysqlpath, conf.Mysqlport, conf.Mysqlpwd, conf.Mysqluser)

	if *path == "error" || *cmd == "error" {
		println("cunlllsba")
		os.Exit(1)
	}

}
