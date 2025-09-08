package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//This defines a CRD called Memcached with a Size spec

type MemcachedSpec struct {
	Size int32 `json:"size"`
}

type MemcachedStatus struct {
	Nodes []string `json:"nodes,omitempty"`
}

type Memcached struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MemcachedSpec   `json:"spec,omitempty"`
	Status MemcachedStatus `json:"status,omitempty"`
}

type MemcachedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Memcached `json:"items:omitempty"`
}
