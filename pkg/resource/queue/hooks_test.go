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
	"reflect"
	"testing"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	svcapitypes "github.com/aws-controllers-k8s/sqs-controller/apis/v1alpha1"
)

func strPtr(s string) *string {
	return &s
}

func TestComparePolicy(t *testing.T) {
	basePolicyJSON := `{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"AWS":"*"},"Action":"sqs:SendMessage","Resource":"arn:aws:sqs:us-east-1:123456789012:MyQueue"}]}`
	equivalentPolicyJSON := `{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Effect":    "Allow",
				"Principal": { "AWS": "*" },
				"Action":    "sqs:SendMessage",
				"Resource":  "arn:aws:sqs:us-east-1:123456789012:MyQueue"
			}
		]
	}`
	differentPolicyJSON := `{"Version":"2012-10-17","Statement":[{"Effect":"Deny","Principal":{"AWS":"*"},"Action":"sqs:SendMessage","Resource":"arn:aws:sqs:us-east-1:123456789012:MyQueue"}]}`
	invalidJSON := `{"Version": "2012-10-17", "Statement": [`

	tests := []struct {
		name       string
		policyA    *string
		policyB    *string
		expectDiff bool
	}{
		{
			name:       "both nil",
			policyA:    nil,
			policyB:    nil,
			expectDiff: false,
		},
		{
			name:       "a nil, b not nil",
			policyA:    nil,
			policyB:    strPtr(basePolicyJSON),
			expectDiff: true,
		},
		{
			name:       "both equal pointers",
			policyA:    strPtr(basePolicyJSON),
			policyB:    strPtr(basePolicyJSON),
			expectDiff: false,
		},
		{
			name:       "semantically equivalent",
			policyA:    strPtr(basePolicyJSON),
			policyB:    strPtr(equivalentPolicyJSON),
			expectDiff: false,
		},
		{
			name:       "different policies",
			policyA:    strPtr(basePolicyJSON),
			policyB:    strPtr(differentPolicyJSON),
			expectDiff: true,
		},
		{
			name:       "a invalid, b valid",
			policyA:    strPtr(invalidJSON),
			policyB:    strPtr(basePolicyJSON),
			expectDiff: true,
		},
		{
			name:       "both invalid",
			policyA:    strPtr(invalidJSON),
			policyB:    strPtr(`{"another": "invalid`),
			expectDiff: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			delta := ackcompare.NewDelta()
			resA := &resource{ko: &svcapitypes.Queue{Spec: svcapitypes.QueueSpec{Policy: tt.policyA}}}
			resB := &resource{ko: &svcapitypes.Queue{Spec: svcapitypes.QueueSpec{Policy: tt.policyB}}}

			comparePolicy(delta, resA, resB)

			diffCount := len(delta.Differences)

			if tt.expectDiff {
				if diffCount != 1 {
					t.Fatalf("Expected exactly one difference, got %d", diffCount)
				}
				diff := delta.Differences[0]
				if diff == nil {
					t.Fatalf("Difference object should not be nil when diff expected")
				}
				expectedPath := ackcompare.NewPath("Spec.Policy")
				if !reflect.DeepEqual(expectedPath, diff.Path) {
					t.Errorf("Expected path %v, got %v", expectedPath, diff.Path)
				}
				if !reflect.DeepEqual(tt.policyA, diff.A) {
					t.Errorf("Expected diff.A %v, got %v", tt.policyA, diff.A)
				}
				if !reflect.DeepEqual(tt.policyB, diff.B) {
					t.Errorf("Expected diff.B %v, got %v", tt.policyB, diff.B)
				}
			} else {
				if diffCount != 0 {
					t.Errorf("Expected no differences, got %d", diffCount)
				}
			}
		})
	}
}

func TestCompareRedrivePolicy(t *testing.T) {
	basePolicyJSON := `{"deadLetterTargetArn":"arn:aws:sqs:us-east-1:123456789012:dlq","maxReceiveCount":10}`
	equivalentPolicyJSON := `{
		"deadLetterTargetArn": "arn:aws:sqs:us-east-1:123456789012:dlq",
		"maxReceiveCount": 10
	}`
	differentPolicyJSON := `{"deadLetterTargetArn":"arn:aws:sqs:us-east-1:123456789012:other-dlq","maxReceiveCount":5}`
	invalidJSON := `{"deadLetterTargetArn":`

	tests := []struct {
		name       string
		policyA    *string
		policyB    *string
		expectDiff bool
	}{
		{
			name:       "both nil",
			policyA:    nil,
			policyB:    nil,
			expectDiff: false,
		},
		{
			name:       "a nil, b not nil",
			policyA:    nil,
			policyB:    strPtr(basePolicyJSON),
			expectDiff: true,
		},
		{
			name:       "both equal pointers",
			policyA:    strPtr(basePolicyJSON),
			policyB:    strPtr(basePolicyJSON),
			expectDiff: false,
		},
		{
			name:       "semantically equivalent",
			policyA:    strPtr(basePolicyJSON),
			policyB:    strPtr(equivalentPolicyJSON),
			expectDiff: false,
		},
		{
			name:       "different policies",
			policyA:    strPtr(basePolicyJSON),
			policyB:    strPtr(differentPolicyJSON),
			expectDiff: true,
		},
		{
			name:       "a invalid, b valid",
			policyA:    strPtr(invalidJSON),
			policyB:    strPtr(basePolicyJSON),
			expectDiff: true,
		},
		{
			name:       "both invalid",
			policyA:    strPtr(invalidJSON),
			policyB:    strPtr(`{"another": "invalid`),
			expectDiff: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			delta := ackcompare.NewDelta()
			resA := &resource{ko: &svcapitypes.Queue{Spec: svcapitypes.QueueSpec{RedrivePolicy: tt.policyA}}}
			resB := &resource{ko: &svcapitypes.Queue{Spec: svcapitypes.QueueSpec{RedrivePolicy: tt.policyB}}}

			compareRedrivePolicy(delta, resA, resB)

			diffCount := len(delta.Differences)

			if tt.expectDiff {
				if diffCount != 1 {
					t.Fatalf("Expected exactly one difference, got %d", diffCount)
				}
				diff := delta.Differences[0]
				if diff == nil {
					t.Fatalf("Difference object should not be nil when diff expected")
				}
				expectedPath := ackcompare.NewPath("Spec.RedrivePolicy")
				if !reflect.DeepEqual(expectedPath, diff.Path) {
					t.Errorf("Expected path %v, got %v", expectedPath, diff.Path)
				}
				if !reflect.DeepEqual(tt.policyA, diff.A) {
					t.Errorf("Expected diff.A %v, got %v", tt.policyA, diff.A)
				}
				if !reflect.DeepEqual(tt.policyB, diff.B) {
					t.Errorf("Expected diff.B %v, got %v", tt.policyB, diff.B)
				}
			} else {
				if diffCount != 0 {
					t.Errorf("Expected no differences, got %d", diffCount)
				}
			}
		})
	}
}
