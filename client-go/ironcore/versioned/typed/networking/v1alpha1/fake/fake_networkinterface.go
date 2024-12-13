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

// FakeNetworkInterfaces implements NetworkInterfaceInterface
type FakeNetworkInterfaces struct {
	Fake *FakeNetworkingV1alpha1
	ns   string
}

var networkinterfacesResource = v1alpha1.SchemeGroupVersion.WithResource("networkinterfaces")

var networkinterfacesKind = v1alpha1.SchemeGroupVersion.WithKind("NetworkInterface")

// Get takes name of the networkInterface, and returns the corresponding networkInterface object, and an error if there is any.
func (c *FakeNetworkInterfaces) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NetworkInterface, err error) {
	emptyResult := &v1alpha1.NetworkInterface{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(networkinterfacesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.NetworkInterface), err
}

// List takes label and field selectors, and returns the list of NetworkInterfaces that match those selectors.
func (c *FakeNetworkInterfaces) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NetworkInterfaceList, err error) {
	emptyResult := &v1alpha1.NetworkInterfaceList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(networkinterfacesResource, networkinterfacesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkInterfaceList{ListMeta: obj.(*v1alpha1.NetworkInterfaceList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkInterfaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkInterfaces.
func (c *FakeNetworkInterfaces) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(networkinterfacesResource, c.ns, opts))

}

// Create takes the representation of a networkInterface and creates it.  Returns the server's representation of the networkInterface, and an error, if there is any.
func (c *FakeNetworkInterfaces) Create(ctx context.Context, networkInterface *v1alpha1.NetworkInterface, opts v1.CreateOptions) (result *v1alpha1.NetworkInterface, err error) {
	emptyResult := &v1alpha1.NetworkInterface{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(networkinterfacesResource, c.ns, networkInterface, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.NetworkInterface), err
}

// Update takes the representation of a networkInterface and updates it. Returns the server's representation of the networkInterface, and an error, if there is any.
func (c *FakeNetworkInterfaces) Update(ctx context.Context, networkInterface *v1alpha1.NetworkInterface, opts v1.UpdateOptions) (result *v1alpha1.NetworkInterface, err error) {
	emptyResult := &v1alpha1.NetworkInterface{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(networkinterfacesResource, c.ns, networkInterface, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.NetworkInterface), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkInterfaces) UpdateStatus(ctx context.Context, networkInterface *v1alpha1.NetworkInterface, opts v1.UpdateOptions) (result *v1alpha1.NetworkInterface, err error) {
	emptyResult := &v1alpha1.NetworkInterface{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(networkinterfacesResource, "status", c.ns, networkInterface, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.NetworkInterface), err
}

// Delete takes name of the networkInterface and deletes it. Returns an error if one occurs.
func (c *FakeNetworkInterfaces) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(networkinterfacesResource, c.ns, name, opts), &v1alpha1.NetworkInterface{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkInterfaces) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(networkinterfacesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkInterfaceList{})
	return err
}

// Patch applies the patch and returns the patched networkInterface.
func (c *FakeNetworkInterfaces) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetworkInterface, err error) {
	emptyResult := &v1alpha1.NetworkInterface{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(networkinterfacesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.NetworkInterface), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied networkInterface.
func (c *FakeNetworkInterfaces) Apply(ctx context.Context, networkInterface *networkingv1alpha1.NetworkInterfaceApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.NetworkInterface, err error) {
	if networkInterface == nil {
		return nil, fmt.Errorf("networkInterface provided to Apply must not be nil")
	}
	data, err := json.Marshal(networkInterface)
	if err != nil {
		return nil, err
	}
	name := networkInterface.Name
	if name == nil {
		return nil, fmt.Errorf("networkInterface.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.NetworkInterface{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(networkinterfacesResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.NetworkInterface), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeNetworkInterfaces) ApplyStatus(ctx context.Context, networkInterface *networkingv1alpha1.NetworkInterfaceApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.NetworkInterface, err error) {
	if networkInterface == nil {
		return nil, fmt.Errorf("networkInterface provided to Apply must not be nil")
	}
	data, err := json.Marshal(networkInterface)
	if err != nil {
		return nil, err
	}
	name := networkInterface.Name
	if name == nil {
		return nil, fmt.Errorf("networkInterface.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.NetworkInterface{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(networkinterfacesResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.NetworkInterface), err
}
