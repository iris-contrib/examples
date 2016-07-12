package main

import (
	"github.com/iris-contrib/template/html"
	"github.com/kataras/iris"
)

type mypage struct {
	Title   string
	Message string
}

func main() {

	iris.UseEngine(html.New()).Directory("./templates", ".html")

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("mypage.html", mypage{"My Page title", "Hello world!"}, iris.Map{"gzip": true})
	})

	iris.Get("/hi_json", func(c *iris.Context) {
		c.JSON(200, iris.Map{
			"Name": "Iris",
			"Age":  2,
		})
	})

	iris.Listen(":8080")
}
