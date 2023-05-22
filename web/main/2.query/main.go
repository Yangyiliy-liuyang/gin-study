package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/query", _query)
	r.GET("Parma/user_id", _parma)
	r.POST("/from", _from)
	r.POST("/raw", _raw)
	r.Run(":8080")
}

//原始参数
func _raw(context *gin.Context) {
	body, _ := context.GetRawData()
	fmt.Println(string(body))
}

//表单参数
func _from(context *gin.Context) {
	name := context.PostForm("name")
	fmt.Println(name)
}

//动态参数
func _parma(context *gin.Context) {
	fmt.Println(context.Param("user_id"))
}

//查询参数
func _query(context *gin.Context) {
	user := context.Query("user")
	_, ok := context.GetQuery("user")
	fmt.Print(ok, user)
	fmt.Println(context.QueryArray(user))
}
