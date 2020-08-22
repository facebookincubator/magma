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

package handlers

import (
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

// GatewayModel describes models that represent a certain type of gateway.
// For example, an LTE gateway, that can be read/updated/deleted.
type GatewayModel interface{}

// GatewayModels keyed by gateway ID.
type GatewayModels map[string]GatewayModel

// PartialGatewayModel describe models that represents a portion of network
// entity that can be read and updated.
type PartialGatewayModel interface {
	serde.ValidatableModel
	// FromBackendModels the same PartialGatewayModel from the configurator
	// entities attached to the networkID and gatewayID.
	FromBackendModels(networkID string, gatewayID string) error
	// ToUpdateCriteria returns a list of EntityUpdateCriteria needed to apply
	// the change in the model.
	ToUpdateCriteria(networkID string, gatewayID string) ([]configurator.EntityUpdateCriteria, error)
}

type MakeTypedGateways func(
	networkID string,
	entsByTK map[storage.TypeAndKey]configurator.NetworkEntity,
	devicesByID map[string]interface{},
	statusesByID map[string]*models.GatewayStatus,
) (GatewayModels, error)

// GetPartialGatewayHandlers returns both GET and PUT handlers for modifying the portion of a
// network entity specified by the model.
// - path : the url at which the handler will be registered.
// - model: the input and output of the handler and it also provides FromBackendModels
//   and ToUpdateCriteria to go between the configurator model.
func GetPartialGatewayHandlers(path string, model PartialGatewayModel) []obsidian.Handler {
	return []obsidian.Handler{
		GetPartialReadGatewayHandler(path, model),
		GetPartialUpdateGatewayHandler(path, model),
	}
}

// GetPartialReadGatewayHandler returns a GET obsidian handler at the specified path.
// This function loads a portion of the gateway specified by the model's FromBackendModels function.
// Example:
//      (m *MagmadGatewayConfigs) FromBackendModels(networkID, gatewayID) (PartialGatewayModel, error) {
// 			return configurator.LoadEntityConfig(networkID, orc8r.MagmadGatewayType, gatewayID)
// 		}
// 		getMagmadConfigsHandler := handlers.GetPartialReadGatewayHandler(URL, &models.MagmadGatewayConfigs{})
//
//      would return a GET handler that can read the magmad gateway config of a gw with the specified ID.
func GetPartialReadGatewayHandler(path string, model PartialGatewayModel) obsidian.Handler {
	return obsidian.Handler{
		Path:    path,
		Methods: obsidian.GET,
		HandlerFunc: func(c echo.Context) error {
			networkID, gatewayID, nerr := obsidian.GetNetworkAndGatewayIDs(c)
			if nerr != nil {
				return nerr
			}

			err := model.FromBackendModels(networkID, gatewayID)
			if err == merrors.ErrNotFound {
				return obsidian.HttpError(err, http.StatusNotFound)
			} else if err != nil {
				return obsidian.HttpError(err, http.StatusInternalServerError)
			}
			return c.JSON(http.StatusOK, model)
		},
	}
}

// GetPartialUpdateGatewayHandler returns a PUT obsidian handler at the specified path.
// This function updates a portion of the network entity specified by the model's ToUpdateCriteria function.
// Example:
//      (m *MagmadGatewayConfigs) ToUpdateCriteria(networkID, gatewayID) ([]configurator.EntityUpdateCriteria, error) {
// 			return []configurator.EntityUpdateCriteria{
//				{
// 					Key: gatewayID,
//					Type: orc8r.MagmadGatewayType,
//					NewConfig: m,
//				}
//          }
// 		}
// 		updateMagmadConfigsHandler := handlers.GetPartialUpdateGatewayHandler(URL, &models.MagmadGatewayConfigs{})
//
//      would return a PUT handler that updates the magmad gateway config of a gw with the specified ID.
func GetPartialUpdateGatewayHandler(path string, model PartialGatewayModel) obsidian.Handler {
	return obsidian.Handler{
		Path:    path,
		Methods: obsidian.PUT,
		HandlerFunc: func(c echo.Context) error {
			networkID, gatewayID, nerr := obsidian.GetNetworkAndGatewayIDs(c)
			if nerr != nil {
				return nerr
			}

			requestedUpdate, nerr := GetAndValidatePayload(c, model)
			if nerr != nil {
				return nerr
			}

			updates, err := requestedUpdate.(PartialGatewayModel).ToUpdateCriteria(networkID, gatewayID)
			if err != nil {
				return obsidian.HttpError(err, http.StatusBadRequest)
			}
			_, err = configurator.UpdateEntities(networkID, updates)
			if err != nil {
				return obsidian.HttpError(err, http.StatusInternalServerError)
			}
			return c.NoContent(http.StatusNoContent)
		},
	}
}

// GetGatewayDeviceHandlers returns GET and PUT handlers to read and update the
// device attached to the gateway.
func GetGatewayDeviceHandlers(path string) []obsidian.Handler {
	return []obsidian.Handler{
		GetReadGatewayDeviceHandler(path),
		GetUpdateGatewayDeviceHandler(path),
	}
}

// GetReadGatewayDeviceHandler returns a GET handler to read the gateway record
// of the gateway.
func GetReadGatewayDeviceHandler(path string) obsidian.Handler {
	return obsidian.Handler{
		Path:    path,
		Methods: obsidian.GET,
		HandlerFunc: func(c echo.Context) error {
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
			device, err := device.GetDevice(networkID, orc8r.AccessGatewayRecordType, physicalID)
			if err == merrors.ErrNotFound {
				return obsidian.HttpError(err, http.StatusNotFound)
			} else if err != nil {
				return obsidian.HttpError(err, http.StatusInternalServerError)
			}

			return c.JSON(http.StatusOK, device)
		},
	}
}

// GetUpdateGatewayDeviceHandler returns a PUT handler to update the gateway
// record of the gateway.
func GetUpdateGatewayDeviceHandler(path string) obsidian.Handler {
	return obsidian.Handler{
		Path:    path,
		Methods: obsidian.PUT,
		HandlerFunc: func(c echo.Context) error {
			networkID, gatewayID, nerr := obsidian.GetNetworkAndGatewayIDs(c)
			if nerr != nil {
				return nerr
			}
			update, nerr := GetAndValidatePayload(c, &models.GatewayDevice{})
			if nerr != nil {
				return nerr
			}

			physicalID, err := configurator.GetPhysicalIDOfEntity(networkID, orc8r.MagmadGatewayType, gatewayID)
			if err == merrors.ErrNotFound {
				return obsidian.HttpError(err, http.StatusNotFound)
			} else if err != nil {
				return obsidian.HttpError(err, http.StatusInternalServerError)
			}
			err = device.UpdateDevice(networkID, orc8r.AccessGatewayRecordType, physicalID, update)
			if err != nil {
				return obsidian.HttpError(err, http.StatusInternalServerError)
			}
			return c.NoContent(http.StatusNoContent)
		},
	}
}

func GetListGatewaysHandler(path string, gatewayType string, makeTypedGateways MakeTypedGateways) obsidian.Handler {
	return obsidian.Handler{
		Path:    path,
		Methods: obsidian.GET,
		HandlerFunc: func(c echo.Context) error {
			nid, nerr := obsidian.GetNetworkId(c)
			if nerr != nil {
				return nerr
			}

			ids, err := configurator.ListEntityKeys(nid, gatewayType)
			if err != nil {
				return obsidian.HttpError(err, http.StatusInternalServerError)
			}
			// for each ID, we want to load the gateway and the magmad gateway
			magmadTKs := make([]storage.TypeAndKey, 0, len(ids))
			for _, id := range ids {
				magmadTKs = append(
					magmadTKs,
					storage.TypeAndKey{Type: orc8r.MagmadGatewayType, Key: id},
				)
			}
			gwTKs := make([]storage.TypeAndKey, 0, len(ids))
			for _, id := range ids {
				gwTKs = append(
					gwTKs,
					storage.TypeAndKey{Type: gatewayType, Key: id},
				)
			}

			// we need the two calls because LoadEntities only takes one type filter per call
			// and we need the type filters in case the TKs are empty
			magmadGWEnts, _, err := configurator.LoadEntities(nid, swag.String(orc8r.MagmadGatewayType), nil, nil, magmadTKs, configurator.FullEntityLoadCriteria())
			gwEnts, _, err := configurator.LoadEntities(nid, swag.String(gatewayType), nil, nil, gwTKs, configurator.FullEntityLoadCriteria())
			ents := configurator.NetworkEntities{}
			for _, m := range magmadGWEnts {
				ents = append(ents, m)
			}
			for _, g := range gwEnts {
				ents = append(ents, g)
			}
			entsByTK := ents.ToEntitiesByID()

			// for each magmad gateway, we have to load its corresponding device and
			// its reported status
			deviceIDs := make([]string, 0, len(ids))
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

			gateways, err := makeTypedGateways(nid, entsByTK, devicesByID, statusesByID)
			if err != nil {
				return obsidian.HttpError(err, http.StatusInternalServerError)
			}
			return c.JSON(http.StatusOK, gateways)
		},
	}
}

func GetDeleteGatewayHandler(gateway GatewaySubtype) func(c echo.Context) error {
	f := func(c echo.Context) error {
		nid, gid, nerr := obsidian.GetNetworkAndGatewayIDs(c)
		if nerr != nil {
			return nerr
		}

		err := gateway.Load(nid, gid)
		switch {
		case err == merrors.ErrNotFound:
			return echo.ErrNotFound
		case err != nil:
			return obsidian.HttpError(errors.Wrap(err, "failed to load gateway"), http.StatusInternalServerError)
		}
		magmadGateway := gateway.GetMagmadGateway()

		var toDelete []storage.TypeAndKey
		toDelete = append(toDelete, magmadGateway.GetAdditionalDeletes()...)
		switch gateway.(type) {
		case *models.MagmadGateway:
			break
		default:
			toDelete = append(toDelete, gateway.GetAdditionalDeletes()...)
		}

		err = configurator.DeleteEntities(nid, toDelete)
		if err != nil {
			return obsidian.HttpError(errors.Wrap(err, "error deleting gateway"), http.StatusInternalServerError)
		}

		// Now we delete the associated device. Even though we error out
		// request if this fails, failing on this specific step is non-
		// blocking because gateway registration handles the case where a
		// device already exists and is unassigned.
		physicalID := magmadGateway.Device.HardwareID
		if physicalID != "" {
			err = device.DeleteDevice(nid, orc8r.AccessGatewayRecordType, physicalID)
			if err != nil {
				return obsidian.HttpError(errors.Wrap(err, "failed to delete device for gateway. no further action is required"), http.StatusInternalServerError)
			}
		}

		return c.NoContent(http.StatusNoContent)
	}

	return f
}
