package e2e

import (
	"flag"
	"os"
	"testing"

	"github.com/giantswarm/conformance/framework"
)

func handleFlags() {
	framework.RegisterCommonFlags(flag.CommandLine)
	framework.RegisterProviderFlags(flag.CommandLine)
	flag.Parse()
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
