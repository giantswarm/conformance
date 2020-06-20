package azure

import (
	"github.com/giantswarm/conformance/framework"
)

// Provider is a structure to handle Azure cloud for e2e testing
type Provider struct {
	framework.Provider
}

func init() {
	framework.RegisterProvider("azure", newProvider)
}

func newProvider() (framework.ProviderInterface, error) {
	return &framework.Provider{}, nil
}
