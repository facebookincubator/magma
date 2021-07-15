#!/usr/bin/env python3

"""
Copyright 2020 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
"""

import collections
import os

from pyparsing import (
    Group,
    Optional,
    Suppress,
    Word,
    ZeroOrMore,
    alphanums,
    alphas,
    cppStyleComment,
    empty,
    nums,
)

"""
Helper script to parse C header files and generate the corresponding python
types of the same name. Not intended to be general purpose C language parser.
All C defines are extracted as consts
All C enums are converted to classes derived from Enum
All C Structs are converted to c_types and are derived from the C compatible
python Structure type.

Headers need to be parsed in a specific order such that there are no unknown
types. Currently doesn't support forward declaration syntax (as it is unused
in the s1aptester headers).
"""

HEADER_FILES = [
    'fw_api_int.h',
    'fw_api_int.x',
    'trfgen.x',
]

OUT_FILE = 's1ap_types.py'

# syntax we don't want to see in the final parse tree
_lcurl = Suppress('{')
_rcurl = Suppress('}')
_lsq_bracket = Suppress('[')
_rsq_bracket = Suppress(']')
_equal = Suppress('=')
_comma = Suppress(',')
_semi_colon = Suppress(';')
_enum = Suppress('enum')
_struct = Suppress('struct')
_define = Suppress('#define')
_union = Suppress('union')

# Conversion from s1AP Types to c_types
py_type_map = {
    'Bool': 'c_ubyte',
    'S8': 'char',
    'U8': 'c_ubyte',
    'S16': 'c_short',
    'U16': 'c_ushort',
    'S32': 'c_int',
    'U32': 'c_uint',
    'int': 'c_int',
}

defines = {}
enums = {}
c_struct = {}
anon_types = {}


def _write_copyright(out_fname):
    """
    Write out copyright string
    """
    cc_str = '''"""
Copyright 2020 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
"""
'''

    import_str = '''
from enum import Enum
from ctypes import *

'''

    warn_str = '''
"""
Autogenerated code based on types parsed from s1aptester headers.
Hand edit with care
"""
'''
    with open(out_fname, 'w') as fd:
        fd.write(cc_str)
        fd.write(import_str)
        fd.write(warn_str)


def _write_constants(out_fname, const_map):
    """
    Write out a map of constant values to a python file
    Args:
        out_fname: The output filename to append constants to
        const_map: The map of constant name to value
    """
    if not const_map:
        return
    const_str = '''
""" Auto-generated list of constants """
'''
    with open(out_fname, 'a') as fd:
        fd.write(const_str)
        for name, val in const_map.items():
            fd.write("%s = %s\n" % (name, val))
        fd.write('\n\n')


def extract_defines(in_fname, out_fname):
    """
    Parse #defines and convert them to python constants.
    Doesn't support macros.
    Args:
        in_fname: file name to read
        out_fname: file name to append read defines to
    Returns
        map of constant name and their value
    """
    global defines
    identifier = Word(alphas, alphanums + '_')
    value = Word(alphanums + '_')
    with open(in_fname, 'r') as header:
        sample = header.read()
    define_def = _define + identifier("name") + empty + value("value")
    for item, _, _ in define_def.scanString(sample):  # item, start, stop
        defines[item.name] = item.value

    _write_constants(out_fname, defines)

    return defines


def _write_enums(out_fname, enum_map):
    """
    Write out enums in python style
    Args:
        enum_map: A nested map of enums where each enum is represented as a map
        of values
    """
    if not enum_map:
        return
    enum_str = '''
""" Auto-generated list of enums """
'''
    with open(out_fname, 'a') as fd:
        fd.write(enum_str)
        for enum_name, enum_struct in enum_map.items():
            class_str = "class %s(Enum):\n" % enum_name
            fd.write(class_str)
            for enum_prop, enum_val in enum_struct.items():
                prop_str = "    %s = %s\n" % (enum_prop, enum_val)
                fd.write(prop_str)
            fd.write('\n\n')


