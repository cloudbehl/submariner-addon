package v1alpha1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_ManagedClusterInfo = map[string]string{
	"clusterName": "ClusterName represents the name of the managed cluster.",
	"vendor":      "Vendor represents the kubernetes vendor of the managed cluster.",
	"platform":    "Platform represents the cloud provider of the managed cluster.",
	"region":      "Region represents the cloud region of the managed cluster.",
	"infraId":     "InfraId represents the infrastructure id of the managed cluster.",
}

func (ManagedClusterInfo) SwaggerDoc() map[string]string {
	return map_ManagedClusterInfo
}

var map_SubmarinerConfig = map[string]string{
	"":       "SubmarinerConfig represents the configuration for Submariner, the submariner-addon will use it to configure the Submariner.",
	"spec":   "Spec defines the configuration of the Submariner",
	"status": "Status represents the current status of submariner configuration",
}

func (SubmarinerConfig) SwaggerDoc() map[string]string {
	return map_SubmarinerConfig
}

var map_SubmarinerConfigList = map[string]string{
	"":         "SubmarinerConfigList is a collection of SubmarinerConfig.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"items":    "Items is a list of SubmarinerConfig.",
}

func (SubmarinerConfigList) SwaggerDoc() map[string]string {
	return map_SubmarinerConfigList
}

var map_SubmarinerConfigSpec = map[string]string{
	"":                   "SubmarinerConfigSpec describes the configuration of the Submariner",
	"cableDriver":        "CableDriver represents the submariner cable driver implementation. Available options are libreswan (default) strongswan, and wireguard.",
	"IPSecIKEPort":       "IPSecIKEPort represents IPsec IKE port (default 500).",
	"IPSecNATTPort":      "IPSecNATTPort represents IPsec NAT-T port (default 4500).",
	"credentialsSecret":  "CredentialsSecret is a reference to the secret with a certain cloud platform credentials, the supported platform includes AWS, GCP, Azure, ROKS and OSD. The submariner-addon will use these credentials to prepare Submariner cluster environment. If the submariner cluster environment requires submariner-addon preparation, this field should be specified.",
	"subscriptionConfig": "SubscriptionConfig represents a Submariner subscription. SubscriptionConfig can be used to customize the Submariner subscription.",
	"imagePullSpecs":     "ImagePullSpecs represents the desired images of submariner components installed on the managed cluster. If not specified, the default submariner images that was defined by submariner operator will be used.",
}

func (SubmarinerConfigSpec) SwaggerDoc() map[string]string {
	return map_SubmarinerConfigSpec
}

var map_SubmarinerConfigStatus = map[string]string{
	"":                   "SubmarinerConfigStatus represents the current status of submariner configuration.",
	"conditions":         "Conditions contain the different condition statuses for this configuration.",
	"managedClusterInfo": "ManagedClusterInfo represents the information of a managed cluster.",
}

func (SubmarinerConfigStatus) SwaggerDoc() map[string]string {
	return map_SubmarinerConfigStatus
}

var map_SubmarinerImagePullSpecs = map[string]string{
	"submarinerImagePullSpec":           "SubmarinerImagePullSpec represents the desired image of submariner.",
	"lighthouseAgentImagePullSpec":      "LighthouseAgentImagePullSpec represents the desired image of the lighthouse agent.",
	"lighthouseCoreDNSImagePullSpec":    "LighthouseCoreDNSImagePullSpec represents the desired image of lighthouse coredns.",
	"submarinerRouteAgentImagePullSpec": "SubmarinerRouteAgentImagePullSpec represents the desired image of the submariner route agent.",
}

func (SubmarinerImagePullSpecs) SwaggerDoc() map[string]string {
	return map_SubmarinerImagePullSpecs
}

var map_SubscriptionConfig = map[string]string{
	"":                "SubscriptionConfig contains configuration specified for a submariner subscription.",
	"source":          "Source represents the catalog source of a submariner subscription. The default value is redhat-operators",
	"sourceNamespace": "SourceNamespace represents the catalog source namespace of a submariner subscription. The default value is openshift-marketplace",
	"channel":         "Channel represents the channel of a submariner subscription. The default value is alpha",
	"startingCSV":     "StartingCSV represents the startingCSV of a submariner subscription. The default value is submariner.v0.8",
}

func (SubscriptionConfig) SwaggerDoc() map[string]string {
	return map_SubscriptionConfig
}

// AUTO-GENERATED FUNCTIONS END HERE
