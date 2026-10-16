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
	"testing"
)

func TestRedrivePolicyEqual(t *testing.T) {
	const dlq = "arn:aws:sqs:us-west-2:111122223333:dlq"
	tests := []struct {
		name string
		a    string
		b    string
		want bool
	}{
		{
			name: "string vs number maxReceiveCount",
			a:    `{"deadLetterTargetArn":"` + dlq + `","maxReceiveCount":"5"}`,
			b:    `{"deadLetterTargetArn":"` + dlq + `","maxReceiveCount":5}`,
			want: true,
		},
		{
			name: "whitespace and key order",
			a:    `{ "maxReceiveCount": 5, "deadLetterTargetArn": "` + dlq + `" }`,
			b:    `{"deadLetterTargetArn":"` + dlq + `","maxReceiveCount":5}`,
			want: true,
		},
		{
			name: "different maxReceiveCount",
			a:    `{"deadLetterTargetArn":"` + dlq + `","maxReceiveCount":"5"}`,
			b:    `{"deadLetterTargetArn":"` + dlq + `","maxReceiveCount":10}`,
			want: false,
		},
		{
			name: "different deadLetterTargetArn",
			a:    `{"deadLetterTargetArn":"` + dlq + `","maxReceiveCount":5}`,
			b:    `{"deadLetterTargetArn":"` + dlq + `2","maxReceiveCount":5}`,
			want: false,
		},
		{
			name: "invalid JSON",
			a:    `{"deadLetterTargetArn":`,
			b:    `{"deadLetterTargetArn":"` + dlq + `","maxReceiveCount":5}`,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := redrivePolicyEqual(tt.a, tt.b); got != tt.want {
				t.Errorf("redrivePolicyEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
