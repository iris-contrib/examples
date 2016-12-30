package routes

import "gopkg.in/kataras/iris.v5"

func Profile(ctx *iris.Context) {
	username := ctx.Param("username")
	ctx.MustRender("profile.html", iris.Map{"Username": username})
}
