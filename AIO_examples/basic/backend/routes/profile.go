package routes

import "gopkg.in/kataras/iris.v4"

func Profile(ctx *iris.Context) {
	username := ctx.Param("username")
	ctx.MustRender("profile.html", iris.Map{"Username": username})
}
