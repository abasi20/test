package main

import api "./api"

// "io/ioutil"
// "gopkg.in/yaml.v2"

func main() {
	// var c conf
	// conf := c.getConf()
	// fmt.Println(conf.Host, conf.Table)
	what := "this what"
	api.What(what)
	api.HelloWorld(what)
	println(what)
}
