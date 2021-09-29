/*
 Copyright © 2021 The OpenEBS Authors

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=storagecohort

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type StorageCohort struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the behavior of a cohort.
	// +required
	Spec StorageCohortSpec `json:"spec"`

	// Most recently observed status of the cohort.
	// Populated by the cohort operator or cohort manager.
	// +optional
	Status StorageCohortStatus `json:"status,omitempty"`
}

// StorageCohortSpec describes the attributes that a cohort is created with.
type StorageCohortSpec struct {
	// NodeSelector is used to specify the cohort to be considered
	// +optional
	NodeSelector *metav1.LabelSelector `json:"nodeSelector,omitempty"`

	// CohortManager stores all the details about the cohort manager responsible for
	// managing the cohort
	// +optional
	CohortManager interface{} `json:"cohortManager,omitempty"`

	// DefaultStorageProvisioner is the default provisioner for the cohort which can be used
	// for provisioning pools or volumes when no provisioner is specified in the storage pool
	// For example: "openebs.io/scp-lvm-provisioner" or "openebs.io/scp-device-provisioner".
	// +optional
	DefaultStorageProvisioner string `json:"defaultStorageProvisioner,omitempty"`
}

// StorageCohortStatus stores information about the current status of a storage cohort.
type StorageCohortStatus struct {
	// Components is an array of different component's conditions which the cohort is comprised of.
	// +optional
	Components []ComponentStatus `json:"components,omitempty"`

	// Capabilities represent capabilities that a cohort consists of
	// +optional
	Capabilities Capabilities `json:"capabilities,omitempty"`
}

// ComponentStatus stores information about the current status of storage cohort's components.
// Note: For scheduling purpose, the scheduler will only br concerned with the CohortCondition
// to make scheduling decisions. Other components status can be used for monitoring purposes
// or troubleshooting purpose.
type ComponentStatus struct {
	// CohortCondition is an array of current observed cohort conditions.
	// The Cohort is deemed to be fully functional when its Ready and Schedulable
	// condition types are `true`. All other types status declares a cohort
	// to be non-functional.
	// +optional
	CohortCondition []CohortCondition `json:"cohortCondition,omitempty"`

	// CohortManagerCondition is an array of current observed cohort manager conditions.
	// +optional
	CohortManagerCondition []ComponentCondition `json:"cohortManagerCondition,omitempty"`

	// NodeCondition is an array of current observed cohort's individual nodes conditions.
	// +optional
	NodeCondition []CohortNodeCondition `json:"nodeCondition,omitempty"`
}

// CohortNodeCondition contains the latest status information for some or all the
// nodes that the cohort is comprised of.
type CohortNodeCondition struct {
	// Name of the node. This must be a DNS_LABEL.
	// For example: "virtual-node-1"
	// +optional
	Name string `json:"name"`

	// Condition is an array of current observed node conditions.
	// +optional
	Condition []ComponentCondition `json:"condition,omitempty"`
}

// CohortCondition contains condition information for a storage cohort.
type CohortCondition struct {
	// Type of component condition.
	Type CohortConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status corev1.ConditionStatus `json:"status"`
	// Last time we got an update on a given condition.
	LastHeartbeatTime metav1.Time `json:"lastHeartbeatTime,omitempty"`
	// Last time the condition transit from one status to another.
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// (brief) reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`
	// Human readable message indicating details about last transition.
	Message string `json:"message,omitempty"`
}

// ComponentCondition contains condition information for a cohort's individual component.
type ComponentCondition struct {
	// Type of component condition.
	Type string `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status corev1.ConditionStatus `json:"status"`
	// Last time we got an update on a given condition.
	LastHeartbeatTime metav1.Time `json:"lastHeartbeatTime,omitempty"`
	// Last time the condition transit from one status to another.
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// (brief) reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`
	// Human readable message indicating details about last transition.
	Message string `json:"message,omitempty"`
}

type CohortConditionType string

// These are valid conditions of cohort.
// In the future, we can add more. The current set of conditions are:
// CohortConditionTypeReady, CohortConditionTypeSchedulable.
const (
	// CohortConditionTypeReady means cohort is healthy and ready to perform its task.
	CohortConditionTypeReady CohortConditionType = "Ready"
	// CohortConditionTypeSchedulable means the cohort is healthy and schedulable.
	CohortConditionTypeSchedulable CohortConditionType = "Schedulable"
	// TODO add more types if necessary
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=storagecohorts
type StorageCohortList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []StorageVolume `json:"items"`
}
