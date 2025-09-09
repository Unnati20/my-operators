package main

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/unnati20/my-operators/operators/memcached/api/v1alpha1"
	"github.com/unnati20/my-operators/operators/memcached/controller"
	"k8s.io/apimachinery/pkg/runtime"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func main() {
	fmt.Println("Starting DIY Memcached Operator...")

	// 1️⃣ Kubernetes config (in-cluster or fallback)
	cfg, err := rest.InClusterConfig()
	if err != nil {
		cfg, err = config.GetConfig() // uses local kubeconfig
		if err != nil {
			panic(err)
		}
	}

	// 2️⃣ client-go clientset
	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}

	// 3️⃣ controller-runtime client
	scheme := runtime.NewScheme()
	_ = v1.AddToScheme(scheme)

	kubeClient, err := client.New(cfg, client.Options{Scheme: scheme})
	if err != nil {
		panic(err)
	}

	// 4️⃣ Reconcile all Memcached CRs in the default namespace
	for {
		memList := &v1.MemcachedList{}
		err := kubeClient.List(context.Background(), memList, client.InNamespace("default"))
		if err != nil {
			fmt.Println("Error listing Memcacheds:", err)
		} else {
			for i := range memList.Items {
				mem := &memList.Items[i]
				err := controller.ReconcileMemcached(mem, kubeClient, clientset)
				if err != nil {
					fmt.Println("Error reconciling Memcached:", err)
				}
			}
		}
		time.Sleep(15 * time.Second)
	}
}
