package file
import (
	"github.com/gin-gonic/gin"
	"blog/services/file"
	"fmt"
)
func FileUpload(c *gin.Context) {
	file, _ := c.FormFile("file")
	assoId := c.PostForm("assoId")
	err := fileService.UploadFile(file, assoId)
	// fmt.Println(file)
	if err != nil {
		fmt.Printf("Error:%s", err.Error())
		c.JSON(500, gin.H{
			"status": 500,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H {
		"status": 200,
		"message": "文件上传成功",
	})
}
