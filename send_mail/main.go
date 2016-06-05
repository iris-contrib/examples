package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/config"
)

func main() {
	// change these to your settings

	iris.Config().Mail = config.Mail{
		Host:     "smtp.mailgun.org",
		Username: "postmaster@sandbox661c307650f04e909150b37c0f3b2f09.mailgun.org",
		Password: "38304272b8ee5c176d5961dc155b2417",
		Port:     587,
	}
	// change these to your e-mail to check if that works

	var to = []string{"kataras2006@hotmail.com", "social@ideopod.com"}

	iris.Get("/send", func(ctx *iris.Context) {
		content := `<h1>Hello From Iris web framework</h1> <br/><br/> <span style="color:blue"> This is the rich message body </span>`

		err := iris.Mail().Send(to, "iris e-mail just t3st", content)

		if err != nil {
			ctx.WriteHTML(200, "<b> Problem while sending the e-mail: "+err.Error())
		} else {
			ctx.WriteHTML(200, "<h1> SUCCESS </h1>")
		}
	})

	// send a body by template
	iris.Get("/send/template", func(ctx *iris.Context) {
		content, _ := ctx.RenderString("body.html", iris.Map{
			"Message": " his is the rich message body sent by a template!!",
			"Footer":  "The footer of this e-mail!",
		})

		err := iris.Mail().Send(to, "iris e-mail just t3st", content)

		if err != nil {
			ctx.WriteHTML(200, "<b> Problem while sending the e-mail: "+err.Error())
		} else {
			ctx.WriteHTML(200, "<h1> SUCCESS </h1>")
		}
	})
	iris.Listen(":8080")
}
