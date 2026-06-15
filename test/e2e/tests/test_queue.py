# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
# 	 http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

"""Integration tests for the Queue API.
"""

import json
import pytest
import time
import logging
import boto3

from acktest.resources import random_suffix_name
from acktest.k8s import resource as k8s, condition
from acktest import tags
from acktest import adoption as adoption
from e2e import service_marker, CRD_GROUP, CRD_VERSION, load_sqs_resource
from e2e.replacement_values import REPLACEMENT_VALUES
from e2e.bootstrap_resources import get_bootstrap_resources
from e2e import sqsqueue

RESOURCE_KIND = "Queue"
RESOURCE_PLURAL = "queues"

CREATE_WAIT_AFTER_SECONDS = 5
MODIFY_WAIT_AFTER_SECONDS = 10
DELETE_WAIT_AFTER_SECONDS = 20

@pytest.fixture(scope="module")
def fifo_queue():
    resource_name = random_suffix_name("sqs-queue", 24) + ".fifo"

    replacements = REPLACEMENT_VALUES.copy()
    replacements["QUEUE_NAME"] = resource_name

    # Load Queue CR
    resource_data = load_sqs_resource(
        "queue_fifo",
        additional_replacements=replacements,
    )
    logging.debug(resource_data)

    # Create k8s resource
    ref = k8s.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, RESOURCE_PLURAL,
        resource_name, namespace="default",
    )
    k8s.create_custom_resource(ref, resource_data)
    cr = k8s.wait_resource_consumed_by_controller(ref)

    sqsqueue.wait_until_exists(resource_name)

    yield cr, ref

    # Delete k8s resource
    _, deleted = k8s.delete_custom_resource(
        ref,
        period_length=DELETE_WAIT_AFTER_SECONDS,
    )
    assert deleted

    sqsqueue.wait_until_deleted(resource_name)


@pytest.fixture(scope="module")
def simple_queue():
    resource_name = random_suffix_name("sqs-queue", 24)

    resources = get_bootstrap_resources()
    logging.debug(resources)

    replacements = REPLACEMENT_VALUES.copy()
    replacements["QUEUE_NAME"] = resource_name

    # Load Queue CR
    resource_data = load_sqs_resource(
        "queue",
        additional_replacements=replacements,
    )
    logging.debug(resource_data)

    # Create k8s resource
    ref = k8s.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, RESOURCE_PLURAL,
        resource_name, namespace="default",
    )
    k8s.create_custom_resource(ref, resource_data)
    cr = k8s.wait_resource_consumed_by_controller(ref)

    sqsqueue.wait_until_exists(resource_name)

    yield cr, ref

    # Delete k8s resource
    _, deleted = k8s.delete_custom_resource(
        ref,
        period_length=DELETE_WAIT_AFTER_SECONDS,
    )
    assert deleted

    sqsqueue.wait_until_deleted(resource_name)


@service_marker
@pytest.mark.canary
class TestQueue:
    def test_crud(self, simple_queue):
        res, ref = simple_queue

        time.sleep(CREATE_WAIT_AFTER_SECONDS)

        # Update the tags associated with the Queue and verify the update is
        # reflected in the SQS API
        cr = k8s.get_resource(ref)
        assert cr is not None
        assert 'spec' in cr
        assert 'delaySeconds' in cr['spec']
        assert cr['spec']['delaySeconds'] == "0"
        assert 'status' in cr
        assert 'queueURL' in cr['status']
        queue_url = cr['status']['queueURL']

        latest_attrs = sqsqueue.get_attributes(queue_url)
        assert 'DelaySeconds' in latest_attrs
        assert latest_attrs['DelaySeconds'] == "0"

        # Test updating one of the attributes...
        new_delay = "10"
        redrive_policy = '{"redrivePermission":"denyAll"}'

        updates = {
            "spec": {
                "delaySeconds": new_delay,
                "redriveAllowPolicy": redrive_policy,
            },
        }
        k8s.patch_custom_resource(ref, updates)
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)

        latest_attrs = sqsqueue.get_attributes(queue_url)
        assert 'DelaySeconds' in latest_attrs
        assert latest_attrs['DelaySeconds'] == new_delay
        assert latest_attrs['RedriveAllowPolicy'] == redrive_policy

        # Test updating tags...
        assert 'tags' in cr['spec']
        assert len(cr['spec']['tags']) == 1
        assert 'key1' in cr['spec']['tags']
        assert cr['spec']['tags']['key1'] == 'val1'

        expect_before_update_tags = {
            "key1": "val1",
        }
        latest_tags = sqsqueue.get_tags(queue_url)
        tags.assert_equal_without_ack_tags(
            expect_before_update_tags, latest_tags,
        )

        new_tags = {
            "key1": None,  # Unfortunately, this is the only way to remove a key
            "key2": "val2",
        }
        updates = {
            "spec": {"tags": new_tags},
        }
        k8s.patch_custom_resource(ref, updates)
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)

        expect_after_update_tags = {
            "key2": "val2",
        }
        latest_tags = sqsqueue.get_tags(queue_url)
        tags.assert_equal_without_ack_tags(
            expect_after_update_tags, latest_tags,
        )

    def test_fifo_crud(self, fifo_queue):
        res, ref = fifo_queue

        time.sleep(CREATE_WAIT_AFTER_SECONDS)
        assert k8s.wait_on_condition(ref, condition.CONDITION_TYPE_RESOURCE_SYNCED, "True", wait_periods=5)
        
        cr = k8s.get_resource(ref)
        assert cr is not None
        assert 'status' in cr
        assert 'queueURL' in cr['status']
        queue_url = cr['status']['queueURL']

        latest_attrs = sqsqueue.get_attributes(queue_url)
        assert 'FifoQueue' in latest_attrs
        assert latest_attrs['FifoQueue'] == "true"
        assert 'DeduplicationScope' in latest_attrs
        assert latest_attrs['DeduplicationScope'] == "queue"
        assert 'FifoThroughputLimit' in latest_attrs
        assert latest_attrs['FifoThroughputLimit'] == "perQueue"

        updates = {
            "spec": {
                "deduplicationScope": "messageGroup",
                "fifoThroughputLimit": "perMessageGroupId",
            },
        }

        k8s.patch_custom_resource(ref, updates)
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)
        assert k8s.wait_on_condition(ref, condition.CONDITION_TYPE_RESOURCE_SYNCED, "True", wait_periods=5)

        latest_attrs = sqsqueue.get_attributes(queue_url)
        assert 'DeduplicationScope' in latest_attrs
        assert latest_attrs['DeduplicationScope'] == "messageGroup"
        assert 'FifoThroughputLimit' in latest_attrs
        assert latest_attrs['FifoThroughputLimit'] == "perMessageGroupId"


