// Copyright 2024 Authors of infrastructure-io
// SPDX-License-Identifier: Apache-2.0

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	context "context"
	time "time"

	apistopohubinfrastructureiov1beta1 "github.com/infrastructure-io/topohub/pkg/k8s/apis/topohub.infrastructure.io/v1beta1"
	versioned "github.com/infrastructure-io/topohub/pkg/k8s/client/clientset/versioned"
	internalinterfaces "github.com/infrastructure-io/topohub/pkg/k8s/client/informers/externalversions/internalinterfaces"
	topohubinfrastructureiov1beta1 "github.com/infrastructure-io/topohub/pkg/k8s/client/listers/topohub.infrastructure.io/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// HostStatusInformer provides access to a shared informer and lister for
// HostStatuses.
type HostStatusInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() topohubinfrastructureiov1beta1.HostStatusLister
}

type hostStatusInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewHostStatusInformer constructs a new informer for HostStatus type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHostStatusInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredHostStatusInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredHostStatusInformer constructs a new informer for HostStatus type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHostStatusInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TopohubV1beta1().HostStatuses().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TopohubV1beta1().HostStatuses().Watch(context.TODO(), options)
			},
		},
		&apistopohubinfrastructureiov1beta1.HostStatus{},
		resyncPeriod,
		indexers,
	)
}

func (f *hostStatusInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredHostStatusInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *hostStatusInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apistopohubinfrastructureiov1beta1.HostStatus{}, f.defaultInformer)
}

func (f *hostStatusInformer) Lister() topohubinfrastructureiov1beta1.HostStatusLister {
	return topohubinfrastructureiov1beta1.NewHostStatusLister(f.Informer().GetIndexer())
}
