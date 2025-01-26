// Copyright 2024 Authors of elf-io
// SPDX-License-Identifier: Apache-2.0

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	context "context"
	time "time"

	apisbmcspidernetiov1beta1 "github.com/infrastructure-io/topohub/pkg/k8s/apis/topohub.infrastructure.io/v1beta1"
	versioned "github.com/infrastructure-io/topohub/pkg/k8s/client/clientset/versioned"
	internalinterfaces "github.com/infrastructure-io/topohub/pkg/k8s/client/informers/externalversions/internalinterfaces"
	bmcspidernetiov1beta1 "github.com/infrastructure-io/topohub/pkg/k8s/client/listers/topohub.infrastructure.io/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterAgentInformer provides access to a shared informer and lister for
// ClusterAgents.
type ClusterAgentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() bmcspidernetiov1beta1.ClusterAgentLister
}

type clusterAgentInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterAgentInformer constructs a new informer for ClusterAgent type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterAgentInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterAgentInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterAgentInformer constructs a new informer for ClusterAgent type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterAgentInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BmcV1beta1().ClusterAgents().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BmcV1beta1().ClusterAgents().Watch(context.TODO(), options)
			},
		},
		&apisbmcspidernetiov1beta1.ClusterAgent{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterAgentInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterAgentInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterAgentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisbmcspidernetiov1beta1.ClusterAgent{}, f.defaultInformer)
}

func (f *clusterAgentInformer) Lister() bmcspidernetiov1beta1.ClusterAgentLister {
	return bmcspidernetiov1beta1.NewClusterAgentLister(f.Informer().GetIndexer())
}
