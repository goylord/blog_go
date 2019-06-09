package user

import (
	"github.com/gin-gonic/gin"
	"blog/services/user"
	"fmt"
)



func UserRegister(c *gin.Context) {
	var httpStatus = 200
	var errMessage = "注册成功"
	username := c.Query("username")
	password := c.Query("password")
	loginId := c.Query("loginId")
	if username == "" {
		httpStatus = 400
		errMessage = "请填写用户名"
	}
	if password == "" {
		httpStatus = 400
		errMessage = "请填写密码"
	} else {
		if len(password) < 6 {
			httpStatus = 400
			errMessage = "密码过短"
		}
	}
	if loginId == "" {
		httpStatus = 400
		errMessage = "请填写登陆名"
	}
	if httpStatus == 200 {
		err := userService.CreateNewUser(username, loginId, password)
		if err != nil {
			fmt.Println(err.Error())
		}
		c.JSON(httpStatus, gin.H{
			"status": httpStatus,
			"message": errMessage,
		})
	} else {
		c.JSON(httpStatus, gin.H{
			"status": httpStatus,
			"message": errMessage,
		})
	}
}
func UserLogin(c *gin.Context) {
	loginId := c.Query("loginId")
	passWord := c.Query("password")
	loginStatus, id, err := userService.UserLogin(loginId, passWord)
	if err != nil {
		fmt.Println("User login failed! error: %s", err.Error())
		c.JSON(500, gin.H{
			"status":500,
			"message": "查询数据库出错",
		})
		return
	}
	if !loginStatus {
		c.JSON(400, gin.H{
			"status":400,
			"message": "登录失败，用户名或密码错误！",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":200,
		"message": "登录成功",
		"token": id,
	})
}
func UserChangeName(c *gin.Context) {
	
}
func UserChangePassword(c *gin.Context) {

}