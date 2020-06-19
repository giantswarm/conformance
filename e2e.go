package e2e

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/giantswarm/conformance/framework"

	// ensure that cloud providers are loaded
	_ "github.com/giantswarm/conformance/framework/providers/aws"
	_ "github.com/giantswarm/conformance/framework/providers/azure"
	_ "github.com/giantswarm/conformance/framework/providers/kvm"
)

func RunE2ETests(t *testing.T) {
	logrus.Printf("Running e2e tests for provider '%s'", framework.TestContext.Provider)
	p, err := framework.SetupProviderConfig(framework.TestContext.Provider)
	if err != nil {
		logrus.Error(err)
	}
	switch os.Args[1] {
	case "create-cluster":
		p.CreateTenantCluster()
	case "update-cluster":
	case "delete-cluster":
	default:
		t.Fatalf("expected 'create-cluster', 'update-cluster' or 'delete-cluster'")
	}
}
