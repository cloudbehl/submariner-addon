package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope="Namespaced"

// SubmarinerConfig represents the configuration for Submariner, the submariner-addon will use it
// to configure the Submariner.
type SubmarinerConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the configuration of the Submariner
	Spec SubmarinerConfigSpec `json:"spec"`

	// Status represents the current status of submariner configuration
	// +optional
	Status SubmarinerConfigStatus `json:"status,omitempty"`
}

// SubmarinerConfigSpec describes the configuration of the Submariner
type SubmarinerConfigSpec struct {
	// CableDriver represents the submariner cable driver implementation.
	// Available options are libreswan (default) strongswan, and wireguard.
	// +optional
	CableDriver string `json:"cableDriver,omitempty"`

	// IPSecIKEPort represents IPsec IKE port (default 500).
	// +optional
	IPSecIKEPort int `json:"IPSecIKEPort,omitempty"`

	// IPSecNATTPort represents IPsec NAT-T port (default 4500).
	// +optional
	IPSecNATTPort int `json:"IPSecNATTPort,omitempty"`

	// CredentialsSecret is a reference to the secret with a certain cloud platform
	// credentials, the supported platform includes AWS, GCP, Azure, ROKS and OSD.
	// The submariner-addon will use these credentials to prepare Submariner cluster
	// environment. If the submariner cluster environment requires submariner-addon
	// preparation, this field should be specified.
	// +optional
	CredentialsSecret *corev1.LocalObjectReference `json:"credentialsSecret,omitempty"`

	// SubscriptionConfig represents a Submariner subscription. SubscriptionConfig
	// can be used to customize the Submariner subscription.
	// +optional
	SubscriptionConfig `json:"subscriptionConfig,omitempty"`

	// ImagePullSpecs represents the desired images of submariner components installed on the managed cluster.
	// If not specified, the default submariner images that was defined by submariner operator will be used.
	// +optional
	ImagePullSpecs SubmarinerImagePullSpecs `json:"imagePullSpecs,omitempty"`
}

// SubscriptionConfig contains configuration specified for a submariner subscription.
type SubscriptionConfig struct {
	// Source represents the catalog source of a submariner subscription.
	// The default value is redhat-operators
	// +optional
	Source string `json:"source,omitempty"`

	// SourceNamespace represents the catalog source namespace of a submariner subscription.
	// The default value is openshift-marketplace
	// +optional
	SourceNamespace string `json:"sourceNamespace,omitempty"`

	// Channel represents the channel of a submariner subscription.
	// The default value is alpha
	// +optional
	Channel string `json:"channel,omitempty"`

	// StartingCSV represents the startingCSV of a submariner subscription.
	// The default value is submariner.v0.8
	// +optional
	StartingCSV string `json:"startingCSV,omitempty"`
}

type SubmarinerImagePullSpecs struct {
	// SubmarinerImagePullSpec represents the desired image of submariner.
	// +optional
	SubmarinerImagePullSpec string `json:"submarinerImagePullSpec"`

	// LighthouseAgentImagePullSpec represents the desired image of the lighthouse agent.
	// +optional
	LighthouseAgentImagePullSpec string `json:"lighthouseAgentImagePullSpec"`

	// LighthouseCoreDNSImagePullSpec represents the desired image of lighthouse coredns.
	// +optional
	LighthouseCoreDNSImagePullSpec string `json:"lighthouseCoreDNSImagePullSpec"`

	// SubmarinerRouteAgentImagePullSpec represents the desired image of the submariner route agent.
	// +optional
	SubmarinerRouteAgentImagePullSpec string `json:"submarinerRouteAgentImagePullSpec"`
}

const (
	// SubmarinerConfigConditionApplied means the configuration has successfully
	// applied.
	SubmarinerConfigConditionApplied string = "SubmarinerConfigApplied"

	// SubmarinerConfigConditionEnvPrepared means the submariner cluster environment
	// is prepared on a specfied cloud platform with the given cloud platform credentials.
	SubmarinerConfigConditionEnvPrepared string = "SubmarinerClusterEnvironmentPrepared"
)

// SubmarinerConfigStatus represents the current status of submariner configuration.
type SubmarinerConfigStatus struct {
	// Conditions contain the different condition statuses for this configuration.
	Conditions []metav1.Condition `json:"conditions"`
	// ManagedClusterInfo represents the information of a managed cluster.
	// +optional
	ManagedClusterInfo ManagedClusterInfo `json:"managedClusterInfo,omitempty"`
}

type ManagedClusterInfo struct {
	// ClusterName represents the name of the managed cluster.
	// +optional
	ClusterName string `json:"clusterName,omitempty"`
	// Vendor represents the kubernetes vendor of the managed cluster.
	// +optional
	Vendor string `json:"vendor,omitempty"`
	// Platform represents the cloud provider of the managed cluster.
	// +optional
	Platform string `json:"platform,omitempty"`
	// Region represents the cloud region of the managed cluster.
	// +optional
	Region string `json:"region,omitempty"`
	// InfraId represents the infrastructure id of the managed cluster.
	// +optional
	InfraId string `json:"infraId,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SubmarinerConfigList is a collection of SubmarinerConfig.
type SubmarinerConfigList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items is a list of SubmarinerConfig.
	Items []SubmarinerConfig `json:"items"`
}
