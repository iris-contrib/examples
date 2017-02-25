package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	app := iris.New()
	// output startup banner and error logs on os.Stdout
	app.Adapt(iris.DevLogger())
	// set the router, you can choose gorillamux too
	app.Adapt(httprouter.New())
	// NOTE:
	// if app.Config.EnablePathEscape = true and
	// Request: http://localhost:8080/details/Project%2FDelta then
	// it will pass the request
	// as http://localhost:8080/project/details/Project/Delta
	// and it will response as NOT FOUND
	// because iris.Config.EnablePathEscape decodes to query the whole path
	// before parsing the path parameters.

	// accepts %2F as slash '/' wih ParamDecoded
	// Request: http://localhost:8080/details/Project%2FDelta then
	// ctx.Param("project") returns the raw named parameter: Project%2FDelta
	// ctx.ParamDecoded("project") returns Project/Delta
	app.Get("/details/:project", func(ctx *iris.Context) {
		projectName := ctx.Param("project")
		projectNameDecoded := ctx.ParamDecoded("project")
		ctx.Writef("Raw: %s\nDecoded: %s", projectName, projectNameDecoded)
	})

	app.Listen(":8080")
}
