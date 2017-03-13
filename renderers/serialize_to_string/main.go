package main

import (
	"bytes"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

// This example is useful when you want to send a rich e-mail
// instead of executing the renderer to the ctx.ResponseWriter itself
// we will create an io.Writer compatible, i.e: &bytes.Buffer{}
// and execute the template or content type renderer, i.e: markdown
// to our writer.
func main() {
	app := iris.New()
	// output startup banner and error logs on os.Stdout
	app.Adapt(iris.DevLogger())
	// set the router, you can choose gorillamux too
	app.Adapt(httprouter.New())

	markdownContents := `## Hello Markdown from Iris

This is an example of Markdown with Iris



Features
--------

All features of Sundown are supported, including:

*   **Compatibility**. The Markdown v1.0.3 test suite passes with
    the --tidy option.  Without --tidy, the differences are
    mostly in whitespace and entity escaping, where blackfriday is
    more consistent and cleaner.

*   **Common extensions**, including table support, fenced code
    blocks, autolinks, strikethroughs, non-strict emphasis, etc.

*   **Safety**. Blackfriday is paranoid when parsing, making it safe
    to feed untrusted user input without fear of bad things
    happening. The test suite stress tests this and there are no
    known inputs that make it crash.  If you find one, please let me
    know and send me the input that does it.

    NOTE: "safety" in this context means *runtime safety only*. In order to
    protect yourself against JavaScript injection in untrusted content, see
    [this example](https://github.com/russross/blackfriday#sanitize-untrusted-content).

*   **Fast processing**. It is fast enough to render on-demand in
    most web applications without having to cache the output.

*   **Thread safety**. You can run multiple parsers in different
    goroutines without ill effect. There is no dependence on global
    shared state.

*   **Minimal dependencies**. Blackfriday only depends on standard
    library packages in Go. The source code is pretty
    self-contained, so it is easy to add to any project, including
    Google App Engine projects.

*   **Standards compliant**. Output successfully validates using the
    W3C validation tool for HTML 4.01 and XHTML 1.0 Transitional.
	
	[this is a link](https://github.com/kataras/iris) `

	app.Get("/", func(ctx *iris.Context) {
		buff := &bytes.Buffer{}
		// use the app's render instead of context in order to write the result on another writer(bytes.Buffer).
		// remember: text/markdown is just a custom Iris type which converted to text/html
		app.Render(buff, "text/markdown",
			markdownContents,
			iris.RenderOptions{"charset": "8859-1"},
		)
		htmlContents := buff.String()
		app.Log(iris.DevMode, htmlContents)
		ctx.Writef("The Raw HTML is:\n%s", htmlContents)
	})

	app.Listen(":8080")
}
