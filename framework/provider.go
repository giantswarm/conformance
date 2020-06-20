package framework

import (
	"fmt"
	"os"
	"sync"

	"github.com/giantswarm/k8sclient/v3/pkg/k8sclient"
	"github.com/pkg/errors"
)

// Factory is a func which operates provider specific behavior.
type Factory func() (ProviderInterface, error)

var (
	providers = make(map[string]Factory)
	mutex     sync.Mutex
)

// RegisterProvider is expected to be called during application init,
// typically by an init function in a provider package.
func RegisterProvider(name string, factory Factory) {
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := providers[name]; ok {
		panic(fmt.Sprintf("provider %s already registered", name))
	}
	providers[name] = factory
}

// SetupProviderConfig validates the chosen provider and creates
// an interface instance for it.
func SetupProviderConfig(providerName string) (ProviderInterface, error) {
	var err error

	mutex.Lock()
	defer mutex.Unlock()
	factory, ok := providers[providerName]
	if !ok {
		return nil, errors.Wrapf(os.ErrNotExist, "The provider %s is unknown.", providerName)
	}
	provider, err := factory()

	return provider, err
}

// ProviderInterface contains the implementation for certain
// provider-specific functionality.
type ProviderInterface interface {
	CreateTenantCluster(c k8sclient.Interface) error
	UpdateTenantCluster() error
	DeleteTenantCluster() error
}

// Provider is the default implementation of the ProviderInterface
// which doesn't do anything.
type Provider struct{}

func (p Provider) CreateTenantCluster(c k8sclient.Interface) error {
	return fmt.Errorf("Provider does not support Create")
}

func (p Provider) UpdateTenantCluster() error {
	return fmt.Errorf("Provider does not support Update")
}

func (p Provider) DeleteTenantCluster() error {
	return fmt.Errorf("Provider does not support Delete")
}

var _ ProviderInterface = Provider{}
