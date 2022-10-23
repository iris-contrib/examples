package main

import (
	"github.com/kataras/iris/v12"
)

// Follow these steps first:
// $ go install github.com/go-bindata/go-bindata/v3/go-bindata@latest
// $ go-bindata -prefix "assets" -fs ./assets/...
// $ go run .
// "physical" files are not used, you can delete the "assets" folder and run the example.
//
// See `file-server/embedding-gzipped-files-into-app-bindata` and
// 'file-server/embedding-files-into-app` examples as well.
func newApp() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.HandleDir("/static", AssetFile())

	/*
		Or if you need to cache them inside the memory (requires the assets folder
		to be located near the executable program):
		app.HandleDir("/static", iris.Dir("./assets"), iris.DirOptions{
			IndexName: "index.html",
			Cache: iris.DirCacheOptions{
				Enable:          true,
				Encodings:       []string{"gzip"},
				CompressIgnore:  iris.MatchImagesAssets,
				CompressMinSize: 30 * iris.B,
			},
		})
	*/
	return app
}

func main() {
	app := newApp()

	// http://localhost:8080/static/css/main.css
	// http://localhost:8080/static/js/main.js
	// http://localhost:8080/static/favicon.ico
	app.Listen(":8080")
}
