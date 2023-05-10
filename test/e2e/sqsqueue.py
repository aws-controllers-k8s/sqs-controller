# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
#	 http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

"""Utilities for working with Queue resources"""

import datetime
import json
import time

import boto3
import pytest

DEFAULT_WAIT_UNTIL_EXISTS_TIMEOUT_SECONDS = 60 * 10
DEFAULT_WAIT_UNTIL_EXISTS_INTERVAL_SECONDS = 15
DEFAULT_WAIT_UNTIL_DELETED_TIMEOUT_SECONDS = 60 * 10
DEFAULT_WAIT_UNTIL_DELETED_INTERVAL_SECONDS = 15


def wait_until_exists(
        queue_name: str,
        timeout_seconds: int = DEFAULT_WAIT_UNTIL_EXISTS_TIMEOUT_SECONDS,
        interval_seconds: int = DEFAULT_WAIT_UNTIL_EXISTS_INTERVAL_SECONDS,
) -> None:
    """Waits until a Queue with a supplied name is returned from SQS GetQueue
    API.

    Usage:
        from e2e.queue import wait_until_exists

        wait_until_exists(queue_name)

    Raises:
        pytest.fail upon timeout
    """
    now = datetime.datetime.now()
    timeout = now + datetime.timedelta(seconds=timeout_seconds)

    while True:
        if datetime.datetime.now() >= timeout:
            pytest.fail(
                "Timed out waiting for Queue to exist "
                "in SQS API"
            )
        time.sleep(interval_seconds)

        latest = get_queue_url(queue_name)
        if latest is not None:
            break


def wait_until_deleted(
        queue_name: str,
        timeout_seconds: int = DEFAULT_WAIT_UNTIL_DELETED_TIMEOUT_SECONDS,
        interval_seconds: int = DEFAULT_WAIT_UNTIL_DELETED_INTERVAL_SECONDS,
) -> None:
    """Waits until a Queue with a supplied ID is no longer returned from
    the SQS API.

    Usage:
        from e2e.queue import wait_until_deleted

        wait_until_deleted(queue_name)

    Raises:
        pytest.fail upon timeout
    """
    now = datetime.datetime.now()
    timeout = now + datetime.timedelta(seconds=timeout_seconds)

    while True:
        if datetime.datetime.now() >= timeout:
            pytest.fail(
                "Timed out waiting for Queue to be "
                "deleted in SQS API"
            )
        time.sleep(interval_seconds)

        latest = get_queue_url(queue_name)
        if latest is None:
            break


def get_queue_url(queue_name):
    """Returns the URL for a supplied Queue name.

    If no such Queue exists, returns None.
    """
    c = boto3.client('sqs')
    try:
        resp = c.get_queue_url(QueueName=queue_name)
        return resp['QueueUrl']
    except c.exceptions.QueueDoesNotExist:
        return None


def get_attributes(queue_url):
    """Returns a dict containing the Queue attributes from the SQS API.

    If no such Queue exists, returns None.
    """
    c = boto3.client('sqs')
    try:
        resp = c.get_queue_attributes(
            QueueUrl=queue_url, AttributeNames=['All'],
        )
        return resp['Attributes']
    except c.exceptions.QueueDoesNotExist:
        return None


def get_tags(queue_url):
    """Returns a dict containing the tags that have been associated to the
    supplied Queue.

    If no such Queue exists, returns None.
    """
    c = boto3.client('sqs')
    try:
        resp = c.list_queue_tags(QueueUrl=queue_url)
        return resp['Tags']
    except c.exceptions.QueueDoesNotExist:
        return None


def create_queue(queue_name) -> str:
    """ create queue with queue_name """
    c = boto3.client('sqs')
    resp = c.create_queue(QueueName=queue_name)
    return resp['QueueUrl']
