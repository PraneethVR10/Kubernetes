package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	//fetching kubeconfig file
	home, _ := os.UserHomeDir() // fetch the current home directory /home/praneeth

	kubeCon := filepath.Join(home, ".kube/config") //join the current home directory and .kube/config path

	//providing context for the cluster

	config, err := clientcmd.BuildConfigFromFlags("", kubeCon) //we use empty string cuz it allows the function to first consider deafault locations and environment variables, then priotize the specific path that we provide

	if err != nil {
		panic(err)
	}

	//create a ClientSet

	client := kubernetes.NewForConfigOrDie(config) // this method creates a new client with the config data from above and handles the error by itself

	// Below method created a new client taking config as an argument and we have to handle it's error.
	/*client , err := kubernetes.NewForConfig(config)

	if err != nil {
		panic(err)
	} */

	pods, err := client.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{}) //The CoreV1().Pods("default") part creates an interface to  interact with pods in the "default" namespace. The List method is called on this interface, with context.Background() to manage the request context and metav1.ListOptions{} for any listing options (none specified here).
	if err != nil {
		panic(err)
	}

	for i, pod := range pods.Items {

		fmt.Printf("Name of the pods %d: %s\n,", i, pod.Name)

	}

}
