package aws

import (
	"github.com/giantswarm/conformance/framework"
	"github.com/sirupsen/logrus"
)

// Provider is a structure to handle AWS cloud for e2e testing
type Provider struct {
	framework.Provider
}

func init() {
	framework.RegisterProvider("aws", newProvider)
}

func newProvider() (framework.ProviderInterface, error) {
	return &Provider{}, nil
}

func (p *Provider) Create() error {
	logrus.Println("AWS provider")
	return nil
}
func (p *Provider) Update() error {
	return nil
}
func (p *Provider) Delete() error {
	return nil
}
