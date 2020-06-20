package aws

import (
	"encoding/json"
	"strings"

	"github.com/giantswarm/conformance/framework"
	"github.com/giantswarm/k8sclient/v3/pkg/k8sclient"
	"github.com/giantswarm/kubectl-gs/pkg/template/cluster"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

func (p *Provider) CreateTenantCluster(c k8sclient.Interface) error {
	masterAZ := strings.Split(framework.TestContext.ProviderConfig.MasterZone, ",")
	config := cluster.Config{
		Credential:     framework.TestContext.ProviderConfig.Credential,
		Domain:         framework.TestContext.ProviderConfig.Domain,
		ExternalSNAT:   framework.TestContext.ProviderConfig.ExternalSNAT,
		MasterAZ:       masterAZ,
		Name:           framework.TestContext.ClusterName,
		PodsCIDR:       framework.TestContext.PodCidr,
		Owner:          framework.TestContext.Owner,
		Region:         framework.TestContext.ProviderConfig.Region,
		ReleaseVersion: framework.TestContext.Release,
	}

	clusterCR, awsClusterCR, g8sControlPlaneCR, aWSControlPlaneCR, err := cluster.NewClusterCRs(config)
	if err != nil {
		return err
	}
	logrus.Println("Dumping CR's:")

	logrus.Println(prettyPrint(clusterCR))
	logrus.Println(prettyPrint(awsClusterCR))
	logrus.Printf(prettyPrint(g8sControlPlaneCR))
	logrus.Printf(prettyPrint(aWSControlPlaneCR))

	pods, _ := c.K8sClient().CoreV1().Pods("giantswarm").List(metav1.ListOptions{})
	for _, pod := range pods.Items {
		logrus.Print(pod.Name)
	}

	return nil
}
func (p *Provider) UpdateTenantCluster() error {
	return nil
}
func (p *Provider) DeleteTenantCluster() error {
	return nil
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
