package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Login() gin.HandlerFunc
	Register() gin.HandlerFunc
	Logout() gin.HandlerFunc
}

type authHandler struct{}

func NewAuthHandler() AuthHandler {
	return &authHandler{}
}

// Login handler
func (a authHandler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		if user := session.Get("user"); user != nil {
			ctx.Redirect(http.StatusFound, "/")
			return
		}

		if ctx.Request.Method == "GET" {
			ctx.HTML(http.StatusOK, "auth_login", gin.H{})
			return
		}
	}
}

// Register
func (a authHandler) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		if user := session.Get("user"); user != nil {
			ctx.Redirect(http.StatusFound, "/")
			return
		}

		if ctx.Request.Method == "GET" {
			ctx.HTML(http.StatusOK, "auth_register", gin.H{})
			return
		}
	}
}

// Logout
func (a authHandler) Logout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		session.Clear()
		session.Delete("user")
		session.Save()

		ctx.Redirect(http.StatusFound, "/auth/login")
	}
}
