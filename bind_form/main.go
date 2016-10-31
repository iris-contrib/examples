// package main contains an example on how to use the ReadForm, but with the same way you can do the ReadJSON & ReadJSON
package main

import (
	"fmt"

	"gopkg.in/kataras/iris.v4"
)

type Visitor struct {
	Username string
	Mail     string
	Data     []string `form:"mydata"`
}

func main() {
	// no need to set a template engine here, because the default is the html with ./templates as directory and .html as files extension

	iris.Get("/", func(ctx *iris.Context) {
		if err := ctx.Render("form.html", nil); err != nil {
			iris.Logger.Printf(err.Error())
		}
	})

	iris.Post("/form_action", func(ctx *iris.Context) {
		visitor := Visitor{}
		err := ctx.ReadForm(&visitor)
		if err != nil {
			fmt.Println("Error when reading form: " + err.Error())
		}
		fmt.Printf("\n Visitor: %#v", visitor)
		ctx.Write("%#v", visitor)
	})

	iris.Listen(":8080")
}
