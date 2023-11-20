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
	corev1alpha1 "github.com/ironcore-dev/ironcore/api/core/v1alpha1"
	v1alpha1 "github.com/ironcore-dev/ironcore/client-go/applyconfigurations/common/v1alpha1"
	v1 "k8s.io/api/core/v1"
)

// VolumeSpecApplyConfiguration represents an declarative configuration of the VolumeSpec type for use
// with apply.
type VolumeSpecApplyConfiguration struct {
	VolumeClassRef     *v1.LocalObjectReference                      `json:"volumeClassRef,omitempty"`
	VolumePoolSelector map[string]string                             `json:"volumePoolSelector,omitempty"`
	VolumePoolRef      *v1.LocalObjectReference                      `json:"volumePoolRef,omitempty"`
	ClaimRef           *v1alpha1.LocalUIDReferenceApplyConfiguration `json:"claimRef,omitempty"`
	Resources          *corev1alpha1.ResourceList                    `json:"resources,omitempty"`
	Image              *string                                       `json:"image,omitempty"`
	ImagePullSecretRef *v1.LocalObjectReference                      `json:"imagePullSecretRef,omitempty"`
	Unclaimable        *bool                                         `json:"unclaimable,omitempty"`
	Tolerations        []v1alpha1.TolerationApplyConfiguration       `json:"tolerations,omitempty"`
	Encryption         *VolumeEncryptionApplyConfiguration           `json:"encryption,omitempty"`
}

// VolumeSpecApplyConfiguration constructs an declarative configuration of the VolumeSpec type for use with
// apply.
func VolumeSpec() *VolumeSpecApplyConfiguration {
	return &VolumeSpecApplyConfiguration{}
}

// WithVolumeClassRef sets the VolumeClassRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VolumeClassRef field is set to the value of the last call.
func (b *VolumeSpecApplyConfiguration) WithVolumeClassRef(value v1.LocalObjectReference) *VolumeSpecApplyConfiguration {
	b.VolumeClassRef = &value
	return b
}

// WithVolumePoolSelector puts the entries into the VolumePoolSelector field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the VolumePoolSelector field,
// overwriting an existing map entries in VolumePoolSelector field with the same key.
func (b *VolumeSpecApplyConfiguration) WithVolumePoolSelector(entries map[string]string) *VolumeSpecApplyConfiguration {
	if b.VolumePoolSelector == nil && len(entries) > 0 {
		b.VolumePoolSelector = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.VolumePoolSelector[k] = v
	}
	return b
}

// WithVolumePoolRef sets the VolumePoolRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VolumePoolRef field is set to the value of the last call.
func (b *VolumeSpecApplyConfiguration) WithVolumePoolRef(value v1.LocalObjectReference) *VolumeSpecApplyConfiguration {
	b.VolumePoolRef = &value
	return b
}

// WithClaimRef sets the ClaimRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClaimRef field is set to the value of the last call.
func (b *VolumeSpecApplyConfiguration) WithClaimRef(value *v1alpha1.LocalUIDReferenceApplyConfiguration) *VolumeSpecApplyConfiguration {
	b.ClaimRef = value
	return b
}

// WithResources sets the Resources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resources field is set to the value of the last call.
func (b *VolumeSpecApplyConfiguration) WithResources(value corev1alpha1.ResourceList) *VolumeSpecApplyConfiguration {
	b.Resources = &value
	return b
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *VolumeSpecApplyConfiguration) WithImage(value string) *VolumeSpecApplyConfiguration {
	b.Image = &value
	return b
}

// WithImagePullSecretRef sets the ImagePullSecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ImagePullSecretRef field is set to the value of the last call.
func (b *VolumeSpecApplyConfiguration) WithImagePullSecretRef(value v1.LocalObjectReference) *VolumeSpecApplyConfiguration {
	b.ImagePullSecretRef = &value
	return b
}

// WithUnclaimable sets the Unclaimable field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Unclaimable field is set to the value of the last call.
func (b *VolumeSpecApplyConfiguration) WithUnclaimable(value bool) *VolumeSpecApplyConfiguration {
	b.Unclaimable = &value
	return b
}

// WithTolerations adds the given value to the Tolerations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tolerations field.
func (b *VolumeSpecApplyConfiguration) WithTolerations(values ...*v1alpha1.TolerationApplyConfiguration) *VolumeSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTolerations")
		}
		b.Tolerations = append(b.Tolerations, *values[i])
	}
	return b
}

// WithEncryption sets the Encryption field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Encryption field is set to the value of the last call.
func (b *VolumeSpecApplyConfiguration) WithEncryption(value *VolumeEncryptionApplyConfiguration) *VolumeSpecApplyConfiguration {
	b.Encryption = value
	return b
}
