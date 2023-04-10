package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// 简单的路由组: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", func(c *gin.Context) {

		})
		v1.POST("/submit", func(c *gin.Context) {

		})
		v1.POST("/read", func(c *gin.Context) {

		})
	}

	// 简单的路由组: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", func(c *gin.Context) {

		})
		v2.POST("/submit", func(c *gin.Context) {

		})
		v2.POST("/read", func(c *gin.Context) {

		})
	}

	router.Run(":8080")
}

// 启动后控制台输出
//[GIN-debug] POST   /v1/login                 --> main.main.func1 (3 handlers)
//[GIN-debug] POST   /v1/submit                --> main.main.func2 (3 handlers)
//[GIN-debug] POST   /v1/read                  --> main.main.func3 (3 handlers)
//[GIN-debug] POST   /v2/login                 --> main.main.func4 (3 handlers)
//[GIN-debug] POST   /v2/submit                --> main.main.func5 (3 handlers)
//[GIN-debug] POST   /v2/read                  --> main.main.func6 (3 handlers)
