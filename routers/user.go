package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"blog/controllers/user"
)

func DealWithUserRoute(router *gin.RouterGroup) {
	router.GET("ping", func(c *gin.Context) {
		fmt.Println("The user is currently in user/ping")
	})
	router.POST("login", user.UserLogin)
	router.POST("register", user.UserRegister)
}