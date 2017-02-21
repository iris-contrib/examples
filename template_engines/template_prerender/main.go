// Package main ;PreRenders can be registered to do something before the main context.MustRender/Render
package main

import (
	"time"

	"gopkg.in/kataras/iris.v6"
)

type MyTemplateData struct {
	Values iris.Map // just a shortcat of map[string]interface{}
	Name   string
}

func main() {
	// enable path escape in order to accept something like http://localhost:8080/profile/my%20name
	// as my name when ctx.Param
	// if it's false (as defaulted) and you want to escape a parameter name you can just use
	// ctx.ParamDecoded instead of ctx.Param
	iris.Config.EnablePathEscape = true

	// PreRender is typeof func(*iris.Context, string, interface{},...map[string]interface{}) bool
	// PreRenders helps developers to pass middleware between
	// the route Handler and a context.Render
	// all parameter receivers can be changed before passing it to the actual context's Render
	// so, you can change the filenameOrSource, the page binding, the options,
	// and even add cookies, session value or a flash message through ctx
	// the return value of a PreRender is a boolean, if returns false then the next PreRender will not be executed, keep note
	// that the actual context's Render will be called at any case.
	iris.UsePreRender(func(ctx *iris.Context, filenameOrSource string, binding interface{}, options ...map[string]interface{}) bool {

		if v, ok := binding.(*MyTemplateData); ok {
			ctx.VisitValues(func(key string, value interface{}) {
				if v.Values == nil {
					v.Values = make(iris.Map, ctx.ValuesLen())
				}
				v.Values[key] = value
			})
		}

		return true
	})

	// http://localhost:8080/profile/yourname
	iris.Get("/profile/:username", func(ctx *iris.Context) {
		name := ctx.Param("username")

		ctx.Set("myvalue", "a request lifetime value which passed to the template via our custom PreRender")

		data := &MyTemplateData{Name: name}
		ctx.MustRender("profile.html", data)
	})

	// localhost:8080/profile/my name/post/my post title
	iris.Get("/profile/:username/post/:title", func(ctx *iris.Context) {
		name := ctx.Param("username")
		// the ":title" also lives in context's request scoped values
		// so it will be visible by our MyTemplate.Values["title"] which will be rendered to our /templates/post.html

		ctx.Set("systemDate", time.Now().Format(time.RFC822))

		data := &MyTemplateData{Name: name}
		ctx.MustRender("post.html", data)
	})

	iris.Listen(":8080")
}
