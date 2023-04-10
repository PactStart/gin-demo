package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
import "github.com/fvbock/endless"

func main() {
	// 不使用默认的中间件
	router := gin.New()
	// Default 使用 Logger 和 Recovery 中间件
	//r := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})
	// [...]
	endless.ListenAndServe(":4242", router)

	//ctrl+c控制台关闭进程时，看日志输出如下：
	//^C2023/04/09 05:12:50 20587 Received SIGINT.
	//2023/04/09 05:12:50 20587 Waiting for connections to finish...
	//2023/04/09 05:12:50 20587 Serve() returning...

	// 不能看出优雅停止，收到信号后，会先等所有的连接完成工作，然后才退出

}
