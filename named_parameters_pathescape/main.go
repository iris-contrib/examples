package main

import "github.com/kataras/iris"

func main() {
	iris.Set(iris.OptionDisablePathEscape(true))
	// accepts parameters with slash '/'
	// Request: http://localhost:8080/details/Project%2FDelta
	// ctx.Param("project") returns the raw named parameter: Project%2FDelta
	// ctx.ParamDecoded("project") returns Project/Delta
	iris.Get("/details/:project", func(ctx *iris.Context) {
		projectName := ctx.Param("project")
		projectNameDecoded := ctx.ParamDecoded("project")
		ctx.Writef("Raw: %s\nDecoded: %s", projectName, projectNameDecoded)
	})

	iris.Listen(":8080")
}
