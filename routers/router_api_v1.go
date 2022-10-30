package routers

import (
	"github.com/aZ4ziL/bloggo/handlers"
	"github.com/gin-gonic/gin"
)

func RouterAPIV1(r *gin.Engine) *gin.RouterGroup {
	api := r.Group("v1")

	h := handlers.NewAPIV1()

	api.GET("/article", h.GetAllArticles())

	api.GET("/category", h.GetAllCategories())

	// user
	api.GET("/auth/user", h.GetUserByUsername())
	api.POST("/auth/user/register", h.Register())

	return api
}
