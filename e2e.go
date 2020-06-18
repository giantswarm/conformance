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

//func createClient() kubernetes.Interface {
//	kubeconfig := os.Getenv("KUBECONFIG")
//
//	var cfg *rest.Config
//	var err error
//	if kubeconfig != "" {
//		cfg, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
//	} else {
//		cfg, err = rest.InClusterConfig()
//	}
//	if err != nil {
//		panic(err)
//	}
//
//	kubeClient, err := kubernetes.NewForConfig(cfg)
//	if err != nil {
//		panic(err)
//	}
//	return kubeClient
//}
//
