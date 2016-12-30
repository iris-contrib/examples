package main

import (
	"encoding/xml"
	"gopkg.in/kataras/iris.v5"
)

type myxml struct {
	XMLName xml.Name `xml:"xml_example"`
	First   string   `xml:"first,attr"`
	Second  string   `xml:"second,attr"`
}

func main() {

	iris.Get("/", func(ctx *iris.Context) {
		ctx.XML(iris.StatusOK, iris.Map{"first": "first attr ", "second": "second attr"})
	})

	iris.Get("/alternative_1", func(ctx *iris.Context) {
		ctx.XML(iris.StatusOK, myxml{First: "first attr", Second: "second attr"})
	})

	iris.Get("/alternative_2", func(ctx *iris.Context) {
		ctx.Render("text/xml", myxml{First: "first attr", Second: "second attr"})
	})

	iris.Get("/alternative_3", func(ctx *iris.Context) {
		ctx.RenderWithStatus(iris.StatusOK, "text/xml", myxml{First: "first attr", Second: "second attr"})
	})

	iris.Get("/alternative_4", func(ctx *iris.Context) {
		ctx.Render("text/xml", myxml{First: "first attr", Second: "second attr"}, iris.RenderOptions{"charset": "UTF-8"}) // UTF-8 is the default.
	})

	iris.Get("/alternative_5", func(ctx *iris.Context) {
		// logs if any error and sends http status '500 internal server error' to the client
		ctx.MustRender("text/xml", myxml{First: "first attr", Second: "second attr"}, iris.RenderOptions{"charset": "UTF-8"})
	})

	iris.Listen(":8080")
}
