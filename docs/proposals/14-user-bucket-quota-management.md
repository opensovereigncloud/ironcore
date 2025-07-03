---
title: User Bucket Quota Management

oep-number: 14

creation-date: 2025-05-19

status: proposed

authors:
- @miroslav.soha
- @marek.soha

reviewers:
- TBD

---

# OEP-14: User Bucket Quota Management

## Table of Contents

- [Summary](#summary)
- [Motivation](#motivation)
    - [Goals](#goals)
    - [Non-Goals](#non-goals)
- [Proposal](#proposal)
    - [ResourceQuota Extension](#resourcequota-extension)
    - [New Resource Scopes](#new-resource-scopes)
    - [User-Specific Resource Types](#user-specific-resource-types)
    - [Implementation Details](#implementation-details)
- [Alternatives](#alternatives)
- [Implementation Plan](#implementation-plan)

## Summary

This proposal extends the existing quota management system (IEP-7) to support user-level bucket quotas. It enables administrators to control resource usage by setting limits on storage and object counts at both user and bucket levels using the existing `ResourceQuota` mechanism with new scope selectors for user-specific quotas.

## Motivation

Resource management and cost control are critical for storage systems to:
- Prevent resource exhaustion
- Control storage costs
- Implement fair usage policies
- Support multi-tenant environments with resource isolation

### Goals

- Extend existing `ResourceQuota` to support user-level bucket quotas
- Enable quota enforcement at user level using scope selectors
- Support both user-level and bucket-level quota management
- Maintain backward compatibility with existing quota system
- Provide real-time usage tracking through existing mechanisms
- Support multi-tenant environments with resource isolation

### Non-Goals

- Dynamic quota adjustment during operations
- Complex quota inheritance rules
- Time-based quota limits
- Quota sharing between users
- Creating new quota resource types

## Proposal

### ResourceQuota Extension

This proposal extends the existing `ResourceQuota` system by adding new resource types and scope selectors for user-level bucket management. No new resource types are introduced; instead, we leverage the existing `ResourceQuota` infrastructure.

### New Resource Scopes

Add new resource scopes to target users and user-bucket relationships:

```go
const (
    // ResourceScopeUser refers to the user who owns the bucket
    ResourceScopeUser ResourceScope = "User"
    
    // ResourceScopeUserBucket refers to the user-bucket relationship
    ResourceScopeUserBucket ResourceScope = "User.Bucket"
)
```

These scopes allow quotas to target specific users or user-bucket combinations, enabling fine-grained quota management in multi-tenant environments.

### User-Specific Resource Types

Extend the existing resource types to include user-specific resources:

| Resource Name | Description |
|---------------|-------------|
| `requests.storage` | Total storage limit (existing, extended for user aggregation) |
| `count/buckets.storage.ironcore.dev` | Number of buckets (existing, extended for user aggregation) |
| `requests.objects` | Total object count (existing, extended for user aggregation) |
| `requests.user.storage` | Per-user storage limit |
| `requests.user.objects` | Per-user object count limit |
| `count/user.buckets.storage.ironcore.dev` | Per-user bucket count limit |

### Implementation Details

#### 1. User and Bucket Association

Extend the bucket resource to include user information:

#### 2. User Scope Matching

Extend the bucket evaluator to support user-based scope matching:

#### 3. User-Bucket Scope Matching

Support matching on user-bucket combinations:

#### 4. Usage Calculation

Extend the bucket evaluator's `Usage` method to calculate user-aggregated metrics:

#### 5. Quota Enforcement

Quota enforcement will be handled by the existing admission controller, which will:

- Check bucket creation/update requests against applicable user quotas
- Validate user-level limits during bucket operations
- Reject operations that would exceed user quota limits
- Support both hard and soft quota limits

### Example Manifests

#### Limit total storage per user across all their buckets:

```yaml
apiVersion: core.ironcore.dev/v1alpha1
kind: ResourceQuota
metadata:
  name: user-storage-limit
spec:
  hard:
    requests.user.storage: 100Gi
    requests.user.objects: 1000000
    count/user.buckets.storage.ironcore.dev: 5
  scopeSelector:
    matchExpressions:
    - scopeName: User
      operator: In
      values:
        - user1
        - user2
```

#### Limit storage for a specific user:

```yaml
apiVersion: core.ironcore.dev/v1alpha1
kind: ResourceQuota
metadata:
  name: specific-user-limit
spec:
  hard:
    requests.user.storage: 50Gi
    requests.user.objects: 500000
    count/user.buckets.storage.ironcore.dev: 3
  scopeSelector:
    matchExpressions:
    - scopeName: User
      operator: Equals
      values:
        - premium-user
```

#### Limit storage for a specific user-bucket combination:

```yaml
apiVersion: core.ironcore.dev/v1alpha1
kind: ResourceQuota
metadata:
  name: user-bucket-limit
spec:
  hard:
    requests.storage: 25Gi
    requests.objects: 250000
  scopeSelector:
    matchExpressions:
    - scopeName: User.Bucket
      operator: Equals
      values:
        - user1/bucket1
```

## Alternatives

1. **Use Ceph's Built-in Quota System**
   - Pros: Native integration, better performance
   - Cons: Less flexible, Ceph-specific, doesn't integrate with existing quota system

2. **Create New UserQuota Resource**
   - Pros: Dedicated resource for user quotas
   - Cons: Duplicates existing quota functionality, increases complexity

3. **External Quota Management**
   - Pros: Separation of concerns
   - Cons: Additional complexity, potential sync issues, doesn't integrate with existing system

4. **Simple Size-based Quotas Only**
   - Pros: Simpler implementation, easier to understand
   - Cons: Less granular control, limited functionality

## Implementation Plan

### Phase 1: Core Extensions
1. Add `ResourceScopeUser` and `ResourceScopeUserBucket` constants
2. Extend bucket spec to include user reference
3. Extend bucket evaluator to support user scopes
4. Add user-specific resource types

### Phase 2: Usage Tracking
1. Implement user-aggregated usage calculation
2. Extend bucket evaluator to track user-specific metrics
3. Add user-bucket relationship tracking

### Phase 3: Enforcement and Testing
1. Extend admission controller for user-specific quota enforcement
2. Add integration tests for user quotas
3. Update documentation and examples