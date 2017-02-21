// Package main same example as template_html_4 but with wildcard subdomains
package main

import (
	"gopkg.in/kataras/iris.v6"
)

func main() {

	wildcard := iris.Party("*.")
	{
		wildcard.Get("/mypath", emptyHandler)("dynamic-subdomain1")
		wildcard.Get("/mypath2/:param1/:param2", emptyHandler)("dynamic-subdomain2")
		wildcard.Get("/mypath3/:param1/statichere/:param2", emptyHandler)("dynamic-subdomain3")
		wildcard.Get("/mypath4/:param1/statichere/:param2/:otherparam/*something", emptyHandler)("dynamic-subdomain4")
		wildcard.Get("/mypath5/:param1/:param2/staticParam/:param3AfterStatic", emptyHandler)("dynamic-subdomain5")
	}

	iris.Get("/", func(ctx *iris.Context) {
		// for dynamic_subdomain:8080/mypath5...
		// the first parameter is always the subdomain part
		paramsAsArray := []string{"username5", "theParam1", "theParam2", "theParam3"}

		if err := ctx.Render("page.html", iris.Map{"ParamsAsArray": paramsAsArray}); err != nil {
			panic(err)
		}
	})

	iris.Get("/redirect/:namedRoute/:subdomain", func(ctx *iris.Context) {
		routeName := ctx.Param("namedRoute")
		subdomain := ctx.Param("subdomain")
		println("The full uri of " + routeName + "is: " + iris.URL(routeName, subdomain))
		// if routeName == "dynamic-subdomain1" && subdomain == "username1"
		// prints: The full uri ofd ynamic-subdomain1 is: http://username1.127.0.0.1:8080/mypath
		ctx.RedirectTo(routeName, subdomain) // the second parameter is the arguments, the first argument for dynamic subdomains is the subdomain part, after this, the named parameters
		// http://127.0.0.1:8080/redirect/my-subdomain1 will redirect to ->  http://username1.127.0.0.1:8080/mypath
	})

	iris.Listen("127.0.0.1:8080")
}

func emptyHandler(ctx *iris.Context) {
	ctx.Writef("[SUBDOMAIN: %s]Hello from Path: %s.", ctx.Subdomain(), ctx.Path())
}
