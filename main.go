package main

import (
	"net/http"

	"github.com/seungjinyu/kubelog_go/clusterinfo"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

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
			"status": "health",
		})
	})

	router.GET("/getpods", func(c *gin.Context) {

		datas := clusterinfo.GetPodListInfo()
		clusterinfo.SavePodInfoList(datas)
		c.JSON(http.StatusOK, gin.H{
			"datas": "Sending completed",
		})

		// c.Redirect(http.StatusMovedPermanently, "/results")

	})

	http.ListenAndServe(":8080", router)
}
