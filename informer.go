package main

import (
	"log"
	"path/filepath"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	KubeConfigPath := filepath.Join(homedir.HomeDir(), ".kube", "config")

	config, _ := clientcmd.BuildConfigFromFlags("", KubeConfigPath)

	clientset, _ := kubernetes.NewForConfig(config)

	informerFactory := informers.NewSharedInformerFactory(clientset, time.Second*30)

	podInformer := informerFactory.Core().V1().Pods()

	podInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc:    AddFuncPod,
			UpdateFunc: UpdateFuncPod,
			DeleteFunc: DeleteFuncPod,
		},
	)

	StopCh := make(chan struct{})
	informerFactory.Start(wait.NeverStop)

	go wait.Until(func() {
		log.Println("------- running --------")
	}, time.Second*10, StopCh)
	<-StopCh
}

func AddFuncPod(obj interface{}) {
	if pod_name, err := cache.MetaNamespaceKeyFunc(obj); err != nil {
		log.Printf("Error Get Pod Obj: %s\n", err.Error())
	} else {
		log.Println("AddFunc Pod:", pod_name)
	}
}

func UpdateFuncPod(old, new interface{}) {
	if pod_name, err := cache.MetaNamespaceKeyFunc(new); err != nil {
		log.Printf("Error Get Pod Obj: %s\n", err.Error())
	} else {
		log.Println("UpdateFunc Pod:", pod_name)
	}
}

func DeleteFuncPod(obj interface{}) {
	if pod_name, err := cache.MetaNamespaceKeyFunc(obj); err != nil {
		log.Printf("Error Get Pod Obj: %s\n", err.Error())
	} else {
		log.Println("DeleteFunc Pod:", pod_name)
	}
}
