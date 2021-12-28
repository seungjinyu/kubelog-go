package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/seungjinyu/kubelog-go/auth"
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
		v1.POST("/getpod", getpod)

	}

	r.GET("/", services.V1welcome)
	r.GET("/welcome", services.Welcome)
	r.GET("/health", services.Healthy)
	r.POST("/verifykeytest", auth.VerifyKey)
	r.POST("/")

	r.Run(":" + os.Getenv("PORT"))

}

func getpods(c *gin.Context) {
	datas := clusterinfo.GetPodListInfo(csi.Clientset)
	clusterinfo.SavePodInfoList(datas)
	c.JSON(http.StatusOK, gin.H{
		"datas": "Sending completed",
	})
}

type rbody struct {
	Namespace string `json:"namespace"`
	PodName   string `json:"podname"`
}

func getpod(c *gin.Context) {

	var rbodyi rbody

	err := json.NewDecoder(c.Request.Body).Decode(&rbodyi)
	if err != nil {
		log.Panic(err)
	}

	datas := clusterinfo.GetPodInfo(csi.Clientset, rbodyi.Namespace, rbodyi.PodName)
	// clusterinfo.SavePodInfo(datas)
	loc, _ := time.LoadLocation("UTC")

	podlog := strings.Split(datas.PodLog, "\n")

	c.JSON(http.StatusOK, gin.H{
		"Current Time": time.Now(),
		"UTC Time":     time.Now().In(loc),
		"Pod Name":     datas.PodName,
		"Pod Log":      podlog,
	})

	// c.Redirect(http.StatusMovedPermanently, "/results")

}
