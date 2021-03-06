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

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1alpha1 "kubeform.dev/provider-vultr-api/apis/bare/v1alpha1"
	blockv1alpha1 "kubeform.dev/provider-vultr-api/apis/block/v1alpha1"
	dnsv1alpha1 "kubeform.dev/provider-vultr-api/apis/dns/v1alpha1"
	firewallv1alpha1 "kubeform.dev/provider-vultr-api/apis/firewall/v1alpha1"
	instancev1alpha1 "kubeform.dev/provider-vultr-api/apis/instance/v1alpha1"
	isov1alpha1 "kubeform.dev/provider-vultr-api/apis/iso/v1alpha1"
	kubernetesv1alpha1 "kubeform.dev/provider-vultr-api/apis/kubernetes/v1alpha1"
	loadv1alpha1 "kubeform.dev/provider-vultr-api/apis/load/v1alpha1"
	objectv1alpha1 "kubeform.dev/provider-vultr-api/apis/object/v1alpha1"
	privatev1alpha1 "kubeform.dev/provider-vultr-api/apis/private/v1alpha1"
	reservedv1alpha1 "kubeform.dev/provider-vultr-api/apis/reserved/v1alpha1"
	reversev1alpha1 "kubeform.dev/provider-vultr-api/apis/reverse/v1alpha1"
	snapshotv1alpha1 "kubeform.dev/provider-vultr-api/apis/snapshot/v1alpha1"
	sshv1alpha1 "kubeform.dev/provider-vultr-api/apis/ssh/v1alpha1"
	startupv1alpha1 "kubeform.dev/provider-vultr-api/apis/startup/v1alpha1"
	userv1alpha1 "kubeform.dev/provider-vultr-api/apis/user/v1alpha1"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=bare.vultr.kubeform.com, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("metalservers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Bare().V1alpha1().MetalServers().Informer()}, nil

		// Group=block.vultr.kubeform.com, Version=v1alpha1
	case blockv1alpha1.SchemeGroupVersion.WithResource("storages"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Block().V1alpha1().Storages().Informer()}, nil

		// Group=dns.vultr.kubeform.com, Version=v1alpha1
	case dnsv1alpha1.SchemeGroupVersion.WithResource("domains"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Dns().V1alpha1().Domains().Informer()}, nil
	case dnsv1alpha1.SchemeGroupVersion.WithResource("records"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Dns().V1alpha1().Records().Informer()}, nil

		// Group=firewall.vultr.kubeform.com, Version=v1alpha1
	case firewallv1alpha1.SchemeGroupVersion.WithResource("groups"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Firewall().V1alpha1().Groups().Informer()}, nil
	case firewallv1alpha1.SchemeGroupVersion.WithResource("rules"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Firewall().V1alpha1().Rules().Informer()}, nil

		// Group=instance.vultr.kubeform.com, Version=v1alpha1
	case instancev1alpha1.SchemeGroupVersion.WithResource("instances"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Instance().V1alpha1().Instances().Informer()}, nil
	case instancev1alpha1.SchemeGroupVersion.WithResource("ipv4s"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Instance().V1alpha1().Ipv4s().Informer()}, nil

		// Group=iso.vultr.kubeform.com, Version=v1alpha1
	case isov1alpha1.SchemeGroupVersion.WithResource("privates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Iso().V1alpha1().Privates().Informer()}, nil

		// Group=kubernetes.vultr.kubeform.com, Version=v1alpha1
	case kubernetesv1alpha1.SchemeGroupVersion.WithResource("kuberneteses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubernetes().V1alpha1().Kuberneteses().Informer()}, nil
	case kubernetesv1alpha1.SchemeGroupVersion.WithResource("nodepoolses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubernetes().V1alpha1().NodePoolses().Informer()}, nil

		// Group=load.vultr.kubeform.com, Version=v1alpha1
	case loadv1alpha1.SchemeGroupVersion.WithResource("balancers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Load().V1alpha1().Balancers().Informer()}, nil

		// Group=object.vultr.kubeform.com, Version=v1alpha1
	case objectv1alpha1.SchemeGroupVersion.WithResource("storages"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Object().V1alpha1().Storages().Informer()}, nil

		// Group=private.vultr.kubeform.com, Version=v1alpha1
	case privatev1alpha1.SchemeGroupVersion.WithResource("networks"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Private().V1alpha1().Networks().Informer()}, nil

		// Group=reserved.vultr.kubeform.com, Version=v1alpha1
	case reservedv1alpha1.SchemeGroupVersion.WithResource("ips"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Reserved().V1alpha1().Ips().Informer()}, nil

		// Group=reverse.vultr.kubeform.com, Version=v1alpha1
	case reversev1alpha1.SchemeGroupVersion.WithResource("ipv4s"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Reverse().V1alpha1().Ipv4s().Informer()}, nil
	case reversev1alpha1.SchemeGroupVersion.WithResource("ipv6s"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Reverse().V1alpha1().Ipv6s().Informer()}, nil

		// Group=snapshot.vultr.kubeform.com, Version=v1alpha1
	case snapshotv1alpha1.SchemeGroupVersion.WithResource("fromurls"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Snapshot().V1alpha1().FromURLs().Informer()}, nil
	case snapshotv1alpha1.SchemeGroupVersion.WithResource("snapshots"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Snapshot().V1alpha1().Snapshots().Informer()}, nil

		// Group=ssh.vultr.kubeform.com, Version=v1alpha1
	case sshv1alpha1.SchemeGroupVersion.WithResource("keys"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Ssh().V1alpha1().Keys().Informer()}, nil

		// Group=startup.vultr.kubeform.com, Version=v1alpha1
	case startupv1alpha1.SchemeGroupVersion.WithResource("scripts"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Startup().V1alpha1().Scripts().Informer()}, nil

		// Group=user.vultr.kubeform.com, Version=v1alpha1
	case userv1alpha1.SchemeGroupVersion.WithResource("users"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.User().V1alpha1().Users().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
