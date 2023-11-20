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
	v1alpha1 "github.com/ironcore-dev/ironcore/api/storage/v1alpha1"
	v1 "k8s.io/api/core/v1"
)

// BucketPoolStatusApplyConfiguration represents an declarative configuration of the BucketPoolStatus type for use
// with apply.
type BucketPoolStatusApplyConfiguration struct {
	State                  *v1alpha1.BucketPoolState `json:"state,omitempty"`
	AvailableBucketClasses []v1.LocalObjectReference `json:"availableBucketClasses,omitempty"`
}

// BucketPoolStatusApplyConfiguration constructs an declarative configuration of the BucketPoolStatus type for use with
// apply.
func BucketPoolStatus() *BucketPoolStatusApplyConfiguration {
	return &BucketPoolStatusApplyConfiguration{}
}

// WithState sets the State field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the State field is set to the value of the last call.
func (b *BucketPoolStatusApplyConfiguration) WithState(value v1alpha1.BucketPoolState) *BucketPoolStatusApplyConfiguration {
	b.State = &value
	return b
}

// WithAvailableBucketClasses adds the given value to the AvailableBucketClasses field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AvailableBucketClasses field.
func (b *BucketPoolStatusApplyConfiguration) WithAvailableBucketClasses(values ...v1.LocalObjectReference) *BucketPoolStatusApplyConfiguration {
	for i := range values {
		b.AvailableBucketClasses = append(b.AvailableBucketClasses, values[i])
	}
	return b
}
