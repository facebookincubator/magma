#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass, field
from datetime import datetime
from functools import partial
from numbers import Number
from typing import Any, Callable, List, Mapping, Optional

from dataclasses_json import dataclass_json
from marshmallow import fields as marshmallow_fields

from .datetime_utils import fromisoformat


DATETIME_FIELD = field(
    metadata={
        "dataclasses_json": {
            "encoder": datetime.isoformat,
            "decoder": fromisoformat,
            "mm_field": marshmallow_fields.DateTime(format="iso"),
        }
    }
)


@dataclass_json
@dataclass
class AddCustomerInput:
    name: str
    externalId: Optional[str] = None


@dataclass_json
@dataclass
class AddCustomerMutation:
    __QUERY__ = """
    mutation AddCustomerMutation($input: AddCustomerInput!) {
  addCustomer(input: $input) {
    id
    name
    externalId
  }
}

    """

    @dataclass_json
    @dataclass
    class AddCustomerMutationData:
        @dataclass_json
        @dataclass
        class Customer:
            id: str
            name: str
            externalId: Optional[str] = None

        addCustomer: Optional[Customer] = None

    data: Optional[AddCustomerMutationData] = None
    errors: Optional[Any] = None

    @classmethod
    # fmt: off
    def execute(cls, client, input: AddCustomerInput):
        # fmt: off
        variables = {"input": input}
        response_text = client.call(cls.__QUERY__, variables=variables)
        return cls.from_json(response_text).data
