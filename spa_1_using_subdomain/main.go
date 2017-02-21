package main

import (
	"gopkg.in/kataras/iris.v6"
)

// hosts: add these lines: (use tabs to separate them, if that doesn't works for you)
// 127.0.0.1 mydomain.com
// 127.0.0.1 api.mydomain.com
func main() {

	// OPTIONAL if you want diffent domain for static files
	// (I do this way but you are not forced to do the same)
	//
	// create the subdomain s.mydomain.com to serve our resources files
	// http://s.mydomain.com/resources/css/bootstrap.min.css ...
	// iris.Party("s.").StaticWeb("/", "./www/resources")
	//

	//
	// REGISTER YOUR REST API
	//

	// http://api.mydomain.com/ ...
	// brackers are optional, it's just a visual declaration.
	api := iris.Party("api.")
	{
		// http://api.mydomain.com/users/42
		api.Get("/users/:userid", func(ctx *iris.Context) {
			ctx.Writef("user with id: %s", ctx.Param("userid"))
		})

	}

	//
	// REGISTER THE PAGE AND ALL OTHER STATIC FILES
	// INCLUDING A FAVICON, CSS, JS and so on
	//

	// http://mydomain.com , here should be your index.html
	// which is the SPA frontend page
	iris.StaticWeb("/", "./www")

	//
	// START THE SERVER
	//
	iris.Listen("mydomain.com:80")
}
