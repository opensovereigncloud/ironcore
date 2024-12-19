// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1alpha1 "github.com/ironcore-dev/ironcore/api/compute/v1alpha1"
	computev1alpha1 "github.com/ironcore-dev/ironcore/client-go/applyconfigurations/compute/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMachineClasses implements MachineClassInterface
type FakeMachineClasses struct {
	Fake *FakeComputeV1alpha1
}

var machineclassesResource = v1alpha1.SchemeGroupVersion.WithResource("machineclasses")

var machineclassesKind = v1alpha1.SchemeGroupVersion.WithKind("MachineClass")

// Get takes name of the machineClass, and returns the corresponding machineClass object, and an error if there is any.
func (c *FakeMachineClasses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MachineClass, err error) {
	emptyResult := &v1alpha1.MachineClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(machineclassesResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.MachineClass), err
}

// List takes label and field selectors, and returns the list of MachineClasses that match those selectors.
func (c *FakeMachineClasses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MachineClassList, err error) {
	emptyResult := &v1alpha1.MachineClassList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(machineclassesResource, machineclassesKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MachineClassList{ListMeta: obj.(*v1alpha1.MachineClassList).ListMeta}
	for _, item := range obj.(*v1alpha1.MachineClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested machineClasses.
func (c *FakeMachineClasses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(machineclassesResource, opts))
}

// Create takes the representation of a machineClass and creates it.  Returns the server's representation of the machineClass, and an error, if there is any.
func (c *FakeMachineClasses) Create(ctx context.Context, machineClass *v1alpha1.MachineClass, opts v1.CreateOptions) (result *v1alpha1.MachineClass, err error) {
	emptyResult := &v1alpha1.MachineClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(machineclassesResource, machineClass, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.MachineClass), err
}

// Update takes the representation of a machineClass and updates it. Returns the server's representation of the machineClass, and an error, if there is any.
func (c *FakeMachineClasses) Update(ctx context.Context, machineClass *v1alpha1.MachineClass, opts v1.UpdateOptions) (result *v1alpha1.MachineClass, err error) {
	emptyResult := &v1alpha1.MachineClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(machineclassesResource, machineClass, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.MachineClass), err
}

// Delete takes name of the machineClass and deletes it. Returns an error if one occurs.
func (c *FakeMachineClasses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(machineclassesResource, name, opts), &v1alpha1.MachineClass{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMachineClasses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(machineclassesResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MachineClassList{})
	return err
}

// Patch applies the patch and returns the patched machineClass.
func (c *FakeMachineClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MachineClass, err error) {
	emptyResult := &v1alpha1.MachineClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(machineclassesResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.MachineClass), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied machineClass.
func (c *FakeMachineClasses) Apply(ctx context.Context, machineClass *computev1alpha1.MachineClassApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.MachineClass, err error) {
	if machineClass == nil {
		return nil, fmt.Errorf("machineClass provided to Apply must not be nil")
	}
	data, err := json.Marshal(machineClass)
	if err != nil {
		return nil, err
	}
	name := machineClass.Name
	if name == nil {
		return nil, fmt.Errorf("machineClass.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.MachineClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(machineclassesResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.MachineClass), err
}
