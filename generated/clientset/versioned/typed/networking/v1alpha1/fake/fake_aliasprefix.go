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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/onmetal/onmetal-api/apis/networking/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAliasPrefixes implements AliasPrefixInterface
type FakeAliasPrefixes struct {
	Fake *FakeNetworkingV1alpha1
	ns   string
}

var aliasprefixesResource = schema.GroupVersionResource{Group: "networking.api.onmetal.de", Version: "v1alpha1", Resource: "aliasprefixes"}

var aliasprefixesKind = schema.GroupVersionKind{Group: "networking.api.onmetal.de", Version: "v1alpha1", Kind: "AliasPrefix"}

// Get takes name of the aliasPrefix, and returns the corresponding aliasPrefix object, and an error if there is any.
func (c *FakeAliasPrefixes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AliasPrefix, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(aliasprefixesResource, c.ns, name), &v1alpha1.AliasPrefix{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AliasPrefix), err
}

// List takes label and field selectors, and returns the list of AliasPrefixes that match those selectors.
func (c *FakeAliasPrefixes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AliasPrefixList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(aliasprefixesResource, aliasprefixesKind, c.ns, opts), &v1alpha1.AliasPrefixList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AliasPrefixList{ListMeta: obj.(*v1alpha1.AliasPrefixList).ListMeta}
	for _, item := range obj.(*v1alpha1.AliasPrefixList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aliasPrefixes.
func (c *FakeAliasPrefixes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(aliasprefixesResource, c.ns, opts))

}

// Create takes the representation of a aliasPrefix and creates it.  Returns the server's representation of the aliasPrefix, and an error, if there is any.
func (c *FakeAliasPrefixes) Create(ctx context.Context, aliasPrefix *v1alpha1.AliasPrefix, opts v1.CreateOptions) (result *v1alpha1.AliasPrefix, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(aliasprefixesResource, c.ns, aliasPrefix), &v1alpha1.AliasPrefix{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AliasPrefix), err
}

// Update takes the representation of a aliasPrefix and updates it. Returns the server's representation of the aliasPrefix, and an error, if there is any.
func (c *FakeAliasPrefixes) Update(ctx context.Context, aliasPrefix *v1alpha1.AliasPrefix, opts v1.UpdateOptions) (result *v1alpha1.AliasPrefix, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(aliasprefixesResource, c.ns, aliasPrefix), &v1alpha1.AliasPrefix{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AliasPrefix), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAliasPrefixes) UpdateStatus(ctx context.Context, aliasPrefix *v1alpha1.AliasPrefix, opts v1.UpdateOptions) (*v1alpha1.AliasPrefix, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(aliasprefixesResource, "status", c.ns, aliasPrefix), &v1alpha1.AliasPrefix{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AliasPrefix), err
}

// Delete takes name of the aliasPrefix and deletes it. Returns an error if one occurs.
func (c *FakeAliasPrefixes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(aliasprefixesResource, c.ns, name, opts), &v1alpha1.AliasPrefix{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAliasPrefixes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(aliasprefixesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AliasPrefixList{})
	return err
}

// Patch applies the patch and returns the patched aliasPrefix.
func (c *FakeAliasPrefixes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AliasPrefix, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(aliasprefixesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AliasPrefix{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AliasPrefix), err
}
