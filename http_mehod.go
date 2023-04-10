package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	// 使用默认中间件（logger 和 recovery 中间件）创建 gin 路由
	router := gin.Default()

	router.GET("/someGet", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "this is a http GET request"})
	})
	router.POST("/somePost", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "this is a http POST request"})
	})
	router.PUT("/somePut", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "this is a http PUT request"})
	})
	router.DELETE("/someDelete", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "this is a http DELETE request"})
	})
	router.PATCH("/somePatch", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "this is a http PATCH request"})
	})
	router.HEAD("/someHead", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "this is a http HEAD request"})
	})
	router.OPTIONS("/someOptions", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "this is a http OPTIONS request"})
	})

	// 默认在 8080 端口启动服务，除非定义了一个 PORT 的环境变量。
	router.Run()
	// router.Run(":3000") hardcode 端口号
}
