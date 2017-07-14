package main

import (
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"

	"github.com/didip/tollbooth"
	"github.com/iris-contrib/middleware/tollboothic"
)

func main() {
	app := iris.New()

	// Create a limiter struct.
	limiter := tollbooth.NewLimiter(1, time.Second)

	app.Get("/", tollboothic.LimitHandler(limiter), func(ctx context.Context) {
		ctx.HTML("<b>Hello, world!</b>")
	})

	app.Run(iris.Addr(":8080"))
}

// Read more at: https://github.com/didip/tollbooth
