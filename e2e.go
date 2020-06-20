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
		t.Fatal(err)
	}

	c, err := framework.LoadClient()
	if err != nil {
		t.Fatal(err)
	}

	switch os.Args[1] {
	case "create-cluster":
		err := p.CreateTenantCluster(c)
		if err != nil {
			t.Fatal(err)
		}
	case "update-cluster":
	case "delete-cluster":
	default:
		t.Fatalf("expected 'create-cluster', 'update-cluster' or 'delete-cluster'")
	}
}
