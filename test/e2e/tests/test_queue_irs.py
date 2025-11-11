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

"""Integration tests for ECR Cross Account Resource Management.
Ideally we want these tests to be in the ACK runtime, but we don't have a way
to run them there yet. So we'll run them here for now.
"""

import pytest
import time
import logging
import boto3
import os

from acktest.resources import random_suffix_name
from acktest.k8s import resource as k8s
from e2e import service_marker, CRD_GROUP, CRD_VERSION, load_sqs_resource
from e2e.replacement_values import REPLACEMENT_VALUES
from e2e.fixtures import create_iam_role_selector

CREATE_WAIT_AFTER_SECONDS = 5

RESOURCE_PLURAL = "queues"

CREATE_WAIT_AFTER_SECONDS = 10
UPDATE_WAIT_AFTER_SECONDS = 10
DELETE_WAIT_AFTER_SECONDS = 10

TESTING_NAMESPACE = "carm-testing"
TESTING_ACCOUNT = "637423602339"
TESTTING_ASSUME_ROLE = "arn:aws:iam::637423602339:role/ack-carm-role-DO-NOT-DELETE"

@service_marker
@pytest.mark.canary
class TestCARM:
    def get_queue_url(self, queue_name: str) -> dict:
        sqs_client = boto3.client(
            "sqs",
            aws_access_key_id=os.environ["CARM_AWS_ACCESS_KEY_ID"],
            aws_secret_access_key=os.environ["CARM_AWS_SECRET_ACCESS_KEY"],
            aws_session_token=os.environ["CARM_AWS_SESSION_TOKEN"],
        )
        try:
            resp = sqs_client.get_queue_url(QueueName=queue_name)
        except Exception as e:
            logging.debug(e)
            return None

        return resp

    def queue_exists(self, queue_name: str) -> bool:
        return self.get_queue_url(queue_name) is not None

    def test_basic_queue(self):
        k8s.create_k8s_namespace(
            TESTING_NAMESPACE,
            annotations={}
        )
        time.sleep(CREATE_WAIT_AFTER_SECONDS)
        create_iam_role_selector(
            TESTING_NAMESPACE,
            "ack-role-selector",
            TESTTING_ASSUME_ROLE
        )

        time.sleep(CREATE_WAIT_AFTER_SECONDS)

        resource_name = random_suffix_name("sqs-queue", 24)

        replacements = REPLACEMENT_VALUES.copy()
        replacements["QUEUE_NAME"] = resource_name
        replacements["NAMESPACE"] = TESTING_NAMESPACE
        # Load ECR CR
        resource_data = load_sqs_resource(
            "queue_carm",
            additional_replacements=replacements,
        )
        logging.debug(resource_data)

        # Create k8s resource
        ref = k8s.CustomResourceReference(
            CRD_GROUP, CRD_VERSION, RESOURCE_PLURAL,
            resource_name, namespace=TESTING_NAMESPACE,
        )

        k8s.create_custom_resource(ref, resource_data)
        cr = k8s.wait_resource_consumed_by_controller(ref)
        assert cr is not None
        assert k8s.get_resource_exists(ref)

        time.sleep(CREATE_WAIT_AFTER_SECONDS)

        assert k8s.wait_on_condition(ref, "ACK.ResourceSynced", "True", wait_periods=5)

        # Check SQS queue exists
        resp = self.get_queue_url(resource_name)
        print(os.environ["CARM_AWS_ACCESS_KEY_ID"])
        print(os.environ["CARM_AWS_SECRET_ACCESS_KEY"])
        print(os.environ["CARM_AWS_SESSION_TOKEN"])
        print(resp)
        assert resp != None

        # Delete k8s resource
        _, deleted = k8s.delete_custom_resource(ref)
        assert deleted is True

        time.sleep(DELETE_WAIT_AFTER_SECONDS)

        # Check SQS queue doesn't exists
        exists = self.queue_exists(resource_name)
        assert not exists