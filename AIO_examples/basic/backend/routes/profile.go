package routes

import "gopkg.in/kataras/iris.v6"

func Profile(ctx *iris.Context) {
	username := ctx.Param("username")
	ctx.MustRender("profile.html", iris.Map{"Username": username})
}
