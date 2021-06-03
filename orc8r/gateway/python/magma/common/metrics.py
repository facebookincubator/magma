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
from prometheus_client import Counter

STREAMER_RESPONSES = Counter(
    'streamer_responses',
    'The number of responses by label',
    ['result'],
)

SERVICE_ERRORS = Counter(
    'service_errors',
    'The number of errors logged',
)
