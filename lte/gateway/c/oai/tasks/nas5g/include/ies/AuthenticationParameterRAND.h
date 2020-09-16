/*
	 Copyright 2020 The Magma Authors.
	 This source code is licensed under the BSD-style license found in the
	 LICENSE file in the root directory of this source tree.
	 Unless required by applicable law or agreed to in writing, software
	 distributed under the License is distributed on an "AS IS" BASIS,
	 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	 See the License for the specific language governing permissions and
	 limitations under the License.
 */

#pragma once
#include <sstream>
#include <cstdint>

using namespace std;

namespace magma5g
{
	// AuthenticationParameterRANDM IE Class
	class AuthenticationParameterRANDMsg
	{
		public:
      #define RAND_MIN_LEN 16
      #define RAND_MAX_LEN 16
			uint8_t iei;
			std::string randval;

			AuthenticationParameterRANDMsg();
			~AuthenticationParameterRANDMsg();
			int EncodeAuthenticationParameterRANDMsg (AuthenticationParameterRANDMsg *authenticationparameterrand, uint8_t iei, uint8_t * buffer, uint32_t len);
			int DecodeAuthenticationParameterRANDMsg ( AuthenticationParameterRANDMsg *authenticationparameterrand, uint8_t iei, uint8_t * buffer, uint32_t len);
	};
}
