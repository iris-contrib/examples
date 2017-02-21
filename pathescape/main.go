package main

import (
	"gopkg.in/kataras/iris.v6"
)

type MyTemplateData struct {
	Values iris.Map // just a shortcat of map[string]interface{}
	Name   string
}

func main() {
	// enable path escape in order
	// to accept something like http://localhost:8080/profile/my%20name
	// as http://localhost:8080/profile/my name
	// if it's false (as defaulted) and you want to escape a parameter name you can just use
	// ctx.ParamDecoded instead of ctx.Param, see 'named_parameters_pathescape' example
	iris.Config.EnablePathEscape = true

	// http://localhost:8080/profile/yourname
	iris.Get("/profile/:username", func(ctx *iris.Context) {
		name := ctx.Param("username")

		ctx.Writef(name)
	})

	iris.Listen(":8080")
}
