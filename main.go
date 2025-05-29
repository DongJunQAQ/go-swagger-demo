package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-swagger-demo/api"
	"net/http"
	"strings"
	"time"
)

var users []*api.User

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowHeaders:     []string{"authorization", "content-type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "yourDomain.com")
		},
		MaxAge: 12 * time.Hour,
	}))
	r.POST("/user", Create)
	r.GET("/user/:name", Get) //路径参数路由
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

func Create(c *gin.Context) { //创建用户
	var newUser api.User                           //新添加的用户
	if err := c.ShouldBind(&newUser); err != nil { //ShouldBind:根据请求的Content-Type（如 application/json），自动将请求体中的数据解析到newUser结构体中
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "code": 10001})
		return
	}
	for _, u := range users {
		if u.Name == newUser.Name { //检查新添加的用户是否已存在
			c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("用户%s已存在", newUser.Name), "code": 10002})
			return
		}
	}
	users = append(users, &newUser)
	c.JSON(http.StatusOK, newUser)
}

func Get(c *gin.Context) { //查询用户
	username := c.Param("name") //提取路径参数
	for _, u := range users {
		if u.Name == username {
			c.JSON(http.StatusOK, u)
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("用户%s不存在", username), "code": 10003})
}
