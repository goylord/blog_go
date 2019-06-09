package routers

import (
	"github.com/gin-gonic/gin"
	"blog/controllers/file"
)

func DealWithFileRoute(router *gin.RouterGroup) {
	router.POST("upload", file.FileUpload)
}