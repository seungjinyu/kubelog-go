package main

import (
	"fmt"
	"log"
	"os"

	"github.com/seungjinyu/kubelog-go/clusterinfo"
	"github.com/seungjinyu/kubelog-go/middleware"
	"github.com/seungjinyu/kubelog-go/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("VERSION 1.1.1")

	var csi clusterinfo.ClientSetInstance

	err := godotenv.Load(".env")

	if err != nil {
		log.Println(err.Error())
	}

	appenv := os.Getenv("APP_ENV")
	if appenv != "OUT" {
		err = csi.CreateInClientSet()
		if err != nil {
			log.Println(err.Error())
		}

	} else {
		err = csi.CreateOutClientSet()
		if err != nil {
			log.Println(err.Error())
		}
	}

	r := gin.Default()

	r.GET("/health", services.Healthy)

	basicService := r.Group("v1")
	basicService.Use(middleware.ClientSet(csi.Clientset))
	basicService.Use(middleware.AuthenticationForBasic)
	{
		basicService.GET("/", services.RedirectToWelcome)
		basicService.GET("/welcome", services.V1welcome)
	}

	getpodService := r.Group("v2")

	getpodService.Use(middleware.AuthenticationForPod)
	{
		getpodService.POST("/getpods", services.Getpods)
		getpodService.POST("/getpod", services.Getpod)
	}

	// r.POST("/verifykeytest", auth.VerifyKey)

	r.Run(":" + os.Getenv("PORT"))

}
