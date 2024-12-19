// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1alpha1 "github.com/ironcore-dev/ironcore/api/networking/v1alpha1"
	networkingv1alpha1 "github.com/ironcore-dev/ironcore/client-go/applyconfigurations/networking/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVirtualIPs implements VirtualIPInterface
type FakeVirtualIPs struct {
	Fake *FakeNetworkingV1alpha1
	ns   string
}

var virtualipsResource = v1alpha1.SchemeGroupVersion.WithResource("virtualips")

var virtualipsKind = v1alpha1.SchemeGroupVersion.WithKind("VirtualIP")

// Get takes name of the virtualIP, and returns the corresponding virtualIP object, and an error if there is any.
func (c *FakeVirtualIPs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VirtualIP, err error) {
	emptyResult := &v1alpha1.VirtualIP{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(virtualipsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.VirtualIP), err
}

// List takes label and field selectors, and returns the list of VirtualIPs that match those selectors.
func (c *FakeVirtualIPs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VirtualIPList, err error) {
	emptyResult := &v1alpha1.VirtualIPList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(virtualipsResource, virtualipsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VirtualIPList{ListMeta: obj.(*v1alpha1.VirtualIPList).ListMeta}
	for _, item := range obj.(*v1alpha1.VirtualIPList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualIPs.
func (c *FakeVirtualIPs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(virtualipsResource, c.ns, opts))

}

// Create takes the representation of a virtualIP and creates it.  Returns the server's representation of the virtualIP, and an error, if there is any.
func (c *FakeVirtualIPs) Create(ctx context.Context, virtualIP *v1alpha1.VirtualIP, opts v1.CreateOptions) (result *v1alpha1.VirtualIP, err error) {
	emptyResult := &v1alpha1.VirtualIP{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(virtualipsResource, c.ns, virtualIP, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.VirtualIP), err
}

// Update takes the representation of a virtualIP and updates it. Returns the server's representation of the virtualIP, and an error, if there is any.
func (c *FakeVirtualIPs) Update(ctx context.Context, virtualIP *v1alpha1.VirtualIP, opts v1.UpdateOptions) (result *v1alpha1.VirtualIP, err error) {
	emptyResult := &v1alpha1.VirtualIP{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(virtualipsResource, c.ns, virtualIP, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.VirtualIP), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVirtualIPs) UpdateStatus(ctx context.Context, virtualIP *v1alpha1.VirtualIP, opts v1.UpdateOptions) (result *v1alpha1.VirtualIP, err error) {
	emptyResult := &v1alpha1.VirtualIP{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(virtualipsResource, "status", c.ns, virtualIP, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.VirtualIP), err
}

// Delete takes name of the virtualIP and deletes it. Returns an error if one occurs.
func (c *FakeVirtualIPs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(virtualipsResource, c.ns, name, opts), &v1alpha1.VirtualIP{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualIPs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(virtualipsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VirtualIPList{})
	return err
}

// Patch applies the patch and returns the patched virtualIP.
func (c *FakeVirtualIPs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VirtualIP, err error) {
	emptyResult := &v1alpha1.VirtualIP{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(virtualipsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.VirtualIP), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied virtualIP.
func (c *FakeVirtualIPs) Apply(ctx context.Context, virtualIP *networkingv1alpha1.VirtualIPApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.VirtualIP, err error) {
	if virtualIP == nil {
		return nil, fmt.Errorf("virtualIP provided to Apply must not be nil")
	}
	data, err := json.Marshal(virtualIP)
	if err != nil {
		return nil, err
	}
	name := virtualIP.Name
	if name == nil {
		return nil, fmt.Errorf("virtualIP.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.VirtualIP{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(virtualipsResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.VirtualIP), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeVirtualIPs) ApplyStatus(ctx context.Context, virtualIP *networkingv1alpha1.VirtualIPApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.VirtualIP, err error) {
	if virtualIP == nil {
		return nil, fmt.Errorf("virtualIP provided to Apply must not be nil")
	}
	data, err := json.Marshal(virtualIP)
	if err != nil {
		return nil, err
	}
	name := virtualIP.Name
	if name == nil {
		return nil, fmt.Errorf("virtualIP.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.VirtualIP{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(virtualipsResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.VirtualIP), err
}
