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
	r.Use(cors.New(cors.Config{ //全局中间件，每个请求都会先经过这里，默认允许的请求方法为GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS
		AllowHeaders:     []string{"authorization", "content-type"}, //指定哪些请求头可以在跨域请求中使用
		AllowCredentials: true,                                      //是否允许跨源请求带有凭证（如cookies）
		AllowOriginFunc: func(origin string) bool { //动态决定是否允许特定的源访问资源，origin是发起请求的URL
			if strings.HasPrefix(origin, "http://localhost") { //开发环境，检查请求的源是否以http://localhost开头
				return true //允许该请求源
			}
			return strings.Contains(origin, "yourDomain.com") //生产环境，检查请求源是否包含生产环境下的域名
		},
		MaxAge: 12 * time.Hour, //preflight请求的缓存时间，使得浏览器在12小时内不需要重新发送preflight请求，而是可以使用缓存的结果来决定是否允许实际请求
	}))
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
