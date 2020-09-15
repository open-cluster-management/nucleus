// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/open-cluster-management/api/client/cluster/clientset/versioned/scheme"
	v1alpha1 "github.com/open-cluster-management/api/cluster/v1alpha1"
	rest "k8s.io/client-go/rest"
)

type ClusterV1alpha1Interface interface {
	RESTClient() rest.Interface
	ManagedClusterSetsGetter
	ManagedClusterSetBindingsGetter
}

// ClusterV1alpha1Client is used to interact with features provided by the cluster.open-cluster-management.io group.
type ClusterV1alpha1Client struct {
	restClient rest.Interface
}

func (c *ClusterV1alpha1Client) ManagedClusterSets() ManagedClusterSetInterface {
	return newManagedClusterSets(c)
}

func (c *ClusterV1alpha1Client) ManagedClusterSetBindings(namespace string) ManagedClusterSetBindingInterface {
	return newManagedClusterSetBindings(c, namespace)
}

// NewForConfig creates a new ClusterV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*ClusterV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ClusterV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new ClusterV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ClusterV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ClusterV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *ClusterV1alpha1Client {
	return &ClusterV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ClusterV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}