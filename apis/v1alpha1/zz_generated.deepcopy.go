//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchResultErrorEntry) DeepCopyInto(out *BatchResultErrorEntry) {
	*out = *in
	if in.Code != nil {
		in, out := &in.Code, &out.Code
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchResultErrorEntry.
func (in *BatchResultErrorEntry) DeepCopy() *BatchResultErrorEntry {
	if in == nil {
		return nil
	}
	out := new(BatchResultErrorEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChangeMessageVisibilityBatchRequestEntry) DeepCopyInto(out *ChangeMessageVisibilityBatchRequestEntry) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ReceiptHandle != nil {
		in, out := &in.ReceiptHandle, &out.ReceiptHandle
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChangeMessageVisibilityBatchRequestEntry.
func (in *ChangeMessageVisibilityBatchRequestEntry) DeepCopy() *ChangeMessageVisibilityBatchRequestEntry {
	if in == nil {
		return nil
	}
	out := new(ChangeMessageVisibilityBatchRequestEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChangeMessageVisibilityBatchResultEntry) DeepCopyInto(out *ChangeMessageVisibilityBatchResultEntry) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChangeMessageVisibilityBatchResultEntry.
func (in *ChangeMessageVisibilityBatchResultEntry) DeepCopy() *ChangeMessageVisibilityBatchResultEntry {
	if in == nil {
		return nil
	}
	out := new(ChangeMessageVisibilityBatchResultEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeleteMessageBatchRequestEntry) DeepCopyInto(out *DeleteMessageBatchRequestEntry) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ReceiptHandle != nil {
		in, out := &in.ReceiptHandle, &out.ReceiptHandle
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeleteMessageBatchRequestEntry.
func (in *DeleteMessageBatchRequestEntry) DeepCopy() *DeleteMessageBatchRequestEntry {
	if in == nil {
		return nil
	}
	out := new(DeleteMessageBatchRequestEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeleteMessageBatchResultEntry) DeepCopyInto(out *DeleteMessageBatchResultEntry) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeleteMessageBatchResultEntry.
func (in *DeleteMessageBatchResultEntry) DeepCopy() *DeleteMessageBatchResultEntry {
	if in == nil {
		return nil
	}
	out := new(DeleteMessageBatchResultEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Message) DeepCopyInto(out *Message) {
	*out = *in
	if in.Body != nil {
		in, out := &in.Body, &out.Body
		*out = new(string)
		**out = **in
	}
	if in.MD5OfBody != nil {
		in, out := &in.MD5OfBody, &out.MD5OfBody
		*out = new(string)
		**out = **in
	}
	if in.MD5OfMessageAttributes != nil {
		in, out := &in.MD5OfMessageAttributes, &out.MD5OfMessageAttributes
		*out = new(string)
		**out = **in
	}
	if in.MessageID != nil {
		in, out := &in.MessageID, &out.MessageID
		*out = new(string)
		**out = **in
	}
	if in.ReceiptHandle != nil {
		in, out := &in.ReceiptHandle, &out.ReceiptHandle
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Message.
func (in *Message) DeepCopy() *Message {
	if in == nil {
		return nil
	}
	out := new(Message)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MessageAttributeValue) DeepCopyInto(out *MessageAttributeValue) {
	*out = *in
	if in.DataType != nil {
		in, out := &in.DataType, &out.DataType
		*out = new(string)
		**out = **in
	}
	if in.StringValue != nil {
		in, out := &in.StringValue, &out.StringValue
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MessageAttributeValue.
func (in *MessageAttributeValue) DeepCopy() *MessageAttributeValue {
	if in == nil {
		return nil
	}
	out := new(MessageAttributeValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MessageSystemAttributeValue) DeepCopyInto(out *MessageSystemAttributeValue) {
	*out = *in
	if in.DataType != nil {
		in, out := &in.DataType, &out.DataType
		*out = new(string)
		**out = **in
	}
	if in.StringValue != nil {
		in, out := &in.StringValue, &out.StringValue
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MessageSystemAttributeValue.
func (in *MessageSystemAttributeValue) DeepCopy() *MessageSystemAttributeValue {
	if in == nil {
		return nil
	}
	out := new(MessageSystemAttributeValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Queue) DeepCopyInto(out *Queue) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Queue.
func (in *Queue) DeepCopy() *Queue {
	if in == nil {
		return nil
	}
	out := new(Queue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Queue) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueList) DeepCopyInto(out *QueueList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Queue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueList.
func (in *QueueList) DeepCopy() *QueueList {
	if in == nil {
		return nil
	}
	out := new(QueueList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QueueList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueSpec) DeepCopyInto(out *QueueSpec) {
	*out = *in
	if in.ContentBasedDeduplication != nil {
		in, out := &in.ContentBasedDeduplication, &out.ContentBasedDeduplication
		*out = new(string)
		**out = **in
	}
	if in.DelaySeconds != nil {
		in, out := &in.DelaySeconds, &out.DelaySeconds
		*out = new(string)
		**out = **in
	}
	if in.FifoQueue != nil {
		in, out := &in.FifoQueue, &out.FifoQueue
		*out = new(string)
		**out = **in
	}
	if in.KMSDataKeyReusePeriodSeconds != nil {
		in, out := &in.KMSDataKeyReusePeriodSeconds, &out.KMSDataKeyReusePeriodSeconds
		*out = new(string)
		**out = **in
	}
	if in.KMSMasterKeyID != nil {
		in, out := &in.KMSMasterKeyID, &out.KMSMasterKeyID
		*out = new(string)
		**out = **in
	}
	if in.MaximumMessageSize != nil {
		in, out := &in.MaximumMessageSize, &out.MaximumMessageSize
		*out = new(string)
		**out = **in
	}
	if in.MessageRetentionPeriod != nil {
		in, out := &in.MessageRetentionPeriod, &out.MessageRetentionPeriod
		*out = new(string)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
	if in.QueueARN != nil {
		in, out := &in.QueueARN, &out.QueueARN
		*out = new(string)
		**out = **in
	}
	if in.QueueName != nil {
		in, out := &in.QueueName, &out.QueueName
		*out = new(string)
		**out = **in
	}
	if in.ReceiveMessageWaitTimeSeconds != nil {
		in, out := &in.ReceiveMessageWaitTimeSeconds, &out.ReceiveMessageWaitTimeSeconds
		*out = new(string)
		**out = **in
	}
	if in.RedrivePolicy != nil {
		in, out := &in.RedrivePolicy, &out.RedrivePolicy
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.VisibilityTimeout != nil {
		in, out := &in.VisibilityTimeout, &out.VisibilityTimeout
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueSpec.
func (in *QueueSpec) DeepCopy() *QueueSpec {
	if in == nil {
		return nil
	}
	out := new(QueueSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueStatus) DeepCopyInto(out *QueueStatus) {
	*out = *in
	if in.ACKResourceMetadata != nil {
		in, out := &in.ACKResourceMetadata, &out.ACKResourceMetadata
		*out = new(corev1alpha1.ResourceMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*corev1alpha1.Condition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1alpha1.Condition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.QueueURL != nil {
		in, out := &in.QueueURL, &out.QueueURL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueStatus.
func (in *QueueStatus) DeepCopy() *QueueStatus {
	if in == nil {
		return nil
	}
	out := new(QueueStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SendMessageBatchRequestEntry) DeepCopyInto(out *SendMessageBatchRequestEntry) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MessageBody != nil {
		in, out := &in.MessageBody, &out.MessageBody
		*out = new(string)
		**out = **in
	}
	if in.MessageDeduplicationID != nil {
		in, out := &in.MessageDeduplicationID, &out.MessageDeduplicationID
		*out = new(string)
		**out = **in
	}
	if in.MessageGroupID != nil {
		in, out := &in.MessageGroupID, &out.MessageGroupID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SendMessageBatchRequestEntry.
func (in *SendMessageBatchRequestEntry) DeepCopy() *SendMessageBatchRequestEntry {
	if in == nil {
		return nil
	}
	out := new(SendMessageBatchRequestEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SendMessageBatchResultEntry) DeepCopyInto(out *SendMessageBatchResultEntry) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MD5OfMessageAttributes != nil {
		in, out := &in.MD5OfMessageAttributes, &out.MD5OfMessageAttributes
		*out = new(string)
		**out = **in
	}
	if in.MD5OfMessageBody != nil {
		in, out := &in.MD5OfMessageBody, &out.MD5OfMessageBody
		*out = new(string)
		**out = **in
	}
	if in.MD5OfMessageSystemAttributes != nil {
		in, out := &in.MD5OfMessageSystemAttributes, &out.MD5OfMessageSystemAttributes
		*out = new(string)
		**out = **in
	}
	if in.MessageID != nil {
		in, out := &in.MessageID, &out.MessageID
		*out = new(string)
		**out = **in
	}
	if in.SequenceNumber != nil {
		in, out := &in.SequenceNumber, &out.SequenceNumber
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SendMessageBatchResultEntry.
func (in *SendMessageBatchResultEntry) DeepCopy() *SendMessageBatchResultEntry {
	if in == nil {
		return nil
	}
	out := new(SendMessageBatchResultEntry)
	in.DeepCopyInto(out)
	return out
}
