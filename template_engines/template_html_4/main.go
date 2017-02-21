// Package main an example on how to naming your routes & use the custom 'url' HTML Template Engine, same for other template engines
// we don't need to import the kataras/go-template/html because iris uses this as the default engine if no other template engine has been registered.
package main

import (
	"gopkg.in/kataras/iris.v6"
)

func main() {

	iris.Get("/mypath", emptyHandler)("my-page1")
	iris.Get("/mypath2/:param1/:param2", emptyHandler)("my-page2")
	iris.Get("/mypath3/:param1/statichere/:param2", emptyHandler)("my-page3")
	iris.Get("/mypath4/:param1/statichere/:param2/:otherparam/*something", emptyHandler)("my-page4")

	// same with Handle/Func
	iris.HandleFunc("GET", "/mypath5/:param1/statichere/:param2/:otherparam/anything/*anything", emptyHandler)("my-page5")

	iris.Get("/mypath6/:param1/:param2/staticParam/:param3AfterStatic", emptyHandler)("my-page6")

	iris.Get("/", func(ctx *iris.Context) {
		// for /mypath6...
		paramsAsArray := []string{"theParam1", "theParam2", "theParam3"}

		if err := ctx.Render("page.html", iris.Map{"ParamsAsArray": paramsAsArray}); err != nil {
			panic(err)
		}
	})

	iris.Get("/redirect/:namedRoute", func(ctx *iris.Context) {
		routeName := ctx.Param("namedRoute")

		println("The full uri of " + routeName + "is: " + iris.URL(routeName))
		// if routeName == "my-page1"
		// prints: The full uri of my-page1 is: http://127.0.0.1:8080/mypath
		ctx.RedirectTo(routeName)
		// http://127.0.0.1:8080/redirect/my-page1 will redirect to -> http://127.0.0.1:8080/mypath
	})

	iris.Listen(":8080")
}

func emptyHandler(ctx *iris.Context) {
	ctx.Writef("Hello from %s.", ctx.Path())

}
