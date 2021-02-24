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

package swagger

import (
	"io/ioutil"

	swagger_lib "magma/orc8r/cloud/go/swagger"

	"github.com/golang/glog"
	"github.com/pkg/errors"
)

var (
	commonSpecPath = "/etc/magma/configs/orc8r/swagger_specs/common/swagger-common.yml"
)

// GetCombinedSpec polls every servicer registered with
// a Swagger spec and merges them together to return a combined spec.
func GetCombinedSpec(yamlCommon string) (string, error) {
	servicers, err := GetSpecServicers()
	if err != nil {
		return "", err
	}

	var yamlSpecs []string
	for _, s := range servicers {
		yamlSpec, err := s.GetSpec()
		if err != nil {
			// Swallow GetSpec error because the polling should continue
			// even if it fails to receive from a single servicer
			err = errors.Wrapf(err, "get Swagger spec from %s service", s.GetService())
			glog.Error(err)
		} else {
			yamlSpecs = append(yamlSpecs, yamlSpec)
		}
	}

	combined, warnings, err := swagger_lib.Combine(yamlCommon, yamlSpecs)
	if err != nil {
		return "", err
	}
	if warnings != nil {
		glog.Infof("Some Swagger spec traits were overwritten or unable to be read: %+v", warnings)
	}

	return combined, nil
}

// GetCombinedSpecFromService polls a service for its spec and combines
// it with the common spec to return a combined spec.
func GetCombinedSpecFromService(yamlCommon string, service string) (string, error) {
	remoteSpec := NewRemoteSpec(service)
	yamlSpec, err := remoteSpec.GetSpec()
	if err != nil {
		return "", err
	}

	combined, warnings, err := swagger_lib.Combine(yamlCommon, []string{yamlSpec})
	if err != nil {
		return "", err
	}
	if warnings != nil {
		glog.Infof("Some Swagger spec traits were overwritten or unable to be read: %+v", warnings)
	}

	return combined, nil
}

// GetCommonSpec returns the Swagger common spec as a YAML string.
func GetCommonSpec() (string, error) {
	data, err := ioutil.ReadFile(commonSpecPath)
	if err != nil {
		err = errors.Wrapf(err, "get common Swagger spec")
		return "", err
	}
	return string(data), nil
}
