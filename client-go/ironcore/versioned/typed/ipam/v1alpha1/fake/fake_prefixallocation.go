// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1alpha1 "github.com/ironcore-dev/ironcore/api/ipam/v1alpha1"
	ipamv1alpha1 "github.com/ironcore-dev/ironcore/client-go/applyconfigurations/ipam/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePrefixAllocations implements PrefixAllocationInterface
type FakePrefixAllocations struct {
	Fake *FakeIpamV1alpha1
	ns   string
}

var prefixallocationsResource = v1alpha1.SchemeGroupVersion.WithResource("prefixallocations")

var prefixallocationsKind = v1alpha1.SchemeGroupVersion.WithKind("PrefixAllocation")

// Get takes name of the prefixAllocation, and returns the corresponding prefixAllocation object, and an error if there is any.
func (c *FakePrefixAllocations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PrefixAllocation, err error) {
	emptyResult := &v1alpha1.PrefixAllocation{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(prefixallocationsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.PrefixAllocation), err
}

// List takes label and field selectors, and returns the list of PrefixAllocations that match those selectors.
func (c *FakePrefixAllocations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PrefixAllocationList, err error) {
	emptyResult := &v1alpha1.PrefixAllocationList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(prefixallocationsResource, prefixallocationsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PrefixAllocationList{ListMeta: obj.(*v1alpha1.PrefixAllocationList).ListMeta}
	for _, item := range obj.(*v1alpha1.PrefixAllocationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested prefixAllocations.
func (c *FakePrefixAllocations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(prefixallocationsResource, c.ns, opts))

}

// Create takes the representation of a prefixAllocation and creates it.  Returns the server's representation of the prefixAllocation, and an error, if there is any.
func (c *FakePrefixAllocations) Create(ctx context.Context, prefixAllocation *v1alpha1.PrefixAllocation, opts v1.CreateOptions) (result *v1alpha1.PrefixAllocation, err error) {
	emptyResult := &v1alpha1.PrefixAllocation{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(prefixallocationsResource, c.ns, prefixAllocation, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.PrefixAllocation), err
}

// Update takes the representation of a prefixAllocation and updates it. Returns the server's representation of the prefixAllocation, and an error, if there is any.
func (c *FakePrefixAllocations) Update(ctx context.Context, prefixAllocation *v1alpha1.PrefixAllocation, opts v1.UpdateOptions) (result *v1alpha1.PrefixAllocation, err error) {
	emptyResult := &v1alpha1.PrefixAllocation{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(prefixallocationsResource, c.ns, prefixAllocation, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.PrefixAllocation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePrefixAllocations) UpdateStatus(ctx context.Context, prefixAllocation *v1alpha1.PrefixAllocation, opts v1.UpdateOptions) (result *v1alpha1.PrefixAllocation, err error) {
	emptyResult := &v1alpha1.PrefixAllocation{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(prefixallocationsResource, "status", c.ns, prefixAllocation, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.PrefixAllocation), err
}

// Delete takes name of the prefixAllocation and deletes it. Returns an error if one occurs.
func (c *FakePrefixAllocations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(prefixallocationsResource, c.ns, name, opts), &v1alpha1.PrefixAllocation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePrefixAllocations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(prefixallocationsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PrefixAllocationList{})
	return err
}

// Patch applies the patch and returns the patched prefixAllocation.
func (c *FakePrefixAllocations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PrefixAllocation, err error) {
	emptyResult := &v1alpha1.PrefixAllocation{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(prefixallocationsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.PrefixAllocation), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied prefixAllocation.
func (c *FakePrefixAllocations) Apply(ctx context.Context, prefixAllocation *ipamv1alpha1.PrefixAllocationApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.PrefixAllocation, err error) {
	if prefixAllocation == nil {
		return nil, fmt.Errorf("prefixAllocation provided to Apply must not be nil")
	}
	data, err := json.Marshal(prefixAllocation)
	if err != nil {
		return nil, err
	}
	name := prefixAllocation.Name
	if name == nil {
		return nil, fmt.Errorf("prefixAllocation.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.PrefixAllocation{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(prefixallocationsResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.PrefixAllocation), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakePrefixAllocations) ApplyStatus(ctx context.Context, prefixAllocation *ipamv1alpha1.PrefixAllocationApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.PrefixAllocation, err error) {
	if prefixAllocation == nil {
		return nil, fmt.Errorf("prefixAllocation provided to Apply must not be nil")
	}
	data, err := json.Marshal(prefixAllocation)
	if err != nil {
		return nil, err
	}
	name := prefixAllocation.Name
	if name == nil {
		return nil, fmt.Errorf("prefixAllocation.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.PrefixAllocation{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(prefixallocationsResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.PrefixAllocation), err
}
