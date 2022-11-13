package main

import (
	"golang-trupentool/controllers"
	"golang-trupentool/initializers"
	"golang-trupentool/seeders"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVar()
	initializers.DbConnect()
}

func main() {
	argLen := len(os.Args)
	if argLen > 1 {
		arg := os.Args[1]
		if arg == "seeders" {
			seeders.Main()
			return
		}
		return
	}

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", controllers.GetHome)
	router.GET("/commanders", controllers.GetCommanders)
	router.GET("/troops", controllers.GetTroops)
	router.GET("/troops/train", controllers.GetTroopsTrain)
	router.POST("/troops/train", controllers.PostTroopTrain)
	router.Run()
}
