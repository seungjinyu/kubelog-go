package main

import (
	"net/http"

	"github.com/seungjinyu/kubelog_go/clusterinfo"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	clientset, err := clusterinfo.CreateOutClientSet()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/welcome")
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

		namespace := ""
		requestPodName := "cvpammanager"
		datas := clusterinfo.GetPodInfo(clientset, namespace, requestPodName)
		// clusterinfo.SavePodInfo(datas)
		c.JSON(http.StatusOK, gin.H{
			"Pod Name": datas.PodName,
			"Pod Log":  datas.PodLog,
		})

		// c.Redirect(http.StatusMovedPermanently, "/results")

	})

	http.ListenAndServe(":8080", router)
}
