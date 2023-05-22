package main

import "github.com/gin-gonic/gin"

type UserInfo struct {
	Name string `json:"name" from:"name" uri:"name"`
	Age  int    `json:"age" from:"age" uri:"age"`
	Sex  string `json:"sex" from:"sex" uri:"sex"`
}

func main() {
	r := gin.Default()
	r.POST("/bind", func(context *gin.Context) {
		var userInfo UserInfo
		err := context.ShouldBindJSON(userInfo)
		if err != nil {
			context.JSON(200, gin.H{"msg": "err"})
			return
		}
		context.JSON(200, gin.H{
			"userInfo": userInfo,
		})
	})
	r.POST("/query", func(context *gin.Context) {
		var userInfo UserInfo
		err := context.ShouldBindQuery(userInfo)
		if err != nil {
			context.JSON(200, gin.H{"msg": "err"})
			return
		}
		context.JSON(200, userInfo)
	})
	r.POST("/uri", func(context *gin.Context) {
		var userInfo UserInfo
		err := context.ShouldBindUri(userInfo)
		if err != nil {
			context.JSON(200, gin.H{"msg": "err"})
			return
		}
		context.JSON(200, userInfo)
	})
	r.POST("/form", func(context *gin.Context) {
		var userInfo UserInfo
		err := context.ShouldBind(userInfo)
		if err != nil {
			context.JSON(200, gin.H{"msg": "err"})
			return
		}
		context.JSON(200, userInfo)
	})
	r.Run(":80")
}
