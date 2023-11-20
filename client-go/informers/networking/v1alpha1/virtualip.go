/*
 * Copyright (c) 2022 by the IronCore authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	networkingv1alpha1 "github.com/ironcore-dev/ironcore/api/networking/v1alpha1"
	internalinterfaces "github.com/ironcore-dev/ironcore/client-go/informers/internalinterfaces"
	ironcore "github.com/ironcore-dev/ironcore/client-go/ironcore"
	v1alpha1 "github.com/ironcore-dev/ironcore/client-go/listers/networking/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VirtualIPInformer provides access to a shared informer and lister for
// VirtualIPs.
type VirtualIPInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.VirtualIPLister
}

type virtualIPInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewVirtualIPInformer constructs a new informer for VirtualIP type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVirtualIPInformer(client ironcore.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVirtualIPInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredVirtualIPInformer constructs a new informer for VirtualIP type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVirtualIPInformer(client ironcore.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1alpha1().VirtualIPs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1alpha1().VirtualIPs(namespace).Watch(context.TODO(), options)
			},
		},
		&networkingv1alpha1.VirtualIP{},
		resyncPeriod,
		indexers,
	)
}

func (f *virtualIPInformer) defaultInformer(client ironcore.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVirtualIPInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *virtualIPInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkingv1alpha1.VirtualIP{}, f.defaultInformer)
}

func (f *virtualIPInformer) Lister() v1alpha1.VirtualIPLister {
	return v1alpha1.NewVirtualIPLister(f.Informer().GetIndexer())
}
