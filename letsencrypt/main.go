// Package main provide one-line integration with letsencrypt.org
package main

import "github.com/kataras/iris"

func main() {
	iris.Get("/", func(ctx *iris.Context) {
		ctx.Write("Hello from SECURE SERVER!")
	})

	iris.Get("/test2", func(ctx *iris.Context) {
		ctx.Write("Welcome to secure server from /test2!")
	})

	// This will provide you automatic certification & key from letsencrypt.org's servers
	// it also starts a second 'http://' server which will redirect all 'http://$PATH' requests to 'https://$PATH'
	iris.ListenLETSENCRYPT("127.0.0.1:443")
}
