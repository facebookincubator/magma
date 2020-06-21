#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass
from datetime import datetime
from gql.gql.datetime_utils import DATETIME_FIELD
from gql.gql.graphql_client import GraphqlClient
from gql.gql.client import OperationException
from gql.gql.reporter import FailedOperationException
from functools import partial
from numbers import Number
from typing import Any, Callable, List, Mapping, Optional, Dict
from time import perf_counter
from dataclasses_json import DataClassJsonMixin

from ..fragment.customer import CustomerFragment, QUERY as CustomerFragmentQuery
from ..fragment.property import PropertyFragment, QUERY as PropertyFragmentQuery
QUERY: List[str] = CustomerFragmentQuery + PropertyFragmentQuery + ["""
fragment ServiceFragment on Service {
  id
  name
  externalId
  serviceType {
    id
    name
  }
  customer {
    ...CustomerFragment
  }
  properties {
    ...PropertyFragment
  }
}

"""]

@dataclass
class ServiceFragment(DataClassJsonMixin):
    @dataclass
    class ServiceType(DataClassJsonMixin):
        id: str
        name: str

    @dataclass
    class Customer(CustomerFragment):
        pass

    @dataclass
    class Property(PropertyFragment):
        pass

    id: str
    name: str
    externalId: Optional[str]
    serviceType: ServiceType
    customer: Optional[Customer]
    properties: List[Property]