@pytest.fixture(scope="module")
def queue_with_policy():
    resource_name = random_suffix_name("sqs-policy", 24)

    # Placeholder ARNs for the policy fixture
    queue_arn = f"arn:aws:sqs:us-west-2:123456789012:{resource_name}"
    topic_arn = "arn:aws:sns:us-west-2:123456789012:test-topic"

    replacements = REPLACEMENT_VALUES.copy()
    replacements["QUEUE_NAME"] = resource_name
    replacements["QUEUE_ARN"] = queue_arn
    replacements["TOPIC_ARN"] = topic_arn

    resource_data = load_sqs_resource(
        "queue_policy",
        additional_replacements=replacements,
    )
    logging.debug(resource_data)

    ref = k8s.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, RESOURCE_PLURAL,
        resource_name, namespace="default",
    )
    k8s.create_custom_resource(ref, resource_data)
    cr = k8s.wait_resource_consumed_by_controller(ref)

    sqsqueue.wait_until_exists(resource_name)

    yield cr, ref

    _, deleted = k8s.delete_custom_resource(
        ref,
        period_length=DELETE_WAIT_AFTER_SECONDS,
    )
    assert deleted

    sqsqueue.wait_until_deleted(resource_name)


@service_marker
class TestQueuePolicy:
    """Tests for Queue Policy comparison (community#2597).

    Validates that the controller does not trigger unnecessary updates when the
    Policy field is semantically equivalent but textually different (e.g. Action
    as a string vs single-element array, or different JSON key ordering).
    """

    def test_policy_no_drift_on_action_format(self, queue_with_policy):
        """Verify no infinite reconcile when AWS returns Action as string
        but spec has Action as array (community#2597).
        """
        res, ref = queue_with_policy

        time.sleep(CREATE_WAIT_AFTER_SECONDS)

        cr = k8s.get_resource(ref)
        assert cr is not None
        assert 'status' in cr
        assert 'queueURL' in cr['status']

        # Record the initial generation
        initial_generation = cr['metadata']['generation']

        # Wait and check that generation doesn't keep incrementing
        # (which would indicate constant reconcile loops)
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)

        cr = k8s.get_resource(ref)
        assert cr is not None
        current_generation = cr['metadata']['generation']

        # The generation should not have changed, meaning no spurious updates
        assert current_generation == initial_generation, (
            f"Queue generation changed from {initial_generation} to "
            f"{current_generation}, indicating the controller is performing "
            f"unnecessary updates (likely due to Policy comparison drift)"
        )

        # Verify the resource stays synced
        condition.assert_synced(ref)

    def test_policy_update_detected(self, queue_with_policy):
        """Verify that actual policy changes ARE detected and applied."""
        res, ref = queue_with_policy

        cr = k8s.get_resource(ref)
        assert cr is not None
        assert 'status' in cr
        assert 'queueURL' in cr['status']
        queue_url = cr['status']['queueURL']

        # Update the policy with a genuinely different policy
        new_policy = json.dumps({
            "Version": "2012-10-17",
            "Statement": [{
                "Effect": "Allow",
                "Principal": {"AWS": "*"},
                "Action": "sqs:SendMessage",
                "Resource": "*",
            }],
        })

        updates = {
            "spec": {"policy": new_policy},
        }
        k8s.patch_custom_resource(ref, updates)
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)

        # Verify the policy was actually updated in AWS
        latest_attrs = sqsqueue.get_attributes(queue_url)
        assert 'Policy' in latest_attrs
        latest_policy = json.loads(latest_attrs['Policy'])
        assert latest_policy['Statement'][0]['Principal'] == {"AWS": "*"}
        assert latest_policy['Statement'][0]['Action'] == "sqs:SendMessage"
