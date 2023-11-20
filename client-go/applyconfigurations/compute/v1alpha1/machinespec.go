/*
 * Copyright (c) 2022 by the IronCore authors.
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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/ironcore-dev/ironcore/api/compute/v1alpha1"
	commonv1alpha1 "github.com/ironcore-dev/ironcore/client-go/applyconfigurations/common/v1alpha1"
	v1 "k8s.io/api/core/v1"
)

// MachineSpecApplyConfiguration represents an declarative configuration of the MachineSpec type for use
// with apply.
type MachineSpecApplyConfiguration struct {
	MachineClassRef     *v1.LocalObjectReference                            `json:"machineClassRef,omitempty"`
	MachinePoolSelector map[string]string                                   `json:"machinePoolSelector,omitempty"`
	MachinePoolRef      *v1.LocalObjectReference                            `json:"machinePoolRef,omitempty"`
	Power               *v1alpha1.Power                                     `json:"power,omitempty"`
	Image               *string                                             `json:"image,omitempty"`
	ImagePullSecretRef  *v1.LocalObjectReference                            `json:"imagePullSecret,omitempty"`
	NetworkInterfaces   []NetworkInterfaceApplyConfiguration                `json:"networkInterfaces,omitempty"`
	Volumes             []VolumeApplyConfiguration                          `json:"volumes,omitempty"`
	IgnitionRef         *commonv1alpha1.SecretKeySelectorApplyConfiguration `json:"ignitionRef,omitempty"`
	EFIVars             []EFIVarApplyConfiguration                          `json:"efiVars,omitempty"`
	Tolerations         []commonv1alpha1.TolerationApplyConfiguration       `json:"tolerations,omitempty"`
}

// MachineSpecApplyConfiguration constructs an declarative configuration of the MachineSpec type for use with
// apply.
func MachineSpec() *MachineSpecApplyConfiguration {
	return &MachineSpecApplyConfiguration{}
}

// WithMachineClassRef sets the MachineClassRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MachineClassRef field is set to the value of the last call.
func (b *MachineSpecApplyConfiguration) WithMachineClassRef(value v1.LocalObjectReference) *MachineSpecApplyConfiguration {
	b.MachineClassRef = &value
	return b
}

// WithMachinePoolSelector puts the entries into the MachinePoolSelector field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the MachinePoolSelector field,
// overwriting an existing map entries in MachinePoolSelector field with the same key.
func (b *MachineSpecApplyConfiguration) WithMachinePoolSelector(entries map[string]string) *MachineSpecApplyConfiguration {
	if b.MachinePoolSelector == nil && len(entries) > 0 {
		b.MachinePoolSelector = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.MachinePoolSelector[k] = v
	}
	return b
}

// WithMachinePoolRef sets the MachinePoolRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MachinePoolRef field is set to the value of the last call.
func (b *MachineSpecApplyConfiguration) WithMachinePoolRef(value v1.LocalObjectReference) *MachineSpecApplyConfiguration {
	b.MachinePoolRef = &value
	return b
}

// WithPower sets the Power field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Power field is set to the value of the last call.
func (b *MachineSpecApplyConfiguration) WithPower(value v1alpha1.Power) *MachineSpecApplyConfiguration {
	b.Power = &value
	return b
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *MachineSpecApplyConfiguration) WithImage(value string) *MachineSpecApplyConfiguration {
	b.Image = &value
	return b
}

// WithImagePullSecretRef sets the ImagePullSecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ImagePullSecretRef field is set to the value of the last call.
func (b *MachineSpecApplyConfiguration) WithImagePullSecretRef(value v1.LocalObjectReference) *MachineSpecApplyConfiguration {
	b.ImagePullSecretRef = &value
	return b
}

// WithNetworkInterfaces adds the given value to the NetworkInterfaces field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NetworkInterfaces field.
func (b *MachineSpecApplyConfiguration) WithNetworkInterfaces(values ...*NetworkInterfaceApplyConfiguration) *MachineSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithNetworkInterfaces")
		}
		b.NetworkInterfaces = append(b.NetworkInterfaces, *values[i])
	}
	return b
}

// WithVolumes adds the given value to the Volumes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Volumes field.
func (b *MachineSpecApplyConfiguration) WithVolumes(values ...*VolumeApplyConfiguration) *MachineSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithVolumes")
		}
		b.Volumes = append(b.Volumes, *values[i])
	}
	return b
}

// WithIgnitionRef sets the IgnitionRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IgnitionRef field is set to the value of the last call.
func (b *MachineSpecApplyConfiguration) WithIgnitionRef(value *commonv1alpha1.SecretKeySelectorApplyConfiguration) *MachineSpecApplyConfiguration {
	b.IgnitionRef = value
	return b
}

// WithEFIVars adds the given value to the EFIVars field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the EFIVars field.
func (b *MachineSpecApplyConfiguration) WithEFIVars(values ...*EFIVarApplyConfiguration) *MachineSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithEFIVars")
		}
		b.EFIVars = append(b.EFIVars, *values[i])
	}
	return b
}

// WithTolerations adds the given value to the Tolerations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tolerations field.
func (b *MachineSpecApplyConfiguration) WithTolerations(values ...*commonv1alpha1.TolerationApplyConfiguration) *MachineSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTolerations")
		}
		b.Tolerations = append(b.Tolerations, *values[i])
	}
	return b
}
