package api

import (
	"github.com/iris-contrib/examples/AIO_examples/mongo/backend/db"
	"github.com/iris-contrib/examples/AIO_examples/mongo/backend/models"
	"gopkg.in/kataras/iris.v6"
)

type CustomAPI struct {
	*iris.Context
}

func (CustomAPI) Serve(ctx *iris.Context) {

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
