package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/view"
	"gopkg.in/kataras/iris.v6/middleware/logger"

	"github.com/iris-contrib/examples/AIO_examples/mongo/backend/api"
	"github.com/iris-contrib/examples/AIO_examples/mongo/backend/db"
	"github.com/iris-contrib/examples/AIO_examples/mongo/backend/routes"
)

var (
	app *iris.Framework
)

func init() {
	app = iris.New()
}

func main() {
	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())
	// set the template engine
	app.Adapt(view.HTML("../frontend/templates", ".html").Layout("layout.html"))
	// set the favicon
	app.Favicon("../frontend/public/images/favicon.ico")

	// set static folder(s)
	app.StaticWeb("/public", "../frontend/public")

	// set the global http request middleware
	app.Use(logger.New())

	// set the custom errors
	app.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.Render("errors/404.html", iris.Map{"Title": iris.StatusText(iris.StatusNotFound)})
	})

	app.OnError(iris.StatusInternalServerError, func(ctx *iris.Context) {
		ctx.Render("errors/500.html", nil, iris.RenderOptions{"layout": iris.NoLayout})
	})

	// DB Main
	DbMain()
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
	// this is other way to declare routes using the 'API'
	auth := new(api.AuthAPI)

	// Custom handler
	app.Handle("GET", "/v1/blog/news", api.CustomAPI{})
	// Function handler
	app.Post("/v1/auth/login", auth.Login)
	app.Post("/v1/auth/register", auth.Register)
	app.Get("/v1/auth/check", auth.Check)
	app.Get("/v1/auth/session", auth.Session)
	// Api handler
	users := app.Party("/v1/users")
	{
		users.Get("/", api.GetAllUsers)
		users.Get("/:userid", api.GetUserByID)
		users.Put("/", api.UpdateUser)
		users.Post("/:userid", api.IsertUser)
		users.Delete("/", api.DeleteUser)
	}

}

func DbMain() {
	// Database Main Conexion
	Db := db.MgoDb{}
	Db.Init()
	// index keys
	keys := []string{"email"}
	Db.Index("auth", keys)

}
