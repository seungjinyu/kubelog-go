package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/seungjinyu/kubelog-go/clusterinfo"
	"github.com/seungjinyu/kubelog-go/middleware"
	"github.com/seungjinyu/kubelog-go/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var csi clusterinfo.ClientSetInstance

func main() {

	fmt.Println("VERSION 1.0.2")

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

	v1 := r.Group("v1")
	v1.Use(middleware.Authenticate)
	{
		v1.GET("/v1welcome", services.V1welcome)
		v1.POST("/getpods", getpods)
	}

	r.GET("/", services.V1welcome)
	r.GET("/welcome", services.Welcome)
	r.GET("/health", services.Healthy)

	r.Run(":" + os.Getenv("PORT"))

}

func getpods(c *gin.Context) {
	datas := clusterinfo.GetPodListInfo(csi.Clientset)
	clusterinfo.SavePodInfoList(datas)
	c.JSON(http.StatusOK, gin.H{
		"datas": "Sending completed",
	})
}
