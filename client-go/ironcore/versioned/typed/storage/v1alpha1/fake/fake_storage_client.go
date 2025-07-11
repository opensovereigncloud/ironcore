// SPDX-FileCopyrightText: 2025 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/ironcore-dev/ironcore/client-go/ironcore/versioned/typed/storage/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeStorageV1alpha1 struct {
	*testing.Fake
}

func (c *FakeStorageV1alpha1) Buckets(namespace string) v1alpha1.BucketInterface {
	return newFakeBuckets(c, namespace)
}

func (c *FakeStorageV1alpha1) BucketClasses() v1alpha1.BucketClassInterface {
	return newFakeBucketClasses(c)
}

func (c *FakeStorageV1alpha1) BucketPools() v1alpha1.BucketPoolInterface {
	return newFakeBucketPools(c)
}

func (c *FakeStorageV1alpha1) Volumes(namespace string) v1alpha1.VolumeInterface {
	return newFakeVolumes(c, namespace)
}

func (c *FakeStorageV1alpha1) VolumeClasses() v1alpha1.VolumeClassInterface {
	return newFakeVolumeClasses(c)
}

func (c *FakeStorageV1alpha1) VolumePools() v1alpha1.VolumePoolInterface {
	return newFakeVolumePools(c)
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeStorageV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
