#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass
from datetime import datetime
from gql.gql.datetime_utils import DATETIME_FIELD
from gql.gql.graphql_client import GraphqlClient
from functools import partial
from numbers import Number
from typing import Any, Callable, List, Mapping, Optional

from dataclasses_json import DataClassJsonMixin

from .add_customer_input import AddCustomerInput


@dataclass
class AddCustomerMutation(DataClassJsonMixin):
    __QUERY__: str = """
    mutation AddCustomerMutation($input: AddCustomerInput!) {
  addCustomer(input: $input) {
    id
    name
    externalId
  }
}

    """

    @dataclass
    class AddCustomerMutationData(DataClassJsonMixin):
        @dataclass
        class Customer(DataClassJsonMixin):
            id: str
            name: str
            externalId: Optional[str] = None

        addCustomer: Optional[Customer] = None

    data: Optional[AddCustomerMutationData] = None

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient, input: AddCustomerInput):
        # fmt: off
        variables = {"input": input}
        response_text = client.call(cls.__QUERY__, variables=variables)
        return cls.from_json(response_text).data
