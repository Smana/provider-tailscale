/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AutorizationObservation struct {

	// Whether or not the device is authorized
	Authorized *bool `json:"authorized,omitempty" tf:"authorized,omitempty"`

	// The device to set as authorized
	DeviceID *string `json:"deviceId,omitempty" tf:"device_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AutorizationParameters struct {

	// Whether or not the device is authorized
	// +kubebuilder:validation:Optional
	Authorized *bool `json:"authorized,omitempty" tf:"authorized,omitempty"`

	// The device to set as authorized
	// +kubebuilder:validation:Optional
	DeviceID *string `json:"deviceId,omitempty" tf:"device_id,omitempty"`
}

// AutorizationSpec defines the desired state of Autorization
type AutorizationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AutorizationParameters `json:"forProvider"`
}

// AutorizationStatus defines the observed state of Autorization.
type AutorizationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AutorizationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Autorization is the Schema for the Autorizations API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tailscale}
type Autorization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.authorized)",message="authorized is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.deviceId)",message="deviceId is a required parameter"
	Spec   AutorizationSpec   `json:"spec"`
	Status AutorizationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AutorizationList contains a list of Autorizations
type AutorizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Autorization `json:"items"`
}

// Repository type metadata.
var (
	Autorization_Kind             = "Autorization"
	Autorization_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Autorization_Kind}.String()
	Autorization_KindAPIVersion   = Autorization_Kind + "." + CRDGroupVersion.String()
	Autorization_GroupVersionKind = CRDGroupVersion.WithKind(Autorization_Kind)
)

func init() {
	SchemeBuilder.Register(&Autorization{}, &AutorizationList{})
}
