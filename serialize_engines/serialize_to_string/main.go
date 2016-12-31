package main

import "github.com/kataras/iris"

func main() {

	// SerializeToString gives you the result of the serialize engine(serializer)'s work, it doesn't renders to the client but you can use
	// this function to collect the end result and send it via e-mail to the user, or anything you can imagine.

	// Note that: iris.SerializeToString is called outside of the context, using your iris $instance (iris. is the default)

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

	iris.Get("/", func(ctx *iris.Context) {
		// let's see
		// convert markdown string to html and print it to the logger
		// THIS WORKS WITH ALL serialize engine(serializer)S, but I am not doing the same example for all engines again :) (the same you can do with templates using the iris.TemplateString)
		htmlContents := iris.SerializeToString("text/markdown", markdownContents, iris.RenderOptions{"charset": "8859-1"}) // default is the iris.Config.Charset, which is UTF-8

		ctx.Log(htmlContents)
		ctx.Writef("The Raw HTML is:\n%s", htmlContents)
	})

	iris.Listen(":8080")
}
