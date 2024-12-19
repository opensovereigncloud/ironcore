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

// FakeNetworkPolicies implements NetworkPolicyInterface
type FakeNetworkPolicies struct {
	Fake *FakeNetworkingV1alpha1
	ns   string
}

var networkpoliciesResource = v1alpha1.SchemeGroupVersion.WithResource("networkpolicies")

var networkpoliciesKind = v1alpha1.SchemeGroupVersion.WithKind("NetworkPolicy")

// Get takes name of the networkPolicy, and returns the corresponding networkPolicy object, and an error if there is any.
func (c *FakeNetworkPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NetworkPolicy, err error) {
	emptyResult := &v1alpha1.NetworkPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(networkpoliciesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.NetworkPolicy), err
}

// List takes label and field selectors, and returns the list of NetworkPolicies that match those selectors.
func (c *FakeNetworkPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NetworkPolicyList, err error) {
	emptyResult := &v1alpha1.NetworkPolicyList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(networkpoliciesResource, networkpoliciesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkPolicyList{ListMeta: obj.(*v1alpha1.NetworkPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkPolicies.
func (c *FakeNetworkPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(networkpoliciesResource, c.ns, opts))

}

// Create takes the representation of a networkPolicy and creates it.  Returns the server's representation of the networkPolicy, and an error, if there is any.
func (c *FakeNetworkPolicies) Create(ctx context.Context, networkPolicy *v1alpha1.NetworkPolicy, opts v1.CreateOptions) (result *v1alpha1.NetworkPolicy, err error) {
	emptyResult := &v1alpha1.NetworkPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(networkpoliciesResource, c.ns, networkPolicy, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.NetworkPolicy), err
}

// Update takes the representation of a networkPolicy and updates it. Returns the server's representation of the networkPolicy, and an error, if there is any.
func (c *FakeNetworkPolicies) Update(ctx context.Context, networkPolicy *v1alpha1.NetworkPolicy, opts v1.UpdateOptions) (result *v1alpha1.NetworkPolicy, err error) {
	emptyResult := &v1alpha1.NetworkPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(networkpoliciesResource, c.ns, networkPolicy, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.NetworkPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkPolicies) UpdateStatus(ctx context.Context, networkPolicy *v1alpha1.NetworkPolicy, opts v1.UpdateOptions) (result *v1alpha1.NetworkPolicy, err error) {
	emptyResult := &v1alpha1.NetworkPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(networkpoliciesResource, "status", c.ns, networkPolicy, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.NetworkPolicy), err
}

// Delete takes name of the networkPolicy and deletes it. Returns an error if one occurs.
func (c *FakeNetworkPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(networkpoliciesResource, c.ns, name, opts), &v1alpha1.NetworkPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(networkpoliciesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkPolicyList{})
	return err
}

// Patch applies the patch and returns the patched networkPolicy.
func (c *FakeNetworkPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetworkPolicy, err error) {
	emptyResult := &v1alpha1.NetworkPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(networkpoliciesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.NetworkPolicy), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied networkPolicy.
func (c *FakeNetworkPolicies) Apply(ctx context.Context, networkPolicy *networkingv1alpha1.NetworkPolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.NetworkPolicy, err error) {
	if networkPolicy == nil {
		return nil, fmt.Errorf("networkPolicy provided to Apply must not be nil")
	}
	data, err := json.Marshal(networkPolicy)
	if err != nil {
		return nil, err
	}
	name := networkPolicy.Name
	if name == nil {
		return nil, fmt.Errorf("networkPolicy.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.NetworkPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(networkpoliciesResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.NetworkPolicy), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeNetworkPolicies) ApplyStatus(ctx context.Context, networkPolicy *networkingv1alpha1.NetworkPolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.NetworkPolicy, err error) {
	if networkPolicy == nil {
		return nil, fmt.Errorf("networkPolicy provided to Apply must not be nil")
	}
	data, err := json.Marshal(networkPolicy)
	if err != nil {
		return nil, err
	}
	name := networkPolicy.Name
	if name == nil {
		return nil, fmt.Errorf("networkPolicy.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.NetworkPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(networkpoliciesResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.NetworkPolicy), err
}
