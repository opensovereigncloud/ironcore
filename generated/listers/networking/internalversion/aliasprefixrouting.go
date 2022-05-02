/*
 * Copyright (c) 2021 by the OnMetal authors.
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
// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	networking "github.com/onmetal/onmetal-api/apis/networking"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AliasPrefixRoutingLister helps list AliasPrefixRoutings.
// All objects returned here must be treated as read-only.
type AliasPrefixRoutingLister interface {
	// List lists all AliasPrefixRoutings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*networking.AliasPrefixRouting, err error)
	// AliasPrefixRoutings returns an object that can list and get AliasPrefixRoutings.
	AliasPrefixRoutings(namespace string) AliasPrefixRoutingNamespaceLister
	AliasPrefixRoutingListerExpansion
}

// aliasPrefixRoutingLister implements the AliasPrefixRoutingLister interface.
type aliasPrefixRoutingLister struct {
	indexer cache.Indexer
}

// NewAliasPrefixRoutingLister returns a new AliasPrefixRoutingLister.
func NewAliasPrefixRoutingLister(indexer cache.Indexer) AliasPrefixRoutingLister {
	return &aliasPrefixRoutingLister{indexer: indexer}
}

// List lists all AliasPrefixRoutings in the indexer.
func (s *aliasPrefixRoutingLister) List(selector labels.Selector) (ret []*networking.AliasPrefixRouting, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*networking.AliasPrefixRouting))
	})
	return ret, err
}

// AliasPrefixRoutings returns an object that can list and get AliasPrefixRoutings.
func (s *aliasPrefixRoutingLister) AliasPrefixRoutings(namespace string) AliasPrefixRoutingNamespaceLister {
	return aliasPrefixRoutingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AliasPrefixRoutingNamespaceLister helps list and get AliasPrefixRoutings.
// All objects returned here must be treated as read-only.
type AliasPrefixRoutingNamespaceLister interface {
	// List lists all AliasPrefixRoutings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*networking.AliasPrefixRouting, err error)
	// Get retrieves the AliasPrefixRouting from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*networking.AliasPrefixRouting, error)
	AliasPrefixRoutingNamespaceListerExpansion
}

// aliasPrefixRoutingNamespaceLister implements the AliasPrefixRoutingNamespaceLister
// interface.
type aliasPrefixRoutingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AliasPrefixRoutings in the indexer for a given namespace.
func (s aliasPrefixRoutingNamespaceLister) List(selector labels.Selector) (ret []*networking.AliasPrefixRouting, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*networking.AliasPrefixRouting))
	})
	return ret, err
}

// Get retrieves the AliasPrefixRouting from the indexer for a given namespace and name.
func (s aliasPrefixRoutingNamespaceLister) Get(name string) (*networking.AliasPrefixRouting, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(networking.Resource("aliasprefixrouting"), name)
	}
	return obj.(*networking.AliasPrefixRouting), nil
}
