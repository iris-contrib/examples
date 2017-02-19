package routes

import (
	"gopkg.in/kataras/iris.v6"
)

type index struct {
	Title   string
	Message string
}

// Index index is the page for GET: / route
func Index() *index {
	return &index{
		Title:   "iris sample - index",
		Message: "This is just a sample index, it's empty because Iris doesnt wants influences!",
	}
}

func (i *index) Serve(ctx *iris.Context) {
	if err := ctx.Render("index.html", i); err != nil {
		// ctx.EmitError(iris.StatusInternalServerError) =>
		ctx.Panic()
	}
}
