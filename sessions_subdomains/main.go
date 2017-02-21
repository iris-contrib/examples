package main

import (
	"gopkg.in/kataras/iris.v6"
)

func main() {
	domain := "localhost.com"

	// TO DISABLE SUBDOMAIN PERSISTANCE USE THAT:
	// iris.Config.Sessions.DisableSubdomainPersistence = true

	// set the subdomain from anywhere, here we do from domain
	iris.Get("/", func(ctx *iris.Context) {
		//set the session
		myusername := "iris"
		ctx.Session().Set("username", myusername)
		ctx.Writef("Username setted to %s, go to /user or mysubdomain.%s:8080/user to view the session's username's value", myusername, domain)
	})

	iris.Get("/user", func(ctx *iris.Context) {
		// get the session
		myusername := ctx.Session().GetString("username")
		ctx.Writef("Hello, your username is %s", myusername)
	})

	subdomain := iris.Party("mysubdomain.")
	{
		// set the session from anywhere either from subdomain either from the domain, here we do from subdomain
		subdomain.Get("/", func(ctx *iris.Context) {
			//set the session
			myusername := "iris"
			ctx.Session().Set("username", myusername)
			ctx.Writef("Username setted FROM THE SUBDOMAIN to %s, go to /user or mysubdomain.%s/user to view the session's username's value", myusername, domain)
		})

		subdomain.Get("/user", func(ctx *iris.Context) {
			// get the session
			myusername := ctx.Session().GetString("username")
			ctx.Writef("Hello from subdomain, your username is %s", myusername)
		})
	}

	iris.Listen(domain + ":8080")
}
