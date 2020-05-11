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

from ..fragment.location import LocationFragment, QUERY as LocationFragmentQuery
from ..fragment.page_info import PageInfoFragment, QUERY as PageInfoFragmentQuery

QUERY: List[str] = LocationFragmentQuery + PageInfoFragmentQuery + ["""
query GetLocationsQuery($after: Cursor, $first: Int) {
  locations(after: $after, first: $first) {
    edges {
      node {
        ...LocationFragment
      }
    }
    pageInfo {
      ...PageInfoFragment
    }
  }
}

"""]

@dataclass
class GetLocationsQuery(DataClassJsonMixin):
    @dataclass
    class GetLocationsQueryData(DataClassJsonMixin):
        @dataclass
        class LocationConnection(DataClassJsonMixin):
            @dataclass
            class LocationEdge(DataClassJsonMixin):
                @dataclass
                class Location(LocationFragment):
                    pass

                node: Optional[Location]

            @dataclass
            class PageInfo(PageInfoFragment):
                pass

            edges: List[LocationEdge]
            pageInfo: PageInfo

        locations: Optional[LocationConnection]

    data: GetLocationsQueryData

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient, after: Optional[str] = None, first: Optional[int] = None) -> Optional[GetLocationsQueryData.LocationConnection]:
        # fmt: off
        variables = {"after": after, "first": first}
        try:
            network_start = perf_counter()
            response_text = client.call(''.join(set(QUERY)), variables=variables)
            decode_start = perf_counter()
            res = cls.from_json(response_text).data
            decode_time = perf_counter() - decode_start
            network_time = decode_start - network_start
            client.reporter.log_successful_operation("GetLocationsQuery", variables, network_time, decode_time)
            return res.locations
        except OperationException as e:
            raise FailedOperationException(
                client.reporter,
                e.err_msg,
                e.err_id,
                "GetLocationsQuery",
                variables,
            )
