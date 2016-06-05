// Package main register static subdomains, simple as parties, check ./hosts if you use windows
package main

import (
	"github.com/kataras/iris"
)

func main() {
	api := iris.New()

	// first the subdomains.
	admin := api.Party("admin.127.0.0.1")
	{
		// admin.127.0.0.1.com
		admin.Get("/", func(c *iris.Context) {
			c.Write("HEY FROM admin.127.0.0.1")
		})
		// admin.127.0.0.1/hey
		admin.Get("/hey", func(c *iris.Context) {
			c.Write("HEY FROM admin.127.0.0.1/hey")
		})
		// admin.127.0.0.1/hey2
		admin.Get("/hey2", func(c *iris.Context) {
			c.Write("HEY SECOND FROM admin.127.0.0.1/hey")
		})
	}

	// 127.0.0.1/hey
	api.Get("/hey", func(c *iris.Context) {
		c.Write("HEY FROM no-subdomain hey")
	})

	api.Listen("127.0.0.1:80")
}
