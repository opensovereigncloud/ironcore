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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/ironcore-dev/ironcore/client-go/informers/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Machines returns a MachineInformer.
	Machines() MachineInformer
	// MachineClasses returns a MachineClassInformer.
	MachineClasses() MachineClassInformer
	// MachinePools returns a MachinePoolInformer.
	MachinePools() MachinePoolInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Machines returns a MachineInformer.
func (v *version) Machines() MachineInformer {
	return &machineInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MachineClasses returns a MachineClassInformer.
func (v *version) MachineClasses() MachineClassInformer {
	return &machineClassInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// MachinePools returns a MachinePoolInformer.
func (v *version) MachinePools() MachinePoolInformer {
	return &machinePoolInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
