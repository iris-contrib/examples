package main

import (
	"github.com/dgrijalva/jwt-go"
	// jwtmiddleware is a community middleware use with caution or select other middleware
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	app := iris.New()
	// output startup banner and error logs on os.Stdout
	app.Adapt(iris.DevLogger())
	// set the router, you can choose gorillamux too
	app.Adapt(httprouter.New())

	myJwtMiddleware := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("My Secret"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	app.Get("/ping", PingHandler)

	app.Get("/secured/ping", myJwtMiddleware.Serve, SecuredPingHandler)
	app.Listen(":8080")

}

type Response struct {
	Text string `json:"text"`
}

func PingHandler(ctx *iris.Context) {
	response := Response{"All good. You don't need to be authenticated to call this"}
	ctx.JSON(iris.StatusOK, response)
}

func SecuredPingHandler(ctx *iris.Context) {
	response := Response{"All good. You only get this message if you're authenticated"}
	// get the *jwt.Token which contains user information using:
	// user:= myJwtMiddleware.Get(ctx) or context.Get("jwt").(*jwt.Token)
	ctx.JSON(iris.StatusOK, response)
}
