package main

import "github.com/kataras/iris"

func main() {
	iris.Config.DisablePathEscape = true
	// accepts parameters with slash '/'
	// Request: http://localhost:8080/details/Project%2FDelta
	// ctx.Param("project") returns the raw named parameter: Project%2FDelta
	// which you can escape it manually with net/url: projectName, _ := url.QueryUnescape(c.Param("project"))
	// With DisablePathEscape = false this will redirect to 404 not found error because of the Project/Delta
	// Look here: https://github.com/kataras/iris/issues/135
	iris.Get("/details/:project", func(ctx *iris.Context) {
		projectName := ctx.Param("project")
		ctx.Write("%s", projectName)
	})

	iris.Listen(":8080")
}
