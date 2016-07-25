package main

import (
	"github.com/iris-contrib/response/markdown"
	"github.com/kataras/iris"
)

func main() {
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

	//first example
	// this is one of the reasons you need to import a default engine,(template engine or response engine)
	/*
		type Config struct {
			MarkdownSanitize bool
		}
	*/
	iris.UseResponse(markdown.New(), markdown.ContentType)
	// you can use anything as the second parameter,
	// the markdown.ContentType is the string "text/markdown",
	// the context.Markdown renders with this engine's key.

	handlerWithRender := func(ctx *iris.Context) {
		// you can also change the charset for a specific render action with RenderOptions
		ctx.Render("text/markdown", markdownContents, iris.RenderOptions{"charset": "8859-1"})
	}

	//second example,
	// but we also want a different renderer, but again "text/markdown" as 'content type' (this is converted to text/html behind the scenes), with MarkdownSanitize option setted to true:
	iris.UseResponse(markdown.New(markdown.Config{MarkdownSanitize: true}), "markdown2")("text/markdown")
	// yes the UseResponse returns a function which you can map the content type if it's not declared on the key
	handlerMarkdown2 := func(ctx *iris.Context) {
		ctx.Render("markdown2", markdownContents, iris.RenderOptions{"gzip": true})
	}

	iris.Get("/", handlerWithRender)

	iris.Get("/markdown2", handlerMarkdown2)

	iris.Listen(":8080")
}
