package routes

import (
	"github.com/iris-contrib/examples/AIO_examples/basic/backend/api/user"
	"gopkg.in/kataras/iris.v6"
)

type (
	Page struct {
		Title string
		Users []userAPI.User
	}
)

func UserList(ctx *iris.Context) {
	page := Page{"All users", userAPI.MyUsers}

	if err := ctx.Render("userlist.html", page); err != nil {
		ctx.Log(iris.DevMode, err.Error())
		ctx.Panic()
	}
}
