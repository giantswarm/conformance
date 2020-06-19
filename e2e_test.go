package e2e

import (
	"flag"
	"os"
	"testing"

	"github.com/giantswarm/conformance/framework"
)

func handleFlags() {
	createCmd := flag.NewFlagSet("create-cluster", flag.ExitOnError)
	framework.RegisterCommonFlags(createCmd)
	framework.RegisterProviderFlags(createCmd)
	switch os.Args[1] {
	case "create-cluster":
		createCmd.Parse(os.Args[2:])
	case "update-cluster":
	case "delete-cluster":
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}

//TestMain is the entrypoint for running e2e tests
func TestMain(m *testing.M) {
	// Register test flags, then parse flags.
	handleFlags()
	os.Exit(m.Run())
}

//TestE2E executes the real tests
func TestE2E(t *testing.T) {
	RunE2ETests(t)
}
