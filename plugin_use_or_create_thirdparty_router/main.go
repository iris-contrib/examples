package main

import (
	"net/http"
	"sort"

	"github.com/gorilla/mux"
	"github.com/kataras/iris"
)

// NOTE: Iris has this as plugin on https://github.com/iris-contrib/plugin/gorillamux
// The time this example writing is the time I'm writing the plugin also :)
// it's a good oportunity to learn how I do these staff, and most importantly what you, as community
// can do with an extensible framework like Iris. I changed zero lines on kataras/iris but look what we can do:

/*
Here you will see the Iris; power of customizable itself.
I'm sure that, these you will see below you never saw it before (on other framework).
Again Iris is first by distance, we make them to follow us.
*/

// Register your routes as you do with Iris, with iris' Handlers and middleware
func main() {
	iris.Plugins.Add(&GorillaMux{})

	// CUSTOM HTTP ERRORS ARE SUPPORTED
	// NOTE: Gorilla mux allows customization only on StatusNotFound(404)
	// Iris allows for everything, so you can register any other custom http error
	// but you have to call it manually from ctx.EmitError(status_code) // 500 for example
	// this will work because it's StatusNotFound:
	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.HTML(iris.StatusNotFound, "<h1> CUSTOM NOT FOUND ERROR PAGE </h1>")
	})

	// GLOBAL/PARTY MIDDLEWARE IS SUPPORTED
	iris.UseFunc(func(ctx *iris.Context) {
		println("Request: " + ctx.Path())
		ctx.Next()
	})

	// http://mydomain.com
	iris.Get("/", func(ctx *iris.Context) {
		ctx.Writef("Hello from index")
	})

	// GORILLA MUX PARAMETERS(regexp) ARE SUPPORTED
	// http://mydomain.com/api/users/42
	iris.Get("/api/users/{userid:[0-9]+}", func(ctx *iris.Context) {
		ctx.Writef("User with id: %s", ctx.Param("userid"))
	})

	// PER-ROUTE MIDDLEWARE ARE SUPPORTED
	// http://mydomain.com/other
	iris.Get("/other", func(ctx *iris.Context) {
		ctx.Writef("/other 1 middleware \n")
		ctx.Next()
	}, func(ctx *iris.Context) {
		ctx.HTML(iris.StatusOK, "<b>Hello from /other</b>")
	})

	// SUBDOMAINS ARE SUPPORTED
	// http://admin.mydomain.com
	iris.Party("admin.").Get("/", func(ctx *iris.Context) {
		ctx.Writef("Hello from admin. subdomain!")
	})

	// WILDCARD SUBDOMAINS ARE SUPPORTED
	// http://api.mydomain.com/hi
	// http://admin.mydomain.com/hi
	// http://x.mydomain.com/hi
	// [depends on your host configuration,
	// you will see an example(win) outside of this folder].
	iris.Party("*.").Get("/hi", func(ctx *iris.Context) {
		ctx.Writef("Hello from wildcard subdomain: %s", ctx.Subdomain())
	})

	// DOMAIN NAMING IS SUPPORTED
	iris.Listen("mydomain.com")
	// iris.Listen(":80")
}

// order matters in gorilla mux
type bySubdomain []iris.Route

// Sorting happens when the mux's request handler initialized
func (s bySubdomain) Len() int {
	return len(s)
}
func (s bySubdomain) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s bySubdomain) Less(i, j int) bool {
	return len(s[i].Subdomain()) > len(s[j].Subdomain())
}

// GorillaMux is the plugin
type GorillaMux struct{}

// At PreBuild state because on PreLookup the iris.UseGlobal/UseGlobalFunc may not be catched.
func (g GorillaMux) PreBuild(s *iris.Framework) {
	router := mux.NewRouter()
	routes := s.Lookups()
	// gorilla mux order matters, so order them by subdomain before looping
	sort.Sort(bySubdomain(routes))

	for i := range routes {
		route := routes[i]
		registerRoute(route, router, s)
	}
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := s.AcquireCtx(w, r)
		// to catch custom 404 not found http errors may registered by user
		ctx.EmitError(iris.StatusNotFound)
		s.ReleaseCtx(ctx)
	})
	s.Router = router
}

// so easy:
func registerRoute(route iris.Route, gorillaRouter *mux.Router, s *iris.Framework) {

	if route.IsOnline() {
		handler := func(w http.ResponseWriter, r *http.Request) {

			ctx := s.AcquireCtx(w, r)

			if params := mux.Vars(r); params != nil && len(params) > 0 {
				// set them with ctx.Set in order to be accesible by ctx.Param in the user's handler
				for k, v := range params {
					ctx.Set(k, v)
				}
			}
			// including the iris.Use/UseFunc and the route's middleware,
			// main handler and any done handlers.
			ctx.Middleware = route.Middleware()
			ctx.Do()

			s.ReleaseCtx(ctx)
		}
		// remember, we get a new iris.Route foreach of the HTTP Methods, so this should be work
		gorillaRoute := gorillaRouter.HandleFunc(route.Path(), handler).Methods(route.Method()).Name(route.Name())
		subdomain := route.Subdomain()
		if subdomain != "" {
			if subdomain == "*." {
				// it's an iris wildcard subdomain
				// so register it as wildcard on gorilla mux too (hopefuly, it supports these things)
				subdomain = "{subdomain}."
			} else {
				// it's a static subdomain (which contains the dot)
			}
			// host = subdomain  + listening host
			gorillaRoute.Host(subdomain + s.Config.VHost)
		}
	}

	// AUTHOR NOTE:
	// the only feature I can think right now that is missing is the
	// iris offline routing (which after a little research can be done with gorilla's mux BuildOnly
	// but I don't know how can I activate that again, I probably need to get its handler and execute dynamically
	// this will slow down the things, so I don't do it, until I think a better way,
	// offline routing is a 2-day feature so I suppose no many people know about that yet,
	// and if they want offline routing they should not change to a custom router (yet). so we are ok.
}

/* HOSTS FILE LINES TO RUN THIS EXAMPLE:

127.0.0.1		mydomain.com
127.0.0.1		admin.mydomain.com
127.0.0.1		api.mydomain.com
127.0.0.1		x.mydomain.com

*/
