// Package main register static subdomains, simple as parties, check ./hosts if you use windows
// this example shows you how to set subdomains at a real domain for example: 'localhost.com'
package main

import (
	"github.com/kataras/iris"
)

func main() {
	api := iris.New()

	// first the subdomains.
	admin := api.Party("admin.localhost.com")
	{
		// admin.localhost.com
		admin.Get("/", func(c *iris.Context) {
			c.Write("HEY FROM admin.localhost.com")
		})
		// admin.localhost.com/hey
		admin.Get("/hey", func(c *iris.Context) {
			c.Write("HEY FROM admin.localhost.com/hey")
		})
		// admin.localhost.com/hey2
		admin.Get("/hey2", func(c *iris.Context) {
			c.Write("HEY SECOND FROM admin.localhost.com/hey")
		})
	}

	// localhost/hey
	api.Get("/hey", func(c *iris.Context) {
		c.Write("HEY FROM no-subdomain hey")
	})

	api.Listen("localhost:80")
}
