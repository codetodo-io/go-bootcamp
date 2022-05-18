package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 提供 unicode 实体
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b> 你好, JSON!</b>",
		})
	})

	// 提供字面字符
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>你好, PureJSON!</b>",
		})
	})

	// 监听 8080 端口，启动服务.
	r.Run(":8080")
}
