"""
Copyright (c) 2016-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""
from copy import deepcopy
import redis
import redis_collections
from typing import Iterator, MutableMapping, TypeVar

from magma.common.redis.serializers import RedisSerde
from orc8r.protos.redis_pb2 import RedisState

# NOTE: these containers replace the serialization methods exposed by
# the redis-collection objects. Although the methods are hinted to be
# privately scoped, the method replacement is encouraged in the library's
# docs: http://redis-collections.readthedocs.io/en/stable/usage-notes.html

T = TypeVar('T')

class RedisList(redis_collections.List):
    """
    List-like interface serializing elements to a Redis datastore.

    Notes:
        - Provides persistence across sessions
        - Mutable elements handled correctly
        - Not expected to be thread safe, but could be extended
    """

    def __init__(self, client, key, serialize, deserialize):
        """
        Initialize instance.

        Args:
            client (redis.Redis): Redis client object
            key (str): key where this container's elements are stored in Redis
            serialize (function (any) -> bytes):
                function called to serialize an element
            deserialize (function (bytes) -> any):
                function called to deserialize an element
        Returns:
            redis_list (redis_collections.List): persistent list-like interface
        """
        self._pickle = serialize
        self._unpickle = deserialize
        super().__init__(redis=client, key=key, writeback=True)

    def __copy__(self):
        return [elt for elt in self]

    def __deepcopy__(self, memo):
        return [deepcopy(elt, memo) for elt in self]


class RedisSet(redis_collections.Set):
    """
    Set-like interface serializing elements to a Redis datastore.

    Notes:
        - Provides persistence across sessions
        - Mutable elements _not_ handled correctly:
            - Get/set mutable elements supported
            - Don't update the contents of a mutable element and
              expect things to go well
        - Expected to be thread safe, but not tested
    """

    def __init__(self, client, key, serialize, deserialize):
        """
        Initialize instance.

        Args:
            client (redis.Redis): Redis client object
            key (str): key where this container's elements are stored in Redis
            serialize (function (any) -> bytes):
                function called to serialize an element
            deserialize (function (bytes) -> any):
                function called to deserialize an element
        Returns:
            redis_set (redis_collections.Set): persistent set-like interface
        """
        # NOTE: redis_collections.Set doesn't have a writeback option, causing
        # issue when mutable elements are updated in-place.
        self._pickle = serialize
        self._unpickle = deserialize
        super().__init__(redis=client, key=key)

    def __copy__(self):
        return {elt for elt in self}

    def __deepcopy__(self, memo):
        return {deepcopy(elt, memo) for elt in self}


class RedisHashDict(redis_collections.DefaultDict):
    """
    Dict-like interface serializing elements to a Redis datastore. This dict
    utilizes Redis's hashmap functionality

    Notes:
        - Keys must be string-like and are serialized to plaintext (UTF-8)
        - Provides persistence across sessions
        - Mutable elements handled correctly
        - Not expected to be thread safe, but could be extended
        - Keys are serialized in plaintext
    """

    @staticmethod
    def serialize_key(key):
        """ Serialize key to plaintext. """
        return key

    @staticmethod
    def deserialize_key(serialized):
        """ Deserialize key from plaintext encoded as UTF-8 bytes. """
        return serialized.decode('utf-8')  # Redis returns bytes

    def __init__(
        self, client, key, serialize, deserialize,
        default_factory=None,
        writeback=False,
    ):
        """
        Initialize instance.

        Args:
            client (redis.Redis): Redis client object
            key (str): key where this container's elements are stored in Redis
            serialize (function (any) -> bytes):
                function called to serialize a value
            deserialize (function (bytes) -> any):
                function called to deserialize a value
            writeback (bool): if writeback is set to true, dict maintains a
                local cache of values and the `sync` method can be called to
                store these values. NOTE: only use this option if syncing
                between services is not important.

        Returns:
            redis_dict (redis_collections.Dict): persistent dict-like interface
        """
        # Key serialization (to/from plaintext)
        self._pickle_key = RedisHashDict.serialize_key
        self._unpickle_key = RedisHashDict.deserialize_key
        # Value serialization
        self._pickle_value = serialize
        self._unpickle = deserialize
        super().__init__(
            default_factory, redis=client, key=key, writeback=writeback)

    def __setitem__(self, key, value):
        """Set ``d[key]`` to *value*.

        Override in order to increment version on each update
        """
        version = self.get_version(key)
        pickled_key = self._pickle_key(key)
        pickled_value = self._pickle_value(value, version + 1)
        self.redis.hset(self.key, pickled_key, pickled_value)

        if self.writeback:
            self.cache[key] = value

    def __copy__(self):
        return {key: self[key] for key in self}

    def __deepcopy__(self, memo):
        return {key: deepcopy(self[key], memo) for key in self}

    def get_version(self, key):
        """Return the version of the value for key *key*. Returns 0 if
        key is not in the map
        """
        try:
            value = self.cache[key]
        except KeyError:
            pickled_key = self._pickle_key(key)
            value = self.redis.hget(self.key, pickled_key)
            if value is None:
                return 0

        proto_wrapper = RedisState()
        proto_wrapper.ParseFromString(value)
        return proto_wrapper.version


class RedisFlatDict(MutableMapping[str, T]):
    """
    Dict-like interface serializing elements to a Redis datastore. This
    dict stores key directly (i.e. without a hashmap).
    """

    def __init__(self, client: redis.Redis, serde: RedisSerde[T]):
        """
        Args:
            client (redis.Redis): Redis client object
            serde (): RedisSerde for de/serializing the object stored
        """
        super().__init__()
        self.redis = client
        self.serde = serde
        self.redis_type = serde.redis_type

    def __len__(self) -> int:
        """Return the number of items in the dictionary."""
        type_pattern = "*:" + self.redis_type
        return len(self.redis.keys(pattern=type_pattern))

    def __iter__(self) -> Iterator[str]:
        """Return an iterator over the keys of the dictionary."""
        type_pattern = "*:" + self.redis_type
        for k in self.redis.keys(pattern=type_pattern):
            try:
                deserialized_key = k.decode('utf-8')
                split_key = deserialized_key.split(":", 1)
            except AttributeError:
                split_key = k.split(":", 1)
            yield split_key[0]

    def __contains__(self, key: str) -> bool:
        """Return ``True`` if *key* is present, else ``False``."""
        composite_key = self._make_composite_key(key)
        return bool(self.redis.exists(composite_key))

    def __getitem__(self, key: str) -> T:
        """Return the item of dictionary with key *key:type*. Raises a
        :exc:`KeyError` if *key:type* is not in the map.
        """
        if ':' in key:
            raise ValueError("Key %s cannot contain ':' char" % key)
        composite_key = self._make_composite_key(key)
        serialized_value = self.redis.get(composite_key)
        if serialized_value is None:
            raise KeyError(composite_key)

        value = self.serde.deserialize(serialized_value)
        return value

    def __setitem__(self, key: str, value: T) -> bool:
        """Set ``d[key:type]`` to *value*."""
        if ':' in key:
            raise ValueError("Key %s cannot contain ':' char" % key)
        version = self.get_version(key)
        serialized_value = self.serde.serialize(value, version + 1)
        composite_key = self._make_composite_key(key)
        return self.redis.set(composite_key, serialized_value)

    def __delitem__(self, key: str) -> int:
        """Remove ``d[key:type]`` from dictionary.
        Raises a :func:`KeyError` if *key:type* is not in the map.
        """
        if ':' in key:
            raise ValueError("Key %s cannot contain ':' char" % key)
        composite_key = self._make_composite_key(key)
        deleted_count = self.redis.delete(composite_key)
        if not deleted_count:
            raise KeyError(composite_key)
        return deleted_count

    def clear(self) -> None:
        """
        Clear all keys in the dictionary
        """
        for key in self.keys():
            composite_key = self._make_composite_key(key)
            self.redis.delete(composite_key)

    def get_version(self, key: str) -> int:
        """Return the version of the value for key *key:type*. Returns 0 if
        key is not in the map
        """
        composite_key = self._make_composite_key(key)
        value = self.redis.get(composite_key)
        if value is None:
            return 0

        proto_wrapper = RedisState()
        proto_wrapper.ParseFromString(value)
        return proto_wrapper.version

    def keys(self):
        """Return a copy of the dictionary's list of keys
        Note: for redis *key:type* key is returned
        """
        return list(self.__iter__())

    def _make_composite_key(self, key):
        return key + ":" + self.redis_type
