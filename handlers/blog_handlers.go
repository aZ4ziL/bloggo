package handlers

import (
	"net/http"

	"github.com/aZ4ziL/bloggo/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type BlogHandler interface {
	Index() gin.HandlerFunc
	Detail() gin.HandlerFunc
}

type blogHandler struct{}

func NewBlogHandler() BlogHandler {
	return &blogHandler{}
}

// Index
func (b blogHandler) Index() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		user := session.Get("user")

		if ctx.Request.Method == "GET" {
			categories := models.GetAllCategories()
			// articles := models.GetAllArticles()

			ctx.HTML(http.StatusOK, "index", gin.H{
				"user":       user,
				"categories": categories,
				// "articles":   articles,
			})
			return
		}
	}
}

// Detail
func (b blogHandler) Detail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		user := session.Get("user")

		if ctx.Request.Method == "GET" {
			slug := ctx.Param("slug")

			article, err := models.GetArticleBySlug(slug)
			if err != nil {
				http.Error(ctx.Writer, "404 not found", http.StatusNotFound)
				return
			}

			comments := models.GetAllCommentsByArticleID(article.ID)

			ctx.HTML(http.StatusOK, "detail", gin.H{
				"user":     user,
				"article":  article,
				"comments": comments,
			})
			return
		}
	}
}
