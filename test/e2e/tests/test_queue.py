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

import pytest
import time
import logging

from acktest.resources import random_suffix_name
from acktest.k8s import resource as k8s
from acktest import tags
from e2e import service_marker, CRD_GROUP, CRD_VERSION, load_sqs_resource
from e2e.replacement_values import REPLACEMENT_VALUES
from e2e.bootstrap_resources import get_bootstrap_resources
from e2e import sqsqueue

RESOURCE_PLURAL = "queues"

CREATE_WAIT_AFTER_SECONDS = 5
MODIFY_WAIT_AFTER_SECONDS = 10
DELETE_WAIT_AFTER_SECONDS = 20

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
            "key1": None, # Unfortunately, this is the only way to remove a key
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
