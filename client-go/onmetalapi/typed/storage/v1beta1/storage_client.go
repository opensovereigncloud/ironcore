/*
 * Copyright (c) 2022 by the OnMetal authors.
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
// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"net/http"

	v1beta1 "github.com/onmetal/onmetal-api/api/storage/v1beta1"
	"github.com/onmetal/onmetal-api/client-go/onmetalapi/scheme"
	rest "k8s.io/client-go/rest"
)

type StorageV1beta1Interface interface {
	RESTClient() rest.Interface
	BucketsGetter
	BucketClassesGetter
	BucketPoolsGetter
	VolumesGetter
	VolumeClassesGetter
	VolumePoolsGetter
}

// StorageV1beta1Client is used to interact with features provided by the storage.api.onmetal.de group.
type StorageV1beta1Client struct {
	restClient rest.Interface
}

func (c *StorageV1beta1Client) Buckets(namespace string) BucketInterface {
	return newBuckets(c, namespace)
}

func (c *StorageV1beta1Client) BucketClasses() BucketClassInterface {
	return newBucketClasses(c)
}

func (c *StorageV1beta1Client) BucketPools() BucketPoolInterface {
	return newBucketPools(c)
}

func (c *StorageV1beta1Client) Volumes(namespace string) VolumeInterface {
	return newVolumes(c, namespace)
}

func (c *StorageV1beta1Client) VolumeClasses() VolumeClassInterface {
	return newVolumeClasses(c)
}

func (c *StorageV1beta1Client) VolumePools() VolumePoolInterface {
	return newVolumePools(c)
}

// NewForConfig creates a new StorageV1beta1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*StorageV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new StorageV1beta1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*StorageV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &StorageV1beta1Client{client}, nil
}

// NewForConfigOrDie creates a new StorageV1beta1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *StorageV1beta1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new StorageV1beta1Client for the given RESTClient.
func New(c rest.Interface) *StorageV1beta1Client {
	return &StorageV1beta1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1beta1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *StorageV1beta1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
