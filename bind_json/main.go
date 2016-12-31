package main

import (
	"fmt"

	"github.com/kataras/iris"
)

type Company struct {
	Name  string
	City  string
	Other string
}

func MyHandler(ctx *iris.Context) {
	c := &Company{}
	if err := ctx.ReadJSON(c); err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("Company: %#v\n", c)
		ctx.Writef("Company: %#v\n", c)
	}
}

func main() {
	// use postman or whatever to do a POST request to the localhost:8080/bind_json with BODY: JSON PAYLOAD and HEADERS content type to application/json
	iris.Post("/bind_json", MyHandler)
	iris.Listen(":8080")
}
