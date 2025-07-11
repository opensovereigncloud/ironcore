// SPDX-FileCopyrightText: 2025 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	commonv1alpha1 "github.com/ironcore-dev/ironcore/api/common/v1alpha1"
)

// NATGatewayStatusApplyConfiguration represents a declarative configuration of the NATGatewayStatus type for use
// with apply.
type NATGatewayStatusApplyConfiguration struct {
	IPs []commonv1alpha1.IP `json:"ips,omitempty"`
}

// NATGatewayStatusApplyConfiguration constructs a declarative configuration of the NATGatewayStatus type for use with
// apply.
func NATGatewayStatus() *NATGatewayStatusApplyConfiguration {
	return &NATGatewayStatusApplyConfiguration{}
}

// WithIPs adds the given value to the IPs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the IPs field.
func (b *NATGatewayStatusApplyConfiguration) WithIPs(values ...commonv1alpha1.IP) *NATGatewayStatusApplyConfiguration {
	for i := range values {
		b.IPs = append(b.IPs, values[i])
	}
	return b
}
