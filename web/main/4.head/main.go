package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//请求头的各种获取方式
	router.GET("/", func(context *gin.Context) {
		//首字母大小写不区分
		//用于获取一个请求头
		fmt.Println(context.GetHeader("User-Agent"))
		//Header是map[string][]string
		fmt.Println(context.Request.Header)
		//如果是使用get或者GetHeader，可以不区分大小写，返回第一个
		fmt.Println(context.Request.Header.Get("User-Agent"))
		//如果是map的取值方式，要注意大小写
		fmt.Println(context.Request.Header["User-Agent"])

		//自定义的请求头也可以不区分大小写
		fmt.Println(context.Request.Header.Get("token"))
		context.JSON(200, gin.H{"msg": "成功"})
	})
	//爬虫和用户区分对待
	router.GET("/index", func(context *gin.Context) {
		//.....一些不同的处理
	})
	//设置响应头
	router.GET("/res", func(context *gin.Context) {
		context.Header("Token", "ahsd;a&dskf")
		context.Header("Content-Type", "application/txt; charset=utf-8")
		context.JSON(200, gin.H{"msg": "查看响应头"})
	})
	router.Run(":80")
}
