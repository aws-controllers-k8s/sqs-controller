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

package queue

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	svcsdk "github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go/aws/arn"
	policy "github.com/micahhausler/aws-iam-policy/policy"
)

// syncTags examines the Tags in the supplied Queue and calls the
// ListQueueTags, TagQueue and UntagQueue APIs to ensure that the set of
// associated Tags  stays in sync with the Queue.Spec.Tags
func (rm *resourceManager) syncTags(
	ctx context.Context,
	desired *resource,
	latest *resource,
) (err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.syncTags")
	defer func() { exit(err) }()
	toAdd := map[string]string{}
	toDelete := []string{}

	existingTags := latest.ko.Spec.Tags

	for k, v := range desired.ko.Spec.Tags {
		if ev, found := existingTags[k]; !found || *ev != *v {
			toAdd[k] = *v
		}
	}

	for k, _ := range existingTags {
		if _, found := desired.ko.Spec.Tags[k]; !found {
			deleteKey := k
			toDelete = append(toDelete, deleteKey)
		}
	}

	if len(toAdd) > 0 {
		for k, v := range toAdd {
			rlog.Debug("adding tag to queue", "key", k, "value", v)
		}
		if err = rm.addTags(ctx, desired, toAdd); err != nil {
			return err
		}
	}
	if len(toDelete) > 0 {
		for _, k := range toDelete {
			rlog.Debug("removing tag from queue", "key", k)
		}
		if err = rm.removeTags(ctx, desired, toDelete); err != nil {
			return err
		}
	}

	return nil
}

// getTags returns the list of tags to the Queue
func (rm *resourceManager) getTags(
	ctx context.Context,
	r *resource,
) (map[string]string, error) {
	var err error
	var resp *svcsdk.ListQueueTagsOutput
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.getTags")
	defer func() { exit(err) }()

	input := &svcsdk.ListQueueTagsInput{}
	input.QueueUrl = r.ko.Status.QueueURL

	// NOTE(jaypipes): Unlike many other ListTags APIs, SQS's is not
	// paginated...
	resp, err = rm.sdkapi.ListQueueTags(ctx, input)
	rm.metrics.RecordAPICall("READ_MANY", "ListQueueTags", err)
	if err != nil || resp == nil {
		return nil, err
	}
	// and the output's Tags field is actually a map[string]*string... go
	// figure :)
	return resp.Tags, err
}

// addTags adds the supplied Tags to the supplied Queue resource
func (rm *resourceManager) addTags(
	ctx context.Context,
	r *resource,
	tags map[string]string,
) (err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.addTag")
	defer func() { exit(err) }()

	input := &svcsdk.TagQueueInput{}
	input.QueueUrl = r.ko.Status.QueueURL
	input.Tags = tags

	_, err = rm.sdkapi.TagQueue(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "TagQueue", err)
	return err
}

// removeTags removes the supplied Tags from the supplied Queue resource
func (rm *resourceManager) removeTags(
	ctx context.Context,
	r *resource,
	tagKeys []string, // the set of tag keys to delete
) (err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.removeTag")
	defer func() { exit(err) }()

	input := &svcsdk.UntagQueueInput{}
	input.QueueUrl = r.ko.Status.QueueURL
	input.TagKeys = tagKeys

	_, err = rm.sdkapi.UntagQueue(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "UntagQueue", err)
	return err
}

func (rm *resourceManager) getQueueNameFromARN(tmpARN ackv1alpha1.AWSResourceName) (string, error) {
	queueARN, err := arn.Parse(string(tmpARN))
	if err != nil {
		return "", fmt.Errorf("error parsing queue ARN: %s, error: %w", tmpARN, err)
	}
	return queueARN.Resource, nil
}

// customPreCompare is the entry point for custom comparison logic
func customPreCompare(
	delta *ackcompare.Delta,
	a *resource,
	b *resource,
) {
	comparePolicy(delta, a, b)
	compareRedrivePolicy(delta, a, b)
}

// comparePolicy compares the Policy fields of two resources by unmarshalling
// them into policy.Policy structs and using reflect.DeepEqual.
func comparePolicy(
	delta *ackcompare.Delta,
	a *resource,
	b *resource,
) {
	if a.ko.Spec.Policy == b.ko.Spec.Policy {
		// both are nil or equal
		return
	}
	if a.ko.Spec.Policy == nil || b.ko.Spec.Policy == nil {
		// one is nil and the other is not
		delta.Add("Spec.Policy", a.ko.Spec.Policy, b.ko.Spec.Policy)
		return
	}
	var policyA policy.Policy
	decoderA := json.NewDecoder(bytes.NewBufferString(*a.ko.Spec.Policy))
	decoderA.DisallowUnknownFields()
	errA := decoderA.Decode(&policyA)

	var policyB policy.Policy
	decoderB := json.NewDecoder(bytes.NewBufferString(*b.ko.Spec.Policy))
	decoderB.DisallowUnknownFields()
	errB := decoderB.Decode(&policyB)

	if errA != nil || errB != nil || !reflect.DeepEqual(policyA, policyB) {
		delta.Add("Spec.Policy", a.ko.Spec.Policy, b.ko.Spec.Policy)
	}
}

// compareRedrivePolicy compares the RedrivePolicy fields of two resources by
// unmarshalling them into interface{} and using reflect.DeepEqual.
// since RedrivePolicy is a JSON string, we need to unmarshal it
// into an interface{} and then compare the two interface{}s.
func compareRedrivePolicy(
	delta *ackcompare.Delta,
	a *resource,
	b *resource,
) {
	if a.ko.Spec.RedrivePolicy == b.ko.Spec.RedrivePolicy {
		// both are nil or equal
		return
	}
	if a.ko.Spec.RedrivePolicy == nil || b.ko.Spec.RedrivePolicy == nil {
		// one is nil and the other is not
		delta.Add("Spec.RedrivePolicy", a.ko.Spec.RedrivePolicy, b.ko.Spec.RedrivePolicy)
		return
	}
	var redriveA interface{}
	errA := json.Unmarshal([]byte(*a.ko.Spec.RedrivePolicy), &redriveA)

	var redriveB interface{}
	errB := json.Unmarshal([]byte(*b.ko.Spec.RedrivePolicy), &redriveB)

	if errA != nil || errB != nil || !reflect.DeepEqual(redriveA, redriveB) {
		delta.Add("Spec.RedrivePolicy", a.ko.Spec.RedrivePolicy, b.ko.Spec.RedrivePolicy)
	}
}
