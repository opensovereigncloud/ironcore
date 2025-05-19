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
    - [Quota Resource](#quota-resource)
    - [Implementation Details](#implementation-details)
- [Alternatives](#alternatives)
- [Implementation Plan](#implementation-plan)

## Summary

This proposal outlines the implementation of quota management for the object storage system. It enables administrators to control resource usage by setting limits on bucket size, object count, and individual upload sizes, with configurable enforcement policies and grace periods.

## Motivation

Resource management is crucial for any storage system to prevent:
- Uncontrolled growth of storage usage
- Resource exhaustion
- Cost overruns
- Performance degradation

### Goals

- Enable quota enforcement at bucket level
- Support size-based and object count quotas
- Implement configurable enforcement policies
- Provide quota status monitoring
- Maintain S3 compatibility
- Ensure backward compatibility with existing buckets

### Non-Goals

- User-level quotas (covered in separate proposal)
- Automatic quota adjustment
- Cross-bucket quota management
- Cost-based quota management

## Proposal

### Quota Resource

```yaml
apiVersion: storage.ironcore.dev/v1alpha1
kind: BucketQuota
metadata:
  name: bucket-1-quota
spec:
  bucketRef:
    name: bucket-1
  limits:
    maxSize: 100Gi
    maxObjects: 1000000
status:
  enabled: true
  currentUsage:
    size: 50Gi
    objectCount: 500000
  quotaStatus:
    size: 50%
    objects: 50%
  lastEnforced: "2024-03-19T10:00:00Z"
  warnings:
    - type: SizeWarning
      message: "Bucket size approaching quota limit"
      severity: warning
      timestamp: "2024-03-19T09:55:00Z"
```

## Implementation Details

### 1. Quota Enforcement

- **Size-based Quotas**
  - Total bucket size limit
  - Per-upload size limit
  - Real-time size tracking

- **Object Count Quotas**
  - Maximum objects per bucket
  - Object count tracking
  - Batch operation handling

### 2. Monitoring and Reporting

- **Quota Status**
  - Current usage tracking
  - Percentage-based status
  - Historical usage data

- **Warning System**
  - Threshold-based warnings
  - Multiple warning levels
  - Configurable notification channels

## Alternatives

1. **Use Ceph's Built-in Quota System**
   - Pros: Native integration, better performance
   - Cons: Less flexible, Ceph-specific

2. **Implement at Storage Layer**
   - Pros: More efficient enforcement
   - Cons: Less flexible, harder to manage

3. **Use External Quota Management**
   - Pros: Separation of concerns
   - Cons: Additional complexity, potential sync issues

## Implementation Plan

### Phase 1: Core Quota Implementation
1. Implement BucketQuota resource
2. Implement BucketQuota enabling/disabling
3. Implement size and object count tracking

### Phase 2: Monitoring and Integration
1. Implement quota status monitoring
2. Add historical data collection
3. Set up notification system

### Phase 3: Testing and Validation
1. Performance testing
2. Edge case handling
3. Documentation and examples 