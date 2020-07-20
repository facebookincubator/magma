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

package tests

import (
	"reflect"
	"regexp"
	"testing"

	"magma/orc8r/lib/go/protos"

	"github.com/stretchr/testify/assert"
)

func TestIdentities(t *testing.T) {

	wc := protos.NewGatewayWildcardIdentity()
	assert.NotNil(t, wc.ToCommonName())
	assert.Equal(t, "Id_Wildcard_Gateway", wc.HashString())

	wc = protos.NewNetworkWildcardIdentity()
	assert.NotNil(t, wc.ToCommonName())
	assert.Equal(t, "Id_Wildcard_Network", wc.HashString())

	wc = protos.NewOperatorWildcardIdentity()
	assert.NotNil(t, wc.ToCommonName())
	assert.Equal(t, "Id_Wildcard_Operator", wc.HashString())

	id := new(protos.Identity)
	id.SetGateway(nil)
	switch id.Value.(type) {
	case *protos.Identity_Gateway_:
	default:
		t.Fatal("Failed to set Gateway Identity Type")
	}
	assert.Nil(t, id.GetGateway(), "Failed to set Nil Gateway Identity")

	const expectedNilIdHashStr = "Id_<nil>_"
	hs := (*protos.Identity)(nil).HashString()
	assert.Equal(t, expectedNilIdHashStr, hs,
		"Nil Identity's Hash String '%s' doesn't match '%s'",
		hs, expectedNilIdHashStr)

	const expectedGatewayHashStr = "Id_Gateway_<nil>"
	hs = id.HashString()
	assert.Equal(t, expectedGatewayHashStr, hs,
		"Gateway's Hash String '%s' doesn't match '%s'",
		hs, expectedGatewayHashStr)

	id.SetOperator("me")
	switch id.Value.(type) {
	case *protos.Identity_Operator:
	default:
		t.Fatal("Failed to set Gateway Operator Type")
	}
	assert.Equal(t, id.GetOperator(), "me")

	cn := id.ToCommonName()
	assert.NotNil(t, cn)
	assert.Equal(t, "me", *cn, "Operator CN '%s' doesn't match 'me'", *cn)

	const expectedOperatorHashStr = "Id_Operator_me"
	hs = id.HashString()
	assert.Equal(t, expectedOperatorHashStr, hs,
		"Operator's Hash String '%s' doesn't match '%s'",
		hs, expectedOperatorHashStr)

	idgw := protos.Identity_Gateway{HardwareId: "hwid", NetworkId: "netwrk", LogicalId: "logical id"}
	id.SetGateway(&idgw)
	idgw_out := id.GetGateway()
	assert.NotNil(t, idgw_out)
	assert.Equal(t, idgw, *idgw_out)

	cn = id.ToCommonName()
	assert.NotNil(t, cn)
	assert.Equal(t, *cn, "hwid")

	gatewayIdentity := new(protos.Identity).SetGateway(
		&protos.Identity_Gateway{HardwareId: "hwid", NetworkId: "netwrk", LogicalId: "logical id"},
	)
	operatorIdentity := new(protos.Identity).SetOperator("test.ID")
	networkIdentity := new(protos.Identity).SetNetwork("test.ID")

	assert.Equal(t, "Id_Gateway_hwid", gatewayIdentity.HashString())
	assert.Equal(t, "Id_Operator_test.ID", operatorIdentity.HashString())
	assert.Equal(t, "Id_Network_test.ID", networkIdentity.HashString())

	assert.Equal(t, "Gateway", generateIdentityTypeName(t, gatewayIdentity))
	assert.Equal(t, "Operator", generateIdentityTypeName(t, operatorIdentity))
	assert.Equal(t, "Network", generateIdentityTypeName(t, networkIdentity))

	operWildcard := protos.NewOperatorWildcardIdentity()
	networkWildcard := protos.NewNetworkWildcardIdentity()

	assert.Equal(t, "Id_Wildcard_Operator", operWildcard.HashString())
	assert.Equal(t, "Id_Wildcard_Network", networkWildcard.HashString())

	assert.True(t, operWildcard.Match(operatorIdentity))
	assert.False(t, operWildcard.Match(networkIdentity))
	assert.True(t, networkWildcard.Match(networkIdentity))
	assert.False(t, networkWildcard.Match(operatorIdentity))

	id.Value = protos.CreateTestIdentityImplValue()
	assert.Equal(t, "myTestIdentityImpl", generateIdentityTypeName(t, id))
	assert.Equal(t, "Id_<UNDEFINED>_<nil>", id.HashString())

	// Verify that all Hashable Type Names are valid alphanumeric ASCII strings
	var isAlphaNum = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString
	typeTable := protos.GetHashableIdentitiesTable()
	for k, v := range typeTable {
		assert.True(t, isAlphaNum(v), "Invalid Type Name %s for type %s", v, k)
	}
}

// generateIdentityTypeName derives name of Identity type from Identity
// implementation compliant with std autogenerated proto Identity wrapper
// structs:
//	type Identity_Gateway_ struct {
//		Gateway *Identity_Gateway `protobuf:"bytes,1,opt,name=gateway,oneof"`
//	}
//	type Identity_Operator struct {
//		Operator string `protobuf:"bytes,2,opt,name=operator,oneof"`
//	}
//	type Identity_Network struct {
//		Network string `protobuf:"bytes,3,opt,name=network,oneof"`
//	}
//	etc.
// The 'wrapped' field name becomes the Identity type name
func generateIdentityTypeName(t *testing.T, id *protos.Identity) string {
	if id != nil {
		typ := reflect.TypeOf(id.Value)
		// prtoc wraps oneof types into a generated structs with the first
		// element being the actual declared type of oneof field
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem() // dereference ptr
		}
		if typ.Kind() == reflect.Struct && typ.NumField() == 1 {
			return typ.Field(0).Name
		} else {
			t.Fatalf("Unexpected Identity Value Type: %s", typ)
		}
	}
	return ""
}
