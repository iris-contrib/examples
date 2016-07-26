package routes

import "github.com/kataras/iris"

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
		User{"firstUsername"},
		User{"secondUsername"},
		User{"thirdUsername"},
	}

	page := Page{"All users", users}

	if err := ctx.Render("userlist.html", page); err != nil {
		iris.Logger.Println(err.Error())
		ctx.Panic()
	}
}
