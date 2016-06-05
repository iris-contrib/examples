package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/i18n"
)

func main() {

	iris.Use(i18n.I18nHandler(i18n.Options{Default: "en-US",
		Languages: map[string]string{
			"en-US": "./locales/locale_en-US.ini",
			"el-GR": "./locales/locale_el-GR.ini",
			"zh-CN": "./locales/locale_zh-CN.ini"}}))
	// or iris.UseFunc(i18n.I18n(....))
	// or iris.Get("/",i18n.I18n(....), func (ctx *iris.Context){})

	iris.Get("/", func(ctx *iris.Context) {
		hi := ctx.GetFmt("translate")("hi", "maki") // hi is the key, 'maki' is the %s, the second parameter is optional
		language := ctx.Get("language")             // language is the language key, example 'en-US'

		ctx.Write("From the language %s translated output: %s", language, hi)
	})

	iris.Listen(":8080")

}
