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
	// Fetching kubeconfig file
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	kubeCon := filepath.Join(home, ".kube", "config")

	// Providing context for the cluster
	config, err := clientcmd.BuildConfigFromFlags("", kubeCon)
	if err != nil {
		panic(err)
	}

	// Create a ClientSet
	client := kubernetes.NewForConfigOrDie(config)

	// List pods in the "default" namespace
	pods, err := client.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	// Print pod names
	for i, pod := range pods.Items {
		fmt.Printf("Name of the pod %d: %s\n", i+1, pod.Name)
	}
}
