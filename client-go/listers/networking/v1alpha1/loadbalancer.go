// SPDX-FileCopyrightText: 2025 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	networkingv1alpha1 "github.com/ironcore-dev/ironcore/api/networking/v1alpha1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// LoadBalancerLister helps list LoadBalancers.
// All objects returned here must be treated as read-only.
type LoadBalancerLister interface {
	// List lists all LoadBalancers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*networkingv1alpha1.LoadBalancer, err error)
	// LoadBalancers returns an object that can list and get LoadBalancers.
	LoadBalancers(namespace string) LoadBalancerNamespaceLister
	LoadBalancerListerExpansion
}

// loadBalancerLister implements the LoadBalancerLister interface.
type loadBalancerLister struct {
	listers.ResourceIndexer[*networkingv1alpha1.LoadBalancer]
}

// NewLoadBalancerLister returns a new LoadBalancerLister.
func NewLoadBalancerLister(indexer cache.Indexer) LoadBalancerLister {
	return &loadBalancerLister{listers.New[*networkingv1alpha1.LoadBalancer](indexer, networkingv1alpha1.Resource("loadbalancer"))}
}

// LoadBalancers returns an object that can list and get LoadBalancers.
func (s *loadBalancerLister) LoadBalancers(namespace string) LoadBalancerNamespaceLister {
	return loadBalancerNamespaceLister{listers.NewNamespaced[*networkingv1alpha1.LoadBalancer](s.ResourceIndexer, namespace)}
}

// LoadBalancerNamespaceLister helps list and get LoadBalancers.
// All objects returned here must be treated as read-only.
type LoadBalancerNamespaceLister interface {
	// List lists all LoadBalancers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*networkingv1alpha1.LoadBalancer, err error)
	// Get retrieves the LoadBalancer from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*networkingv1alpha1.LoadBalancer, error)
	LoadBalancerNamespaceListerExpansion
}

// loadBalancerNamespaceLister implements the LoadBalancerNamespaceLister
// interface.
type loadBalancerNamespaceLister struct {
	listers.ResourceIndexer[*networkingv1alpha1.LoadBalancer]
}
