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

// File registry.go provides an indexer registry by forwarding calls to
// the service registry.

package indexer

import (
	"strconv"
	"testing"

	"github.com/pkg/errors"
	"github.com/thoas/go-funk"

	"magma/orc8r/cloud/go/orc8r"
	"magma/orc8r/lib/go/registry"
)

// GetIndexer returns the remote indexer for a desired service.
// Returns nil if not found.
func GetIndexer(serviceName string) (Indexer, error) {
	x, err := getIndexer(serviceName)
	if err != nil {
		return nil, errors.Wrapf(err, "get indexer for service %s", serviceName)
	}
	return x, nil
}

// GetIndexers returns all registered indexers.
func GetIndexers() ([]Indexer, error) {
	indexingServices := registry.FindServices(orc8r.StateIndexerLabel)

	var ret []Indexer
	for _, serviceName := range indexingServices {
		x, err := getIndexer(serviceName)
		if err != nil {
			return nil, err
		}
		ret = append(ret, x)
	}
	return ret, nil
}

// GetIndexersForState returns all registered indexers which handle the passed
// state type.
func GetIndexersForState(stateType string) ([]Indexer, error) {
	indexers, err := GetIndexers()
	if err != nil {
		return nil, err
	}

	var filtered []Indexer
	for _, x := range indexers {
		if funk.Contains(x.GetTypes(), stateType) {
			filtered = append(filtered, x)
		}

	}
	return filtered, nil
}

// getIndexer returns a new remote indexer generated by parsing
// service metadata.
func getIndexer(serviceName string) (Indexer, error) {
	versionVal, err := registry.GetAnnotation(serviceName, orc8r.StateIndexerVersionAnnotation)
	if err != nil {
		return nil, err
	}
	versionInt, err := strconv.Atoi(versionVal)
	if err != nil {
		return nil, errors.Wrapf(err, "convert indexer version %v to int for service %s", versionVal, serviceName)
	}
	version, err := NewIndexerVersion(int64(versionInt))
	if err != nil {
		return nil, err
	}

	types, err := registry.GetAnnotationFields(serviceName, orc8r.StateIndexerTypesAnnotation, orc8r.AnnotationFieldSeparator)
	if err != nil {
		return nil, err
	}

	return NewRemoteIndexer(serviceName, version, types...), nil
}

// DeregisterAllForTest deregisters all previously-registered indexers.
// This should only be called by test code.
func DeregisterAllForTest(t *testing.T) {
	if t == nil {
		panic("for tests only")
	}
	registry.RemoveServicesWithLabel(orc8r.StateIndexerLabel)
}
