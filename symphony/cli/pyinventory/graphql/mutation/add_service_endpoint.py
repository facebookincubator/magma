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
from typing import Any, Callable, List, Mapping, Optional
from time import perf_counter
from dataclasses_json import DataClassJsonMixin

from ..fragment.customer import CustomerFragment, QUERY as CustomerFragmentQuery
from ..fragment.link import LinkFragment, QUERY as LinkFragmentQuery
from ..input.add_service_endpoint import AddServiceEndpointInput


QUERY: List[str] = CustomerFragmentQuery + LinkFragmentQuery + ["""
mutation AddServiceEndpointMutation($input: AddServiceEndpointInput!) {
  addServiceEndpoint(input: $input) {
    id
    name
    externalId
    customer {
      ...CustomerFragment
    }
    endpoints {
      id
      port {
        id
      }
      definition {
        role
      }
    }
    links {
      ...LinkFragment
    }
  }
}

"""]

@dataclass
class AddServiceEndpointMutation(DataClassJsonMixin):
    @dataclass
    class AddServiceEndpointMutationData(DataClassJsonMixin):
        @dataclass
        class Service(DataClassJsonMixin):
            @dataclass
            class Customer(CustomerFragment):
                pass

            @dataclass
            class ServiceEndpoint(DataClassJsonMixin):
                @dataclass
                class EquipmentPort(DataClassJsonMixin):
                    id: str

                @dataclass
                class ServiceEndpointDefinition(DataClassJsonMixin):
                    role: Optional[str]

                id: str
                definition: ServiceEndpointDefinition
                port: Optional[EquipmentPort]

            @dataclass
            class Link(LinkFragment):
                pass

            id: str
            name: str
            endpoints: List[ServiceEndpoint]
            links: List[Link]
            externalId: Optional[str]
            customer: Optional[Customer]

        addServiceEndpoint: Service

    data: AddServiceEndpointMutationData

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient, input: AddServiceEndpointInput) -> AddServiceEndpointMutationData.Service:
        # fmt: off
        variables = {"input": input}
        try:
            start_time = perf_counter()
            response_text = client.call(''.join(set(QUERY)), variables=variables)
            res = cls.from_json(response_text).data
            elapsed_time = perf_counter() - start_time
            client.reporter.log_successful_operation("AddServiceEndpointMutation", variables, elapsed_time)
            return res.addServiceEndpoint
        except OperationException as e:
            raise FailedOperationException(
                client.reporter,
                e.err_msg,
                e.err_id,
                "AddServiceEndpointMutation",
                variables,
            )