def extract_enums(in_fname, out_fname):
    """
    Extract enums from a given file
    Args:
        in_fname: file name to read
        out_fname: file name to append read enums to
    Returns
        nested map of enums and their properties
    """
    global enums
    with open(in_fname, 'r') as header:
        sample = header.read()
    identifier = Word(alphas + '_', alphanums + '_')
    integer = Word("0x" + nums)

    enum_value = Group(
        identifier('name') + Optional(_equal + integer('value')),
    )
    enum_list = Group(enum_value + ZeroOrMore(_comma + enum_value))
    enum = _enum + Optional(identifier('enum_prefix')) + _lcurl + \
           enum_list('list') + _rcurl + Optional(identifier('enum_postfix'))
    enum.ignore(cppStyleComment)

    for item, _, _ in enum.scanString(sample):  # item, start, stop
        enum_name = item.enum_postfix
        enum_props = collections.OrderedDict()
        enum_int_val = 0
        for entry in item.list:
            if entry.value != '':
                enum_int_val = int(entry.value, 0)
            enum_props[entry.name] = enum_int_val
            enum_int_val += 1
        enums[enum_name] = enum_props
    _write_enums(out_fname, enums)
    return enums


def _write_class(out_fname, class_name, class_map, base_class):
    """
    Write a class with a set of C param types to names mapping deriving from a
    C compatible base_class like Structure or Union.
    """
    if not class_map:
        return
    with open(out_fname, 'a') as fd:
        class_str = "class %s(%s):\n" % (class_name, base_class)
        prop_prefix = "    _fields_ = [\n"
        fd.write(class_str)
        fd.write(prop_prefix)
        for struct_prop, struct_type in class_map.items():
            prop_str = ' ' * 8 + '(\"%s\", %s),\n' \
                                 % (struct_prop, struct_type)
            fd.write(prop_str)
        fd.write("    ]\n")
        fd.write('\n\n')


def _extract_inner_params(parsed_struct):
    """

    """
    params = collections.OrderedDict()
    for param in parsed_struct.inner_param_list:
        param_name, param_type = normalize_param(
            param[0].param_type,
            param[0].param_name, None,
            param[0].length,
        )

        params[param_name] = param_type
    return params


def extract_inner_struct(parsed_struct, out_fname):
    """
    Extract the param type and name for the anonymous struct and write
    out the struct as an explicit python Structure
    Args:
        parsed_struct: The parsed anonymous struct ParseResult object
        out_fname: The filename to write the anonymous struct to.
    Returns:
        param_type, param_name and array length of the anonymous struct
    """
    params = _extract_inner_params(parsed_struct)
    _write_class(out_fname, parsed_struct.inner_struct, params, "Structure")
    return parsed_struct.inner_struct, parsed_struct.inner_param_name, None, \
           parsed_struct.length


def extract_inner_union(parsed_union, out_fname):
    """
    Extract the param type and name for the anonymous union and write
    out the union as an explicit python Union
    Args:
        parsed_union: The parsed anonymous union ParseResult object
        out_fname: The filename to write the anonymous union to

    Returns:
        param_type, param_name and array length for the union
    """
    params = _extract_inner_params(parsed_union)
    union_type_name = parsed_union.inner_param_name.title()
    _write_class(out_fname, union_type_name, params, "Union")
    return union_type_name, parsed_union.inner_param_name, None, \
           parsed_union.length


def extract_simple_param(parsed_struct):
    return parsed_struct[0].param_type, parsed_struct[0].param_name, \
           parsed_struct[0].ptr, parsed_struct[0].length


def normalize_param(param_type, param_name, ptr, length):
    """
    Given a param_type, name, ptr field and length generate a normalized C
    representation for the same to be representated as a Python Structure
    Args:
        param_type: The type of the param
        param_name: The name of the param
        ptr: Set if the param is a pointer
        length: Set if the param is an array, and is the length of the array

    Returns:
        The param name and its normalized type
    """
    p_type = param_type
    if param_type in py_type_map:
        # Replace typedef with corresponding c_type
        p_type = py_type_map[param_type]
    elif param_type in enums:
        # There is no native enum type in python
        p_type = 'c_uint'
    else:
        # Assert we know about this struct or const
        assert p_type in c_struct or p_type in defines or \
               p_type in anon_types or p_type in enums, \
            "{} not in {} or {} or {} or {}".format(
                p_type, c_struct, defines,
                anon_types, enums,
            )
    if ptr:
        p_type = 'POINTER(%s)' % p_type
    if length:
        p_type = '%s * %s' % (p_type, length)
    return param_name, p_type


