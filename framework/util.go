package framework

import (
	"fmt"

	infrastructurev1alpha2 "github.com/giantswarm/apiextensions/pkg/apis/infrastructure/v1alpha2"
	releasev1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/release/v1alpha1"
	"github.com/giantswarm/k8sclient/v3/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	apiv1alpha2 "sigs.k8s.io/cluster-api/api/v1alpha2"
)

// LoadClient returns k8sclient for connecting to kubernetes clusters.
func LoadClient() (k8sclient.Interface, error) {
	var client k8sclient.Interface
	var restConfig *rest.Config
	var err error

	if TestContext.KubeConfig != "" {
		restConfig, err = clientcmd.BuildConfigFromFlags("", TestContext.KubeConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to load key kubeconfig: %v", err)
		}
	} else {
		restConfig, err = rest.InClusterConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to load key kubeconfig: %v", err)
		}
	}
	newLogger, err := micrologger.New(micrologger.Config{})
	if err != nil {
		return nil, err
	}
	c := k8sclient.ClientsConfig{
		SchemeBuilder: k8sclient.SchemeBuilder{
			apiv1alpha2.AddToScheme,
			infrastructurev1alpha2.AddToScheme,
			releasev1alpha1.AddToScheme,
		},
		Logger: newLogger,

		RestConfig: restConfig,
	}

	client, err = k8sclient.NewClients(c)
	if err != nil {
		return nil, microerror.Mask(err)
	}
	return client, nil
}
