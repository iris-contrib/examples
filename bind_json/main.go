package main

import (
	"gopkg.in/kataras/iris.v4"
)

type Company struct {
	Name string
	City string
}

func MyHandler(ctx *iris.Context) {
	c := &Company{}
	if err := ctx.ReadJSON(c); err != nil {
		panic(err.Error())
	} else {
		ctx.Write("Company: %#v", c)
	}
}

func main() {
	// use postman or whatever to do a POST request to the localhost:8080/bind_json with BODY: JSON PAYLOAD and HEADERS content type to application/json
	iris.Post("/bind_json", MyHandler)
	iris.Listen(":8080")
}
