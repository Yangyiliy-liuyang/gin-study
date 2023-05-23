package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()

	// 定义中间件
	r.Use(func(c *gin.Context) {
		fmt.Println("Start middleware")
		start := time.Now()

		// 执行路由处理函数
		c.Next()

		// 在处理函数执行之后执行代码
		elapsed := time.Since(start)
		fmt.Println("Elapsed time:", elapsed)
	})

	// 定义路由 m1 m2为中间件，func为视图
	r.GET("/", m1, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
		//拦截
		c.Abort()
	}, m2)

	// 启动服务
	r.Run(":80")
}

func m2(context *gin.Context) {
	fmt.Println("m2.....")

}

func m1(context *gin.Context) {
	fmt.Println("m1.....")
}
