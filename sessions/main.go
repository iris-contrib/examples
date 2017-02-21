package main

import "gopkg.in/kataras/iris.v6"

func main() {

	/*  These are the optionally fields to configurate the sessions, using the station's Config field (iris.Config.Sessions)

	// Cookie string, the session's client cookie name, for example: "irissessionid"
	Cookie string
	// DecodeCookie set it to true to decode the cookie key with base64 URLEncoding
	// Defaults to false
	DecodeCookie bool
	// Expires the duration of which the cookie must expires (created_time.Add(Expires)).
	// Default infinitive/unlimited life duration(0)
	Expires time.Duration
	// GcDuration every how much duration(GcDuration) the memory should be clear for unused cookies (GcDuration)
	// for example: time.Duration(2)*time.Hour. it will check every 2 hours if cookie hasn't be used for 2 hours,
	// deletes it from backend memory until the user comes back, then the session continue to work as it was
	//
	// Default 2 hours
	GcDuration time.Duration
	// DisableSubdomainPersistence set it to true in order dissallow your iris subdomains to have access to the session cookie
	// defaults to false
	DisableSubdomainPersistence bool
	*/
	iris.Get("/", func(c *iris.Context) {
		c.Writef("You should navigate to the /set, /get, /delete, /clear,/destroy instead")
	})
	iris.Get("/set", func(c *iris.Context) {

		//set session values
		c.Session().Set("name", "iris")

		//test if setted here
		c.Writef("All ok session setted to: %s", c.Session().GetString("name"))
	})

	iris.Get("/get", func(c *iris.Context) {
		// get a specific key, as string, if no found returns just an empty string
		name := c.Session().GetString("name")

		c.Writef("The name on the /set was: %s", name)
	})

	iris.Get("/delete", func(c *iris.Context) {
		// delete a specific key
		c.Session().Delete("name")
	})

	iris.Get("/clear", func(c *iris.Context) {
		// removes all entries
		c.Session().Clear()
	})

	iris.Get("/destroy", func(c *iris.Context) {
		//destroy, removes the entire session and cookie
		c.SessionDestroy()
		c.Log("You have to refresh the page to completely remove the session (on browsers), so the name should NOT be empty NOW, is it?\n ame: %s\n\nAlso check your cookies in your browser's cookies, should be no field for localhost/127.0.0.1 (or what ever you use)", c.Session().GetString("name"))
		c.Writef("You have to refresh the page to completely remove the session (on browsers), so the name should NOT be empty NOW, is it?\nName: %s\n\nAlso check your cookies in your browser's cookies, should be no field for localhost/127.0.0.1 (or what ever you use)", c.Session().GetString("name"))
	})

	iris.Listen(":8080")
	//iris.ListenTLS("0.0.0.0:443", "mycert.cert", "mykey.key")
}
