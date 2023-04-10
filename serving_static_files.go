package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.StaticFS("/more_static", http.Dir("html"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// 监听并在 0.0.0.0:8080 上启动服务
	router.Run(":8080")
}

// 启动后控制台输出：
//[GIN-debug] GET    /assets/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
//[GIN-debug] HEAD   /assets/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
//[GIN-debug] GET    /more_static/*filepath    --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
//[GIN-debug] HEAD   /more_static/*filepath    --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
//[GIN-debug] GET    /favicon.ico              --> github.com/gin-gonic/gin.(*RouterGroup).StaticFile.func1 (3 handlers)
//[GIN-debug] HEAD   /favicon.ico              --> github.com/gin-gonic/gin.(*RouterGroup).StaticFile.func1 (3 handlers)
