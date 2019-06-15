package routers

import (
	"github.com/gin-gonic/gin"
	"blog/controllers/comment"
)

func DealWithCommentRoute(router *gin.RouterGroup) {
	router.POST("create", comment.Create)
	router.POST("edit", comment.Edit)
	router.POST("delete", comment.Delete)
	router.POST("getList", comment.GetCommentByArticleId)

}