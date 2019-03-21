#!/usr/bin/env python3
#
# Copyright (c) 2016-present, Facebook, Inc.
# All rights reserved.
#
# This source code is licensed under the BSD-style license found in the
# LICENSE file in the root directory of this source tree. An additional grant
# of patent rights can be found in the PATENTS file in the same directory.

# This script creates the build context for the orc8r docker builds.
# It first creates a tmp directory, and then copies the cloud directories
# for all modules into it.

import argparse
import glob
import os
import shutil
import subprocess
import yaml
from typing import List

BUILD_CONTEXT = "/tmp/magma_orc8r_build"
MAGMA_ROOT = "../../../."
DEFAULT_MODULES_FILE = os.path.join(MAGMA_ROOT, "modules.yml")
FB_MODULES_FILE = os.path.join(MAGMA_ROOT, "fb/config/modules.yml")


def _get_module_dirs() -> List[str]:
    """ Read the modules config file, and returns the list of module dirs """
    filename = os.environ.get("MAGMA_MODULES_FILE", DEFAULT_MODULES_FILE)
    # Use the FB modules file if the file exists
    if os.path.isfile(FB_MODULES_FILE):
        filename = FB_MODULES_FILE
    module_dirs = []
    with open(filename) as file:
        conf = yaml.safe_load(file)
        for module in conf["native_modules"]:
            module_dirs.append(os.path.join(MAGMA_ROOT, module))
        for ext_modules in conf["external_modules"]:
            # NOTE: the external modules need to be relative to
            # the orc8r/cloud directory
            module_dirs.append(os.path.join(MAGMA_ROOT, "orc8r", "cloud",
                                            ext_modules["host_path"]))
    return module_dirs


def _create_build_context() -> None:
    """ Clear out the build context from the previous run """
    if os.path.exists(BUILD_CONTEXT):
        shutil.rmtree(BUILD_CONTEXT)
    os.mkdir(BUILD_CONTEXT)

    print("Creating build context in '%s'..." % BUILD_CONTEXT)
    modules = []
    for module_dir in _get_module_dirs():
        module = os.path.basename(module_dir)
        _copy_module(module, module_dir)
        modules.append(module)
    print("Context created for modules: %s" % ", ".join(modules))


def _copy_module(module: str, src: str) -> None:
    """ Copy the module dir into the build context  """
    if not os.path.isdir(src):
        print("ERROR: '%s' is not a directory!" % src)
        exit(1)

    # Copy the module to the magma directory
    dst = os.path.join(BUILD_CONTEXT, "magma", module)
    shutil.copytree(os.path.join(src, "cloud"), os.path.join(dst, "cloud"))
    # Copy the tools directory if it exists for the module
    if os.path.isdir(os.path.join(src, "tools")):
        shutil.copytree(os.path.join(src, "tools"), os.path.join(dst, "tools"))

    # Copy the config to the configs directory
    shutil.copytree(
        os.path.join(src, "cloud", "configs"),
        os.path.join(BUILD_CONTEXT, "configs", module))

    # Copy the go.mod file for caching the go downloads
    for filename in glob.iglob(dst + "/**/go.mod", recursive=True):
        gomod = filename.replace(
            dst, os.path.join(BUILD_CONTEXT, "gomod", module))
        os.makedirs(os.path.dirname(gomod))
        shutil.copyfile(filename, gomod)


def _get_mount_volumes() -> List[str]:
    """ Return the volumes argument for docker-compose commands """
    volumes = []
    cwd = os.getcwd()
    for module_dir in _get_module_dirs():
        module = os.path.basename(module_dir)
        volumes.extend(["-v", "%s/%s:/magma/%s" % (cwd, module_dir, module)])
    return volumes


def _run_docker(cmd: List[str]) -> None:
    """ Run the required docker-compose command """
    print("Running 'docker-compose %s'..." % " ".join(cmd))
    try:
        subprocess.run(["docker-compose"] + cmd, check=True)
    except subprocess.CalledProcessError as err:
        exit(err.returncode)


def _parse_args() -> argparse.Namespace:
    """ Parse the command line args """
    parser = argparse.ArgumentParser(description="Orc8r build tool")
    parser.add_argument("--tests", "-t", action="store_true",
                        help="Run unit tests")
    parser.add_argument("--mount", "-m", action="store_true",
                        help="Mount the source code and create a bash shell")
    args = parser.parse_args()
    return args


def main() -> None:
    args = _parse_args()
    if args.mount:
        # Mount the source code and run a contrainer with bash shell
        _run_docker(["run", "--rm"] + _get_mount_volumes() + ["test", "bash"])
    elif args.tests:
        # Run unit tests
        _create_build_context()
        _run_docker(["build", "test"])
        _run_docker(["run", "--rm", "test", "make test"])
    else:
        # Build all containers
        _create_build_context()
        _run_docker(["build"])


if __name__ == '__main__':
    main()
