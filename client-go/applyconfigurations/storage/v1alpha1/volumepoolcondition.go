// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/ironcore-dev/ironcore/api/storage/v1alpha1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VolumePoolConditionApplyConfiguration represents an declarative configuration of the VolumePoolCondition type for use
// with apply.
type VolumePoolConditionApplyConfiguration struct {
	Type               *v1alpha1.VolumePoolConditionType `json:"type,omitempty"`
	Status             *v1.ConditionStatus               `json:"status,omitempty"`
	Reason             *string                           `json:"reason,omitempty"`
	Message            *string                           `json:"message,omitempty"`
	ObservedGeneration *int64                            `json:"observedGeneration,omitempty"`
	LastTransitionTime *metav1.Time                      `json:"lastTransitionTime,omitempty"`
}

// VolumePoolConditionApplyConfiguration constructs an declarative configuration of the VolumePoolCondition type for use with
// apply.
func VolumePoolCondition() *VolumePoolConditionApplyConfiguration {
	return &VolumePoolConditionApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *VolumePoolConditionApplyConfiguration) WithType(value v1alpha1.VolumePoolConditionType) *VolumePoolConditionApplyConfiguration {
	b.Type = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *VolumePoolConditionApplyConfiguration) WithStatus(value v1.ConditionStatus) *VolumePoolConditionApplyConfiguration {
	b.Status = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *VolumePoolConditionApplyConfiguration) WithReason(value string) *VolumePoolConditionApplyConfiguration {
	b.Reason = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *VolumePoolConditionApplyConfiguration) WithMessage(value string) *VolumePoolConditionApplyConfiguration {
	b.Message = &value
	return b
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *VolumePoolConditionApplyConfiguration) WithObservedGeneration(value int64) *VolumePoolConditionApplyConfiguration {
	b.ObservedGeneration = &value
	return b
}

// WithLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTransitionTime field is set to the value of the last call.
func (b *VolumePoolConditionApplyConfiguration) WithLastTransitionTime(value metav1.Time) *VolumePoolConditionApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}
