package main

import (
	"html/template"

	"github.com/kataras/go-template/pug"
	"gopkg.in/kataras/iris.v6"
)

type Person struct {
	Name   string
	Age    int
	Emails []string
	Jobs   []*Job
}

type Job struct {
	Employer string
	Role     string
}

func main() {
	// set the configuration for this template engine  (all template engines has its configuration)
	cfg := pug.DefaultConfig()
	cfg.Funcs["bold"] = func(content string) (template.HTML, error) {
		return template.HTML("<b>" + content + "</b>"), nil
	}

	iris.UseTemplate(pug.New(cfg)).
		Directory("./templates", ".jade")

	iris.Get("/", func(ctx *iris.Context) {

		job1 := Job{Employer: "Super Employer", Role: "Team leader"}
		job2 := Job{Employer: "Fast Employer", Role: "Project managment"}

		person := Person{
			Name:   "name1",
			Age:    50,
			Emails: []string{"email1@something.gr", "email2.anything@gmail.com"},
			Jobs:   []*Job{&job1, &job2},
		}
		ctx.MustRender("page.jade", person)

	})

	iris.Listen(":8080")
}
