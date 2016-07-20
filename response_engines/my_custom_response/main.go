// You can create a custom response engine using a func or an interface which implements the
// iris.ResponseEngine which contains a simple function: Response(val interface{}, options ...map[string]interface{}) ([]byte, error)

// A custom engine can be used to register a totally new content writer for a known ContentType or for a custom ContentType

// Let's do a 'trick' here, which works for all other response engines, custom or not:

// say for example, that you want a static'footer/suffix' on your content, without the need to create & register a middleware for that, per route or globally
// you want to be even more organised.
//
// IF a response engine has the same key and the same content type then the contents are appended and the final result will be rendered to the client.

// Enough with my 'bad' english, let's code something small:

package main

import (
	"github.com/iris-contrib/response/text"
	"github.com/kataras/iris"
)

// Let's do this with ` text/plain` content type, because you can see its results easly, the first engine will use this "text/plain" as key,
// the second & third will use the same, as firsts, key, which is the ContentType also.
func main() {
	// we are registering the default text/plain,  and after we will register the 'appender' only
	// we have to register the default because we will add more response engines with the same content,
	// iris will not register this by-default if other response engine with the corresponding ContentType already exists
	iris.UseResponse(text.New(), text.ContentType) // it's the key which happens to be a valid content-type also, "text/plain" so this will be used as the ContentType header

	// register a response engine: iris.ResponseEngine
	iris.UseResponse(&CustomTextEngine{}, text.ContentType)
	// register a response engine with func
	iris.UseResponse(iris.ResponseEngineFunc(func(val interface{}, options ...map[string]interface{}) ([]byte, error) {
		return []byte("\nThis is the static SECOND AND LAST suffix!"), nil
	}), text.ContentType)

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Text(iris.StatusOK, "Hello!") // or ctx.Render(text.ContentType," Hello!")
	})

	iris.Listen(":8080")
}

// This is the way you create one with raw iris.ResponseEngine implementation:

// CustomTextEngine the response engine which appends a simple string on the default's text engine
type CustomTextEngine struct{}

// Implement the iris.ResponseEngine
func (e *CustomTextEngine) Response(val interface{}, options ...map[string]interface{}) ([]byte, error) {
	// we don't need the val, because we want only to append, so what we should do?
	// just return the []byte we want to be appended after the first registered text/plain engine

	return []byte("\nThis is the static FIRST suffix!"), nil
}
