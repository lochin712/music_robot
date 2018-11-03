package main

import (
	"music_robot/handlers"
	"music_robot/util"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/search", handlers.Search)
	util.SetProcessID("robot.pid")

	r.Run(":1024")
}
