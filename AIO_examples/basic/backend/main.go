package main

import (
	"github.com/iris-contrib/iris-command-assets/basic/backend/api"
	"github.com/iris-contrib/iris-command-assets/basic/backend/routes"

	"github.com/iris-contrib/middleware/logger"
	"github.com/kataras/iris"
)

func main() {
	// set the configs
	iris.Config.Render.Template.Directory = "../frontend/templates"
	iris.Config.Render.Template.Layout = "layout.html"

	// set the favicon
	iris.Favicon("../frontend/public/images/favicon.ico")

	// set static folder(s)
	iris.Static("/public", "../frontend/public", 1)

	// set the global middlewares
	iris.Use(logger.New(iris.Logger))

	// set the custom errors
	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.Render("errors/404.html", iris.Map{"Title": iris.StatusText(iris.StatusNotFound)})
	})

	iris.OnError(iris.StatusInternalServerError, func(ctx *iris.Context) {
		ctx.Render("errors/500.html", nil, iris.NoLayout)
	})

	// register the routes & the public API
	registerRoutes()
	registerAPI()

	// start the server
	iris.Listen("127.0.0.1:80")
}

func registerRoutes() {
	// register index using a 'Handler'
	iris.Handle("GET", "/", routes.Index())

	// this is other way to declare a route
	// using a 'HandlerFunc'
	iris.Get("/about", routes.About)

	// Dynamic route

	iris.Get("/profile/:username", routes.Profile)("user-profile")
	// user-profile is the custom,optional, route's Name: with this we can use the {{ url "user-profile" $username}} inside userlist.html

	iris.Get("/all", routes.UserList)
}

func registerAPI() {
	// this is other way to declare routes using the 'API'
	iris.API("/users", api.UserAPI{})
}
