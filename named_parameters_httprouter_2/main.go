package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	app := iris.New()
	// output startup banner and error logs on os.Stdout
	app.Adapt(iris.DevLogger())
	// set the router, you can choose gorillamux too
	app.Adapt(httprouter.New())

	app.UseFunc(func(ctx *iris.Context) {
		ctx.Log(iris.DevMode, ctx.Method()+": "+ctx.Path()+"\n")
		ctx.Next()
	})

	app.Get("/healthcheck", h)
	app.Get("/games/:gameID/clans", h)
	app.Get("/games/:gameID/clans/clan/:publicID", h)
	app.Get("/games/:gameID/clans/search", h)

	app.Put("/games/:gameID/players/:publicID", h)
	app.Put("/games/:gameID/clans/clan/:publicID", h)

	app.Post("/games/:gameID/clans", h)
	app.Post("/games/:gameID/players", h)
	app.Post("/games/:gameID/clans/:publicID/leave", h)
	app.Post("/games/:gameID/clans/:clanPublicID/memberships/application", h)
	app.Post("/games/:gameID/clans/:clanPublicID/memberships/application/:action", h)
	app.Post("/games/:gameID/clans/:clanPublicID/memberships/invitation", h)
	app.Post("/games/:gameID/clans/:clanPublicID/memberships/invitation/:action", h)
	app.Post("/games/:gameID/clans/:clanPublicID/memberships/delete", h)
	app.Post("/games/:gameID/clans/:clanPublicID/memberships/promote", h)
	app.Post("/games/:gameID/clans/:clanPublicID/memberships/demote", h)

	app.Listen(":80")

	/*
		gameID  = 1
		publicID = 2
		clanPublicID = 22
		action = 3

		GET
		http://localhost/healthcheck
		http://localhost/games/1/clans
		http://localhost/games/1/clans/clan/2
		http://localhost/games/1/clans/search

		PUT
		http://localhost/games/1/players/2
		http://localhost/games/1/clans/clan/2

		POST
		http://localhost/games/1/clans
		http://localhost/games/1/players
		http://localhost/games/1/clans/2/leave
		http://localhost/games/1/clans/22/memberships/application -> 494
		http://localhost/games/1/clans/22/memberships/application/3- > 404
		http://localhost/games/1/clans/22/memberships/invitation
		http://localhost/games/1/clans/22/memberships/invitation/3
		http://localhost/games/1/clans/2/memberships/delete
		http://localhost/games/1/clans/22/memberships/promote
		http://localhost/games/1/clans/22/memberships/demote

	*/
}

func h(ctx *iris.Context) {
	ctx.HTML(iris.StatusOK, "<h1>Path<h1/>"+ctx.Path())
}
