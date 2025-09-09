package controller

import (
	"context"
	"fmt"

	v1 "github.com/unnati20/my-operators/operators/memcached/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime/pkg/client"
)

// ReconcileMemcached ensures Pods match Spec.Size and updates Status.Nodes
func ReconcileMemcached(mem *v1.Memcached, kubeClient ctrl.Client, clientset *kubernetes.Clientset) error {
	// Create a deep copy of the Memcached CR to avoid modifying the cache object
	memCopy := mem.DeepCopy()

	namespace := memCopy.Namespace
	createdPods := []string{}

	for i := int32(0); i < memCopy.Spec.Size; i++ {
		podName := fmt.Sprintf("%s-%d", memCopy.Name, i)

		pod := &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      podName,
				Namespace: namespace,
				Labels:    map[string]string{"app": memCopy.Name},
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{Name: "memcached", Image: "memcached:1.6.6"},
				},
			},
		}

		// Check if pod exists
		_, err := clientset.CoreV1().Pods(namespace).Get(context.Background(), podName, metav1.GetOptions{})
		if err != nil {
			_, err := clientset.CoreV1().Pods(namespace).Create(context.Background(), pod, metav1.CreateOptions{})
			if err != nil {
				fmt.Println("Error creating pod:", err)
			} else {
				fmt.Println("Created pod:", podName)
			}
		}

		createdPods = append(createdPods, podName)
	}

	// Update Status.Nodes
	memCopy.Status.Nodes = createdPods
	if err := kubeClient.Status().Update(context.Background(), memCopy); err != nil {
		fmt.Println("Error updating Memcached status:", err)
		return err
	} else {
		fmt.Println("Updated Memcached status:", createdPods)
	}

	return nil
}
