package main

import (
	encodingXML "encoding/xml"

	"github.com/iris-contrib/response/xml"
	"github.com/kataras/iris"
)

type myxml struct {
	XMLName encodingXML.Name `xml:"xml_example"`
	First   string           `xml:"first,attr"`
	Second  string           `xml:"second,attr"`
}

func main() {
	iris.Config.Charset = "UTF-8" // this is the default, which you can change

	//first example
	// this is one of the reasons you need to import a default engine,(template engine or response engine)
	/*
		type Config struct {
			Indent bool
			Prefix []byte
		}
	*/
	iris.UseResponse(xml.New(xml.Config{
		Indent: true,
	}), xml.ContentType)
	// you can use anything as the second parameter,
	// the jsonp.ContentType is the string "text/xml",
	// the context.XML renders with this engine's key.

	handlerSimple := func(ctx *iris.Context) {
		ctx.XML(iris.StatusOK, myxml{First: "first attr", Second: "second attr"})
	}

	handlerWithRender := func(ctx *iris.Context) {
		// you can also change the charset for a specific render action with RenderOptions
		ctx.Render("text/xml", myxml{First: "first attr", Second: "second attr"}, iris.RenderOptions{"charset": "8859-1"})
	}

	//second example,
	// but we also want a different renderer, but again "text/xml" as content type, with prefix option setted by configuration:
	iris.UseResponse(xml.New(xml.Config{Prefix: []byte("")}), "xml2")("text/xml") // if you really use a PREFIX it will be not valid xml, use it only for special cases
	// yes the UseResponse returns a function which you can map the content type if it's not declared on the key
	handlerXML2 := func(ctx *iris.Context) {
		ctx.Render("xml2", myxml{First: "first attr", Second: "second attr"})
	}

	iris.Get("/", handlerSimple)

	iris.Get("/render", handlerWithRender)

	iris.Get("/xml2", handlerXML2)

	iris.Listen(":8080")
}
