package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-swagger-demo/api"
	"net/http"
)

var users []*api.User

func main() {
	r := gin.Default()
	r.POST("/user", Create)
	r.GET("/user/:name", Get) //路径参数路由
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

func Create(c *gin.Context) {
	var newUser api.User //新添加的用户
	if err := c.ShouldBind(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "code": 10001})
		return
	}
	for _, u := range users {
		if u.Name == newUser.Name {
			c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("用户%s已存在", newUser.Name), "code": 10002})
			return
		}
	}
	users = append(users, &newUser)
	c.JSON(http.StatusOK, newUser)
}

func Get(c *gin.Context) {
	username := c.Param("name") //提取路径参数
	for _, u := range users {
		if u.Name == username {
			c.JSON(http.StatusOK, u)
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("用户%s不存在", username), "code": 10003})
}
