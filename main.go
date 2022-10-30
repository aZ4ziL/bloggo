package main

import (
	"github.com/aZ4ziL/bloggo/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	routers.RouterAPIV1(r)

	r.Run(":8000")
}
