---
title: Direct Ceph Object Storage Integration

oep-number: 11

creation-date: 2025-05-19

status: proposed

authors:
- @miroslav.soha
- @marek.soha

reviewers:
- TBD

---

# OEP-11: Direct Ceph Object Storage Integration

## Table of Contents

- [Summary](#summary)
- [Motivation](#motivation)
    - [Goals](#goals)
    - [Non-Goals](#non-goals)
- [Proposal](#proposal)
    - [Architecture Overview](#architecture-overview)
    - [API Resources](#api-resources)
    - [Implementation Details](#implementation-details)
- [Alternatives](#alternatives)
- [Migration Path](#migration-path)

## Summary

This proposal outlines the transition from Rook-managed Ceph object storage to direct Ceph integration for S3-compatible object storage in IronCore. The new implementation will interact directly with the Ceph RADOS Gateway (RGW) API, eliminating the Rook operator dependency while maintaining S3 compatibility.

## Motivation

The current implementation relies on Rook for managing Ceph object storage, which adds an additional layer of complexity and potential points of failure. By integrating directly with Ceph's RGW, we can:
- Reduce operational complexity
- Improve performance by removing an abstraction layer
- Have more direct control over bucket lifecycle and configuration
- Enable more granular monitoring and management capabilities

### Goals

- Remove Rook dependency for object storage management
- Maintain S3 compatibility
- Direct integration with Ceph RGW API
- Simplified bucket lifecycle management
- Improved error handling and status reporting
- Maintain existing bucket provisioning workflow

### Non-Goals

- Support for non-Ceph storage backends
- Implementation of bucket user management (covered in separate proposal)
- Implementation of usage monitoring and quotas (covered in separate proposal)
- Changes to the existing bucket API structure

## Proposal

### Architecture Overview

The new architecture will consist of:
1. A dedicated controller for managing Ceph RGW integration
2. Direct communication with Ceph RGW API
3. Enhanced status reporting and error handling
4. Simplified bucket lifecycle management

### API Resources

The existing `Bucket`, `BucketClass`, and `BucketPool` resources will be maintained.

#### Bucket

```yaml
apiVersion: storage.ironcore.dev/v1alpha1
kind: Bucket
metadata:
  name: bucket-1
spec:
  bucketClassRef:
    name: slow
  bucketPoolSelector:
    matchLabels:
      key: db
      foo: bar
      zone: default
  bucketPoolRef:
    name: ceph-pool
status:
  access:
    endpoint: foo.bar.example.org 
    secretRef:
      name: 000225194345f27a40257c5777c96a03ce219f96731f22afc45b7dfda7d077d
  state: Available
```
### BucketClass

```yaml
apiVersion: storage.ironcore.dev/v1alpha1
kind: BucketClass
metadata:
  name: slow
capabilities:
  iops: 10
  tps: 20Mi
```

#### BucketPool

```yaml
apiVersion: storage.ironcore.dev/v1alpha1
kind: BucketPool
metadata:
  name: ceph-pool
spec:
  providerID: ceph://pool
  taints: []
status:
  availableBucketClasses:
    - name: fast
    - name: slow
  state: Available
```

### Implementation Details

1. **Ceph RGW Controller**
   - New controller component for managing Ceph RGW integration
   - Handles bucket creation, deletion, and updates
   - Manages Ceph-specific configurations
   - Provides enhanced status reporting

2. **Direct API Integration**
   - Direct communication with Ceph RGW REST API
   - Support for Ceph's native bucket operations
   - Enhanced error handling and status reporting

3. **Configuration Management**
   - Ceph-specific configuration options
   - Support for multiple RGW endpoints
   - Zone and placement configuration

4. **Status Reporting**
   - Enhanced status information
   - Detailed error reporting
   - Bucket synchronization status

## Alternatives

1. **Keep Rook Integration**
   - Pros: Maintains existing architecture, less development effort
   - Cons: Additional complexity, potential performance impact

2. **Use Ceph Operator Instead of Rook**
   - Pros: More Ceph-specific features, better Ceph integration
   - Cons: Still adds complexity, potential migration challenges

3. **Implement Custom Ceph Client**
   - Pros: Full control over implementation
   - Cons: Significant development effort, maintenance overhead

## Migration Path

1. **Phase 1: Preparation**
   - Update API resources

1. **Phase 2: Implementation**
   - Implement new Ceph RGW controller
   - Add Ceph-specific configurations

2. **Phase 3: Testing**
   - Test with new implementation
   - Validate S3 compatibility
   - Performance testing

3. **Phase 4: Migration**
   - Gradual migration of existing buckets
   - Validation of migrated buckets
   - Removal of Rook dependency

4. **Phase 5: Cleanup**
   - Remove Rook-related code
   - Update documentation
   - Final validation 