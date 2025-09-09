// +k8s:deepcopy-gen=package
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Define GroupVersion for your CRD
var GroupVersion = schema.GroupVersion{
	Group:   "mygroup.example.com",
	Version: "v1alpha1",
}

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
	Items           []Memcached `json:"items"`
}

func (in *Memcached) DeepCopy() *Memcached {
	if in == nil {
		return nil
	}
	out := new(Memcached)
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = *in.ObjectMeta.DeepCopy() // use ObjectMeta.DeepCopy
	out.Spec = in.Spec                         // shallow copy is fine if no pointers
	out.Status = in.Status                     // shallow copy is fine
	return out
}

func (in *Memcached) DeepCopyObject() runtime.Object {
	if in == nil {
		return nil
	}
	out := new(Memcached)
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = *in.ObjectMeta.DeepCopy()
	out.Spec = in.Spec
	out.Status = in.Status
	return out
}

func (in *MemcachedList) DeepCopyObject() runtime.Object {
	if in == nil {
		return nil
	}
	out := new(MemcachedList)
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		out.Items = make([]Memcached, len(in.Items))
		for i := range in.Items {
			out.Items[i] = *in.Items[i].DeepCopy()
		}
	}
	return out
}

var (
	SchemeGroupVersion = GroupVersion
	SchemeBuilder      = runtime.NewSchemeBuilder(addKnownTypes)
)

func AddToScheme(scheme *runtime.Scheme) error {
	return SchemeBuilder.AddToScheme(scheme)
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Memcached{},
		&MemcachedList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
