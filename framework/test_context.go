package framework

import (
	"flag"
)

type TestContextType struct {
	// Provider identifies the infrastructure provider (aws, azure, kvm)
	Provider string

	CloudConfig CloudConfig
}

// CloudConfig holds the cloud configuration for e2e test suites.
type CloudConfig struct {
}

// ProviderConfig holds the provider configuration for e2e test suites.
type ProviderConfig struct {
	Provider ProviderInterface
}

// TestContext should be used by all tests to access common context data.
var TestContext TestContextType

// RegisterClusterFlags registers flags specific to the cluster e2e test suite.
func RegisterClusterFlags(flags *flag.FlagSet) {
}

// RegisterProviderFlags registers flags specific to the provider e2e test suite.
func RegisterProviderFlags(flags *flag.FlagSet) {
	flags.StringVar(&TestContext.Provider, "provider", "", "The name of the Kubernetes provider (aws, azure, kvm)")
}
