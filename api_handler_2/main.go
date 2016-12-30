// Package main same as api_handler_1 but with a common middleware
// NOTE, as listed on the book, .API is not a method I like to discuss, I coded this only for newcomers, it's slow and not recomemnded to use
// Use .Get/.Post/.Put/.Delete/.Head/.Connect/.Options/.Any instead.
package main

import (
	"gopkg.in/kataras/iris.v5"
)

type UserAPI struct {
	*iris.Context
}

// GET /users
func (u UserAPI) Get() {
	u.Write("Get from /users")
	// u.JSON(iris.StatusOK,myDb.AllUsers())
}

// GET /:param1 which its value passed to the id argument
func (u UserAPI) GetBy(id string) { // id equals to u.Param("param1")
	u.Write("Get from /users/%s", id)
	// u.JSON(iris.StatusOK, myDb.GetUserById(id))

}

// PUT /users
func (u UserAPI) Put() {
	name := u.FormValue("name")
	// myDb.InsertUser(...)
	println(string(name))
	println("Put from /users")
}

// POST /users/:param1
func (u UserAPI) PostBy(id string) {

	name := u.PostValue("name") // you can still use the whole Context's features!
	// myDb.UpdateUser(...)
	println(string(name))
	println("Post from /users/" + id)
}

// DELETE /users/:param1
func (u UserAPI) DeleteBy(id string) {
	// myDb.DeleteUser(id)
	println("Delete from /" + id)
}

func main() {
	iris.API("/users", UserAPI{}, myUsersMiddleware1, myUsersMiddleware2)
	iris.Listen(":8080")
}

func myUsersMiddleware1(ctx *iris.Context) {
	println("From users middleware 1 ")
	ctx.Next()
}
func myUsersMiddleware2(ctx *iris.Context) {
	println("From users middleware 2 ")
	ctx.Next()
}

// NOTE, as listed on the book, .API is not a method I like to discuss, I coded this only for newcomers, it's slow and not recomemnded to use
// Use .Get/.Post/.Put/.Delete/.Head/.Connect/.Options/.Any instead.
