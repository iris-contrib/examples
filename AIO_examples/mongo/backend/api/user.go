package api

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/mgo.v2/bson"

	"github.com/iris-contrib/examples/AIO_examples/mongo/backend/db"
	"github.com/iris-contrib/examples/AIO_examples/mongo/backend/models"
)

func GetAllUsers(ctx *iris.Context) {
	Db := db.MgoDb{}
	Db.Init()

	result := []models.User{}
	if err := Db.C("people").Find(nil).All(&result); err != nil {
		ctx.JSON(iris.StatusOK, models.Err("1"))
		return
	}

	ctx.JSON(iris.StatusOK, &result)

	Db.Close()
}

func GetUserByID(ctx *iris.Context) {
	id := ctx.Param("userid")

	Db := db.MgoDb{}
	Db.Init()

	result := models.User{}

	if err := Db.C("people").Find(bson.M{"id": id}).One(&result); err != nil {
		ctx.JSON(iris.StatusOK, models.Err("1"))
		return
	}

	ctx.JSON(iris.StatusOK, &result)

	Db.Close()
}

func UpdateUser(ctx *iris.Context) {
	newUsername := string(ctx.FormValue("username"))
	// myDb.InsertUser(newUsername)
	println(newUsername + " has been inserted to database")
	ctx.JSON(iris.StatusOK, iris.Map{"response": true})

	// // Update
	// colQuerier := bson.M{"name": "Ale"}
	// change := bson.M{"$set": bson.M{"phone": "+86 99 8888 7777", "timestamp": time.Now()}}
	// err = c.Update(colQuerier, change)
	// if err != nil {
	// 	panic(err)
	// }
}

func InsertUser(ctx *iris.Context) {
	id := ctx.Param("userid") //?
	usr := models.User{}
	err := ctx.ReadForm(&usr)

	if err != nil {
		ctx.JSON(iris.StatusOK, models.Err("4"))
		panic(err.Error())
	}

	usr.Id = id // ?

	Db := db.MgoDb{}
	Db.Init()

	// Insert
	if err := Db.C("people").Insert(&usr); err != nil {
		ctx.JSON(iris.StatusOK, models.Err("5"))
	} else {
		ctx.JSON(iris.StatusOK, iris.Map{"response": true})
	}

	Db.Close()
}

func DeleteUser(ctx *iris.Context) {
	// if _, err := db.Col.RemoveAll(bson.M{"id": id}); err != nil {
	// 	ctx.JSON(iris.StatusOK, models.Err("1"))
	// 	return
	// }

	// ctx.JSON(iris.StatusOK, iris.Map{"response": true})
}

// // Get Params example code
// var _name = string(ctx.FormValue("name"))
// var _grender = string(ctx.FormValue("gender"))
// var _age = string(ctx.FormValue("age"))
// _newage, _ := strconv.Atoi(_age)
// var _id = string(ctx.FormValue("id"))

// usr := models.User{
// 	Name:   _name,
// 	Email:   "ivancduran@gmail.com",
// 	Gender: false,
// 	Birth:  "1989-08-21",
// 	Id:     _id,
// }
