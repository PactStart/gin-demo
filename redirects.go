package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 外部重定向
	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})

	r.GET("/bar", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/foo")
	})

	r.GET("/foo", func(c *gin.Context) {
		c.JSON(200, gin.H{"foo": "foo"})
	})

	// 路由重定向

	r.GET("/test2", func(c *gin.Context) {
		c.Request.URL.Path = "/test3"
		r.HandleContext(c)
	})
	r.GET("/test3", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	r.Run()
}
