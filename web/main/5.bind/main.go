package main

import "github.com/gin-gonic/gin"

type UserInfo struct {
	Name string `json:"name" form:"name" uri:"name"`
	Age  int    `json:"age" form:"age" uri:"age"`
	Sex  string `json:"sex" form:"sex" uri:"sex"`
}

func main() {
	r := gin.Default()
	r.POST("/bind", func(c *gin.Context) {
		var user UserInfo
		if c.BindJSON(&user) == nil {
			c.JSON(200, gin.H{"status": "ok", "user": user})
		} else {
			c.JSON(400, gin.H{"status": "error"})
		}
	})
	r.GET("/query", func(context *gin.Context) {
		var userInfo UserInfo
		// Bind query string 数据到userInfo变量
		if err := context.ShouldBindQuery(&userInfo); err == nil {
			context.JSON(200, gin.H{"status": "ok", "userInfo": userInfo})
		} else {
			context.JSON(400, gin.H{"status": "error"})
		}
	})

	r.POST("/uri", func(context *gin.Context) {
		var userInfo UserInfo
		if err := context.ShouldBindUri(&userInfo); err == nil {
			context.JSON(200, gin.H{"status": "ok", "userInfo": userInfo})
		} else {
			context.JSON(400, gin.H{"status": "error"})
		}
	})
	r.POST("/form", func(context *gin.Context) {
		var userInfo UserInfo
		if err := context.ShouldBind(&userInfo); err == nil {
			context.JSON(200, gin.H{"status": "ok", "userInfo": userInfo})
		} else {
			context.JSON(400, gin.H{"status": "error"})
		}
	})
	r.Run(":80")
}
