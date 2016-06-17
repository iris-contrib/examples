package main

import "github.com/kataras/iris"

func main() {

	iris.UseFunc(func(ctx *iris.Context) {
		ctx.Log(ctx.MethodString() + ": " + ctx.PathString() + "\n")
		ctx.Next()
	})

	iris.Get("/healthcheck", h)
	iris.Get("/games/:gameID/clans", h)
	iris.Get("/games/:gameID/clans/clan/:publicID", h)
	iris.Get("/games/:gameID/clans/search", h)

	iris.Put("/games/:gameID/players/:publicID", h)
	iris.Put("/games/:gameID/clans/clan/:publicID", h)

	iris.Post("/games/:gameID/clans", h)
	iris.Post("/games/:gameID/players", h)
	iris.Post("/games/:gameID/clans/:publicID/leave", h)
	iris.Post("/games/:gameID/clans/:clanPublicID/memberships/application", h)
	iris.Post("/games/:gameID/clans/:clanPublicID/memberships/application/:action", h)
	iris.Post("/games/:gameID/clans/:clanPublicID/memberships/invitation", h)
	iris.Post("/games/:gameID/clans/:clanPublicID/memberships/invitation/:action", h)
	iris.Post("/games/:gameID/clans/:clanPublicID/memberships/delete", h)
	iris.Post("/games/:gameID/clans/:clanPublicID/memberships/promote", h)
	iris.Post("/games/:gameID/clans/:clanPublicID/memberships/demote", h)

	iris.Listen(":80")

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
	ctx.Write(ctx.PathString())
}
