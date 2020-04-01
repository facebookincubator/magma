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

from .customer_fragment import CustomerFragment, QUERY as CustomerFragmentQuery
from .property_fragment import PropertyFragment, QUERY as PropertyFragmentQuery
from gql.gql.enum_utils import enum_field
from .service_endpoint_role_enum import ServiceEndpointRole


@dataclass
class AddServiceLinkMutation(DataClassJsonMixin):
    @dataclass
    class AddServiceLinkMutationData(DataClassJsonMixin):
        @dataclass
        class Service(DataClassJsonMixin):
            @dataclass
            class Customer(CustomerFragment):
                pass

            @dataclass
            class ServiceEndpoint(DataClassJsonMixin):
                @dataclass
                class EquipmentPort(DataClassJsonMixin):
                    @dataclass
                    class Property(PropertyFragment):
                        pass

                    @dataclass
                    class EquipmentPortDefinition(DataClassJsonMixin):
                        @dataclass
                        class EquipmentPortType(DataClassJsonMixin):
                            id: str
                            name: str

                        id: str
                        name: str
                        portType: Optional[EquipmentPortType]

                    @dataclass
                    class Link(DataClassJsonMixin):
                        @dataclass
                        class Property(PropertyFragment):
                            pass

                        @dataclass
                        class Service(DataClassJsonMixin):
                            id: str

                        id: str
                        properties: List[Property]
                        services: List[Service]

                    id: str
                    properties: List[Property]
                    definition: EquipmentPortDefinition
                    link: Optional[Link]

                id: str
                port: EquipmentPort
                role: ServiceEndpointRole = enum_field(ServiceEndpointRole)

            @dataclass
            class Link(DataClassJsonMixin):
                @dataclass
                class Property(PropertyFragment):
                    pass

                @dataclass
                class Service(DataClassJsonMixin):
                    id: str

                id: str
                properties: List[Property]
                services: List[Service]

            id: str
            name: str
            endpoints: List[ServiceEndpoint]
            links: List[Link]
            externalId: Optional[str]
            customer: Optional[Customer]

        addServiceLink: Service

    data: AddServiceLinkMutationData

    __QUERY__: str = CustomerFragmentQuery + PropertyFragmentQuery + """
    mutation AddServiceLinkMutation($id: ID!, $linkId: ID!) {
  addServiceLink(id: $id, linkId: $linkId) {
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
        properties {
          ...PropertyFragment
        }
        definition {
          id
          name
          portType {
            id
            name
          }
        }
        link {
          id
          properties {
            ...PropertyFragment
          }
          services {
            id
          }
        }
      }
      role
    }
    links {
      id
      properties {
        ...PropertyFragment
      }
      services {
        id
      }
    }
  }
}

    """

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient, id: str, linkId: str) -> AddServiceLinkMutationData:
        # fmt: off
        variables = {"id": id, "linkId": linkId}
        response_text = client.call(cls.__QUERY__, variables=variables)
        return cls.from_json(response_text).data
