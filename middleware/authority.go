package middleware

import (
	model "blog/models/user"
	"fmt"
	"log"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var store cookie.Store

type User struct {
	Username string
	Id       string
	LoginId  string
}

func init() {
	// 初始化一个session仓库
	store = cookie.NewStore([]byte("secret"))
}

// 用户session中间件
func UserSession(router *gin.Engine) {
	router.Use(sessions.Sessions("systemSession", store))
}

// 用户登陆判断
func UserPermission(c *gin.Context) {
	if c.FullPath() == "/user/login" || c.FullPath() == "/user/register" || c.FullPath() == "" {
		log.Println("用户登陆")
		c.Next()
		return
	}
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		log.Println("用户未登陆！")
		c.JSON(403, gin.H{
			"status":  "403",
			"message": "No Permission!",
		})
		c.Abort()
		return
	}
	log.Println("用户已经登陆-->")
	userInfos := strings.Split(user.(string), "&")
	currentUser := User{
		Username: userInfos[0],
		LoginId:  userInfos[1],
		Id:       userInfos[2],
	}
	log.Println("输出用户-->", currentUser)
	c.Next()
}

// 储存用户session
func SaveUserSession(c *gin.Context, user model.User) {
	session := sessions.Default(c)
	log.Printf("用户登陆-", user.Username)
	session.Set("user", fmt.Sprintf("%s&%s&%s", user.Username, user.LoginId, user.Id))
	session.Save()
}
