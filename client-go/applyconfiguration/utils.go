/*
Copyright 2023 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	schema "k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha2 "sigs.k8s.io/jobset/api/jobset/v1alpha2"
	jobsetv1alpha2 "sigs.k8s.io/jobset/client-go/applyconfiguration/jobset/v1alpha2"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=jobset.x-k8s.io, Version=v1alpha2
	case v1alpha2.SchemeGroupVersion.WithKind("FailurePolicy"):
		return &jobsetv1alpha2.FailurePolicyApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("JobSet"):
		return &jobsetv1alpha2.JobSetApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("JobSetSpec"):
		return &jobsetv1alpha2.JobSetSpecApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("JobSetStatus"):
		return &jobsetv1alpha2.JobSetStatusApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("Network"):
		return &jobsetv1alpha2.NetworkApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("ReplicatedJob"):
		return &jobsetv1alpha2.ReplicatedJobApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("ReplicatedJobStatus"):
		return &jobsetv1alpha2.ReplicatedJobStatusApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("SuccessPolicy"):
		return &jobsetv1alpha2.SuccessPolicyApplyConfiguration{}

	}
	return nil
}
