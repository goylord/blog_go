package comment

import (
	"github.com/gin-gonic/gin"
	"blog/services/comment"
)

func Create(c *gin.Context) {
	httpStatus := 200
	httpMessage := "评论成功"
	commentUSerId := c.Query("userId")
	byCommentUserId := c.Query("byCommentId")
	commentContent := c.Query("content")
	commentArticleId := c.Query("articleId")
	if commentUSerId == "" {
		httpStatus = 400
		httpMessage = "请登陆后评论"
	}
	if commentContent == "" {
		httpStatus = 400
		httpMessage = "请输入评论内容"
	}
	if commentArticleId == "" {
		httpStatus = 400
		httpMessage = "请在文章内评论"
	}
	if httpStatus != 200 {
		c.JSON(httpStatus, gin.H{
			"status": httpStatus,
			"message": httpMessage,
		})
		return
	}
	comment, err := commentService.CreateComment(commentUSerId, byCommentUserId, commentContent, commentArticleId)
	if err != nil {
		c.JSON(500, gin.H{
			"status": 500,
			"message": "评论失败",
		})
		return
	}
	c.JSON(httpStatus, gin.H{
		"status": httpStatus,
		"message": httpMessage,
		"data": comment,
	})
}
func GetCommentByArticleId(c *gin.Context) {
	httpStatus := 200
	httpMessage := "查询成功"
	articleId := c.Query("articleId")
	if articleId == "" {
		httpStatus = 400
		httpMessage = "请传入文章ID"
	}
	commentList, err := commentService.GetCommentListByArticleId(articleId)
	if err != nil {
		c.JSON(500, gin.H{
			"status": 500,
			"message": "查询失败",
		})
		return
	}
	c.JSON(httpStatus, gin.H{
		"status": httpStatus,
		"message": httpMessage,
		"data": commentList,
	})
	
}
func Delete(c *gin.Context) {
}
func Edit(c *gin.Context) {
}