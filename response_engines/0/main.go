package main

import (
	"encoding/xml"

	"github.com/kataras/iris"
)

type ExampleXml struct {
	XMLName xml.Name `xml:"example"`
	One     string   `xml:"one,attr"`
	Two     string   `xml:"two,attr"`
}

func main() {
	iris.Get("/data", func(ctx *iris.Context) {
		ctx.Data(iris.StatusOK, []byte("Some binary data here."))
	})

	iris.Get("/text", func(ctx *iris.Context) {
		ctx.Text(iris.StatusOK, "Plain text here")
	})

	iris.Get("/json", func(ctx *iris.Context) {
		ctx.JSON(iris.StatusOK, map[string]string{"hello": "json"}) // or myjsonStruct{hello:"json}
	})

	iris.Get("/jsonp", func(ctx *iris.Context) {
		ctx.JSONP(iris.StatusOK, "callbackName", map[string]string{"hello": "jsonp"})
	})

	iris.Get("/xml", func(ctx *iris.Context) {
		ctx.XML(iris.StatusOK, ExampleXml{One: "hello", Two: "xml"}) // or iris.Map{"One":"hello"...}
	})

	iris.Get("/markdown", func(ctx *iris.Context) {
		ctx.Markdown(iris.StatusOK, "# Hello Dynamic Markdown Iris")
	})

	iris.Listen(":8080")
}
