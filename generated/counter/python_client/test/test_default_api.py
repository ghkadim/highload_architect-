"""
    OTUS Highload Architect

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 1.2.0
    Generated by: https://openapi-generator.tech
"""


import unittest

import openapi_client_counter
from openapi_client_counter.api.default_api import DefaultApi  # noqa: E501


class TestDefaultApi(unittest.TestCase):
    """DefaultApi unit test stubs"""

    def setUp(self):
        self.api = DefaultApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_counter_counter_id_get(self):
        """Test case for counter_counter_id_get

        """
        pass


if __name__ == '__main__':
    unittest.main()
