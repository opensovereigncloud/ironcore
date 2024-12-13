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

// FakeNATGateways implements NATGatewayInterface
type FakeNATGateways struct {
	Fake *FakeNetworkingV1alpha1
	ns   string
}

var natgatewaysResource = v1alpha1.SchemeGroupVersion.WithResource("natgateways")

var natgatewaysKind = v1alpha1.SchemeGroupVersion.WithKind("NATGateway")

// Get takes name of the nATGateway, and returns the corresponding nATGateway object, and an error if there is any.
func (c *FakeNATGateways) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NATGateway, err error) {
	emptyResult := &v1alpha1.NATGateway{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(natgatewaysResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.NATGateway), err
}

// List takes label and field selectors, and returns the list of NATGateways that match those selectors.
func (c *FakeNATGateways) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NATGatewayList, err error) {
	emptyResult := &v1alpha1.NATGatewayList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(natgatewaysResource, natgatewaysKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NATGatewayList{ListMeta: obj.(*v1alpha1.NATGatewayList).ListMeta}
	for _, item := range obj.(*v1alpha1.NATGatewayList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nATGateways.
func (c *FakeNATGateways) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(natgatewaysResource, c.ns, opts))

}

// Create takes the representation of a nATGateway and creates it.  Returns the server's representation of the nATGateway, and an error, if there is any.
func (c *FakeNATGateways) Create(ctx context.Context, nATGateway *v1alpha1.NATGateway, opts v1.CreateOptions) (result *v1alpha1.NATGateway, err error) {
	emptyResult := &v1alpha1.NATGateway{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(natgatewaysResource, c.ns, nATGateway, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.NATGateway), err
}

// Update takes the representation of a nATGateway and updates it. Returns the server's representation of the nATGateway, and an error, if there is any.
func (c *FakeNATGateways) Update(ctx context.Context, nATGateway *v1alpha1.NATGateway, opts v1.UpdateOptions) (result *v1alpha1.NATGateway, err error) {
	emptyResult := &v1alpha1.NATGateway{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(natgatewaysResource, c.ns, nATGateway, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.NATGateway), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNATGateways) UpdateStatus(ctx context.Context, nATGateway *v1alpha1.NATGateway, opts v1.UpdateOptions) (result *v1alpha1.NATGateway, err error) {
	emptyResult := &v1alpha1.NATGateway{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(natgatewaysResource, "status", c.ns, nATGateway, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.NATGateway), err
}

// Delete takes name of the nATGateway and deletes it. Returns an error if one occurs.
func (c *FakeNATGateways) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(natgatewaysResource, c.ns, name, opts), &v1alpha1.NATGateway{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNATGateways) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(natgatewaysResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NATGatewayList{})
	return err
}

// Patch applies the patch and returns the patched nATGateway.
func (c *FakeNATGateways) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NATGateway, err error) {
	emptyResult := &v1alpha1.NATGateway{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(natgatewaysResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.NATGateway), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied nATGateway.
func (c *FakeNATGateways) Apply(ctx context.Context, nATGateway *networkingv1alpha1.NATGatewayApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.NATGateway, err error) {
	if nATGateway == nil {
		return nil, fmt.Errorf("nATGateway provided to Apply must not be nil")
	}
	data, err := json.Marshal(nATGateway)
	if err != nil {
		return nil, err
	}
	name := nATGateway.Name
	if name == nil {
		return nil, fmt.Errorf("nATGateway.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.NATGateway{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(natgatewaysResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.NATGateway), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeNATGateways) ApplyStatus(ctx context.Context, nATGateway *networkingv1alpha1.NATGatewayApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.NATGateway, err error) {
	if nATGateway == nil {
		return nil, fmt.Errorf("nATGateway provided to Apply must not be nil")
	}
	data, err := json.Marshal(nATGateway)
	if err != nil {
		return nil, err
	}
	name := nATGateway.Name
	if name == nil {
		return nil, fmt.Errorf("nATGateway.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.NATGateway{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(natgatewaysResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.NATGateway), err
}
