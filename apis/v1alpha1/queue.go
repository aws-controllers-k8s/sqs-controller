// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// QueueSpec defines the desired state of Queue.
type QueueSpec struct {
	ContentBasedDeduplication    *string                                  `json:"contentBasedDeduplication,omitempty"`
	DelaySeconds                 *string                                  `json:"delaySeconds,omitempty"`
	FIFOQueue                    *string                                  `json:"fifoQueue,omitempty"`
	KMSDataKeyReusePeriodSeconds *string                                  `json:"kmsDataKeyReusePeriodSeconds,omitempty"`
	KMSMasterKeyID               *string                                  `json:"kmsMasterKeyID,omitempty"`
	KMSMasterKeyRef              *ackv1alpha1.AWSResourceReferenceWrapper `json:"kmsMasterKeyRef,omitempty"`
	MaximumMessageSize           *string                                  `json:"maximumMessageSize,omitempty"`
	MessageRetentionPeriod       *string                                  `json:"messageRetentionPeriod,omitempty"`
	Policy                       *string                                  `json:"policy,omitempty"`
	PolicyRef                    *ackv1alpha1.AWSResourceReferenceWrapper `json:"policyRef,omitempty"`
	// +kubebuilder:validation:Required
	QueueName                     *string `json:"queueName"`
	ReceiveMessageWaitTimeSeconds *string `json:"receiveMessageWaitTimeSeconds,omitempty"`
	RedriveAllowPolicy            *string `json:"redriveAllowPolicy,omitempty"`
	RedrivePolicy                 *string `json:"redrivePolicy,omitempty"`
	SQSManagedSSEEnabled          *string `json:"sqsManagedSSEEnabled,omitempty"`
	// Add cost allocation tags to the specified Amazon SQS queue. For an overview,
	// see Tagging Your Amazon SQS Queues (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-queue-tags.html)
	// in the Amazon SQS Developer Guide.
	//
	// When you use queue tags, keep the following guidelines in mind:
	//
	//   - Adding more than 50 tags to a queue isn't recommended.
	//
	//   - Tags don't have any semantic meaning. Amazon SQS interprets tags as
	//     character strings.
	//
	//   - Tags are case-sensitive.
	//
	//   - A new tag with a key identical to that of an existing tag overwrites
	//     the existing tag.
	//
	// For a full list of tag restrictions, see Quotas related to queues (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-limits.html#limits-queues)
	// in the Amazon SQS Developer Guide.
	//
	// To be able to tag a queue on creation, you must have the sqs:CreateQueue
	// and sqs:TagQueue permissions.
	//
	// Cross-account permissions don't apply to this action. For more information,
	// see Grant cross-account permissions to a role and a username (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name)
	// in the Amazon SQS Developer Guide.
	Tags              map[string]*string `json:"tags,omitempty"`
	VisibilityTimeout *string            `json:"visibilityTimeout,omitempty"`
}

// QueueStatus defines the observed state of Queue
type QueueStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// +kubebuilder:validation:Optional
	QueueARN *string `json:"queueARN,omitempty"`
	// The URL of the created Amazon SQS queue.
	// +kubebuilder:validation:Optional
	QueueURL *string `json:"queueURL,omitempty"`
}

// Queue is the Schema for the Queues API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="DelaySeconds",type=string,priority=0,JSONPath=`.spec.delaySeconds`
// +kubebuilder:printcolumn:name="maximumMessageSize",type=string,priority=1,JSONPath=`.spec.maximumMessageSize`
// +kubebuilder:printcolumn:name="messageRetentionPeriod",type=string,priority=1,JSONPath=`.spec.messageRetentionPeriod`
// +kubebuilder:printcolumn:name="receiveMessageWaitTimeSeconds",type=string,priority=1,JSONPath=`.spec.receiveMessageWaitTimeSeconds`
// +kubebuilder:printcolumn:name="visibilityTimeout",type=string,priority=0,JSONPath=`.spec.visibilityTimeout`
// +kubebuilder:printcolumn:name="Synced",type="string",priority=0,JSONPath=".status.conditions[?(@.type==\"ACK.ResourceSynced\")].status"
// +kubebuilder:printcolumn:name="Age",type="date",priority=0,JSONPath=".metadata.creationTimestamp"
type Queue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              QueueSpec   `json:"spec,omitempty"`
	Status            QueueStatus `json:"status,omitempty"`
}

// QueueList contains a list of Queue
// +kubebuilder:object:root=true
type QueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Queue `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Queue{}, &QueueList{})
}
