package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	r := gin.Default()
	r.GET("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		fmt.Println(file.Header)
		fmt.Println(file.Size) //单位字节
		c.SaveUploadedFile(file, "web/main/6.upfile/uploads/1.png")
		c.JSON(200, gin.H{"msg": "上传成功"})
	})
	r.GET("/upload2", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		//读取文件
		readerfile, _ := file.Open()
		//创建一个新的文件
		writerfile, _ := os.Create("web/main/6.upfile/uploads/" + file.Filename)
		//关闭文件
		defer writerfile.Close()
		//拷贝文件
		io.Copy(writerfile, readerfile)
		c.JSON(200, gin.H{"msg": "上传成功"})
	})
	r.Run(":80")
}
