package utils

import (
	"text/template"

	"github.com/gin-contrib/multitemplate"
)

func CreateMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	// Auth
	r.AddFromFiles("auth_login", "views/auth/login.tpl")
	r.AddFromFiles("auth_register", "views/auth/register.tpl")

	// Blog
	r.AddFromFilesFuncs(
		"index",
		template.FuncMap{},
		"views/blogs/base.tpl",
		"views/blogs/index.tpl",
	)
	r.AddFromFilesFuncs(
		"detail",
		template.FuncMap{
			"markdown": Markdown,
		},
		"views/blogs/detail.tpl",
	)

	return r
}
