package services

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seungjinyu/kubelog-go/clusterinfo"
	"k8s.io/client-go/kubernetes"
)

type rbody struct {
	Namespace string `json:"namespace"`
	PodName   string `json:"podname"`
}

func Healthy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}

func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"greetings": "welcome",
	})
}
func RedirectToWelcome(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/welcome")
}

func V1welcome(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"auth": true,
	})
}

func Getpod(c *gin.Context) {

	var rbodyi rbody
	body := c.Request.Body

	err := json.NewDecoder(body).Decode(&rbodyi)
	defer body.Close()

	if err != nil {
		log.Println(err)
	}

	datas := clusterinfo.GetPodInfo(c.Keys["clientset"].(*kubernetes.Clientset), rbodyi.Namespace, rbodyi.PodName)
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

func Getpods(c *gin.Context) {
	datas := clusterinfo.GetPodListInfo(c.Keys["clientset"].(*kubernetes.Clientset))
	clusterinfo.SavePodInfoList(datas)
	c.JSON(http.StatusOK, gin.H{
		"datas": "Sending completed",
	})
}
