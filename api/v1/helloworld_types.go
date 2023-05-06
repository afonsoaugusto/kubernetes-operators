package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HelloWorldSpec define o estado desejado de HelloWorld
type HelloWorldSpec struct {
	Message string `json:"message,omitempty"`
}

// HelloWorldStatus define o estado observado de HelloWorld
type HelloWorldStatus struct {
	// INSERIR CAMPO DE STATUS ADICIONAL - definir o estado observado do cluster
	// Importante: execute "make" para regenerar o código após modificar este arquivo
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// HelloWorld é o esquema para a API helloworlds
type HelloWorld struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HelloWorldSpec   `json:"spec,omitempty"`
	Status HelloWorldStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HelloWorldList contém uma lista de HelloWorld
type HelloWorldList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HelloWorld `json:"items"`
}