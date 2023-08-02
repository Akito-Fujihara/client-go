package main

import (
	"context"
	"fmt"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	defaultKubeConfigPath := filepath.Join(homedir.HomeDir(), ".kube", "config")

	config, _ := clientcmd.BuildConfigFromFlags("", defaultKubeConfigPath)

	clientset, _ := kubernetes.NewForConfig(config)

	pods, _ := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})

	fmt.Println("NAMESPACE\tNAME")
	for _, pod := range pods.Items {
		fmt.Printf("%s\t%s\n", pod.GetNamespace(), pod.GetName())
	}
}
