package routers

import (
	"github.com/aZ4ziL/bloggo/handlers"
	"github.com/gin-gonic/gin"
)

func BlogRouter(r *gin.Engine) {
	h := handlers.NewBlogHandler()

	r.GET("", h.Index())
	r.GET("/detail/:slug", h.Detail())
}
