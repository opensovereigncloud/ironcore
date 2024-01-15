// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "github.com/ironcore-dev/ironcore/api/ipam/v1alpha1"
	ipamv1alpha1 "github.com/ironcore-dev/ironcore/client-go/applyconfigurations/ipam/v1alpha1"
	scheme "github.com/ironcore-dev/ironcore/client-go/ironcore/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PrefixAllocationsGetter has a method to return a PrefixAllocationInterface.
// A group's client should implement this interface.
type PrefixAllocationsGetter interface {
	PrefixAllocations(namespace string) PrefixAllocationInterface
}

// PrefixAllocationInterface has methods to work with PrefixAllocation resources.
type PrefixAllocationInterface interface {
	Create(ctx context.Context, prefixAllocation *v1alpha1.PrefixAllocation, opts v1.CreateOptions) (*v1alpha1.PrefixAllocation, error)
	Update(ctx context.Context, prefixAllocation *v1alpha1.PrefixAllocation, opts v1.UpdateOptions) (*v1alpha1.PrefixAllocation, error)
	UpdateStatus(ctx context.Context, prefixAllocation *v1alpha1.PrefixAllocation, opts v1.UpdateOptions) (*v1alpha1.PrefixAllocation, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.PrefixAllocation, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.PrefixAllocationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PrefixAllocation, err error)
	Apply(ctx context.Context, prefixAllocation *ipamv1alpha1.PrefixAllocationApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.PrefixAllocation, err error)
	ApplyStatus(ctx context.Context, prefixAllocation *ipamv1alpha1.PrefixAllocationApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.PrefixAllocation, err error)
	PrefixAllocationExpansion
}

// prefixAllocations implements PrefixAllocationInterface
type prefixAllocations struct {
	client rest.Interface
	ns     string
}

// newPrefixAllocations returns a PrefixAllocations
func newPrefixAllocations(c *IpamV1alpha1Client, namespace string) *prefixAllocations {
	return &prefixAllocations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the prefixAllocation, and returns the corresponding prefixAllocation object, and an error if there is any.
func (c *prefixAllocations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PrefixAllocation, err error) {
	result = &v1alpha1.PrefixAllocation{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("prefixallocations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PrefixAllocations that match those selectors.
func (c *prefixAllocations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PrefixAllocationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PrefixAllocationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("prefixallocations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested prefixAllocations.
func (c *prefixAllocations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("prefixallocations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a prefixAllocation and creates it.  Returns the server's representation of the prefixAllocation, and an error, if there is any.
func (c *prefixAllocations) Create(ctx context.Context, prefixAllocation *v1alpha1.PrefixAllocation, opts v1.CreateOptions) (result *v1alpha1.PrefixAllocation, err error) {
	result = &v1alpha1.PrefixAllocation{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("prefixallocations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(prefixAllocation).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a prefixAllocation and updates it. Returns the server's representation of the prefixAllocation, and an error, if there is any.
func (c *prefixAllocations) Update(ctx context.Context, prefixAllocation *v1alpha1.PrefixAllocation, opts v1.UpdateOptions) (result *v1alpha1.PrefixAllocation, err error) {
	result = &v1alpha1.PrefixAllocation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("prefixallocations").
		Name(prefixAllocation.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(prefixAllocation).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *prefixAllocations) UpdateStatus(ctx context.Context, prefixAllocation *v1alpha1.PrefixAllocation, opts v1.UpdateOptions) (result *v1alpha1.PrefixAllocation, err error) {
	result = &v1alpha1.PrefixAllocation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("prefixallocations").
		Name(prefixAllocation.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(prefixAllocation).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the prefixAllocation and deletes it. Returns an error if one occurs.
func (c *prefixAllocations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("prefixallocations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *prefixAllocations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("prefixallocations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched prefixAllocation.
func (c *prefixAllocations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PrefixAllocation, err error) {
	result = &v1alpha1.PrefixAllocation{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("prefixallocations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied prefixAllocation.
func (c *prefixAllocations) Apply(ctx context.Context, prefixAllocation *ipamv1alpha1.PrefixAllocationApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.PrefixAllocation, err error) {
	if prefixAllocation == nil {
		return nil, fmt.Errorf("prefixAllocation provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(prefixAllocation)
	if err != nil {
		return nil, err
	}
	name := prefixAllocation.Name
	if name == nil {
		return nil, fmt.Errorf("prefixAllocation.Name must be provided to Apply")
	}
	result = &v1alpha1.PrefixAllocation{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("prefixallocations").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *prefixAllocations) ApplyStatus(ctx context.Context, prefixAllocation *ipamv1alpha1.PrefixAllocationApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.PrefixAllocation, err error) {
	if prefixAllocation == nil {
		return nil, fmt.Errorf("prefixAllocation provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(prefixAllocation)
	if err != nil {
		return nil, err
	}

	name := prefixAllocation.Name
	if name == nil {
		return nil, fmt.Errorf("prefixAllocation.Name must be provided to Apply")
	}

	result = &v1alpha1.PrefixAllocation{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("prefixallocations").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
