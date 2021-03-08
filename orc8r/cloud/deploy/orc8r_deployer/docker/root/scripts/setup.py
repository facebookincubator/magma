# Copyright 2021 The Magma Authors.

# This source code is licensed under the BSD-style license found in the
# LICENSE file in the root directory of this source tree.

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
from setuptools import setup

setup(
    name='deployer',
    version='0.1',
    py_modules=['configlib', 'orcl'],
    install_requires=[
        'Click',
    ],
    entry_points='''
        [console_scripts]
        orcl=orcl:cli
    ''',
)