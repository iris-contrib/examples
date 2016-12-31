// Package main register static subdomains, simple as parties, check ./hosts if you use windows
package main

import (
	"github.com/kataras/iris"
)

func main() {
	api := iris.New()

	// no order, you can register subdomains at the end also.
	admin := api.Party("admin.")
	{
		// admin.mydomain.com
		admin.Get("/", func(c *iris.Context) {
			c.Writef("INDEX FROM admin.mydomain.com")
		})
		// admin.mydomain.com/hey
		admin.Get("/hey", func(c *iris.Context) {
			c.Writef("HEY FROM admin.mydomain.com/hey")
		})
		// admin.mydomain.com/hey2
		admin.Get("/hey2", func(c *iris.Context) {
			c.Writef("HEY SECOND FROM admin.mydomain.com/hey")
		})
	}

	// mydomain.com/
	api.Get("/", func(c *iris.Context) {
		c.Writef("INDEX FROM no-subdomain hey")
	})

	// mydomain.com/hey
	api.Get("/hey", func(c *iris.Context) {
		c.Writef("HEY FROM no-subdomain hey")
	})

	api.Listen("mydomain.com:80")
}
