package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		//data :=map[string]interface{}{"name":"king"}
		data := gin.H{"name":"king","age":18}
		c.JSON(http.StatusOK, data)
	})

	type person struct {
		Name string
		Age int
	}
	r.GET("/test", func(c *gin.Context) {
		data :=person{"kou",20}
		c.JSON(http.StatusOK, data)
	})
	r.GET("/web", func(c *gin.Context) {
		// 获取浏览器那边携带请求参数
		name:=c.Query("name")
		c.JSON(http.StatusOK,name)

	})

	err := r.Run()
	if err != nil {
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}
