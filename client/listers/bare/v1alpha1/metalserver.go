/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-vultr-api/apis/bare/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MetalServerLister helps list MetalServers.
// All objects returned here must be treated as read-only.
type MetalServerLister interface {
	// List lists all MetalServers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MetalServer, err error)
	// MetalServers returns an object that can list and get MetalServers.
	MetalServers(namespace string) MetalServerNamespaceLister
	MetalServerListerExpansion
}

// metalServerLister implements the MetalServerLister interface.
type metalServerLister struct {
	indexer cache.Indexer
}

// NewMetalServerLister returns a new MetalServerLister.
func NewMetalServerLister(indexer cache.Indexer) MetalServerLister {
	return &metalServerLister{indexer: indexer}
}

// List lists all MetalServers in the indexer.
func (s *metalServerLister) List(selector labels.Selector) (ret []*v1alpha1.MetalServer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MetalServer))
	})
	return ret, err
}

// MetalServers returns an object that can list and get MetalServers.
func (s *metalServerLister) MetalServers(namespace string) MetalServerNamespaceLister {
	return metalServerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MetalServerNamespaceLister helps list and get MetalServers.
// All objects returned here must be treated as read-only.
type MetalServerNamespaceLister interface {
	// List lists all MetalServers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MetalServer, err error)
	// Get retrieves the MetalServer from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.MetalServer, error)
	MetalServerNamespaceListerExpansion
}

// metalServerNamespaceLister implements the MetalServerNamespaceLister
// interface.
type metalServerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MetalServers in the indexer for a given namespace.
func (s metalServerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MetalServer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MetalServer))
	})
	return ret, err
}

// Get retrieves the MetalServer from the indexer for a given namespace and name.
func (s metalServerNamespaceLister) Get(name string) (*v1alpha1.MetalServer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("metalserver"), name)
	}
	return obj.(*v1alpha1.MetalServer), nil
}
