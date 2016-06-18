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
			c.Write("HEY FROM dashboard")
		})
	}
	system := app.Party("system.")
	{
		system.Get("/", func(c *iris.Context) {
			c.Write("HEY FROM system")
		})
	}

	app.Get("/", func(c *iris.Context) {
		c.Write("HEY FROM frontend /")
	})

	app.Listen("domain.local:80")
}
