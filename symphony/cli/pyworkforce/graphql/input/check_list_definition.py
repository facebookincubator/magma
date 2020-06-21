#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass
from datetime import datetime
from functools import partial
from gql.gql.datetime_utils import DATETIME_FIELD
from numbers import Number
from typing import Any, Callable, List, Mapping, Optional

from dataclasses_json import DataClassJsonMixin

from gql.gql.enum_utils import enum_field
from ..enum.check_list_item_enum_selection_mode import CheckListItemEnumSelectionMode
from ..enum.check_list_item_type import CheckListItemType

@dataclass
class CheckListDefinitionInput(DataClassJsonMixin):
    title: str
    type: CheckListItemType = enum_field(CheckListItemType)
    id: Optional[str] = None
    index: Optional[int] = None
    enumValues: Optional[str] = None
    enumSelectionMode: Optional[CheckListItemEnumSelectionMode] = None
    helpText: Optional[str] = None

