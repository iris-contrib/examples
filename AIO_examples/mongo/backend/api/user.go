package api

import (
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2/bson"

	"github.com/iris-contrib/examples/AIO_examples/mongo/backend/db"
	"github.com/iris-contrib/examples/AIO_examples/mongo/backend/models"
)

type UserAPI struct {
	*iris.Context
}

// GET /users
func (this UserAPI) Get() {

	Db := db.MgoDb{}
	Db.Init()

	result := []models.User{}
	if err := Db.C("people").Find(nil).All(&result); err != nil {
		this.JSON(iris.StatusOK, models.Err("1"))
		return
	} else {
		this.JSON(iris.StatusOK, &result)
	}

	Db.Close()

}

// GET /users/:param1
func (this UserAPI) GetBy(id string) {

	Db := db.MgoDb{}
	Db.Init()

	result := models.User{}

	if err := Db.C("people").Find(bson.M{"id": id}).One(&result); err != nil {
		this.JSON(iris.StatusOK, models.Err("1"))
		return
	} else {
		this.JSON(iris.StatusOK, &result)
	}

	Db.Close()

}

// PUT /users
func (this UserAPI) Put() {

	newUsername := string(this.FormValue("username"))
	// myDb.InsertUser(newUsername)
	println(newUsername + " has been inserted to database")
	this.JSON(iris.StatusOK, iris.Map{"response": true})

	// // Update
	// colQuerier := bson.M{"name": "Ale"}
	// change := bson.M{"$set": bson.M{"phone": "+86 99 8888 7777", "timestamp": time.Now()}}
	// err = c.Update(colQuerier, change)
	// if err != nil {
	// 	panic(err)
	// }

}

// POST /users/:param1
func (this UserAPI) PostBy(id string) {

	usr := models.User{}
	err := this.ReadForm(&usr)

	if err != nil {
		this.JSON(iris.StatusOK, models.Err("4"))
		panic(err.Error())
	}

	usr.Id = id

	Db := db.MgoDb{}
	Db.Init()

	// Insert
	if err := Db.C("people").Insert(&usr); err != nil {
		this.JSON(iris.StatusOK, models.Err("5"))
	} else {
		this.JSON(iris.StatusOK, iris.Map{"response": true})
	}

	Db.Close()

}

// DELETE /users/:param1
func (this UserAPI) DeleteBy(id string) {

	// if _, err := db.Col.RemoveAll(bson.M{"id": id}); err != nil {
	// 	this.JSON(iris.StatusOK, models.Err("1"))
	// 	return
	// }

	// this.JSON(iris.StatusOK, iris.Map{"response": true})

}

// // Get Params example code
// var _name = string(this.FormValue("name"))
// var _grender = string(this.FormValue("gender"))
// var _age = string(this.FormValue("age"))
// _newage, _ := strconv.Atoi(_age)
// var _id = string(this.FormValue("id"))

// usr := models.User{
// 	Name:   _name,
// 	Email:   "ivancduran@gmail.com",
// 	Gender: false,
// 	Birth:  "1989-08-21",
// 	Id:     _id,
// }
