package main

import (
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
)

func main() {

	myJwtMiddleware := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("My Secret"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	iris.Get("/ping", PingHandler)

	iris.Get("/secured/ping", myJwtMiddleware.Serve, SecuredPingHandler)
	iris.Listen(":8080")

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
