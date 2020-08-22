/*
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// File gateway_handlers.go provides generic gateway handlers, with hooks for
// specific gateway types.
//
// These handlers do not support updating a gateway's ID.

package handlers

import (
	"fmt"
	"net/http"

	"magma/orc8r/cloud/go/obsidian"
	"magma/orc8r/cloud/go/orc8r"
	"magma/orc8r/cloud/go/serde"
	"magma/orc8r/cloud/go/services/configurator"
	"magma/orc8r/cloud/go/services/device"
	"magma/orc8r/cloud/go/services/orchestrator/obsidian/models"
	"magma/orc8r/cloud/go/services/state/wrappers"
	"magma/orc8r/cloud/go/storage"
	merrors "magma/orc8r/lib/go/errors"

	"github.com/go-openapi/swag"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

// GatewaySubtype represents a subtype of the Magmad gateway.
// The gateway subtype wraps the base Magmad gateway with additional fields by
// creating and associating further network entities with the Magmad gateway.
//
// Note: since the default Magmad gateway model implements this interface as
// well, DO NOT return the base Magmad model from any of the "get additional"
// methods.
type GatewaySubtype interface {
	serde.ValidatableModel

	// Load the gateway model's contents from backing store.
	Load(networkID, gatewayID string) error

	// GetMagmadGateway returns the Magmad gateways wrapped by the subtype.
	GetMagmadGateway() *models.MagmadGateway

	// GetAdditionalWritesOnCreate returns extra write operations to perform
	// during creation.
	// The writes are performed in the same backend transaction with creation
	// of the Magmad gateway.
	GetAdditionalWritesOnCreate() []configurator.EntityWriteOperation

	// GetAdditionalLoadsOnUpdate returns a list of additional entity keys to
	// load during an update.
	// The entities loaded during this operation will be passed to
	// GetAdditionalWritesOnUpdate.
	GetAdditionalLoadsOnUpdate() []storage.TypeAndKey

	// GetAdditionalWritesOnUpdate returns extra write operations to perform
	// during an update.
	// The writes are performed in the same backend transaction with the update
	// of the Magmad gateway.
	GetAdditionalWritesOnUpdate(loadedEntities map[storage.TypeAndKey]configurator.NetworkEntity) ([]configurator.EntityWriteOperation, error)

	// GetAdditionalDeletes returns a list of additional entity keys to delete
	// during deletion.
	// The deletes are performed in the same backend transaction with the
	// deletion of the Magmad gateway.
	GetAdditionalDeletes() []storage.TypeAndKey
}

func CreateGateway(c echo.Context, model GatewaySubtype) *echo.HTTPError {
	nid, nerr := obsidian.GetNetworkId(c)
	if nerr != nil {
		return nerr
	}

	payload, nerr := GetAndValidatePayload(c, model)
	if nerr != nil {
		return nerr
	}
	subGateway := payload.(GatewaySubtype)
	mdGateway := subGateway.GetMagmadGateway()

	// Must associate to an existing tier
	tierExists, err := configurator.DoesEntityExist(nid, orc8r.UpgradeTierEntityType, string(mdGateway.Tier))
	if err != nil {
		return obsidian.HttpError(errors.Wrap(err, "failed to check for tier existence"), http.StatusInternalServerError)
	}
	if !tierExists {
		return echo.NewHTTPError(http.StatusBadRequest, "requested tier does not exist")
	}

	// If the device is already registered, throw an error if it's already
	// assigned to an entity
	// If the device exists but is unassigned, update it to the payload
	// If the device doesn't exist, create it and move on
	deviceID := mdGateway.Device.HardwareID
	_, err = device.GetDevice(nid, orc8r.AccessGatewayRecordType, deviceID)
	switch {
	case err == merrors.ErrNotFound:
		err = device.RegisterDevice(nid, orc8r.AccessGatewayRecordType, deviceID, mdGateway.Device)
		if err != nil {
			return obsidian.HttpError(errors.Wrap(err, "failed to register physical device"), http.StatusInternalServerError)
		}
		break
	case err != nil:
		return obsidian.HttpError(errors.Wrap(err, "failed to check if physical device is already registered"), http.StatusConflict)
	default: // err == nil
		assignedEnt, err := configurator.LoadEntityForPhysicalID(deviceID, configurator.EntityLoadCriteria{})
		switch {
		case err == nil:
			return obsidian.HttpError(errors.Errorf("device %s is already mapped to gateway %s", deviceID, assignedEnt.Key), http.StatusBadRequest)
		case err != merrors.ErrNotFound:
			return obsidian.HttpError(errors.Wrap(err, "failed to check for existing device assignment"), http.StatusInternalServerError)
		}

		if err := device.UpdateDevice(nid, orc8r.AccessGatewayRecordType, deviceID, mdGateway.Device); err != nil {
			return obsidian.HttpError(errors.Wrap(err, "failed to update device record"), http.StatusInternalServerError)
		}
	}

	// Create the magmad gateway, update the tier, perform additional writes
	// as necessary
	var writes []configurator.EntityWriteOperation
	writes = append(writes, mdGateway.GetAdditionalWritesOnCreate()...)
	writes = append(writes, configurator.EntityUpdateCriteria{
		Type:              orc8r.UpgradeTierEntityType,
		Key:               string(mdGateway.Tier),
		AssociationsToAdd: []storage.TypeAndKey{{Type: orc8r.MagmadGatewayType, Key: string(mdGateway.ID)}},
	})
	// These type switches aren't great but it's the best I could think of
	switch payload.(type) {
	case *models.MagmadGateway:
		break
	default:
		writes = append(writes, subGateway.GetAdditionalWritesOnCreate()...)
	}

	if err = configurator.WriteEntities(nid, writes...); err != nil {
		return obsidian.HttpError(errors.Wrap(err, "error creating gateway"), http.StatusInternalServerError)
	}
	return nil
}

func UpdateGateway(c echo.Context, nid string, gid string, model GatewaySubtype) *echo.HTTPError {
	payload, nerr := GetAndValidatePayload(c, model)
	if nerr != nil {
		return nerr
	}
	subGateway := payload.(GatewaySubtype)
	mdGateway := subGateway.GetMagmadGateway()

	if gid != string(mdGateway.ID) {
		err := fmt.Errorf("gateway ID cannot be updated: gateway ID from parameter (%s) and payload (%s) must match", gid, mdGateway.ID)
		return obsidian.HttpError(err, http.StatusBadRequest)
	}

	var entsToLoad []storage.TypeAndKey
	entsToLoad = append(entsToLoad, mdGateway.GetAdditionalLoadsOnUpdate()...)
	switch payload.(type) {
	case *models.MagmadGateway:
		break
	default:
		entsToLoad = append(entsToLoad, subGateway.GetAdditionalLoadsOnUpdate()...)
	}

	loadedEnts, _, err := configurator.LoadEntities(
		nid,
		nil, nil, nil,
		entsToLoad,
		configurator.FullEntityLoadCriteria(),
	)
	if err != nil {
		return obsidian.HttpError(errors.Wrap(err, "failed to load gateway before update"), http.StatusInternalServerError)
	}

	writes, nerr := getUpdateWrites(subGateway, loadedEnts)
	if nerr != nil {
		return nerr
	}

	err = configurator.WriteEntities(nid, writes...)
	if err != nil {
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}

	// Device info is cheap to update, so just do it all the time if
	// configurator write was successful
	err = device.UpdateDevice(nid, orc8r.AccessGatewayRecordType, mdGateway.Device.HardwareID, mdGateway.Device)
	if err != nil {
		return obsidian.HttpError(errors.Wrap(err, "failed to update device info"), http.StatusInternalServerError)
	}

	return nil
}

func GetStateHandler(c echo.Context) error {
	networkID, gatewayID, nerr := obsidian.GetNetworkAndGatewayIDs(c)
	if nerr != nil {
		return nerr
	}

	physicalID, err := configurator.GetPhysicalIDOfEntity(networkID, orc8r.MagmadGatewayType, gatewayID)
	if err == merrors.ErrNotFound {
		return obsidian.HttpError(err, http.StatusNotFound)
	} else if err != nil {
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}

	st, err := wrappers.GetGatewayStatus(networkID, physicalID)
	if err == merrors.ErrNotFound {
		return obsidian.HttpError(err, http.StatusNotFound)
	} else if err != nil {
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, st)
}

func listGatewaysHandler(c echo.Context) error {
	nid, nerr := obsidian.GetNetworkId(c)
	if nerr != nil {
		return nerr
	}

	ents, _, err := configurator.LoadEntities(nid, swag.String(orc8r.MagmadGatewayType), nil, nil, nil, configurator.FullEntityLoadCriteria())
	if err != nil {
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}
	entsByTK := ents.ToEntitiesByID()

	// For each magmad gateway, we have to load its corresponding device and
	// its reported status
	deviceIDs := make([]string, 0, len(entsByTK))
	for tk, ent := range entsByTK {
		if tk.Type == orc8r.MagmadGatewayType && ent.PhysicalID != "" {
			deviceIDs = append(deviceIDs, ent.PhysicalID)
		}
	}

	devicesByID, err := device.GetDevices(nid, orc8r.AccessGatewayRecordType, deviceIDs)
	if err != nil {
		return obsidian.HttpError(errors.Wrap(err, "failed to load devices"), http.StatusInternalServerError)
	}
	statusesByID, err := wrappers.GetGatewayStatuses(nid, deviceIDs)
	if err != nil {
		return obsidian.HttpError(errors.Wrap(err, "failed to load statuses"), http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, makeGateways(entsByTK, devicesByID, statusesByID))
}

func createGatewayHandler(c echo.Context) error {
	if nerr := CreateGateway(c, &models.MagmadGateway{}); nerr != nil {
		return nerr
	}
	return c.NoContent(http.StatusCreated)
}

func getGatewayHandler(c echo.Context) error {
	networkID, gatewayID, nerr := obsidian.GetNetworkAndGatewayIDs(c)
	if nerr != nil {
		return nerr
	}

	gateway := &models.MagmadGateway{}
	err := gateway.Load(networkID, gatewayID)
	if err == merrors.ErrNotFound {
		return echo.ErrNotFound
	}
	if err != nil {
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, gateway)
}

func updateGatewayHandler(c echo.Context) error {
	nid, gid, nerr := obsidian.GetNetworkAndGatewayIDs(c)
	if nerr != nil {
		return nerr
	}

	if nerr = UpdateGateway(c, nid, gid, &models.MagmadGateway{}); nerr != nil {
		return nerr
	}
	return c.NoContent(http.StatusNoContent)
}

func deleteGatewayHandler(c echo.Context) error {
	nid, gid, nerr := obsidian.GetNetworkAndGatewayIDs(c)
	if nerr != nil {
		return nerr
	}

	existingEnt, err := configurator.LoadEntity(
		nid, orc8r.MagmadGatewayType, gid,
		configurator.EntityLoadCriteria{LoadMetadata: true, LoadAssocsToThis: true},
	)
	switch {
	case err == merrors.ErrNotFound:
		return echo.ErrNotFound
	case err != nil:
		return obsidian.HttpError(errors.Wrap(err, "failed to load gateway"), http.StatusInternalServerError)
	}

	err = configurator.DeleteEntity(nid, orc8r.MagmadGatewayType, gid)
	if err != nil {
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}

	if existingEnt.PhysicalID != "" {
		err = device.DeleteDevice(nid, orc8r.AccessGatewayRecordType, existingEnt.PhysicalID)
		if err != nil {
			return obsidian.HttpError(errors.Wrap(err, "failed to delete device for gateway"), http.StatusInternalServerError)
		}
	}

	return c.NoContent(http.StatusNoContent)
}

func getUpdateWrites(payload GatewaySubtype, loadedEnts configurator.NetworkEntities) ([]configurator.EntityWriteOperation, *echo.HTTPError) {
	var writes []configurator.EntityWriteOperation
	loadedEntsByID := loadedEnts.ToEntitiesByID()

	mdGwWrites, err := payload.GetMagmadGateway().GetAdditionalWritesOnUpdate(loadedEntsByID)
	switch {
	case err == merrors.ErrNotFound:
		return writes, echo.ErrNotFound
	case err != nil:
		return writes, obsidian.HttpError(errors.Wrap(err, "failed to get update operations from magmad model"), http.StatusInternalServerError)
	}

	// Short circuit if this is the magmad gateway
	switch payload.(type) {
	case *models.MagmadGateway:
		return mdGwWrites, nil
	}

	payloadWrites, err := payload.GetAdditionalWritesOnUpdate(loadedEntsByID)
	switch {
	case err == merrors.ErrNotFound:
		return writes, echo.ErrNotFound
	case err != nil:
		return writes, obsidian.HttpError(errors.Wrap(err, "failed to get update operations from payload model"), http.StatusInternalServerError)
	}

	writes = append(writes, mdGwWrites...)
	writes = append(writes, payloadWrites...)
	return writes, nil
}

func makeGateways(
	entsByTK map[storage.TypeAndKey]configurator.NetworkEntity,
	devicesByID map[string]interface{},
	statusesByID map[string]*models.GatewayStatus,
) map[string]*models.MagmadGateway {
	gatewayEntsByKey := map[string]*models.MagmadGateway{}
	for tk, ent := range entsByTK {
		hwID := ent.PhysicalID
		var devCasted *models.GatewayDevice
		if devicesByID[hwID] != nil {
			devCasted = devicesByID[hwID].(*models.GatewayDevice)
		}
		gatewayEntsByKey[tk.Key] = (&models.MagmadGateway{}).FromBackendModels(ent, devCasted, statusesByID[hwID])
	}
	return gatewayEntsByKey
}
