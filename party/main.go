package main

import "github.com/kataras/iris"

func main() {
	// Let's party
	admin := iris.Party("/admin")
	{
		// add a silly middleware
		admin.UseFunc(func(c *iris.Context) {
			//your authentication logic here...
			println("from ", c.Path())
			authorized := true
			if authorized {
				c.Next()
			} else {
				c.Text(401, c.Path()+" is not authorized for you")
			}

		})
		admin.Get("/", func(c *iris.Context) {
			c.Writef("Hello World")
			c.Writef("from /admin/ or /admin if you pathcorrection on")
		})
		admin.Get("/dashboard", func(c *iris.Context) {
			c.Writef("/admin/dashboard")
		})
		admin.Delete("/delete/:userId", func(c *iris.Context) {
			c.Writef("admin/delete/%s", c.Param("userId"))
		})
	}

	beta := admin.Party("/beta")
	beta.Get("/hey", func(c *iris.Context) { c.Writef("hey from /admin/beta/hey") })

	//for subdomains goto: ../subdomains_1/main.go

	iris.Listen(":8080")

}
