---
title: Object Storage Quota Management

oep-number: 12

creation-date: 2025-05-19

status: proposed

authors:
- @miroslav.soha
- @marek.soha

reviewers:
- TBD

---

# OEP-12: Object Storage Quota Management

## Table of Contents

- [Summary](#summary)
- [Motivation](#motivation)
    - [Goals](#goals)
    - [Non-Goals](#non-goals)
- [Proposal](#proposal)
    - [ResourceQuota Extension](#resourcequota-extension)
    - [New Resource Scope](#new-resource-scope)
    - [New Scope Selector Operator](#new-scope-selector-operator)
    - [Object Storage Resource Types](#object-storage-resource-types)
    - [Implementation Details](#implementation-details)
- [Alternatives](#alternatives)
- [Implementation Plan](#implementation-plan)

## Summary

This proposal extends the existing quota management system (IEP-7) to support object storage quotas for buckets. It enables administrators to control resource usage by setting limits on bucket size, object count, and individual upload sizes using the existing `ResourceQuota` mechanism with new scope selectors for bucket-specific quotas.

## Motivation

Resource management is crucial for any storage system to prevent:
- Uncontrolled growth of storage usage
- Resource exhaustion
- Cost overruns
- Performance degradation

The existing quota system (IEP-7) provides a solid foundation for resource management but lacks support for object storage-specific quotas and bucket-level targeting.

### Goals

- Extend existing `ResourceQuota` to support object storage quotas
- Enable quota enforcement at bucket level using scope selectors
- Support size-based and object count quotas for buckets
- Maintain backward compatibility with existing quota system
- Ensure S3 compatibility
- Provide quota status monitoring through existing mechanisms

### Non-Goals

- User-level quotas (covered in separate proposal)
- Automatic quota adjustment
- Cross-bucket quota management
- Cost-based quota management
- Creating new quota resource types

## Proposal

### ResourceQuota Extension

This proposal extends the existing `ResourceQuota` system by adding new resource types and scope selectors for object storage. No new resource types are introduced; instead, we leverage the existing `ResourceQuota` infrastructure.

### New Resource Scope

Add a new resource scope to target specific buckets:

```go
const (
    // ResourceScopeBucketName refers to the name of a specific bucket
    ResourceScopeBucketName ResourceScope = "Bucket.Name"
)
```

This scope allows quotas to target specific buckets by name, enabling fine-grained quota management.

### New Scope Selector Operator

Add a new operator for exact matching:

```go
const (
    ResourceScopeSelectorOperatorEquals ResourceScopeSelectorOperator = "Equals"
)
```

This operator enables exact matching for bucket names and other resources, complementing the existing `In`, `NotIn`, `Exists`, and `DoesNotExist` operators.

### Object Storage Resource Types

Extend the existing resource types to include object storage-specific resources:

| Resource Name | Description |
|---------------|-------------|
| `requests.storage` | Total bucket size limit (existing, extended for buckets) |
| `count/buckets.storage.ironcore.dev` | Number of buckets (existing) |
| `requests.objects` | Total object count per bucket |


### Implementation Details

#### 1. Bucket Evaluator Extension

Extend the existing bucket evaluator in `internal/quota/evaluator/storage/bucket.go` to support:

- Object count tracking
- Upload size limits
- Bucket name scope matching

#### 2. Usage Calculation

Extend the bucket evaluator's `Usage` method to calculate object storage metrics:

#### 3. Quota Enforcement

Quota enforcement will be handled by the existing admission controller, which will:

- Check bucket creation/update requests against applicable quotas
- Validate upload size limits during object upload operations
- Reject operations that would exceed quota limits

### Example Manifests

#### Limit total storage across all buckets in a namespace:

```yaml
apiVersion: core.ironcore.dev/v1alpha1
kind: ResourceQuota
metadata:
  name: bucket-storage-limit
spec:
  hard:
    requests.storage: 100Gi
    count/buckets.storage.ironcore.dev: 10
```

#### Limit storage for a specific bucket:

```yaml
apiVersion: core.ironcore.dev/v1alpha1
kind: ResourceQuota
metadata:
  name: specific-bucket-limit
spec:
  hard:
    requests.storage: 50Gi
    requests.objects: 1000000
  scopeSelector:
    matchExpressions:
    - scopeName: Bucket.Name
      operator: Equals
      values:
        - my-bucket
```

#### Limit storage for buckets of a specific class:

```yaml
apiVersion: core.ironcore.dev/v1alpha1
kind: ResourceQuota
metadata:
  name: premium-bucket-limit
spec:
  hard:
    requests.storage: 200Gi
    requests.objects: 5000000
  scopeSelector:
    matchExpressions:
    - scopeName: BucketClass
      operator: In
      values:
        - premium
```

## Alternatives

1. **Use Ceph's Built-in Quota System**
   - Pros: Native integration, better performance
   - Cons: Less flexible, Ceph-specific, doesn't integrate with existing quota system

2. **Implement at Storage Layer**
   - Pros: More efficient enforcement
   - Cons: Less flexible, harder to manage, doesn't leverage existing quota infrastructure

3. **Create New BucketQuota Resource**
   - Pros: Dedicated resource for bucket quotas
   - Cons: Duplicates existing quota functionality, increases complexity

4. **Use External Quota Management**
   - Pros: Separation of concerns
   - Cons: Additional complexity, potential sync issues

## Implementation Plan

### Phase 1: Core Extensions
1. Add `ResourceScopeBucketName` constant
2. Add `ResourceScopeSelectorOperatorEquals` constant
3. Extend bucket evaluator to support new scope and operator
4. Add object count and upload size resource types

### Phase 2: Usage Tracking
1. Implement bucket usage calculation from storage provider
2. Extend bucket evaluator to track object counts and sizes
3. Add metrics collection for quota monitoring

### Phase 3: Enforcement and Testing
1. Extend admission controller for bucket-specific quota enforcement
2. Add integration tests for bucket quotas
3. Update documentation and examples 