def extract_structs(in_fname, out_fname):
    """
    Extract a C struct and write its corresponding C compatible python object
    Args:
        in_fname: Filename to read structs from
        out_fname: Filename to write structs to.

    Returns:
        A dict of c_structs written.
    """
    global anon_types
    global c_struct
    with open(in_fname, 'r') as header:
        sample = header.read()
    identifier = Word(alphas + '_', alphanums + '_')
    integer = Word(alphanums + '_')
    struct_array = _lsq_bracket + integer('length') + _rsq_bracket
    pointer = Word('*', max=1)
    struct_param = Group(
        identifier('param_type') + Optional(pointer('ptr')) +
        identifier('param_name') + Optional(struct_array),
    )

    simple_param = Group(struct_param + _semi_colon)

    # Anonymous structs (nested)
    inner_param_list = Group(ZeroOrMore(simple_param('inner_list')))
    anonymous_struct_param = Group(
        _struct +
        Optional(identifier('inner_struct')) +
        _lcurl +
        inner_param_list('inner_param_list') +
        _rcurl + identifier('inner_param_name') +
        Optional(struct_array) + _semi_colon,
    )
    anonymous_struct_param.ignore(cppStyleComment)

    # Anonymous unions (nested)
    anonymous_union_param = Group(
        _union +
        Optional(identifier('inner_union')) +
        _lcurl +
        inner_param_list('inner_param_list') +
        _rcurl + identifier('inner_param_name') +
        Optional(struct_array) + _semi_colon,
    )

    # A struct can contain a list of
    # 1. simple param types params name;
    # 2. Anonymous nested structs.
    # 3. Anonymous nested unions.
    struct_list = Group(
        ZeroOrMore(
            simple_param('simple_param') |
            anonymous_struct_param('anon_struct') |
            anonymous_union_param('anon_union'),
        ),
    )

    struct = _struct + Optional(identifier('struct_prefix')) + _lcurl + \
             struct_list('list') + _rcurl + \
             Optional(identifier('struct_postfix')) + _semi_colon
    struct.ignore(cppStyleComment)

    for item, _, _ in struct.scanString(sample):  # item, start, stop
        struct_name = item.struct_postfix
        params = collections.OrderedDict()
        for parsed_entry in item.list:

            # Classify the struct member.
            if (parsed_entry.getName() == "anon_struct"):
                param_type, param_name, ptr, length = \
                    extract_inner_struct(parsed_entry, out_fname)
                anon_types[param_type] = parsed_entry
            elif (parsed_entry.getName() == "anon_union"):
                param_type, param_name, ptr, length = \
                    extract_inner_union(parsed_entry, out_fname)
                anon_types[param_type] = parsed_entry
            elif (parsed_entry.getName() == "simple_param"):
                param_type, param_name, ptr, length = \
                    extract_simple_param(parsed_entry)
            else:
                assert False, "Invalid parse"

            # Normalize the parameter type
            param_name, p_type = normalize_param(
                param_type, param_name,
                ptr, length,
            )
            params[param_name] = p_type
            c_struct[struct_name] = params

        # Write out the struct
        _write_class(out_fname, struct_name, params, "Structure")
    return c_struct


if __name__ == "__main__":
    s1ap_hdr_path = os.path.join(os.environ['S1AP_TESTER_ROOT'], 'bin')
    output_file = os.path.join(s1ap_hdr_path, OUT_FILE)

    _write_copyright(output_file)
    for header_file in HEADER_FILES:
        print("Parsing file %s" % header_file)
        f_name = os.path.join(s1ap_hdr_path, header_file)
        defines.update(extract_defines(f_name, output_file))
        enums.update(extract_enums(f_name, output_file))
        extract_structs(f_name, output_file)
