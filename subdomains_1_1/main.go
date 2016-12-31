package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	/*
	 * Setup static files
	 */

	app.Static("/assets", "./public/assets", 1)
	app.Static("/upload_resources", "./public/upload_resources", 1)

	dashboard := app.Party("dashboard.")
	{
		dashboard.Get("/", func(c *iris.Context) {
			c.Writef("HEY FROM dashboard")
		})
	}
	system := app.Party("system.")
	{
		system.Get("/", func(c *iris.Context) {
			c.Writef("HEY FROM system")
		})
	}

	app.Get("/", func(c *iris.Context) {
		c.Writef("HEY FROM frontend /")
	})
	/* test this on firefox, because the domain is not real (because of .local), on firefox this will fail, but you can test it with other domain */
	app.Listen("domain.local:80")
}
