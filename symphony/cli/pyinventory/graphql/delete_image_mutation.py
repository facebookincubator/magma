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
from .image_entity_enum import ImageEntity


@dataclass
class DeleteImageMutation(DataClassJsonMixin):
    __QUERY__: str = """
    mutation DeleteImageMutation(
  $entityType: ImageEntity!
  $entityId: ID!
  $id: ID!
) {
  deleteImage(entityType: $entityType, entityId: $entityId, id: $id) {
    id
    fileName
  }
}

    """

    @dataclass
    class DeleteImageMutationData(DataClassJsonMixin):
        @dataclass
        class File(DataClassJsonMixin):
            id: str
            fileName: str

        deleteImage: Optional[File] = None

    data: Optional[DeleteImageMutationData] = None

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient, entityType: ImageEntity, entityId: str, id: str):
        # fmt: off
        variables = {"entityType": entityType, "entityId": entityId, "id": id}
        response_text = client.call(cls.__QUERY__, variables=variables)
        return cls.from_json(response_text).data
