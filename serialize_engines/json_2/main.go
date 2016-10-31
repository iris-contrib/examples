package main

import (
	"github.com/kataras/go-serializer/json"
	"gopkg.in/kataras/iris.v4"
)

type myjson struct {
	Name string `json:"name"`
}

func main() {
	iris.Config.Charset = "UTF-8" // this is the default, which you can change

	//first example
	// use the json's Config, we need the import of the json serialize engine(serializer) in order to change its internal configs
	// this is one of the reasons you need to import a default engine,(template engine or serialize engine(serializer))
	/*
		type Config struct {
			Indent        bool
			UnEscapeHTML  bool
			Prefix        []byte
			StreamingJSON bool
		}
	*/
	iris.UseSerializer(json.ContentType, json.New(json.Config{
		Prefix: []byte("MYPREFIX"),
	})) // you can use anything as the second parameter, the json.ContentType is the string "application/json", the context.JSON renders with this engine's key.

	jsonHandlerSimple := func(ctx *iris.Context) {
		ctx.JSON(iris.StatusOK, myjson{Name: "iris"})
	}

	jsonHandlerWithRender := func(ctx *iris.Context) {
		// you can also change the charset for a specific render action with RenderOptions
		ctx.Render("application/json", myjson{Name: "iris"}, iris.RenderOptions{"charset": "8859-1"})
	}

	//second example,
	// imagine that we need the context.JSON to be listening to our "application/json" serialize engine(serializer) with a custom prefix (we did that before)
	// but we also want a different renderer, but again application/json content type, with Indent option setted to true:
	iris.UseSerializer("json2", json.New(json.Config{Indent: true}))
	json2Handler := func(ctx *iris.Context) {
		ctx.Render("json2", myjson{Name: "My iris"})
		ctx.SetContentType("application/json")
	}

	iris.Get("/", jsonHandlerSimple)

	iris.Get("/render", jsonHandlerWithRender)

	iris.Get("/json2", json2Handler)

	iris.Listen(":8080")
}
