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
class RemoveEquipmentTypeMutation:
    __QUERY__ = """
    mutation RemoveEquipmentTypeMutation($id: ID!) {
  removeEquipmentType(id: $id)
}

    """

    @dataclass_json
    @dataclass
    class RemoveEquipmentTypeMutationData:
        removeEquipmentType: str

    data: Optional[RemoveEquipmentTypeMutationData] = None
    errors: Any = None

    @classmethod
    # fmt: off
    def execute(cls, client, id: str):
        # fmt: off
        variables = {"id": id}
        response_text = client.call(cls.__QUERY__, variables=variables)
        return cls.from_json(response_text).data
