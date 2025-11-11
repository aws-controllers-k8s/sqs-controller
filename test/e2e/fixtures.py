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

"""Fixtures common to all SQS controller tests"""

from acktest.k8s import resource as k8s

CRD_GROUP = "services.k8s.aws"
CRD_VERSION = "v1alpha1"
RESOURCE_PLURAL = "iamroleselectors"

def create_iam_role_selector(namespace: str, name: str, role: str):
    iam_role_selector = {
        "apiVersion": "services.k8s.aws/v1alpha1",
        "kind": "IAMRoleSelector",
        "metadata": {
            "name": name,
        },
        "spec": {
            "namespaceSelector": {
                "names": [namespace]
            },
            "arn": role
        }
    }

    ref = k8s.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, RESOURCE_PLURAL,
        name, namespace=None,
    )
    k8s.create_custom_resource(ref, iam_role_selector)
