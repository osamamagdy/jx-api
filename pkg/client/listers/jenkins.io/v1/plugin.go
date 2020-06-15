// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/jenkins-x/jx-api/pkg/apis/jenkins.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PluginLister helps list Plugins.
type PluginLister interface {
	// List lists all Plugins in the indexer.
	List(selector labels.Selector) (ret []*v1.Plugin, err error)
	// Plugins returns an object that can list and get Plugins.
	Plugins(namespace string) PluginNamespaceLister
	PluginListerExpansion
}

// pluginLister implements the PluginLister interface.
type pluginLister struct {
	indexer cache.Indexer
}

// NewPluginLister returns a new PluginLister.
func NewPluginLister(indexer cache.Indexer) PluginLister {
	return &pluginLister{indexer: indexer}
}

// List lists all Plugins in the indexer.
func (s *pluginLister) List(selector labels.Selector) (ret []*v1.Plugin, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Plugin))
	})
	return ret, err
}

// Plugins returns an object that can list and get Plugins.
func (s *pluginLister) Plugins(namespace string) PluginNamespaceLister {
	return pluginNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PluginNamespaceLister helps list and get Plugins.
type PluginNamespaceLister interface {
	// List lists all Plugins in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.Plugin, err error)
	// Get retrieves the Plugin from the indexer for a given namespace and name.
	Get(name string) (*v1.Plugin, error)
	PluginNamespaceListerExpansion
}

// pluginNamespaceLister implements the PluginNamespaceLister
// interface.
type pluginNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Plugins in the indexer for a given namespace.
func (s pluginNamespaceLister) List(selector labels.Selector) (ret []*v1.Plugin, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Plugin))
	})
	return ret, err
}

// Get retrieves the Plugin from the indexer for a given namespace and name.
func (s pluginNamespaceLister) Get(name string) (*v1.Plugin, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("plugin"), name)
	}
	return obj.(*v1.Plugin), nil
}
