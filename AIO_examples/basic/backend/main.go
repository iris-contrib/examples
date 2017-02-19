package main

import (
	"log"

	"github.com/iris-contrib/examples/AIO_examples/basic/backend/api/user"
	"github.com/iris-contrib/examples/AIO_examples/basic/backend/routes"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/view"
	"gopkg.in/kataras/iris.v6/middleware/logger"
)

var app *iris.Framework

func init() {
	app = iris.New(iris.Configuration{Gzip: false, Charset: "UTF-8"})
}

func main() {
	// set the router we want to use
	app.Adapt(httprouter.New())

	// adapt a new logger which will print the dev messages(mostly errors)
	// and panic on production messages (by-default only fatal errors are printed via ProdMode)
	app.Adapt(iris.LoggerPolicy(func(mode iris.LogMode, msg string) {
		if mode == iris.DevMode {
			log.Printf(msg)
		} else if mode == iris.ProdMode {
			panic(msg)
		}
	})) // or use app.Adapt(iris.DevLogger()) to print only DevMode messages to the os.Stdout

	// set the template engine
	app.Adapt(view.HTML("../frontend/templates", ".html").Layout("layout.html"))
	// set the favicon
	app.Favicon("../frontend/public/images/favicon.ico")

	// set static folder(s)
	app.StaticWeb("/public", "../frontend/public")

	// set the global middlewares
	app.Use(logger.New())

	// set the custom errors
	app.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.Render("errors/404.html", iris.Map{"Title": iris.StatusText(iris.StatusNotFound)})
	})

	app.OnError(iris.StatusInternalServerError, func(ctx *iris.Context) {
		ctx.Render("errors/500.html", nil, iris.RenderOptions{"layout": iris.NoLayout})
	})

	// register the routes & the public API
	registerRoutes()
	registerAPI()

	// start the server
	app.Listen("127.0.0.1:8080")
}

func registerRoutes() {
	// register index using a 'Handler'
	app.Handle("GET", "/", routes.Index())

	// this is other way to declare a route
	// using a 'HandlerFunc'
	app.Get("/about", routes.About)

	// Dynamic route

	app.Get("/profile/:username", routes.Profile).ChangeName("user-profile")
	// user-profile is the custom,optional, route's Name: with this we can use the {{ url "user-profile" $username}} inside userlist.html

	app.Get("/all", routes.UserList)
}

func registerAPI() {
	p := app.Party("/users")
	{
		p.Get("/", userAPI.GetAll)
		p.Get("/:id", userAPI.GetByID)
		p.Put("/", userAPI.Insert)
		p.Post("/:id", userAPI.Update)
		p.Delete("/:id", userAPI.DeleteByID)
	}
}
