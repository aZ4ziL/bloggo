package routers

import (
	"github.com/aZ4ziL/bloggo/handlers"
	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.Engine) {
	auth := r.Group("/auth")

	h := handlers.NewAuthHandler()

	auth.GET("/login", h.Login())
	auth.GET("/logout", h.Logout())
	auth.GET("/register", h.Register())
}
