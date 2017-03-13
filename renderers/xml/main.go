package main

import (
	"encoding/xml"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

type myxml struct {
	XMLName xml.Name `xml:"xml_example"`
	First   string   `xml:"first,attr"`
	Second  string   `xml:"second,attr"`
}

func main() {
	app := iris.New()
	// output startup banner and error logs on os.Stdout
	app.Adapt(iris.DevLogger())
	// set the router, you can choose gorillamux too
	app.Adapt(httprouter.New())

	app.Get("/", func(ctx *iris.Context) {
		ctx.XML(iris.StatusOK, iris.Map{"first": "first attr ", "second": "second attr"})
	})

	app.Get("/alternative_1", func(ctx *iris.Context) {
		ctx.XML(iris.StatusOK, myxml{First: "first attr", Second: "second attr"})
	})

	app.Get("/alternative_2", func(ctx *iris.Context) {
		ctx.Render("text/xml", myxml{First: "first attr", Second: "second attr"})
	})

	app.Get("/alternative_3", func(ctx *iris.Context) {
		ctx.RenderWithStatus(iris.StatusOK, "text/xml", myxml{First: "first attr", Second: "second attr"})
	})

	app.Get("/alternative_4", func(ctx *iris.Context) {
		ctx.Render("text/xml", myxml{First: "first attr", Second: "second attr"}, iris.RenderOptions{"charset": "UTF-8"}) // UTF-8 is the default.
	})

	app.Get("/alternative_5", func(ctx *iris.Context) {
		// logs if any error and sends http status '500 internal server error' to the client
		ctx.MustRender("text/xml", myxml{First: "first attr", Second: "second attr"}, iris.RenderOptions{"charset": "UTF-8"})
	})

	app.Listen(":8080")
}
