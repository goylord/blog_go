package article

import (
	"github.com/gin-gonic/gin"
	"blog/services/article"
)

func Publish(c *gin.Context) {
	httpStatus := 200
	httpMessage := "文章新建成功"
	title := c.Query("title")
	content := c.Query("content")
	previewContent := c.Query("previewContent")
	creatorId := c.Query("creatorId")
	classId := c.Query("classId")
	if title == "" {
		httpStatus = 400
		httpMessage = "请输入标题"
	}
	if content == "" {
		httpStatus = 400
		httpMessage = "请输入内容"
	}
	if previewContent == "" {
		httpStatus = 400
		httpMessage = "请输入预览文本"
	}
	if creatorId == "" {
		var err error
		creatorId, err = c.Cookie("token")
		if err != nil {
			creatorId = "01"
		}
	}
	if (httpStatus != 200) {
		c.JSON(httpStatus, gin.H{
			"status": httpStatus,
			"message": httpMessage,
		})	
		return
	}
	articleId, err := articleService.CreateNewArticle(title, content, previewContent, creatorId, classId)
	if err != nil {
		c.JSON(500, gin.H{
			"status": 500,
			"message": "文章新建失败" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 200,
		"message": "文章新建成功",
		"id": articleId,
	})

}
func Delete(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(400, gin.H{
			"status": 400,
			"message": "请传入删除文章id",
		})
	}
	_, err := articleService.DeleteArticle(id)
	if err != nil {
		c.JSON(500, gin.H{
			"status": 500,
			"message": "删除失败，" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 200,
		"message": "删除成功",
	})
}
func Edit(c *gin.Context) {
}
func GetArticleList(c *gin.Context) {
	var articleList []map[string]interface{}
	pageNo := c.Query("pageNo")
	pageSize := c.Query("pageSize")
	if pageNo == "" {
		pageNo = "1"
	}
	if pageSize == "" {
		pageSize = "10"
	}
	articleList, err := articleService.GetList(pageNo, pageSize)
	if err != nil {
		c.JSON(500, gin.H {
			"status": 500,
			"message": "查询出错, error:" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 200,
		"message": "查询成功",
		"data": articleList,
	})
}