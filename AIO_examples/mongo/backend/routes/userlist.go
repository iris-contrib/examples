package routes

import "gopkg.in/kataras/iris.v6"

type (
	User struct {
		Username string
	}

	Page struct {
		Title string
		Users []User
	}
)

func UserList(ctx *iris.Context) {
	users := []User{
		{"firstUsername"},
		{"secondUsername"},
		{"thirdUsername"},
	}

	page := Page{"All users", users}

	if err := ctx.Render("userlist.html", page); err != nil {
		ctx.Log(iris.DevMode, err.Error())
		ctx.Panic()
	}
}
