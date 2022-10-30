package utils

import (
	"text/template"

	"github.com/gin-contrib/multitemplate"
)

func CreateMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	// Auth
	r.AddFromFiles("auth_login", "views/auth/login.tmpl")
	r.AddFromFiles("auth_register", "views/auth/register.tmpl")

	// Blog
	r.AddFromFilesFuncs(
		"index",
		template.FuncMap{},
		"views/blogs/base.tmpl",
		"views/blogs/index.tmpl",
	)
	r.AddFromFilesFuncs(
		"detail",
		template.FuncMap{
			"markdown": Markdown,
		},
		"views/blogs/detail.tmpl",
	)

	return r
}
