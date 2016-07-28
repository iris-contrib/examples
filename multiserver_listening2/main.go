package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/config"
)

/*
	Previous example: multiserver_listening
*/
func main() {
	host := "127.0.0.1:443"
	iris.Get("/", func(ctx *iris.Context) {
		ctx.Write("Hello from the SECURE server")
	})

	iris.Get("/mypath", func(ctx *iris.Context) {
		ctx.Write("Hello from the SECURE server on path /mypath")
	})

	// start a secondary server (HTTP) on port 80, this is a non-blocking func
	// redirects all http to the main server which is tls/ssl on port :443

	iris.AddServer(config.Server{ListeningAddr: ":80", RedirectTo: "https://" + host})

	// start the MAIN server (HTTPS) on port 443, this is a blocking func
	iris.ListenTLS(host, "mycert.cert", "mykey.key")

	// now if you navigate to http://127.0.0.1/mypath it will send you back to https://127.0.0.1:443/mypath (https://127.0.0.1/mypath)
	//
	// go to the letsencrypt example to view how you can integrade your server to get automatic certification and key from the letsencrypt.org 's servers.
}
