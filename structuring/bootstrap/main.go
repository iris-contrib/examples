package main

import (
	"github.com/iris-contrib/examples/structuring/bootstrap/bootstrap"
	"github.com/iris-contrib/examples/structuring/bootstrap/middleware/identity"
	"github.com/iris-contrib/examples/structuring/bootstrap/routes"
)

var app = bootstrap.New("Awesome App", "kataras2006@hotmail.com",
	identity.Configure,
	routes.Configure,
)

func init() {
	app.Bootstrap()
}

func main() {
	app.Listen(":8080")
}
