package main

/* this example doesn't contains any line for  customize the redis you can take a look at the iris book.
   but this does the work: import "github.com/kataras/iris/sessions/providers/redis"; redis.Config.Addr = "localhost:1222"; 	iris.Config().Session.Provider = "redis"
*/
import (
	"github.com/kataras/iris"
)

func main() {
	iris.Config.Sessions.Provider = "redis"

	iris.Get("/set", func(c *iris.Context) {

		//set session values
		c.Session().Set("name", "iris")

		//test if setted here
		c.Write("All ok session setted to: %s", c.Session().GetString("name"))
	})

	iris.Get("/get", func(c *iris.Context) {
		name := c.Session().GetString("name")

		c.Write("The name on the /set was: %s", name)
	})

	iris.Get("/delete", func(c *iris.Context) {
		//get the session for this context

		c.Session().Delete("name")

	})

	iris.Get("/clear", func(c *iris.Context) {

		// removes all entries
		c.Session().Clear()
	})

	iris.Get("/destroy", func(c *iris.Context) {
		//destroy, removes the entire session and cookie
		c.SessionDestroy()
	})

	iris.Listen(":8080")
}
