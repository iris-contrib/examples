// You can create a custom serialize engine(serializer) using a func or an interface which implements the
// serializer.Serializer which contains a simple function: Serialize(val interface{}, options ...map[string]interface{}) ([]byte, error)

// A custom engine can be used to register a totally new content writer for a known ContentType or for a custom ContentType

// Let's do a 'trick' here, which works for all other serialize engine(serializer)s, custom or not:

// say for example, that you want a static'footer/suffix' on your content, without the need to create & register a middleware for that, per route or globally
// you want to be even more organised.
//
// IF a serialize engine(serializer) has the same key and the same content type then the contents are appended and the final result will be rendered to the client.

// Enough with my 'bad' english, let's code something small:

package main

import (
	"github.com/kataras/go-serializer"
	"github.com/kataras/go-serializer/text"
	"github.com/kataras/iris"
)

// Let's do this with ` text/plain` content type, because you can see its results easly, the first engine will use this "text/plain" as key,
// the second & third will use the same, as firsts, key, which is the ContentType also.
func main() {
	// we are registering the default text/plain,  and after we will register the 'appender' only
	// we have to register the default because we will add more serialize engine(serializer)s with the same content,
	// iris will not register this by-default if other serialize engine(serializer) with the corresponding ContentType already exists
	iris.UseSerializer(text.ContentType, text.New())

	// register a serialize engine(serializer) serializer.Serializer
	iris.UseSerializer(text.ContentType, &CustomTextEngine{})
	// register a serialize engine(serializer) with func
	iris.UseSerializer(text.ContentType, serializer.SerializeFunc(func(val interface{}, options ...map[string]interface{}) ([]byte, error) {
		return []byte("\nThis is the static SECOND AND LAST suffix!"), nil
	}))

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Text(iris.StatusOK, "Hello!") // or ctx.Render(text.ContentType," Hello!")
	})

	iris.Listen(":8080")
}

// This is the way you create one with raw serialiser.Serializer implementation:

// CustomTextEngine the serialize engine(serializer) which appends a simple string on the default's text engine
type CustomTextEngine struct{}

// Implement the serializer.Serializer
func (e *CustomTextEngine) Serialize(val interface{}, options ...map[string]interface{}) ([]byte, error) {
	// we don't need the val, because we want only to append, so what we should do?
	// just return the []byte we want to be appended after the first registered text/plain engine
	return []byte("\nThis is the static FIRST suffix!"), nil
}
