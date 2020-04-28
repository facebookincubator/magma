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


QUERY: List[str] = ["""
mutation RemoveLocationTypeMutation($id: ID!) {
  removeLocationType(id: $id)
}

"""]

@dataclass
class RemoveLocationTypeMutation(DataClassJsonMixin):
    @dataclass
    class RemoveLocationTypeMutationData(DataClassJsonMixin):
        removeLocationType: str

    data: RemoveLocationTypeMutationData

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient, id: str) -> str:
        # fmt: off
        variables = {"id": id}
        try:
            start_time = perf_counter()
            response_text = client.call(''.join(set(QUERY)), variables=variables)
            res = cls.from_json(response_text).data
            elapsed_time = perf_counter() - start_time
            client.reporter.log_successful_operation("RemoveLocationTypeMutation", variables, elapsed_time)
            return res.removeLocationType
        except OperationException as e:
            raise FailedOperationException(
                client.reporter,
                e.err_msg,
                e.err_id,
                "RemoveLocationTypeMutation",
                variables,
            )
