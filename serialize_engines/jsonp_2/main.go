package main

import (
	"github.com/kataras/go-serializer/jsonp"
	"gopkg.in/kataras/iris.v6"
)

type myjson struct {
	Name string `json:"name"`
}

func main() {
	iris.Config.Charset = "UTF-8" // this is the default, which you can change

	//first example
	// this is one of the reasons you need to import a default engine,(template engine or serialize engine(serializer))
	/*
		type Config struct {
			Indent   bool
			Callback string // the callback can be override by the context's options or parameter on context.JSONP
		}
	*/
	iris.UseSerializer(jsonp.ContentType, jsonp.New(jsonp.Config{
		Indent: true,
	}))
	// you can use anything as the second parameter,
	// the jsonp.ContentType is the string "application/javascript",
	// the context.JSONP renders with this engine's key.

	handlerSimple := func(ctx *iris.Context) {
		ctx.JSONP(iris.StatusOK, "callbackName", myjson{Name: "iris"})
	}

	handlerWithRender := func(ctx *iris.Context) {
		// you can also change the charset for a specific render action with RenderOptions
		ctx.Render("application/javascript", myjson{Name: "iris"}, iris.RenderOptions{"callback": "callbackName", "charset": "8859-1"})
	}

	//second example,
	// but we also want a different renderer, but again "application/javascript" as content type, with Callback option setted globaly:
	iris.UseSerializer("jsonp2", jsonp.New(jsonp.Config{Callback: "callbackName"}))
	// yes the UseSerializer returns a function which you can map the content type if it's not declared on the key
	handlerJsonp2 := func(ctx *iris.Context) {
		ctx.Render("jsonp2", myjson{Name: "My iris"})
		ctx.SetContentType("application/javascript")
	}

	iris.Get("/", handlerSimple)

	iris.Get("/render", handlerWithRender)

	iris.Get("/jsonp2", handlerJsonp2)

	iris.Listen(":8080")
}
