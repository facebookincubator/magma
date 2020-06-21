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

from ..fragment.location import LocationFragment, QUERY as LocationFragmentQuery
from ..input.add_location import AddLocationInput


QUERY: List[str] = LocationFragmentQuery + ["""
mutation AddLocationMutation($input: AddLocationInput!) {
  addLocation(input: $input) {
    ...LocationFragment
  }
}

"""]

@dataclass
class AddLocationMutation(DataClassJsonMixin):
    @dataclass
    class AddLocationMutationData(DataClassJsonMixin):
        @dataclass
        class Location(LocationFragment):
            pass

        addLocation: Location

    data: AddLocationMutationData

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient, input: AddLocationInput) -> AddLocationMutationData.Location:
        # fmt: off
        variables: Dict[str, Any] = {"input": input}
        try:
            network_start = perf_counter()
            response_text = client.call(''.join(set(QUERY)), variables=variables)
            decode_start = perf_counter()
            res = cls.from_json(response_text).data
            decode_time = perf_counter() - decode_start
            network_time = decode_start - network_start
            client.reporter.log_successful_operation("AddLocationMutation", variables, network_time, decode_time)
            return res.addLocation
        except OperationException as e:
            raise FailedOperationException(
                client.reporter,
                e.err_msg,
                e.err_id,
                "AddLocationMutation",
                variables,
            )
