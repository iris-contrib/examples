package userAPI

import "gopkg.in/kataras/iris.v6"

// User is just a test User representation
type User struct {
	ID        int
	Firstname string
	Lastname  string
}

// MyUsers is our fake users database/store
var MyUsers = []User{
	{
		ID:        1,
		Firstname: "Gerasimos",
		Lastname:  "Maropoulos",
	},
	{
		ID:        2,
		Firstname: "Hillary",
		Lastname:  "Clinton",
	},
	{
		ID:        3,
		Firstname: "Donald",
		Lastname:  "Trump",
	},
}

// GetAll returns all users as "application/json"
func GetAll(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, MyUsers)
}

// GetByID returns a user by ID as "application/json"
func GetByID(ctx *iris.Context) {
	userID, err := ctx.ParamInt("id")
	if err != nil {
		ctx.EmitError(iris.StatusBadRequest)
		return
	}
	for _, u := range MyUsers {
		if u.ID == userID {
			ctx.JSON(iris.StatusOK, u)
			return
		}
	}
	ctx.EmitError(iris.StatusNotFound)
}

// Insert adds a new user to 'myUsers'
func Insert(ctx *iris.Context) {
	firstname := ctx.FormValue("firstname")
	lastname := ctx.FormValue("lastname")
	if firstname == "" || lastname == "" {
		ctx.EmitError(iris.StatusBadRequest)
		return
	}

	lastUserID := MyUsers[len(MyUsers)-1].ID
	userID := lastUserID + 1
	MyUsers = append(MyUsers, User{
		ID:        userID,
		Firstname: firstname,
		Lastname:  lastname,
	})
	ctx.JSON(iris.StatusOK, iris.Map{"Status": "Success"})
}

// Update updates a user by path param ID
// this may not work but you get the point
func Update(ctx *iris.Context) {
	userID, err := ctx.ParamInt("id")
	if err != nil {
		ctx.EmitError(iris.StatusBadRequest)
	}

	newFirstname := ctx.FormValue("firstname")
	newLastname := ctx.FormValue("lastname")

	if newFirstname == "" && newLastname == "" {
		ctx.EmitError(iris.StatusBadRequest)
		return
	}

	for _, u := range MyUsers {
		if u.ID == userID {
			if newFirstname != "" {
				u.Firstname = newFirstname
			}

			if newLastname != "" {
				u.Lastname = newLastname
			}

			ctx.JSON(iris.StatusOK, iris.Map{"Status": "Success"})
			return
		}
	}

	ctx.EmitError(iris.StatusNotFound)

}

// DeleteByID removes a user from 'myUsers' by ID
func DeleteByID(ctx *iris.Context) {
	userID, err := ctx.ParamInt("id")
	if err != nil {
		ctx.EmitError(iris.StatusBadRequest)
	}
	for i, u := range MyUsers {
		if u.ID == userID {
			MyUsers = append(MyUsers[:0], MyUsers[i:]...)
			ctx.JSON(iris.StatusOK, iris.Map{"Status": "Success"})
			return
		}
	}
	ctx.EmitError(iris.StatusNotFound)
}
