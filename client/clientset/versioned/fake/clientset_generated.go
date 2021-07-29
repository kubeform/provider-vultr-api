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
	clientset "kubeform.dev/provider-vultr-api/client/clientset/versioned"
	barev1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/bare/v1alpha1"
	fakebarev1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/bare/v1alpha1/fake"
	blockv1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/block/v1alpha1"
	fakeblockv1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/block/v1alpha1/fake"
	dnsv1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/dns/v1alpha1"
	fakednsv1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/dns/v1alpha1/fake"
	firewallv1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/firewall/v1alpha1"
	fakefirewallv1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/firewall/v1alpha1/fake"
	instancev1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/instance/v1alpha1"
	fakeinstancev1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/instance/v1alpha1/fake"
	isov1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/iso/v1alpha1"
	fakeisov1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/iso/v1alpha1/fake"
	loadv1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/load/v1alpha1"
	fakeloadv1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/load/v1alpha1/fake"
	objectv1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/object/v1alpha1"
	fakeobjectv1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/object/v1alpha1/fake"
	privatev1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/private/v1alpha1"
	fakeprivatev1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/private/v1alpha1/fake"
	reservedv1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/reserved/v1alpha1"
	fakereservedv1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/reserved/v1alpha1/fake"
	reversev1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/reverse/v1alpha1"
	fakereversev1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/reverse/v1alpha1/fake"
	snapshotv1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/snapshot/v1alpha1"
	fakesnapshotv1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/snapshot/v1alpha1/fake"
	sshv1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/ssh/v1alpha1"
	fakesshv1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/ssh/v1alpha1/fake"
	startupv1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/startup/v1alpha1"
	fakestartupv1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/startup/v1alpha1/fake"
	userv1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/user/v1alpha1"
	fakeuserv1alpha1 "kubeform.dev/provider-vultr-api/client/clientset/versioned/typed/user/v1alpha1/fake"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var _ clientset.Interface = &Clientset{}

// BareV1alpha1 retrieves the BareV1alpha1Client
func (c *Clientset) BareV1alpha1() barev1alpha1.BareV1alpha1Interface {
	return &fakebarev1alpha1.FakeBareV1alpha1{Fake: &c.Fake}
}

// BlockV1alpha1 retrieves the BlockV1alpha1Client
func (c *Clientset) BlockV1alpha1() blockv1alpha1.BlockV1alpha1Interface {
	return &fakeblockv1alpha1.FakeBlockV1alpha1{Fake: &c.Fake}
}

// DnsV1alpha1 retrieves the DnsV1alpha1Client
func (c *Clientset) DnsV1alpha1() dnsv1alpha1.DnsV1alpha1Interface {
	return &fakednsv1alpha1.FakeDnsV1alpha1{Fake: &c.Fake}
}

// FirewallV1alpha1 retrieves the FirewallV1alpha1Client
func (c *Clientset) FirewallV1alpha1() firewallv1alpha1.FirewallV1alpha1Interface {
	return &fakefirewallv1alpha1.FakeFirewallV1alpha1{Fake: &c.Fake}
}

// InstanceV1alpha1 retrieves the InstanceV1alpha1Client
func (c *Clientset) InstanceV1alpha1() instancev1alpha1.InstanceV1alpha1Interface {
	return &fakeinstancev1alpha1.FakeInstanceV1alpha1{Fake: &c.Fake}
}

// IsoV1alpha1 retrieves the IsoV1alpha1Client
func (c *Clientset) IsoV1alpha1() isov1alpha1.IsoV1alpha1Interface {
	return &fakeisov1alpha1.FakeIsoV1alpha1{Fake: &c.Fake}
}

// LoadV1alpha1 retrieves the LoadV1alpha1Client
func (c *Clientset) LoadV1alpha1() loadv1alpha1.LoadV1alpha1Interface {
	return &fakeloadv1alpha1.FakeLoadV1alpha1{Fake: &c.Fake}
}

// ObjectV1alpha1 retrieves the ObjectV1alpha1Client
func (c *Clientset) ObjectV1alpha1() objectv1alpha1.ObjectV1alpha1Interface {
	return &fakeobjectv1alpha1.FakeObjectV1alpha1{Fake: &c.Fake}
}

// PrivateV1alpha1 retrieves the PrivateV1alpha1Client
func (c *Clientset) PrivateV1alpha1() privatev1alpha1.PrivateV1alpha1Interface {
	return &fakeprivatev1alpha1.FakePrivateV1alpha1{Fake: &c.Fake}
}

// ReservedV1alpha1 retrieves the ReservedV1alpha1Client
func (c *Clientset) ReservedV1alpha1() reservedv1alpha1.ReservedV1alpha1Interface {
	return &fakereservedv1alpha1.FakeReservedV1alpha1{Fake: &c.Fake}
}

// ReverseV1alpha1 retrieves the ReverseV1alpha1Client
func (c *Clientset) ReverseV1alpha1() reversev1alpha1.ReverseV1alpha1Interface {
	return &fakereversev1alpha1.FakeReverseV1alpha1{Fake: &c.Fake}
}

// SnapshotV1alpha1 retrieves the SnapshotV1alpha1Client
func (c *Clientset) SnapshotV1alpha1() snapshotv1alpha1.SnapshotV1alpha1Interface {
	return &fakesnapshotv1alpha1.FakeSnapshotV1alpha1{Fake: &c.Fake}
}

// SshV1alpha1 retrieves the SshV1alpha1Client
func (c *Clientset) SshV1alpha1() sshv1alpha1.SshV1alpha1Interface {
	return &fakesshv1alpha1.FakeSshV1alpha1{Fake: &c.Fake}
}

// StartupV1alpha1 retrieves the StartupV1alpha1Client
func (c *Clientset) StartupV1alpha1() startupv1alpha1.StartupV1alpha1Interface {
	return &fakestartupv1alpha1.FakeStartupV1alpha1{Fake: &c.Fake}
}

// UserV1alpha1 retrieves the UserV1alpha1Client
func (c *Clientset) UserV1alpha1() userv1alpha1.UserV1alpha1Interface {
	return &fakeuserv1alpha1.FakeUserV1alpha1{Fake: &c.Fake}
}