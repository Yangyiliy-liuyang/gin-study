package main

import "github.com/gin-gonic/gin"

type SignUserInfo struct {
	//binding:"require" 信息校验
	Name       string   `json:"name" binding:"required"`
	Password   string   `json:"password"`
	RePassword string   `json:"re_password" binding:"required,eqfield=Password"`
	Age        int      `json:"age" binding:"lt=30,gt=19""`
	Sex        string   `json:"sex" binding:"oneof=man woman"`
	LinkList   []string `json:"link_list" binding:"required,dive,startswith=like"`
	IP         string   `json:"IP" binding:"ip"`
	Url        string   `json:"url" binding:"url"`
	Uri        string   `json:"uri" binding:"uri"`
	Date       string   `json:"date" binding:"datetime=2006-01-02 15:04:06"`
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		var user SignUserInfo
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(200, gin.H{"msg": err.Error()})
		} else {
			c.JSON(200, gin.H{"data": user})
		}
	})
	router.Run(":80")
}
