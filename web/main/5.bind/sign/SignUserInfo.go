package main

import "github.com/gin-gonic/gin"

type SignUserInfo struct {
	//binding:"require" 信息校验
	Name       string `json:"name" binding:"required"`
	Password   string `json:"password"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		var user SignUserInfo
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(200, err.Error())
		} else {
			c.JSON(200, gin.H{"data": user})
		}
	})
	router.Run(":80")
}
