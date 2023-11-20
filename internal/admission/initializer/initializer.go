// Copyright 2023 IronCore authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package initializer

import (
	ironcoreinformers "github.com/ironcore-dev/ironcore/client-go/informers"
	"github.com/ironcore-dev/ironcore/client-go/ironcore"
	"github.com/ironcore-dev/ironcore/utils/quota"
	"k8s.io/apiserver/pkg/admission"
)

type initializer struct {
	externalClient    ironcore.Interface
	externalInformers ironcoreinformers.SharedInformerFactory
	quotaRegistry     quota.Registry
}

func New(
	externalClient ironcore.Interface,
	externalInformers ironcoreinformers.SharedInformerFactory,
	quotaRegistry quota.Registry,
) admission.PluginInitializer {
	return &initializer{
		externalClient:    externalClient,
		externalInformers: externalInformers,
		quotaRegistry:     quotaRegistry,
	}
}

func (i *initializer) Initialize(plugin admission.Interface) {
	if wants, ok := plugin.(WantsExternalIronCoreClientSet); ok {
		wants.SetExternalIronCoreClientSet(i.externalClient)
	}

	if wants, ok := plugin.(WantsExternalInformers); ok {
		wants.SetExternalIronCoreInformerFactory(i.externalInformers)
	}

	if wants, ok := plugin.(WantsQuotaRegistry); ok {
		wants.SetQuotaRegistry(i.quotaRegistry)
	}
}

type WantsExternalIronCoreClientSet interface {
	SetExternalIronCoreClientSet(client ironcore.Interface)
	admission.InitializationValidator
}

type WantsExternalInformers interface {
	SetExternalIronCoreInformerFactory(f ironcoreinformers.SharedInformerFactory)
	admission.InitializationValidator
}

type WantsQuotaRegistry interface {
	SetQuotaRegistry(registry quota.Registry)
	admission.InitializationValidator
}
