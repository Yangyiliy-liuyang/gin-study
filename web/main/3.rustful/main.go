package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ArticleModel struct {
	Title   string `json:"title"`
	Context string `json:"context"`
	id      int    `json:"id"`
}
type Response struct {
	Code string `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

//文章列表页
func _getList(c *gin.Context) {
	//包含搜索 分页 关键词
	articleList := []ArticleModel{
		{
			Title:   "golang入门",
			Context: "这是一本关于golang入门的书籍，在这本书中你可以....",
			id:      1,
		}, {
			Title:   "gin框架中文笔记",
			Context: "这是一本关gin框架的书籍，在这本书中你可以....",
			id:      2,
		},
	}
	c.JSON(200, Response{
		Code: "0",
		Data: articleList,
		Msg:  "成功",
	})
}

//文章详情
func _getDetail(c *gin.Context) {
	//获取parma的id
	fmt.Println(c.Param("id"))
	article := ArticleModel{
		Title:   "golang入门",
		Context: "这是一本关于golang入门的书籍，在这本书中你可以....",
		id:      1,
	}
	c.JSON(200, Response{
		Code: "0",
		Data: article,
		Msg:  "成功",
	})
	c.JSON(200, gin.H{})
}

func _bindJson(c *gin.Context, obj any) (err error) {
	body, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		err = json.Unmarshal(body, &obj)
		if err != nil {
			fmt.Printf(err.Error())
			return err
		}
	}
	return nil
}

//创建文章
func _crete(c *gin.Context) {
	//接受前端传过来的json数据
	var article ArticleModel
	err := _bindJson(c, &article)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, Response{
		Code: "0",
		Data: article,
		Msg:  "成功",
	})
}

//编辑文章
func _updata(c *gin.Context) {
	fmt.Println(c.Param("id"))
	var article ArticleModel
	err := _bindJson(c, &article)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, Response{
		Code: "0",
		Data: article,
		Msg:  "修改成功",
	})
}

//删除文章
func _delete(c *gin.Context) {
	fmt.Println(c.Param("id"))
	//........这里按id去数据库查找删除
	c.JSON(200, Response{
		Code: "0",
		Data: map[string]string{},
		Msg:  "删除成功",
	})
}

func main() {
	r := gin.Default()
	r.GET("/articles", _getList)
	r.GET("/articles/:id", _getDetail)
	r.POST("/articles/:id", _crete)
	r.PUT("/articles/:id", _updata)
	r.DELETE("/articles/:id", _delete)
	r.Run(":80")
}
