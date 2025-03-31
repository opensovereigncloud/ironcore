// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	networkingv1alpha1 "github.com/ironcore-dev/ironcore/api/networking/v1alpha1"
	v1 "k8s.io/api/core/v1"
)

// NATGatewaySpecApplyConfiguration represents a declarative configuration of the NATGatewaySpec type for use
// with apply.
type NATGatewaySpecApplyConfiguration struct {
	Type                     *networkingv1alpha1.NATGatewayType `json:"type,omitempty"`
	IPFamily                 *v1.IPFamily                       `json:"ipFamily,omitempty"`
	NetworkRef               *v1.LocalObjectReference           `json:"networkRef,omitempty"`
	PortsPerNetworkInterface *int32                             `json:"portsPerNetworkInterface,omitempty"`
}

// NATGatewaySpecApplyConfiguration constructs a declarative configuration of the NATGatewaySpec type for use with
// apply.
func NATGatewaySpec() *NATGatewaySpecApplyConfiguration {
	return &NATGatewaySpecApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *NATGatewaySpecApplyConfiguration) WithType(value networkingv1alpha1.NATGatewayType) *NATGatewaySpecApplyConfiguration {
	b.Type = &value
	return b
}

// WithIPFamily sets the IPFamily field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IPFamily field is set to the value of the last call.
func (b *NATGatewaySpecApplyConfiguration) WithIPFamily(value v1.IPFamily) *NATGatewaySpecApplyConfiguration {
	b.IPFamily = &value
	return b
}

// WithNetworkRef sets the NetworkRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NetworkRef field is set to the value of the last call.
func (b *NATGatewaySpecApplyConfiguration) WithNetworkRef(value v1.LocalObjectReference) *NATGatewaySpecApplyConfiguration {
	b.NetworkRef = &value
	return b
}

// WithPortsPerNetworkInterface sets the PortsPerNetworkInterface field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PortsPerNetworkInterface field is set to the value of the last call.
func (b *NATGatewaySpecApplyConfiguration) WithPortsPerNetworkInterface(value int32) *NATGatewaySpecApplyConfiguration {
	b.PortsPerNetworkInterface = &value
	return b
}
