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


QUERY: List[str] = ["""
query SearchQuery(
  $name: String!
  $after: Cursor
  $first: Int = 10
  $before: Cursor
  $last: Int
) {
  searchForNode(
    name: $name
    after: $after
    first: $first
    before: $before
    last: $last
  ) {
    edges {
      node {
        typename: __typename
        ... on Location {
          id
          externalId
          name
          locationType {
            name
          }
        }
      }
    }
  }
}

"""]

@dataclass
class SearchQuery(DataClassJsonMixin):
    @dataclass
    class SearchQueryData(DataClassJsonMixin):
        @dataclass
        class SearchNodesConnection(DataClassJsonMixin):
            @dataclass
            class SearchNodeEdge(DataClassJsonMixin):
                @dataclass
                class Node(DataClassJsonMixin):
                    @dataclass
                    class LocationType(DataClassJsonMixin):
                        name: str

                    typename: str
                    id: str
                    externalId: Optional[str]
                    name: str
                    locationType: LocationType

                node: Optional[Node]

            edges: List[SearchNodeEdge]

        searchForNode: SearchNodesConnection

    data: SearchQueryData

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient, name: str, after: Optional[str] = None, first: Optional[int] = 10, before: Optional[str] = None, last: Optional[int] = None) -> SearchQueryData.SearchNodesConnection:
        # fmt: off
        variables: Dict[str, Any] = {"name": name, "after": after, "first": first, "before": before, "last": last}
        try:
            network_start = perf_counter()
            response_text = client.call(''.join(set(QUERY)), variables=variables)
            decode_start = perf_counter()
            res = cls.from_json(response_text).data
            decode_time = perf_counter() - decode_start
            network_time = decode_start - network_start
            client.reporter.log_successful_operation("SearchQuery", variables, network_time, decode_time)
            return res.searchForNode
        except OperationException as e:
            raise FailedOperationException(
                client.reporter,
                e.err_msg,
                e.err_id,
                "SearchQuery",
                variables,
            )
