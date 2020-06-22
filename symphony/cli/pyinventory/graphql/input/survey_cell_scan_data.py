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
from ..enum.cellular_network_type import CellularNetworkType

@dataclass
class SurveyCellScanData(DataClassJsonMixin):
    signalStrength: int
    networkType: CellularNetworkType = enum_field(CellularNetworkType)
    timestamp: Optional[int] = None
    baseStationID: Optional[str] = None
    networkID: Optional[str] = None
    systemID: Optional[str] = None
    cellID: Optional[str] = None
    locationAreaCode: Optional[str] = None
    mobileCountryCode: Optional[str] = None
    mobileNetworkCode: Optional[str] = None
    primaryScramblingCode: Optional[str] = None
    operator: Optional[str] = None
    arfcn: Optional[int] = None
    physicalCellID: Optional[str] = None
    trackingAreaCode: Optional[str] = None
    timingAdvance: Optional[int] = None
    earfcn: Optional[int] = None
    uarfcn: Optional[int] = None
    latitude: Optional[Number] = None
    longitude: Optional[Number] = None

