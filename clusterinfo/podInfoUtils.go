package clusterinfo

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/seungjinyu/kubelog_go/models"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetPodListInfo(clientset *kubernetes.Clientset) models.PodInfoList {

	podlist := givePodList(clientset)
	result := extractDataFromPodList(podlist, clientset)

	return result
}

func GetPodInfo(clientset *kubernetes.Clientset, namespace, requestPodName string) models.PodInfo {

	podInfo, err := givePod(clientset, namespace, requestPodName)
	if err != nil {
		panic(err)
	}
	result := extractDataFromPod(podInfo, clientset)

	return result
}

// GivePodList gives backs the pod instance of the cluster by the kubernetes config file
func givePodList(clientset *kubernetes.Clientset) []v1.Pod {

	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods \n", len(pods.Items))

	items := pods.Items
	// result := model.PodInfo{}
	return items

}

func givePod(clientset *kubernetes.Clientset, namespace, requestPodName string) (v1.Pod, error) {

	if namespace == "" {
		namespace = "default"
	}
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, v := range pods.Items {
		if strings.Contains(v.Name, requestPodName) {
			return v, nil
		}
	}

	return v1.Pod{}, err
}

// Here is what we came up with eventually using client-go library:
func getPodLogs(pod v1.Pod, clientset *kubernetes.Clientset) string {
	podLogOpts := v1.PodLogOptions{}

	// clientset, err := CreateOutClientSet()
	req := clientset.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, &podLogOpts)
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		return "error in opening stream"
	}
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		return "error in copy information from podLogs to buf"
	}
	str := buf.String()

	return str
}

// ExtracDataFromPodList extracts data from the pod list
func extractDataFromPodList(pl []v1.Pod, clientset *kubernetes.Clientset) models.PodInfoList {

	var tmp models.PodInfoList
	repl := pl
	tmp.InfoList = make([]models.PodInfo, len(repl))

	fmt.Println(len(repl))

	for i, value := range repl {
		tmp.InfoList[i] = models.PodInfo{PodName: value.GetName(), PodLog: getPodLogs(value, clientset), PodStatus: value.Status.Message}
		// tmp.InfoList[i] = models.PodInfo{PodName: value.GetName(), PodLog: "nil", PodStatus: value.Status.Message}
	}
	return tmp
}

func extractDataFromPod(pi v1.Pod, clientset *kubernetes.Clientset) models.PodInfo {

	tmp := models.PodInfo{
		PodName:   pi.GetName(),
		PodLog:    getPodLogs(pi, clientset),
		PodStatus: pi.Status.Message,
	}

	fmt.Println(tmp.PodLog, tmp.PodName)
	return tmp

}
func SavePodInfoList(pil models.PodInfoList) {

	for _, v := range pil.InfoList {

		fileName := "./logs/" + v.PodName + ".log"
		tmp, err := os.Create(fileName)
		if err != nil {
			log.Println(err)
		}
		contents := v.PodName + "\n" + v.PodLog
		tmp.WriteString(contents)

		defer tmp.Close()
	}

}

func SavePodInfo(pi models.PodInfo) {

	fileName := "./logs/" + pi.PodName + "log"
	tmp, err := os.Create(fileName)
	if err != nil {
		log.Println(err)
	}
	contents := pi.PodName + "\n" + pi.PodLog
	tmp.WriteString(contents)

	defer tmp.Close()

}
