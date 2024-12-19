// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1alpha1 "github.com/ironcore-dev/ironcore/api/storage/v1alpha1"
	storagev1alpha1 "github.com/ironcore-dev/ironcore/client-go/applyconfigurations/storage/v1alpha1"
	scheme "github.com/ironcore-dev/ironcore/client-go/ironcore/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// VolumePoolsGetter has a method to return a VolumePoolInterface.
// A group's client should implement this interface.
type VolumePoolsGetter interface {
	VolumePools() VolumePoolInterface
}

// VolumePoolInterface has methods to work with VolumePool resources.
type VolumePoolInterface interface {
	Create(ctx context.Context, volumePool *v1alpha1.VolumePool, opts v1.CreateOptions) (*v1alpha1.VolumePool, error)
	Update(ctx context.Context, volumePool *v1alpha1.VolumePool, opts v1.UpdateOptions) (*v1alpha1.VolumePool, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, volumePool *v1alpha1.VolumePool, opts v1.UpdateOptions) (*v1alpha1.VolumePool, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.VolumePool, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.VolumePoolList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VolumePool, err error)
	Apply(ctx context.Context, volumePool *storagev1alpha1.VolumePoolApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.VolumePool, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, volumePool *storagev1alpha1.VolumePoolApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.VolumePool, err error)
	VolumePoolExpansion
}

// volumePools implements VolumePoolInterface
type volumePools struct {
	*gentype.ClientWithListAndApply[*v1alpha1.VolumePool, *v1alpha1.VolumePoolList, *storagev1alpha1.VolumePoolApplyConfiguration]
}

// newVolumePools returns a VolumePools
func newVolumePools(c *StorageV1alpha1Client) *volumePools {
	return &volumePools{
		gentype.NewClientWithListAndApply[*v1alpha1.VolumePool, *v1alpha1.VolumePoolList, *storagev1alpha1.VolumePoolApplyConfiguration](
			"volumepools",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1alpha1.VolumePool { return &v1alpha1.VolumePool{} },
			func() *v1alpha1.VolumePoolList { return &v1alpha1.VolumePoolList{} }),
	}
}
