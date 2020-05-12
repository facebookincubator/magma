// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package authz_test

import (
	"context"
	"testing"

	models2 "github.com/facebookincubator/symphony/graph/authz/models"

	"github.com/facebookincubator/symphony/graph/graphql/models"
	"github.com/facebookincubator/symphony/graph/viewer/viewertest"
)

func TestEquipmentWritePolicyRule(t *testing.T) {
	c := viewertest.NewTestClient(t)
	ctx := viewertest.NewContext(context.Background(), c)
	equipmentType := c.EquipmentType.Create().
		SetName("EquipmentType").
		SaveX(ctx)
	equipment := c.Equipment.Create().
		SetName("Equipment").
		SetType(equipmentType).
		SaveX(ctx)
	createEquipment := func(ctx context.Context) error {
		_, err := c.Equipment.Create().
			SetName("NewEquipment").
			SetType(equipmentType).
			Save(ctx)
		return err
	}
	updateEquipment := func(ctx context.Context) error {
		return c.Equipment.UpdateOne(equipment).
			SetName("NewName").
			Exec(ctx)
	}
	deleteEquipment := func(ctx context.Context) error {
		return c.Equipment.DeleteOne(equipment).
			Exec(ctx)
	}
	runCudPolicyTest(t, cudPolicyTest{
		getCud: func(p *models.PermissionSettings) *models.Cud {
			return p.InventoryPolicy.Equipment
		},
		create: createEquipment,
		update: updateEquipment,
		delete: deleteEquipment,
	})
}

func TestEquipmentTypeWritePolicyRule(t *testing.T) {
	c := viewertest.NewTestClient(t)
	ctx := viewertest.NewContext(context.Background(), c)
	equipmentType := c.EquipmentType.Create().
		SetName("EquipmentType").
		SaveX(ctx)
	createEquipmentType := func(ctx context.Context) error {
		_, err := c.EquipmentType.Create().
			SetName("NewEquipmentType").
			Save(ctx)
		return err
	}
	updateEquipmentType := func(ctx context.Context) error {
		return c.EquipmentType.UpdateOne(equipmentType).
			SetName("NewName").
			Exec(ctx)
	}
	deleteEquipmentType := func(ctx context.Context) error {
		return c.EquipmentType.DeleteOne(equipmentType).
			Exec(ctx)
	}
	runCudPolicyTest(t, cudPolicyTest{
		getCud: func(p *models.PermissionSettings) *models.Cud {
			return p.InventoryPolicy.EquipmentType
		},
		create: createEquipmentType,
		update: updateEquipmentType,
		delete: deleteEquipmentType,
	})
}

func TestEquipmentPortTypeWritePolicyRule(t *testing.T) {
	c := viewertest.NewTestClient(t)
	ctx := viewertest.NewContext(context.Background(), c)
	equipmentType := c.EquipmentPortType.Create().
		SetName("EquipmentPortType").
		SaveX(ctx)
	createEquipmentPortType := func(ctx context.Context) error {
		_, err := c.EquipmentPortType.Create().
			SetName("NewEquipmentPortType").
			Save(ctx)
		return err
	}
	updateEquipmentPortType := func(ctx context.Context) error {
		return c.EquipmentPortType.UpdateOne(equipmentType).
			SetName("NewName").
			Exec(ctx)
	}
	deleteEquipmentPortType := func(ctx context.Context) error {
		return c.EquipmentPortType.DeleteOne(equipmentType).
			Exec(ctx)
	}
	runCudPolicyTest(t, cudPolicyTest{
		getCud: func(p *models.PermissionSettings) *models.Cud {
			return p.InventoryPolicy.PortType
		},
		create: createEquipmentPortType,
		update: updateEquipmentPortType,
		delete: deleteEquipmentPortType,
	})
}

func TestEquipmentPortDefinitionWritePolicyRule(t *testing.T) {
	c := viewertest.NewTestClient(t)
	ctx := viewertest.NewContext(context.Background(), c)
	equipmentType := c.EquipmentType.Create().
		SetName("EquipmentType").
		SaveX(ctx)
	equipmentPortDefinition := c.EquipmentPortDefinition.Create().
		SetName("EquipmentPortDefinition").
		SetEquipmentType(equipmentType).
		SaveX(ctx)
	createEquipmentPortDefinition := func(ctx context.Context) error {
		_, err := c.EquipmentPortDefinition.Create().
			SetName("NewEquipmentPortDefinition").
			SetEquipmentType(equipmentType).
			Save(ctx)
		return err
	}
	updateEquipmentPortDefinition := func(ctx context.Context) error {
		return c.EquipmentPortDefinition.UpdateOne(equipmentPortDefinition).
			SetName("NewName").
			Exec(ctx)
	}
	deleteEquipmentPortDefinition := func(ctx context.Context) error {
		return c.EquipmentPortDefinition.DeleteOne(equipmentPortDefinition).
			Exec(ctx)
	}
	runCudPolicyTest(t, cudPolicyTest{
		appendPermissions: func(p *models.PermissionSettings) {
			p.InventoryPolicy.EquipmentType.Update.IsAllowed = models2.PermissionValueYes
		},
		create: createEquipmentPortDefinition,
		update: updateEquipmentPortDefinition,
		delete: deleteEquipmentPortDefinition,
	})
}

func TestEquipmentCategoryWritePolicyRule(t *testing.T) {
	c := viewertest.NewTestClient(t)
	ctx := viewertest.NewContext(context.Background(), c)
	equipmentCategory := c.EquipmentCategory.Create().
		SetName("EquipmentCategory").
		SaveX(ctx)
	createEquipmentCategory := func(ctx context.Context) error {
		_, err := c.EquipmentCategory.Create().
			SetName("NewEquipmentCategory").
			Save(ctx)
		return err
	}
	updateEquipmentCategory := func(ctx context.Context) error {
		return c.EquipmentCategory.UpdateOne(equipmentCategory).
			SetName("NewName").
			Exec(ctx)
	}
	deleteEquipmentCategory := func(ctx context.Context) error {
		return c.EquipmentCategory.DeleteOne(equipmentCategory).
			Exec(ctx)
	}
	runCudPolicyTest(t, cudPolicyTest{
		appendPermissions: func(p *models.PermissionSettings) {
			p.InventoryPolicy.EquipmentType.Update.IsAllowed = models2.PermissionValueYes
		},
		create: createEquipmentCategory,
		update: updateEquipmentCategory,
		delete: deleteEquipmentCategory,
	})
}
