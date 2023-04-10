package main

import "github.com/gin-gonic/gin"

type myForm struct {
	Colors []string `form:"colors[]"`
}

func main() {

	r := gin.Default()
	r.POST("/submit", func(c *gin.Context) {
		var fakeForm myForm
		c.ShouldBind(&fakeForm)
		c.JSON(200, gin.H{"color": fakeForm.Colors})
	})
	r.Run()

}
