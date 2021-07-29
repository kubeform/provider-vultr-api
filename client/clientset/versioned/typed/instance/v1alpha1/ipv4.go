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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "kubeform.dev/provider-vultr-api/apis/instance/v1alpha1"
	scheme "kubeform.dev/provider-vultr-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// Ipv4sGetter has a method to return a Ipv4Interface.
// A group's client should implement this interface.
type Ipv4sGetter interface {
	Ipv4s(namespace string) Ipv4Interface
}

// Ipv4Interface has methods to work with Ipv4 resources.
type Ipv4Interface interface {
	Create(ctx context.Context, ipv4 *v1alpha1.Ipv4, opts v1.CreateOptions) (*v1alpha1.Ipv4, error)
	Update(ctx context.Context, ipv4 *v1alpha1.Ipv4, opts v1.UpdateOptions) (*v1alpha1.Ipv4, error)
	UpdateStatus(ctx context.Context, ipv4 *v1alpha1.Ipv4, opts v1.UpdateOptions) (*v1alpha1.Ipv4, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Ipv4, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.Ipv4List, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Ipv4, err error)
	Ipv4Expansion
}

// ipv4s implements Ipv4Interface
type ipv4s struct {
	client rest.Interface
	ns     string
}

// newIpv4s returns a Ipv4s
func newIpv4s(c *InstanceV1alpha1Client, namespace string) *ipv4s {
	return &ipv4s{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the ipv4, and returns the corresponding ipv4 object, and an error if there is any.
func (c *ipv4s) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Ipv4, err error) {
	result = &v1alpha1.Ipv4{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ipv4s").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Ipv4s that match those selectors.
func (c *ipv4s) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.Ipv4List, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.Ipv4List{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ipv4s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ipv4s.
func (c *ipv4s) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ipv4s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a ipv4 and creates it.  Returns the server's representation of the ipv4, and an error, if there is any.
func (c *ipv4s) Create(ctx context.Context, ipv4 *v1alpha1.Ipv4, opts v1.CreateOptions) (result *v1alpha1.Ipv4, err error) {
	result = &v1alpha1.Ipv4{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ipv4s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ipv4).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a ipv4 and updates it. Returns the server's representation of the ipv4, and an error, if there is any.
func (c *ipv4s) Update(ctx context.Context, ipv4 *v1alpha1.Ipv4, opts v1.UpdateOptions) (result *v1alpha1.Ipv4, err error) {
	result = &v1alpha1.Ipv4{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ipv4s").
		Name(ipv4.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ipv4).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *ipv4s) UpdateStatus(ctx context.Context, ipv4 *v1alpha1.Ipv4, opts v1.UpdateOptions) (result *v1alpha1.Ipv4, err error) {
	result = &v1alpha1.Ipv4{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ipv4s").
		Name(ipv4.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ipv4).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the ipv4 and deletes it. Returns an error if one occurs.
func (c *ipv4s) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ipv4s").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ipv4s) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ipv4s").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched ipv4.
func (c *ipv4s) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Ipv4, err error) {
	result = &v1alpha1.Ipv4{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ipv4s").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
