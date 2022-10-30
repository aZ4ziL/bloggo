package main

import (
	"encoding/gob"
	"time"

	"github.com/aZ4ziL/bloggo/models"
	"github.com/aZ4ziL/bloggo/routers"
	"github.com/aZ4ziL/bloggo/utils"
	"github.com/gin-contrib/sessions"
	gormsession "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
)

func init() {
	gob.Register(time.Time{})
	gob.Register(map[string]interface{}{})
}

func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// Media Files
	r.Static("/media", "./media")

	// Static Files
	r.Static("/static", "./static")

	r.HTMLRender = utils.CreateMyRender()

	store := gormsession.NewStore(models.GetDB(), true, []byte("mysecretkey"))
	r.Use(sessions.Sessions("ginsessionID", store))

	routers.AuthRouter(r)
	routers.BlogRouter(r)
	routers.RouterAPIV1(r)

	r.Run(":8000")
}
