package main

import (
	"github.com/iris-contrib/response/json"
	"github.com/kataras/iris"
)

type myjson struct {
	Name string `json:"name"`
}

func main() {
	iris.Config.Charset = "UTF-8" // this is the default, which you can change

	//first example
	// use the json's Config, we need the import of the json response engine in order to change its internal configs
	// this is one of the reasons you need to import a default engine,(template engine or response engine)
	/*
		type Config struct {
			Indent        bool
			UnEscapeHTML  bool
			Prefix        []byte
			StreamingJSON bool
		}
	*/
	iris.UseResponse(json.New(json.Config{
		Prefix: []byte("MYPREFIX"),
	}), json.ContentType) // you can use anything as the second parameter, the json.ContentType is the string "application/json", the context.JSON renders with this engine's key.

	jsonHandlerSimple := func(ctx *iris.Context) {
		ctx.JSON(iris.StatusOK, myjson{Name: "iris"})
	}

	jsonHandlerWithRender := func(ctx *iris.Context) {
		// you can also change the charset for a specific render action with RenderOptions
		ctx.Render("application/json", myjson{Name: "iris"}, iris.RenderOptions{"charset": "8859-1"})
	}

	//second example,
	// imagine that we need the context.JSON to be listening to our "application/json" response engine with a custom prefix (we did that before)
	// but we also want a different renderer, but again application/json content type, with Indent option setted to true:
	iris.UseResponse(json.New(json.Config{Indent: true}), "json2")("application/json")
	// yes the UseResponse returns a function which you can map the content type if it's not declared on the key
	json2Handler := func(ctx *iris.Context) {
		ctx.Render("json2", myjson{Name: "My iris"})
	}

	iris.Get("/", jsonHandlerSimple)

	iris.Get("/render", jsonHandlerWithRender)

	iris.Get("/json2", json2Handler)

	iris.Listen(":8080")
}
