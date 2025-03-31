// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	networkingv1alpha1 "github.com/ironcore-dev/ironcore/api/networking/v1alpha1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// NATGatewayLister helps list NATGateways.
// All objects returned here must be treated as read-only.
type NATGatewayLister interface {
	// List lists all NATGateways in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*networkingv1alpha1.NATGateway, err error)
	// NATGateways returns an object that can list and get NATGateways.
	NATGateways(namespace string) NATGatewayNamespaceLister
	NATGatewayListerExpansion
}

// nATGatewayLister implements the NATGatewayLister interface.
type nATGatewayLister struct {
	listers.ResourceIndexer[*networkingv1alpha1.NATGateway]
}

// NewNATGatewayLister returns a new NATGatewayLister.
func NewNATGatewayLister(indexer cache.Indexer) NATGatewayLister {
	return &nATGatewayLister{listers.New[*networkingv1alpha1.NATGateway](indexer, networkingv1alpha1.Resource("natgateway"))}
}

// NATGateways returns an object that can list and get NATGateways.
func (s *nATGatewayLister) NATGateways(namespace string) NATGatewayNamespaceLister {
	return nATGatewayNamespaceLister{listers.NewNamespaced[*networkingv1alpha1.NATGateway](s.ResourceIndexer, namespace)}
}

// NATGatewayNamespaceLister helps list and get NATGateways.
// All objects returned here must be treated as read-only.
type NATGatewayNamespaceLister interface {
	// List lists all NATGateways in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*networkingv1alpha1.NATGateway, err error)
	// Get retrieves the NATGateway from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*networkingv1alpha1.NATGateway, error)
	NATGatewayNamespaceListerExpansion
}

// nATGatewayNamespaceLister implements the NATGatewayNamespaceLister
// interface.
type nATGatewayNamespaceLister struct {
	listers.ResourceIndexer[*networkingv1alpha1.NATGateway]
}
