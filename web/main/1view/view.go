package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func _json(c *gin.Context) {
	type UserInfo struct {
		Username string `json:"user_name"`
		Age      int    `json:"user_age"`
		Password string `json:"-"`
	}
	user := UserInfo{"江城小王子", 22, "123456"}
	c.JSON(200, user)
	userMap := map[string]interface{}{
		"username": "zhangsan",
		"age":      18,
	}
	c.JSON(200, userMap)
}
func _xml(c *gin.Context) {
	c.XML(200, gin.H{
		"username": "江城小王子",
		"age":      18,
		"status":   http.StatusOK,
		"data": gin.H{
			"msg": "hello golang",
		},
	})
}
func _yaml(context *gin.Context) {
	context.YAML(200, gin.H{
		"username": "江城小王子",
		"age":      18,
		"status":   http.StatusOK,
		"data": gin.H{
			"msg": "hello golang",
		},
	})
}

//响应html
func _html(context *gin.Context) {
	context.HTML(200, "index.tmpl", gin.H{
		"name": "江城小王子",
	})
}
func main() {
	r := gin.Default()
	//加载模版目录下所有html
	r.LoadHTMLFiles("templates/*")
	// 网页请求这个静态目录的前缀 第二个参数是一个目录，注意前缀不要重复
	r.StaticFS("/static", http.Dir("web/main/1view/static/static"))
	//在golang中没有相对文件的目录，只要相对项目的路径
	//配置单个文件 relativePath 网页请求的路由（文件访问路径 127.0.0.1/static/background.png 第二个参数filePath是文件路径
	r.StaticFile(`/background.png`, "web/main/1view/static/background.png")

	r.GET("/", func(c *gin.Context) {
		//响应字符串 状态码200
		c.String(http.StatusOK, "你好啊")
		c.JSON(200, gin.H{
			"msg": "hello golang",
		})
	})
	r.GET("/json", _json)
	r.GET("/xml", _xml)
	r.GET("/yaml", _yaml)
	r.GET("/html", _html)

	r.GET("/baidu", _redirect)

	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("userpassword")
		// c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	})
	r.Run(":8080")
}

//重定向
func _redirect(context *gin.Context) {
	context.Redirect(301, "https://www.baidu.com")
}
