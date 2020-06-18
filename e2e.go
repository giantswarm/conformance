package e2e

import (
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/giantswarm/conformance/framework"

	// ensure that cloud providers are loaded
	_ "github.com/giantswarm/conformance/framework/providers/aws"
	_ "github.com/giantswarm/conformance/framework/providers/azure"
	_ "github.com/giantswarm/conformance/framework/providers/kvm"
)

func RunE2ETests(t *testing.T) {
	p, err := framework.SetupProviderConfig(framework.TestContext.Provider)
	if err != nil {
		logrus.Error(err)
	}
	p.Create()
	logrus.Println("Blob")
}
