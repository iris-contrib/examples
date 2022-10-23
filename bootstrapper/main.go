package main

import (
	"github.com/iris-contrib/examples/bootstrapper/bootstrap"
	"github.com/iris-contrib/examples/bootstrapper/middleware/identity"
	"github.com/iris-contrib/examples/bootstrapper/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Awesome App", "kataras2006@hotmail.com")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
