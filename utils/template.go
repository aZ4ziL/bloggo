package utils

import "github.com/gin-contrib/multitemplate"

func CreateMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	// Auth
	r.AddFromFiles("auth_login", "views/auth/login.tpl")
	r.AddFromFiles("auth_register", "views/auth/register.tpl")

	return r
}
