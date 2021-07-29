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
	v1alpha1 "kubeform.dev/provider-vultr-api/apis/instance/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// InstanceLister helps list Instances.
// All objects returned here must be treated as read-only.
type InstanceLister interface {
	// List lists all Instances in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Instance, err error)
	// Instances returns an object that can list and get Instances.
	Instances(namespace string) InstanceNamespaceLister
	InstanceListerExpansion
}

// instanceLister implements the InstanceLister interface.
type instanceLister struct {
	indexer cache.Indexer
}

// NewInstanceLister returns a new InstanceLister.
func NewInstanceLister(indexer cache.Indexer) InstanceLister {
	return &instanceLister{indexer: indexer}
}

// List lists all Instances in the indexer.
func (s *instanceLister) List(selector labels.Selector) (ret []*v1alpha1.Instance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Instance))
	})
	return ret, err
}

// Instances returns an object that can list and get Instances.
func (s *instanceLister) Instances(namespace string) InstanceNamespaceLister {
	return instanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InstanceNamespaceLister helps list and get Instances.
// All objects returned here must be treated as read-only.
type InstanceNamespaceLister interface {
	// List lists all Instances in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Instance, err error)
	// Get retrieves the Instance from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Instance, error)
	InstanceNamespaceListerExpansion
}

// instanceNamespaceLister implements the InstanceNamespaceLister
// interface.
type instanceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Instances in the indexer for a given namespace.
func (s instanceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Instance, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Instance))
	})
	return ret, err
}

// Get retrieves the Instance from the indexer for a given namespace and name.
func (s instanceNamespaceLister) Get(name string) (*v1alpha1.Instance, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("instance"), name)
	}
	return obj.(*v1alpha1.Instance), nil
}
