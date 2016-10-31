package api

import (
	"gopkg.in/iris-contrib/examples.v4/AIO_examples/mongo/backend/db"
	"gopkg.in/iris-contrib/examples.v4/AIO_examples/mongo/backend/models"
	"gopkg.in/kataras/iris.v4"
)

type CustomAPI struct {
	*iris.Context
}

func (this CustomAPI) Serve(ctx *iris.Context) {

	Db := db.MgoDb{}
	Db.Init()

	result := []models.User{}

	if err := Db.C("auth").Find(nil).All(&result); err != nil {
		ctx.JSON(iris.StatusOK, models.Err("5"))
		return
	}

	// u.JSON(iris.StatusOK, models.ErrNo1())

	ctx.JSON(iris.StatusOK, &result)

	Db.Close()

}
