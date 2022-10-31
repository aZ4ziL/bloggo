package main

import (
	"text/template"

	"github.com/gin-contrib/multitemplate"
)

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	// Auth
	r.AddFromFiles("auth_login", "views/auth/login.tmpl")
	r.AddFromFiles("auth_register", "views/auth/register.tmpl")

	// Blog
	r.AddFromFilesFuncs(
		"index",
		template.FuncMap{},
		"views/blogs/base.tmpl",
		"views/blogs/navbar.tmpl",
		"views/blogs/footer.tmpl",
		"views/blogs/index.tmpl",
	)
	r.AddFromFilesFuncs(
		"detail",
		template.FuncMap{
			"markdown":        markdown,
			"getFullNameByID": getFullNameByID,
		},
		"views/blogs/detail.tmpl",
		"views/blogs/navbar.tmpl",
		"views/blogs/footer.tmpl",
	)
	r.AddFromFilesFuncs(
		"search",
		template.FuncMap{},
		"views/blogs/base.tmpl",
		"views/blogs/navbar.tmpl",
		"views/blogs/footer.tmpl",
		"views/blogs/search.tmpl",
	)

	return r
}
