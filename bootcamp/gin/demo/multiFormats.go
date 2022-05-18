package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func main() {
	r := gin.Default()
	// gin.H 是 map[string]interface{} 的一种快捷方式
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "qiaqia",
			"status":  http.StatusOK,
		})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		// 可以使用结构体
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "qiaqia"
		msg.Message = "my name"
		msg.Number = 12
		// 注意 msg.Name 在 JSON 中为了 "user"
		// JSON 输出：{"user": "qiaqia", "Message": "my name", "Number": 12}
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"message": "qiaqia",
			"status":  http.StatusOK,
		})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{
			"message": "qiaqia",
			"status":  http.StatusOK,
		})
	})

	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// protobuf 的具体定义写在 testdata/protoexample 文件中
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// 注意，数据在响应中为二进制格式，将输出被 protoexample.Test protobuf 序列化后的数据
		c.ProtoBuf(http.StatusOK, data)
	})

	// 监听 8080 端口，启动服务.
	r.Run(":8080")
}
