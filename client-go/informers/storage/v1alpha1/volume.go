// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	storagev1alpha1 "github.com/ironcore-dev/ironcore/api/storage/v1alpha1"
	internalinterfaces "github.com/ironcore-dev/ironcore/client-go/informers/internalinterfaces"
	ironcore "github.com/ironcore-dev/ironcore/client-go/ironcore"
	v1alpha1 "github.com/ironcore-dev/ironcore/client-go/listers/storage/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VolumeInformer provides access to a shared informer and lister for
// Volumes.
type VolumeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.VolumeLister
}

type volumeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewVolumeInformer constructs a new informer for Volume type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVolumeInformer(client ironcore.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVolumeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredVolumeInformer constructs a new informer for Volume type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVolumeInformer(client ironcore.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorageV1alpha1().Volumes(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorageV1alpha1().Volumes(namespace).Watch(context.TODO(), options)
			},
		},
		&storagev1alpha1.Volume{},
		resyncPeriod,
		indexers,
	)
}

func (f *volumeInformer) defaultInformer(client ironcore.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVolumeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *volumeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&storagev1alpha1.Volume{}, f.defaultInformer)
}

func (f *volumeInformer) Lister() v1alpha1.VolumeLister {
	return v1alpha1.NewVolumeLister(f.Informer().GetIndexer())
}
