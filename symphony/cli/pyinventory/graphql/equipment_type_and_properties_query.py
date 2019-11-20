#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass, field
from datetime import datetime
from enum import Enum
from functools import partial
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
class EquipmentTypeAndPropertiesQuery:
    __QUERY__ = """
    query EquipmentTypeAndPropertiesQuery($id: ID!) {
  equipment(id: $id) {
    equipmentType {
      name
    }
    properties {
      propertyType {
        id
      }
      stringValue
      intValue
      booleanValue
      floatValue
      latitudeValue
      longitudeValue
    }
  }
}

    """

    @dataclass_json
    @dataclass
    class EquipmentTypeAndPropertiesQueryData:
        @dataclass_json
        @dataclass
        class Equipment:
            @dataclass_json
            @dataclass
            class EquipmentType:
                name: str

            @dataclass_json
            @dataclass
            class Property:
                @dataclass_json
                @dataclass
                class PropertyType:
                    id: str

                propertyType: PropertyType
                stringValue: Optional[str] = None
                intValue: Optional[int] = None
                booleanValue: Optional[bool] = None
                floatValue: Optional[float] = None
                latitudeValue: Optional[float] = None
                longitudeValue: Optional[float] = None

            equipmentType: EquipmentType
            properties: List[Property]

        equipment: Optional[Equipment] = None

    data: Optional[EquipmentTypeAndPropertiesQueryData] = None
    errors: Any = None

    @classmethod
    # fmt: off
    def execute(cls, client, id: str):
        # fmt: off
        variables = {"id": id}
        response_text = client.call(cls.__QUERY__, variables=variables)
        return cls.from_json(response_text).data
