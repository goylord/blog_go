package routers

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouters(router *gin.Engine) {
	userRouter := router.Group("user")
	DealWithUserRoute(userRouter)
	articleRouter := router.Group("article")
	DealWithArticleRoute(articleRouter)
	fileRouter := router.Group("file")
	DealWithFileRoute(fileRouter)
}