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

package fake

import (
	"context"

	v1alpha1 "kubeform.dev/provider-vultr-api/apis/snapshot/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFromURLs implements FromURLInterface
type FakeFromURLs struct {
	Fake *FakeSnapshotV1alpha1
	ns   string
}

var fromurlsResource = schema.GroupVersionResource{Group: "snapshot.vultr.kubeform.com", Version: "v1alpha1", Resource: "fromurls"}

var fromurlsKind = schema.GroupVersionKind{Group: "snapshot.vultr.kubeform.com", Version: "v1alpha1", Kind: "FromURL"}

// Get takes name of the fromURL, and returns the corresponding fromURL object, and an error if there is any.
func (c *FakeFromURLs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FromURL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(fromurlsResource, c.ns, name), &v1alpha1.FromURL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FromURL), err
}

// List takes label and field selectors, and returns the list of FromURLs that match those selectors.
func (c *FakeFromURLs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FromURLList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(fromurlsResource, fromurlsKind, c.ns, opts), &v1alpha1.FromURLList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FromURLList{ListMeta: obj.(*v1alpha1.FromURLList).ListMeta}
	for _, item := range obj.(*v1alpha1.FromURLList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested fromURLs.
func (c *FakeFromURLs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(fromurlsResource, c.ns, opts))

}

// Create takes the representation of a fromURL and creates it.  Returns the server's representation of the fromURL, and an error, if there is any.
func (c *FakeFromURLs) Create(ctx context.Context, fromURL *v1alpha1.FromURL, opts v1.CreateOptions) (result *v1alpha1.FromURL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(fromurlsResource, c.ns, fromURL), &v1alpha1.FromURL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FromURL), err
}

// Update takes the representation of a fromURL and updates it. Returns the server's representation of the fromURL, and an error, if there is any.
func (c *FakeFromURLs) Update(ctx context.Context, fromURL *v1alpha1.FromURL, opts v1.UpdateOptions) (result *v1alpha1.FromURL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(fromurlsResource, c.ns, fromURL), &v1alpha1.FromURL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FromURL), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFromURLs) UpdateStatus(ctx context.Context, fromURL *v1alpha1.FromURL, opts v1.UpdateOptions) (*v1alpha1.FromURL, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(fromurlsResource, "status", c.ns, fromURL), &v1alpha1.FromURL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FromURL), err
}

// Delete takes name of the fromURL and deletes it. Returns an error if one occurs.
func (c *FakeFromURLs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(fromurlsResource, c.ns, name), &v1alpha1.FromURL{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFromURLs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(fromurlsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FromURLList{})
	return err
}

// Patch applies the patch and returns the patched fromURL.
func (c *FakeFromURLs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FromURL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fromurlsResource, c.ns, name, pt, data, subresources...), &v1alpha1.FromURL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FromURL), err
}
