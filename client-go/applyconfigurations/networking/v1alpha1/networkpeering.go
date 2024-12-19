// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// NetworkPeeringApplyConfiguration represents a declarative configuration of the NetworkPeering type for use
// with apply.
type NetworkPeeringApplyConfiguration struct {
	Name       *string                                     `json:"name,omitempty"`
	NetworkRef *NetworkPeeringNetworkRefApplyConfiguration `json:"networkRef,omitempty"`
	Prefixes   []PeeringPrefixApplyConfiguration           `json:"prefixes,omitempty"`
}

// NetworkPeeringApplyConfiguration constructs a declarative configuration of the NetworkPeering type for use with
// apply.
func NetworkPeering() *NetworkPeeringApplyConfiguration {
	return &NetworkPeeringApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *NetworkPeeringApplyConfiguration) WithName(value string) *NetworkPeeringApplyConfiguration {
	b.Name = &value
	return b
}

// WithNetworkRef sets the NetworkRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NetworkRef field is set to the value of the last call.
func (b *NetworkPeeringApplyConfiguration) WithNetworkRef(value *NetworkPeeringNetworkRefApplyConfiguration) *NetworkPeeringApplyConfiguration {
	b.NetworkRef = value
	return b
}

// WithPrefixes adds the given value to the Prefixes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Prefixes field.
func (b *NetworkPeeringApplyConfiguration) WithPrefixes(values ...*PeeringPrefixApplyConfiguration) *NetworkPeeringApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPrefixes")
		}
		b.Prefixes = append(b.Prefixes, *values[i])
	}
	return b
}
