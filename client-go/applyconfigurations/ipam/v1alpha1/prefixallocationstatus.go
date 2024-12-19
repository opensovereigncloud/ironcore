// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/ironcore-dev/ironcore/api/common/v1alpha1"
	ipamv1alpha1 "github.com/ironcore-dev/ironcore/api/ipam/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PrefixAllocationStatusApplyConfiguration represents a declarative configuration of the PrefixAllocationStatus type for use
// with apply.
type PrefixAllocationStatusApplyConfiguration struct {
	Prefix                  *v1alpha1.IPPrefix                  `json:"prefix,omitempty"`
	Phase                   *ipamv1alpha1.PrefixAllocationPhase `json:"phase,omitempty"`
	LastPhaseTransitionTime *v1.Time                            `json:"lastPhaseTransitionTime,omitempty"`
}

// PrefixAllocationStatusApplyConfiguration constructs a declarative configuration of the PrefixAllocationStatus type for use with
// apply.
func PrefixAllocationStatus() *PrefixAllocationStatusApplyConfiguration {
	return &PrefixAllocationStatusApplyConfiguration{}
}

// WithPrefix sets the Prefix field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Prefix field is set to the value of the last call.
func (b *PrefixAllocationStatusApplyConfiguration) WithPrefix(value v1alpha1.IPPrefix) *PrefixAllocationStatusApplyConfiguration {
	b.Prefix = &value
	return b
}

// WithPhase sets the Phase field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Phase field is set to the value of the last call.
func (b *PrefixAllocationStatusApplyConfiguration) WithPhase(value ipamv1alpha1.PrefixAllocationPhase) *PrefixAllocationStatusApplyConfiguration {
	b.Phase = &value
	return b
}

// WithLastPhaseTransitionTime sets the LastPhaseTransitionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastPhaseTransitionTime field is set to the value of the last call.
func (b *PrefixAllocationStatusApplyConfiguration) WithLastPhaseTransitionTime(value v1.Time) *PrefixAllocationStatusApplyConfiguration {
	b.LastPhaseTransitionTime = &value
	return b
}
