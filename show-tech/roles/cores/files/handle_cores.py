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
import argparse
import click
import glob
import gzip
import logging
import os
import shutil
import subprocess
import shlex


logging.basicConfig(format='%(levelname)s: %(message)s' ,level=logging.INFO)
logger = logging.getLogger(__name__)

DEFAULT_CORE_MAP = {
    "mme": {
        "path": "/tmp/core-*/core-*-mme-*",
        "binary": "/usr/local/bin/mme",
    }
}


class ComponentCores(object):
    def __init__(self, cores_map, component, max_age, dest_dir):
        self.component = component
        self.component_data = cores_map[component]
        self.max_age = max_age
        self.dest_dir = dest_dir
        self.cores = glob.glob(self.component_data["path"])
        self.app_binary = self.component_data["binary"]
        self.core_dirs = { os.path.dirname(x) for x in self.cores }

    def get_core_files(self):
        return self.cores

    def get_core_dirs(self):
        return self.core_dirs

    def process_cores(self):
        # Copy them to destination folder
        logger.info(f"Processing cores of component {self.component} dirs {self.get_core_dirs()}")
        for core_dir in self.get_core_dirs():
            dest_core_dir = os.path.join(self.dest_dir, os.path.basename(core_dir))
            logger.info(f"Copying {core_dir} to {dest_core_dir}")
            if os.path.exists(dest_core_dir):
                shutil.rmtree(dest_core_dir)
            shutil.copytree(core_dir, dest_core_dir)
        # Uncompress them on source dir
        for core_file in  self.get_core_files():
            logger.debug(f"Analyzing {core_file}")
            core = CoreFile(core_file, self.app_binary, self.dest_dir)
            core.analyze()


class CoreFile(object):
    def __init__(self, core_file_name, app_binary, dest_dir):
        self.core_file_name = core_file_name
        self.uncompressed_core_file = core_file_name
        self.app_binary = app_binary
        self.dest_dir = dest_dir

    def uncompress_file(self):
        if self.is_compressed():
            uncompressed_core_file = self.core_file_name.replace(".gz", "")
            with gzip.open(self.core_file_name, 'rb') as f_in:
                with open(uncompressed_core_file, 'wb') as f_out:
                    shutil.copyfileobj(f_in, f_out)
            self.uncompressed_core_file = uncompressed_core_file
            os.remove(self.core_file_name)

    def is_compressed(self):
        if self.core_file_name.endswith(".gz"):
            return True
        return False

    def analyze(self):
        self.uncompress_file()
        cmd = f"gdb --batch --quiet -ex 'thread apply all bt full' -ex 'quit'  {self.app_binary} {self.uncompressed_core_file}"
        core_dest_dir = os.path.join(self.dest_dir, os.path.basename(os.path.dirname(self.uncompressed_core_file)))
        dbg_file = os.path.join(core_dest_dir, "dbg.txt")
        os.makedirs(core_dest_dir, exist_ok=True)
        if os.path.isfile(dbg_file):
            # trying to avoid duplicates of analyzing gzipped files
            return
        logger.info(f"dbg_file {dbg_file}")

        with open(dbg_file, 'a') as fout:
            ret = subprocess.run(
                shlex.split(cmd),
                check=True,
                stdout=fout,
                stderr=fout,
                timeout=60
            )


@click.command()
@click.option(
    "--cores-map",
    help="Map of core files to collect and process with binary",
    default=DEFAULT_CORE_MAP
)
@click.option(
    "--component",
    help="Component name like mme",
    required=True,
)
@click.option(
    "--max-age",
    help="Age of core files to process in days",
    default=7,
)
@click.option(
    "--dest-dir",
    help="destination directory to place cores and dbg outputs in",
    required=True,
)
def main(cores_map : dict, component: str, max_age: int, dest_dir: str):
    print(f"processing cores on component {component}")
    c = ComponentCores(cores_map, component, max_age, dest_dir)
    c.process_cores()


if __name__ == "__main__":
    main()
