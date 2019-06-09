package main
import(
	"blog/routers"
	"github.com/gin-gonic/gin"
	"blog/util"
	"log"
)
import (
)

func main()  {
	// Register the sever router
	router := gin.Default()
	log.Println(util.GetCurrentDirectory())
	router.Static("/publish", util.GetCurrentDirectory() + "/publish")
	routers.RegisterRouters(router)
	router.Run()
}