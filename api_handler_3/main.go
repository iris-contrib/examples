// NOTE, as listed on the book, .API is not a method I like to discuss, I coded this only for newcomers, it's slow and not recomemnded to use
// Use .Get/.Post/.Put/.Delete/.Head/.Connect/.Options/.Any instead.
package main

import (
	"github.com/kataras/iris"
)

type UserAPI struct {
	*iris.Context
}

// GET /users
func (u UserAPI) Get() {
	u.Write("Get from /users")
	// u.JSON(iris.StatusOK,myDb.AllUsers())
}

// GET /users/:param1 which its value passed to the id argument
func (u UserAPI) GetBy(id string) { // id equals to u.Param("param1")
	u.Write("Get from /users/%s", id)
	// u.JSON(iris.StatusOK, myDb.GetUserById(id))

}

// POST /users
func (u UserAPI) Post() {
	name := u.FormValue("name")
	// myDb.InsertUser(...)
	println(string(name))
	println("Post from /users")
}

// PUT /users/:param1
func (u UserAPI) PutBy(id string) {
	name := u.FormValue("name") // you can still use the whole Context's features!
	// myDb.UpdateUser(...)
	println(string(name))
	println("Put from /users/" + id)
}

// DELETE /users/:param1
func (u UserAPI) DeleteBy(id string) {
	// myDb.DeleteUser(id)
	println("Delete from /" + id)
}

func main() {

	admin := iris.Party("/om/corp", func(ctx *iris.Context) { println("middleware here"); ctx.Next() })
	{

		admin.Get("/roles", func(ctx *iris.Context) { ctx.HTML(iris.StatusOK, "<h1>/roles</h1>") })
		admin.API("/users", UserAPI{})

	}
	iris.Listen(":8080")
}

// NOTE, as listed on the book, .API is not a method I like to discuss, I coded this only for newcomers, it's slow and not recomemnded to use
// Use .Get/.Post/.Put/.Delete/.Head/.Connect/.Options/.Any instead.
