package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

type MyData struct {
	Sysname   string // this will be the same for all requests
	Version   int    // this will be the same for all requests
	UserAgent string // this will be different for each request
}

type MyHandler struct {
	data MyData
}

func (m *MyHandler) Serve(ctx *iris.Context) {
	data := &m.data
	data.UserAgent = ctx.RequestHeader("User-Agent")
	ctx.Writef("Path: %s", ctx.Path())
	ctx.Writef("\nUser agent: %s", data.UserAgent)
	ctx.Writef("\nData always same: data.Sysname: %s and data.Version: %d", data.Sysname, data.Version)
}

func main() {
	app := iris.New()
	// output startup banner and error logs on os.Stdout
	app.Adapt(iris.DevLogger())
	// set the router, you can choose gorillamux too
	app.Adapt(httprouter.New())

	/*
	 Create the data with the predefined fields that will no change at every request
	*/
	myData := MyData{
		Sysname: "Redhat",
		Version: 1,
	}
	/*
		use a new MyHandler for each route,
		keep the same original myData with the fields no need to change on every request
	*/
	app.Handle("GET", "/", &MyHandler{myData})
	app.Handle("GET", "/about", &MyHandler{myData})

	app.Listen(":8080")
}
