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
	// AuthenticationParameterAUTN IE Class
	class AuthenticationParameterAUTNMsg
	{
		public:
      #define AUTN_MIN_LEN 16
      #define AUTN_MAX_LEN 16
			uint8_t iei;
			std::string AUTN;

			AuthenticationParameterAUTNMsg();
			~AuthenticationParameterAUTNMsg();
			int EncodeAuthenticationParameterAUTNMsg (AuthenticationParameterAUTNMsg *authenticationparameterautn, uint8_t iei, uint8_t * buffer, uint32_t len);
			int DecodeAuthenticationParameterAUTNMsg (AuthenticationParameterAUTNMsg *authenticationparameterautn, uint8_t iei, uint8_t * buffer, uint32_t len);
	};
}
