package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/seungjinyu/kubelog_go/clusterinfo"
	"k8s.io/client-go/kubernetes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err.Error())
	}
	// appenv := "IN"
	appenv := os.Getenv("APP_ENV")
	var clientset *kubernetes.Clientset

	if appenv != "OUT" {
		clientset, err = clusterinfo.CreateInClientSet()
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		clientset, err = clusterinfo.CreateOutClientSet()
		if err != nil {
			log.Println(err.Error())
		}
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Main": "Page",
		})
	})

	router.GET("/welcome", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"greetings": "welcome",
		})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})
	router.GET("/getpods", func(c *gin.Context) {

		datas := clusterinfo.GetPodListInfo(clientset)
		clusterinfo.SavePodInfoList(datas)
		c.JSON(http.StatusOK, gin.H{
			"datas": "Sending completed",
		})

		// c.Redirect(http.StatusMovedPermanently, "/results")

	})
	router.GET("/getpod", func(c *gin.Context) {

		namespace := c.Query("namespace")
		requestPodName := c.Query("podname")
		datas := clusterinfo.GetPodInfo(clientset, namespace, requestPodName)
		// clusterinfo.SavePodInfo(datas)
		loc, _ := time.LoadLocation("UTC")

		c.JSON(http.StatusOK, gin.H{
			"Current Time": time.Now(),
			"UTC Time":     time.Now().In(loc),
			"Pod Name":     datas.PodName,
			"Pod Log":      datas.PodLog,
		})

		// c.Redirect(http.StatusMovedPermanently, "/results")

	})
	router.Run(":" + os.Getenv("PORT"))

}
