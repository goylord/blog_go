package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"blog/controllers/article"
)

func DealWithArticleRoute(router *gin.RouterGroup) {
	router.GET("ping", func(c *gin.Context) {
		fmt.Println("The article is currently in article/ping")
	})
	router.POST("publish", article.Publish)
	router.POST("delete", article.Delete)
	router.POST("edit", article.Edit)
	router.POST("list", article.GetArticleList)
}