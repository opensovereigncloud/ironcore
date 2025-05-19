---
title: Object Storage User Management

oep-number: 13

creation-date: 2025-05-19

status: proposed

authors:
- @miroslav.soha
- @marek.soha

reviewers:
- TBD

---

# OEP-13: Object Storage User Management

## Table of Contents

- [Summary](#summary)
- [Motivation](#motivation)
    - [Goals](#goals)
    - [Non-Goals](#non-goals)
- [Proposal](#proposal)
    - [User Resource](#user-resource)
    - [Implementation Details](#implementation-details)
- [Alternatives](#alternatives)
- [Implementation Plan](#implementation-plan)

## Summary

This proposal outlines the implementation of user management for the object storage system. It provides S3-compatible access key management and user-specific configurations for bucket access.

## Motivation

Secure and flexible access control is essential for any storage system to:
- Manage user access to buckets
- Control resource usage per user
- Implement security best practices
- Support multi-tenant environments

### Goals

- Implement S3-compatible user management
- Maintain S3 compatibility
- Ensure backward compatibility with existing buckets

### Non-Goals

- Support for access key rotation
- Fine-grained access control
- Complex IAM policies
- User federation
- Multi-factor authentication
- Role-based access control (RBAC)

## Proposal

### User Resource

```yaml
apiVersion: storage.ironcore.dev/v1alpha1
kind: BucketUser
metadata:
  name: user-1
spec:
  userId: userID
  displayName: userDisplayName
  email: user@mail.com
  bucketQuota:
    maxSize: 100Gi
    maxObjects: 1000000
  userQuota:
    maxSize: 100Gi
    maxObjects: 1000000
  accessKeys:
    - user: userID
      keyType: s3
      accessKey: AKIAIOSFODNN7EXAMPLE
      secretKey: SLKA409382LKSJ0-SAK35241
      created: "2024-03-19T10:00:00Z"
status:
  state: Active
  lastAccess: "2024-03-19T10:00:00Z"
  usage:
    userQuota:
      size:
        limit: 100Gi
        used: 50Gi
        available: 50Gi
        percentUsed: 50
      objects:
        limit: 1000000
        used: 500000
        available: 500000
        percentUsed: 50

```

## Implementation Details

### 1. User Management

- **User Lifecycle**
  - User creation and deletion
  - State management
  - Access key lifecycle
  - Usage tracking

- **Security Features**
  - Usage monitoring
  - Access logging

### 2. Integration

- **Bucket Integration**
  - Bucket access control
  - User-specific quotas
  - Usage tracking
  - Access logging

## Alternatives

1. **Use External IAM System**
   - Pros: More feature-rich, battle-tested
   - Cons: Additional complexity, potential integration issues

2. **Implement Custom Authentication**
   - Pros: Full control over implementation
   - Cons: Development overhead, maintenance burden

3. **Use Ceph's Built-in User Management**
   - Pros: Native integration, better performance
   - Cons: Less flexible, Ceph-specific

## Implementation Plan

### Phase 1: Core User Management
1. Implement BucketUser resource
2. Add access key management
3. Implement basic access control

### Phase 2: Integration and Testing
1. Integrate with bucket system
2. Add monitoring and logging
3. Perform security testing 