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

package subscriberdb_cache_test

import (
	"testing"
	"time"

	"magma/lte/cloud/go/lte"
	"magma/lte/cloud/go/serdes"
	lte_models "magma/lte/cloud/go/services/lte/obsidian/models"
	lte_test_init "magma/lte/cloud/go/services/lte/test_init"
	"magma/lte/cloud/go/services/subscriberdb/obsidian/models"
	"magma/lte/cloud/go/services/subscriberdb/storage"
	"magma/lte/cloud/go/services/subscriberdb_cache"
	"magma/orc8r/cloud/go/clock"
	"magma/orc8r/cloud/go/services/configurator"
	configurator_test_init "magma/orc8r/cloud/go/services/configurator/test_init"
	"magma/orc8r/cloud/go/sqorc"
	"magma/orc8r/cloud/go/test_utils"

	"github.com/stretchr/testify/assert"
)

func TestSubscriberdbCacheWorker(t *testing.T) {
	db, err := test_utils.GetSharedMemoryDB()
	assert.NoError(t, err)
	flatDigestStore := storage.NewFlatDigestLookup(db, sqorc.GetSqlBuilder())
	assert.NoError(t, flatDigestStore.Initialize())
	perSubDigestStore := storage.NewPerSubDigestLookup(db, sqorc.GetSqlBuilder())
	assert.NoError(t, perSubDigestStore.Initialize())
	serviceConfig := subscriberdb_cache.Config{
		SleepIntervalSecs:  5,
		UpdateIntervalSecs: 300,
	}

	lte_test_init.StartTestService(t)
	configurator_test_init.StartTestService(t)

	allNetworks, err := storage.GetAllNetworks(flatDigestStore)
	assert.NoError(t, err)
	assert.Equal(t, []string{}, allNetworks)
	flatDigest, err := storage.GetDigest(flatDigestStore, "n1")
	assert.NoError(t, err)
	checkFlatDigestEqual(t, "", flatDigest, true)
	perSubDigests, err := storage.GetDigest(perSubDigestStore, "n1")
	assert.NoError(t, err)
	assert.Equal(t, storage.DigestInfos{}, perSubDigests)

	err = configurator.CreateNetwork(configurator.Network{ID: "n1"}, serdes.Network)
	assert.NoError(t, err)

	subscriberdb_cache.RenewDigests(flatDigestStore, perSubDigestStore, serviceConfig)
	flatDigest, err = storage.GetDigest(flatDigestStore, "n1")
	assert.NoError(t, err)
	flatDigestCanon := checkFlatDigestEqual(t, "", flatDigest, false)
	perSubDigests, err = storage.GetDigest(perSubDigestStore, "n1")
	assert.NoError(t, err)
	// The apn resources digest should be the last element in the list of per sub digests,
	// and should not be empty. Since no subscribers exist in this network, the apn resources
	// digest should be the first and last in the list of DigestInfos.
	assert.Equal(t, "_apn_resources", perSubDigests[0].Subscriber)
	assert.NotEqual(t, "", perSubDigests[0].Digest)
	apnDigestCanon := perSubDigests[0].Digest

	// Detect outdated digests and update
	_, err = configurator.CreateEntities(
		"n1",
		[]configurator.NetworkEntity{
			{
				Type: lte.APNEntityType, Key: "apn1",
				Config: &lte_models.ApnConfiguration{},
			},
			{
				Type: lte.SubscriberEntityType, Key: "IMSI99999",
				Config: &models.SubscriberConfig{
					Lte: &models.LteSubscription{State: "ACTIVE"},
				},
			},
			{
				Type: lte.SubscriberEntityType, Key: "IMSI11111",
				Config: &models.SubscriberConfig{
					Lte: &models.LteSubscription{State: "ACTIVE"},
				},
			},
		},
		serdes.Entity,
	)
	assert.NoError(t, err)

	clock.SetAndFreezeClock(t, clock.Now().Add(10*time.Minute))
	subscriberdb_cache.RenewDigests(flatDigestStore, perSubDigestStore, serviceConfig)
	flatDigest, err = storage.GetDigest(flatDigestStore, "n1")
	assert.NoError(t, err)
	checkFlatDigestEqual(t, flatDigestCanon, flatDigest, false)

	perSubDigests, err = storage.GetDigest(perSubDigestStore, "n1")
	assert.NoError(t, err)
	// The individual subscriber digests are ordered by subscriber ID
	assert.Equal(t, "11111", perSubDigests[0].Subscriber)
	assert.NotEqual(t, "", perSubDigests[0].Digest)
	assert.Equal(t, "99999", perSubDigests[1].Subscriber)
	assert.NotEqual(t, "", perSubDigests[1].Digest)
	assert.Equal(t, "_apn_resources", perSubDigests[2].Subscriber)
	assert.Equal(t, apnDigestCanon, perSubDigests[2].Digest)
	clock.UnfreezeClock(t)

	// Detect newly added and removed networks
	err = configurator.CreateNetwork(configurator.Network{ID: "n2"}, serdes.Network)
	assert.NoError(t, err)
	configurator.DeleteNetwork("n1")

	clock.SetAndFreezeClock(t, clock.Now().Add(20*time.Minute))
	subscriberdb_cache.RenewDigests(flatDigestStore, perSubDigestStore, serviceConfig)
	flatDigest, err = storage.GetDigest(flatDigestStore, "n1")
	assert.NoError(t, err)
	checkFlatDigestEqual(t, "", flatDigest, true)
	perSubDigests, err = storage.GetDigest(perSubDigestStore, "n1")
	assert.NoError(t, err)
	assert.Equal(t, storage.DigestInfos{}, perSubDigests)

	flatDigest, err = storage.GetDigest(flatDigestStore, "n2")
	assert.NoError(t, err)
	checkFlatDigestEqual(t, "", flatDigest, false)
	perSubDigests, err = storage.GetDigest(perSubDigestStore, "n2")
	assert.NoError(t, err)
	assert.Equal(t, "_apn_resources", perSubDigests[0].Subscriber)
	assert.NotEqual(t, "", perSubDigests[0].Digest)

	allNetworks, err = storage.GetAllNetworks(flatDigestStore)
	assert.NoError(t, err)
	assert.Equal(t, []string{"n2"}, allNetworks)
	clock.UnfreezeClock(t)
}

func checkFlatDigestEqual(t *testing.T, expected string, digest storage.DigestInfos, equal bool) string {
	// A network has at most 1 flat digest in store
	got := ""
	if len(digest) > 0 {
		assert.Equal(t, 1, len(digest))
		got = digest[0].Digest
	}
	if equal {
		assert.Equal(t, expected, got)
	} else {
		assert.NotEqual(t, expected, got)
	}
	return got
}
