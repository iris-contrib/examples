package routes

import "github.com/kataras/iris"

func Profile(ctx *iris.Context) {
	username := ctx.Param("username")
	ctx.MustRender("profile.html", iris.Map{"Username": username})
}
