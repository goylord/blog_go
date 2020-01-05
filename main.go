package main

import (
	"blog/middleware"
	"blog/routers"
	"blog/util"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Register the sever router
	router := gin.Default()
	middleware.UserSession(router)
	router.Use(middleware.UserPermission)
	log.Println(util.GetCurrentDirectory())
	router.Static("/publish", util.GetCurrentDirectory()+"/publish")
	routers.RegisterRouters(router)
	router.Run()
}
