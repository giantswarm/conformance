package kvm

import (
	"github.com/giantswarm/conformance/framework"
)

// Provider is a structure to handle KVM for e2e testing
type Provider struct {
	framework.Provider
}

func init() {
	framework.RegisterProvider("kvm", newProvider)
}

func newProvider() (framework.ProviderInterface, error) {
	return &framework.Provider{}, nil
}
