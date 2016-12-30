package main

import (
	"crypto/md5"
	"fmt"
	"gopkg.in/kataras/iris.v5"
	"io"
	"os"
	"strconv"
	"time"
)

func main() {

	// Serve the form.html to the user
	iris.Get("/upload", func(ctx *iris.Context) {
		//these are optionaly you can just call RenderFile("form.html",{})
		//create the token
		now := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(now, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		//render the form with the token for any use you like
		ctx.Render("form.html", token)
	})

	// Handle the post request from the form.html to the server
	iris.Post("/upload", func(ctx *iris.Context) {

		// Get the file from the request
		info, err := ctx.FormFile("uploadfile")
		file, err := info.Open()
		defer file.Close()
		fname := info.Filename

		// Create a file with the same name
		// assuming that you have a folder named 'uploads'
		out, err := os.OpenFile("./uploads/"+fname, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer out.Close()

		io.Copy(out, file)

	})
	// set the max request body size of the server, used to be able to send big files to the server
	// 32MB max upload filesize)
	// By default request body size is 4MB.
	iris.Set(iris.OptionMaxRequestBodySize(32 << 20))

	iris.Listen(":8080")

}
