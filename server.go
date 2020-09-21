package main

import (
	"fmt"

	api "./api"
)

// api "helloword/api"

// "github.com/labstack/echo"
func main() {
	// e := echo.New()
	// e.GET("/", api.HelloWorld)
	// e.Logger.Fatal(e.Start(":1323"))
	k := 20
	println(k)
	api.HelloWorld("famcloud")
	type Human struct {
		Name string
	}

	var people = Human{Name: "春生"}
	println(people.Name)
	fmt.Printf("%x\n", 13)
}
