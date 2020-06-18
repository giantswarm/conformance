package framework

import (
	"flag"
)

type TestContextType struct {
	ClusterName    string
	Owner          string
	PodCidr        string
	Release        string
	Provider       string
	ProviderConfig ProviderConfig
}

// ProviderConfig holds the provider configuration for e2e test suites.
type ProviderConfig struct {
	Credential   string
	Domain       string
	ExternalSNAT bool
	MasterZone   string
	Provider     ProviderInterface
	Region       string
}

// TestContext should be used by all tests to access common context data.
var TestContext TestContextType

// RegisterCommonFlags registers flags common to the cluster e2e test suite.
func RegisterCommonFlags(flags *flag.FlagSet) {
	flags.StringVar(&TestContext.ClusterName, "cluster-name", "", "The name of the Kubernetes cluster")
	flags.StringVar(&TestContext.Owner, "owner", "", "The owner of the Kubernetes cluster")
	flags.StringVar(&TestContext.PodCidr, "pod-cidr", "", "The pod cidr IP range")
	flags.StringVar(&TestContext.Release, "release", "", "The release version")
}

// RegisterProviderFlags registers flags specific to the provider e2e test suite.
func RegisterProviderFlags(flags *flag.FlagSet) {
	flags.StringVar(&TestContext.Provider, "provider", "", "The name of the Kubernetes provider (aws, azure, kvm)")
	flags.StringVar(&TestContext.ProviderConfig.Credential, "credential", "", "The provider credential, e.g. AWS cloud credentials that point to the AWS account used to spin up the cluster resources")
	flags.StringVar(&TestContext.ProviderConfig.Domain, "domain", "", "The domain name, e.g. 'gauss.eu-central-1.aws.gigantic.io'")
	flags.BoolVar(&TestContext.ProviderConfig.ExternalSNAT, "external-snat", false, "Enable external SNAT")
	flags.StringVar(&TestContext.ProviderConfig.Domain, "master-zone", "", "Zone of the Kubernetes master, e.g. eu-central-1a")
	flags.StringVar(&TestContext.ProviderConfig.Region, "region", "", "Region of the Kubernetes cluster, e.g. eu-central-1")
}
