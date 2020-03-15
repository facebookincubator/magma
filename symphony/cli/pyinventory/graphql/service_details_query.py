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

from gql.gql.enum_utils import enum_field
from .property_kind_enum import PropertyKind
from .service_endpoint_role_enum import ServiceEndpointRole


@dataclass
class ServiceDetailsQuery(DataClassJsonMixin):
    @dataclass
    class ServiceDetailsQueryData(DataClassJsonMixin):
        @dataclass
        class Node(DataClassJsonMixin):
            @dataclass
            class Customer(DataClassJsonMixin):
                id: str
                name: str
                externalId: Optional[str] = None

            @dataclass
            class ServiceEndpoint(DataClassJsonMixin):
                @dataclass
                class EquipmentPort(DataClassJsonMixin):
                    @dataclass
                    class Property(DataClassJsonMixin):
                        @dataclass
                        class PropertyType(DataClassJsonMixin):
                            id: str
                            name: str
                            type: PropertyKind = enum_field(PropertyKind)
                            index: Optional[int] = None
                            stringValue: Optional[str] = None
                            intValue: Optional[int] = None
                            booleanValue: Optional[bool] = None
                            floatValue: Optional[Number] = None
                            latitudeValue: Optional[Number] = None
                            longitudeValue: Optional[Number] = None
                            isEditable: Optional[bool] = None
                            isInstanceProperty: Optional[bool] = None

                        id: str
                        propertyType: PropertyType
                        stringValue: Optional[str] = None
                        intValue: Optional[int] = None
                        floatValue: Optional[Number] = None
                        booleanValue: Optional[bool] = None
                        latitudeValue: Optional[Number] = None
                        longitudeValue: Optional[Number] = None
                        rangeFromValue: Optional[Number] = None
                        rangeToValue: Optional[Number] = None

                    @dataclass
                    class EquipmentPortDefinition(DataClassJsonMixin):
                        id: str
                        name: str

                    @dataclass
                    class Link(DataClassJsonMixin):
                        @dataclass
                        class Service(DataClassJsonMixin):
                            id: str

                        id: str
                        services: List[Service]

                    id: str
                    properties: List[Property]
                    definition: EquipmentPortDefinition
                    link: Optional[Link] = None

                id: str
                port: EquipmentPort
                role: ServiceEndpointRole = enum_field(ServiceEndpointRole)

            @dataclass
            class Link(DataClassJsonMixin):
                @dataclass
                class Service(DataClassJsonMixin):
                    id: str

                id: str
                services: List[Service]

            id: str
            name: str
            endpoints: List[ServiceEndpoint]
            links: List[Link]
            externalId: Optional[str] = None
            customer: Optional[Customer] = None

        service: Optional[Node] = None

    data: ServiceDetailsQueryData

    __QUERY__: str = """
    query ServiceDetailsQuery($id: ID!) {
  service: node(id: $id) {
    ... on Service {
      id
      name
      externalId
      customer {
        id
        name
        externalId
      }
      endpoints {
        id
        port {
          id
          properties {
            id
            propertyType {
              id
              name
              type
              index
              stringValue
              intValue
              booleanValue
              floatValue
              latitudeValue
              longitudeValue
              isEditable
              isInstanceProperty
            }
            stringValue
            intValue
            floatValue
            booleanValue
            latitudeValue
            longitudeValue
            rangeFromValue
            rangeToValue
          }
          definition {
            id
            name
          }
          link {
            id
            services {
              id
            }
          }
        }
        role
      }
      links {
        id
        services {
          id
        }
      }
    }
  }
}

    """

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient, id: str) -> ServiceDetailsQueryData:
        # fmt: off
        variables = {"id": id}
        response_text = client.call(cls.__QUERY__, variables=variables)
        return cls.from_json(response_text).data
