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

from .property_kind_enum import PropertyKind


DATETIME_FIELD = field(
    metadata={
        "dataclasses_json": {
            "encoder": datetime.isoformat,
            "decoder": fromisoformat,
            "mm_field": marshmallow_fields.DateTime(format="iso"),
        }
    }
)


def enum_field(enum_type):
    def encode_enum(value):
        return value.value

    def decode_enum(t, value):
        return t(value)

    return field(
        metadata={
            "dataclasses_json": {
                "encoder": encode_enum,
                "decoder": partial(decode_enum, enum_type),
            }
        }
    )



@dataclass_json
@dataclass
class EquipmentPortTypesQuery:
    __QUERY__ = """
    query EquipmentPortTypesQuery {
  equipmentPortTypes {
    edges {
      node {
        id
        name
        propertyTypes {
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
        linkPropertyTypes {
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
      }
    }
  }
}

    """

    @dataclass_json
    @dataclass
    class EquipmentPortTypesQueryData:
        @dataclass_json
        @dataclass
        class EquipmentPortTypeConnection:
            @dataclass_json
            @dataclass
            class EquipmentPortTypeEdge:
                @dataclass_json
                @dataclass
                class EquipmentPortType:
                    @dataclass_json
                    @dataclass
                    class PropertyType:
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
                    name: str
                    propertyTypes: List[PropertyType]
                    linkPropertyTypes: List[PropertyType]

                node: Optional[EquipmentPortType] = None

            edges: List[EquipmentPortTypeEdge]

        equipmentPortTypes: EquipmentPortTypeConnection

    data: Optional[EquipmentPortTypesQueryData] = None
    errors: Optional[Any] = None

    @classmethod
    # fmt: off
    def execute(cls, client):
        # fmt: off
        variables = None
        response_text = client.call(cls.__QUERY__, variables=variables)
        return cls.from_json(response_text).data
