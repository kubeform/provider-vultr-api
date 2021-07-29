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

	v1alpha1 "kubeform.dev/provider-vultr-api/apis/ssh/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKeys implements KeyInterface
type FakeKeys struct {
	Fake *FakeSshV1alpha1
	ns   string
}

var keysResource = schema.GroupVersionResource{Group: "ssh.vultr.kubeform.com", Version: "v1alpha1", Resource: "keys"}

var keysKind = schema.GroupVersionKind{Group: "ssh.vultr.kubeform.com", Version: "v1alpha1", Kind: "Key"}

// Get takes name of the key, and returns the corresponding key object, and an error if there is any.
func (c *FakeKeys) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Key, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(keysResource, c.ns, name), &v1alpha1.Key{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Key), err
}

// List takes label and field selectors, and returns the list of Keys that match those selectors.
func (c *FakeKeys) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(keysResource, keysKind, c.ns, opts), &v1alpha1.KeyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KeyList{ListMeta: obj.(*v1alpha1.KeyList).ListMeta}
	for _, item := range obj.(*v1alpha1.KeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested keys.
func (c *FakeKeys) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(keysResource, c.ns, opts))

}

// Create takes the representation of a key and creates it.  Returns the server's representation of the key, and an error, if there is any.
func (c *FakeKeys) Create(ctx context.Context, key *v1alpha1.Key, opts v1.CreateOptions) (result *v1alpha1.Key, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(keysResource, c.ns, key), &v1alpha1.Key{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Key), err
}

// Update takes the representation of a key and updates it. Returns the server's representation of the key, and an error, if there is any.
func (c *FakeKeys) Update(ctx context.Context, key *v1alpha1.Key, opts v1.UpdateOptions) (result *v1alpha1.Key, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(keysResource, c.ns, key), &v1alpha1.Key{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Key), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKeys) UpdateStatus(ctx context.Context, key *v1alpha1.Key, opts v1.UpdateOptions) (*v1alpha1.Key, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(keysResource, "status", c.ns, key), &v1alpha1.Key{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Key), err
}

// Delete takes name of the key and deletes it. Returns an error if one occurs.
func (c *FakeKeys) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(keysResource, c.ns, name), &v1alpha1.Key{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKeys) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(keysResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.KeyList{})
	return err
}

// Patch applies the patch and returns the patched key.
func (c *FakeKeys) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Key, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(keysResource, c.ns, name, pt, data, subresources...), &v1alpha1.Key{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Key), err
}
