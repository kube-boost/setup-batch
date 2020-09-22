/*
Copyright 2017 The Kubernetes Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    api "k8s.io/kubernetes/pkg/apis/core"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SetupJob represents the configuration of a job that runs on every node in
// the kubernetes cluster.
type SetupJob struct {
	metav1.TypeMeta   `json:",inline"`

    // Standard object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`

    // Specification of the desired behaviour of a setup job.
	Spec   SetupJobSpec   `json:"spec"`

    // Current status of a setup job.
	Status SetupJobStatus `json:"status"`
}

// SetupJobSpec describes how the setup job execution will look like.
type SetupJobSpec struct {
	// Optional duration in seconds relative to the startTime that the setup 
    // job may be active before the system tries to terminate it; value must 
    // be positive integer
	ActiveDeadlineSeconds *int64

	// Optional number of retries before marking this job failed.
	// Defaults to 6
	BackoffLimit *int32

	// A label query over pods that should match the pod count.
	// Normally, the system sets this field for you.
	Selector *metav1.LabelSelector

	// Describes the pod that will be created when executing a setup job.
	Template api.PodTemplateSpec

	// ttlSecondsAfterFinished limits the lifetime of a SetupJob that has 
    // finished execution (either Complete or Failed). If this field is set, 
    // ttlSecondsAfterFinished after the SetupJob finishes, it is eligible to 
    // be automatically deleted. When the SetupJob is being deleted, its 
    // lifecycle guarantees (e.g. finalizers) will be honored. 
    // If this field is unset, the SetupJob won't be automatically deleted. 
    // If this field is set to zero, the SetupJob becomes eligible to be 
    // deleted immediately after it finishes. 
    // This field is alpha-level and is only honored by servers that enable 
    // the TTLAfterFinished feature.
	TTLSecondsAfterFinished *int32
}

// SetupJobStatus is the status for a SetupJob resource
type SetupJobStatus struct {
	// The latest available observations of an object's current state.
	// When a setup job fails, one of the conditions will have 
    // type == "Failed".
	Conditions []SetupJobCondition `json:"conditions"`

	// Represents time when the setup job was acknowledged by the setup job 
    // controller. It is not guaranteed to be set in happens-before order 
    // across separate operations. It is represented in RFC3339 form and is 
    // in UTC.
	StartTime *metav1.Time `json:"startTime"`

	// Represents time when the setup job was completed. It is not 
    // guaranteed to be set in happens-before order across separate operations.
	// It is represented in RFC3339 form and is in UTC.
	// The completion time is only set when the setup job finishes 
    // successfully.
	CompletionTime *metav1.Time `json:"completionTime"`

	// The number of actively running pods.
	Active int32 `json:"active"`

	// The number of pods which reached phase Succeeded.
	Succeeded int32 `json:"succeeded"`

	// The number of pods which reached phase Failed.
	Failed int32 `json:"failed"`
}

// SetupJobConditionType is a valid value for SetupJobCondition.Type
type SetupJobConditionType string

// These are valid conditions of a setup job.
const (
	// SetupJobComplete means the setup job has completed its execution.
	SetupJobComplete SetupJobConditionType = "Complete"
	// SetupJobFailed means the setup job has failed its execution.
	SetupJobFailed SetupJobConditionType = "Failed"
)

// SetupJobCondition describes current state of a setup job.
type SetupJobCondition struct {
	// Type of setup job condition, Complete or Failed.
	Type SetupJobConditionType

	// Status of the condition, one of True, False, Unknown.
	Status api.ConditionStatus

	// Last time the condition was checked.
	LastProbeTime metav1.Time

	// Last time the condition transit from one status to another.
	LastTransitionTime metav1.Time

	// (brief) reason for the condition's last transition.
	Reason string

	// Human readable message indicating details about last transition.
	Message string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SetupJobList is a list of SetupJob resources.
type SetupJobList struct {
	metav1.TypeMeta `json:",inline"`

    // Standard list metadata.
	metav1.ListMeta `json:"metadata"`

    // Items is the list of SetupJobs.
	Items []SetupJob `json:"items"`
}
