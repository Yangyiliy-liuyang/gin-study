package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//1、创建一个默认的路由引擎
	r := gin.Default()
	//2、绑定路由规则和和路由函数
	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "hello golang",
		})
	})
	//3、启动监听，gin会把web服务运行在本机的0.0.0.0:8080端口上
	//r.Run(":8080")
	//或者使用原生的http服务的方式 在router.Run中封装ListenAndServe
	http.ListenAndServe(":8080", r)
}
