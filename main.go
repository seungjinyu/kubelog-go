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

	// runtime.GOMAXPROCS(1)

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
	r.Use(middleware.ClientSet(csi.Clientset))

	basicService := r.Group("basic")
	basicService.Use(middleware.AuthenticationForBasic)
	{
		basicService.GET("/", services.RedirectToWelcome)
		basicService.GET("/welcome", services.V1welcome)
	}

	getpodService := r.Group("podservice")
	getpodService.Use(middleware.AuthenticationForPod)
	{
		getpodService.POST("/getpods", services.Getpods)
		getpodService.POST("/getpod", services.Getpod)
	}

	testpodservice := r.Group("test")
	testpodservice.POST("/getpods", services.Getpods)
	testpodservice.POST("/getpod", services.Getpod)
	testpodservice.GET("/sleep", services.Sleep)

	r.Run(":" + os.Getenv("PORT"))

	// http.ListenAndServe(":8080", r)

}
