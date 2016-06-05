package main

import (
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/config"
	"github.com/kataras/iris/sessions"

	_ "github.com/kataras/iris/sessions/providers/redis" // add a store it is auto-registers itself
	/* this example doesn't contains any line for  customize the redis you can take a look at the iris book.
	   but... import "github.com/kataras/iris/sessions/providers/redis"; redis.Config.Addr = "localhost:1222"; 	sess  = sessions.new("redis"...)
	*/)

var sess *sessions.Manager

func init() {
	sessConfig := config.Sessions{
		Provider:   "redis", // if you set it to ""  means that sessions are disabled.
		Cookie:     "yoursessionCOOKIEID",
		Expires:    config.CookieExpireNever,
		GcDuration: time.Duration(2) * time.Hour,
	}
	sess = sessions.New(sessConfig)

}

func main() {

	iris.Get("/set", func(c *iris.Context) {
		//get the session for this context
		session := sess.Start(c)

		//set session values
		session.Set("name", "iris")

		//test if setted here
		c.Write("All ok session setted to: %s", session.GetString("name"))
	})

	iris.Get("/get", func(c *iris.Context) {
		//get the session for this context
		session := sess.Start(c)

		//get the session value
		name := session.GetString("name")

		c.Write("The name on the /set was: %s", name)
	})

	iris.Get("/delete", func(c *iris.Context) {
		//get the session for this context
		session := sess.Start(c)

		session.Delete("name")

	})

	iris.Get("/clear", func(c *iris.Context) {
		//get the session for this context
		session := sess.Start(c)
		// removes all entries
		session.Clear()
	})

	iris.Get("/destroy", func(c *iris.Context) {
		//destroy, removes the entire session and cookie
		sess.Destroy(c)
	})

	println("Server is listening at :8080")
	iris.Listen("8080")

}

// session.GetAll() returns all values a map[interface{}]interface{}
// session.VisitAll(func(key interface{}, value interface{}) { /* loops for each entry */})
