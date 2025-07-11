//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-FileCopyrightText: 2025 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by defaulter-gen. DO NOT EDIT.

package v1alpha1

import (
	computev1alpha1 "github.com/ironcore-dev/ironcore/api/compute/v1alpha1"
	ipamv1alpha1 "github.com/ironcore-dev/ironcore/internal/apis/ipam/v1alpha1"
	networkingv1alpha1 "github.com/ironcore-dev/ironcore/internal/apis/networking/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&computev1alpha1.Machine{}, func(obj interface{}) { SetObjectDefaults_Machine(obj.(*computev1alpha1.Machine)) })
	scheme.AddTypeDefaultingFunc(&computev1alpha1.MachineList{}, func(obj interface{}) { SetObjectDefaults_MachineList(obj.(*computev1alpha1.MachineList)) })
	return nil
}

func SetObjectDefaults_Machine(in *computev1alpha1.Machine) {
	SetDefaults_MachineSpec(&in.Spec)
	for i := range in.Spec.NetworkInterfaces {
		a := &in.Spec.NetworkInterfaces[i]
		if a.NetworkInterfaceSource.Ephemeral != nil {
			if a.NetworkInterfaceSource.Ephemeral.NetworkInterfaceTemplate != nil {
				networkingv1alpha1.SetDefaults_NetworkInterfaceSpec(&a.NetworkInterfaceSource.Ephemeral.NetworkInterfaceTemplate.Spec)
				for j := range a.NetworkInterfaceSource.Ephemeral.NetworkInterfaceTemplate.Spec.IPs {
					b := &a.NetworkInterfaceSource.Ephemeral.NetworkInterfaceTemplate.Spec.IPs[j]
					if b.Ephemeral != nil {
						if b.Ephemeral.PrefixTemplate != nil {
							ipamv1alpha1.SetDefaults_PrefixSpec(&b.Ephemeral.PrefixTemplate.Spec)
						}
					}
				}
				for j := range a.NetworkInterfaceSource.Ephemeral.NetworkInterfaceTemplate.Spec.Prefixes {
					b := &a.NetworkInterfaceSource.Ephemeral.NetworkInterfaceTemplate.Spec.Prefixes[j]
					if b.Ephemeral != nil {
						if b.Ephemeral.PrefixTemplate != nil {
							ipamv1alpha1.SetDefaults_PrefixSpec(&b.Ephemeral.PrefixTemplate.Spec)
						}
					}
				}
			}
		}
	}
	SetDefaults_MachineStatus(&in.Status)
	for i := range in.Status.NetworkInterfaces {
		a := &in.Status.NetworkInterfaces[i]
		SetDefaults_NetworkInterfaceStatus(a)
	}
	for i := range in.Status.Volumes {
		a := &in.Status.Volumes[i]
		SetDefaults_VolumeStatus(a)
	}
}

func SetObjectDefaults_MachineList(in *computev1alpha1.MachineList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Machine(a)
	}
}
