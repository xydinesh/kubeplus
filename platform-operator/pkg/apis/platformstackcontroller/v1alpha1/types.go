package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PlatformStack is specification for a PlatformStack resource
// +k8s:openapi-gen=true
type PlatformStack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PlatformStackSpec   `json:"spec"`
	Status PlatformStackStatus `json:"status"`
}

// PlatformStackSpec is the spec for a PlatformStack resource
// +k8s:openapi-gen=true
type PlatformStackSpec struct {
	// LabelSelector that selects resources of that label
	LabelSelector map[string]string `json:"labelSelector,omitempty"` 
	// List of stack elements that forms this Platform Stack
	StackElements []StackElements `json:"stackElements"`
}

type StackElements struct {
	// Name of Kubernetes Kind - could be Built-in Kind or Custom Kind
	Kind string `json:"kind"`
	// Name of the Resource
	Name string `json:"name"`
	// Namespace of the Resource
	// If not specified then 'default' Namespace is assumed
	// +optional
	Namespace string `json:"namespace"`
	// +optional
	DependsOn []DependsOn `json:"dependsOn,omitempty"`
}

type DependsOn struct {
	// Name of the Resource that the Resource in which dependsOn is included actually dependsOn
	Name string `json:"name"`
}

// PlatformStackStatus is the status for a PlatformStack resource.
// +k8s:openapi-gen=true
type PlatformStackStatus struct {
	Status             string   `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// PlatformStackList is a list of PlatformStack resources
type PlatformStackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []PlatformStack `json:"items"`
}