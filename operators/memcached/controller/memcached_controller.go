package controller

import (
	"context"
	"fmt"

	v1 "github.com/unnati20/my-operators/operators/memcache/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func ReconcileMemcached(mem *v1.Memcached) {
	// connect to cluster

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	for i := int32(0); i < mem.Spec.Size; i++ {
		podName := fmt.Sprintf("%s-%d", mem.Name, i)
		pod := &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      podName,
				Namespace: mem.Namespace,
				Labels: map[string]string{
					"app": "memcached",
				},
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name:  "memcached",
						Image: "memcached:1.6.6",
					},
				},
			},
		}
		_, err := clientset.CoreV1().Pods(mem.Namespace).Create(context.Background(), pod, metav1.CreateOptions{})
		if err != nil {
			fmt.Println("Pod creation error:", err)
		} else {
			fmt.Println("Created pod:", podName)
		}
	}
}
