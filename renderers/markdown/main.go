package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

type myjson struct {
	Name string `json:"name"`
}

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
		ctx.Markdown(iris.StatusOK, markdownContents)
	})

	app.Get("/alternative_1", func(ctx *iris.Context) {
		htmlContents := ctx.MarkdownString(markdownContents)
		ctx.HTML(iris.StatusOK, htmlContents)
	})

	// text/markdown is just the key which the markdown serialize engine(serializer) and ctx.Markdown communicate,
	// it's real content type is text/html
	app.Get("/alternative_2", func(ctx *iris.Context) {
		ctx.Render("text/markdown", markdownContents)
	})

	app.Get("/alternative_3", func(ctx *iris.Context) {
		ctx.RenderWithStatus(iris.StatusOK, "text/markdown", markdownContents)
	})

	app.Get("/alternative_4", func(ctx *iris.Context) {
		ctx.Render("text/markdown", markdownContents, iris.RenderOptions{"charset": "UTF-8"}) // UTF-8 is the default.
	})

	app.Get("/alternative_5", func(ctx *iris.Context) {
		// logs if any error and sends http status '500 internal server error' to the client
		ctx.MustRender("text/markdown", markdownContents, iris.RenderOptions{"charset": "UTF-8"}) // UTF-8 is the default.
	})

	app.Listen(":8080")
}
