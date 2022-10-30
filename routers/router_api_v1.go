package routers

import (
	"github.com/aZ4ziL/bloggo/handlers"
	"github.com/gin-gonic/gin"
)

func RouterAPIV1(r *gin.Engine) *gin.RouterGroup {
	api := r.Group("/api/v1")

	h := handlers.NewAPIV1()

	// Article
	api.GET("/article", h.GetAllArticles())
	api.POST("/article", h.CreateNewArticle())
	api.PUT("/article", h.CreateNewArticle())

	// Comment
	api.GET("/comment", h.GetAllCommentsByArticleID())

	api.GET("/category", h.GetAllCategories())

	// user
	api.GET("/auth/user", h.GetUserByUsername())
	api.POST("/auth/user/register", h.Register())
	api.POST("/auth/user/login", h.Login())
	api.GET("/auth/user/login", h.Login())
	api.GET("/auth/user/logout", h.Logout())

	return api
}
