// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolver

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/symphony/graph/graphql/models"
	"github.com/facebookincubator/symphony/graph/resolverutil"
	"github.com/facebookincubator/symphony/pkg/ent"
	"github.com/facebookincubator/symphony/pkg/ent/checklistitem"
	"github.com/facebookincubator/symphony/pkg/ent/checklistitemdefinition"
	"github.com/facebookincubator/symphony/pkg/ent/equipment"
	"github.com/facebookincubator/symphony/pkg/ent/file"
	"github.com/facebookincubator/symphony/pkg/ent/link"
	"github.com/facebookincubator/symphony/pkg/ent/property"
	"github.com/facebookincubator/symphony/pkg/ent/propertytype"
	"github.com/facebookincubator/symphony/pkg/ent/workorder"
	"github.com/facebookincubator/symphony/pkg/ent/workordertemplate"
	"github.com/facebookincubator/symphony/pkg/ent/workordertype"
	"github.com/facebookincubator/symphony/pkg/viewer"

	"github.com/AlekSi/pointer"
	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"
)

type workOrderDefinitionResolver struct{}

func (checkListCategoryResolver) CheckListCategoryDefinitions(ctx context.Context, obj *ent.WorkOrderType) ([]*ent.CheckListCategoryDefinition, error) {
	return obj.QueryCheckListCategoryDefinitions().All(ctx)
}

func (workOrderTypeResolver) CheckListCategoryDefinitions(ctx context.Context, obj *ent.WorkOrderType) ([]*ent.CheckListCategoryDefinition, error) {
	return obj.QueryCheckListCategoryDefinitions().All(ctx)
}

func (workOrderDefinitionResolver) Type(ctx context.Context, obj *ent.WorkOrderDefinition) (*ent.WorkOrderType, error) {
	typ, err := obj.Edges.TypeOrErr()
	if ent.IsNotLoaded(err) {
		return obj.QueryType().Only(ctx)
	}
	return typ, err
}

type workOrderTypeResolver struct{}

func (workOrderTypeResolver) PropertyTypes(ctx context.Context, obj *ent.WorkOrderType) ([]*ent.PropertyType, error) {
	return obj.QueryPropertyTypes().All(ctx)
}

func (workOrderTypeResolver) NumberOfWorkOrders(ctx context.Context, obj *ent.WorkOrderType) (int, error) {
	return obj.QueryWorkOrders().Count(ctx)
}

type workOrderTemplateResolver struct{}

func (workOrderTemplateResolver) PropertyTypes(ctx context.Context, obj *ent.WorkOrderTemplate) ([]*ent.PropertyType, error) {
	return obj.QueryPropertyTypes().All(ctx)
}

func (workOrderTemplateResolver) CheckListCategoryDefinitions(ctx context.Context, obj *ent.WorkOrderTemplate) ([]*ent.CheckListCategoryDefinition, error) {
	return obj.QueryCheckListCategoryDefinitions().All(ctx)
}

type workOrderResolver struct{}

func (r workOrderResolver) Activities(ctx context.Context, obj *ent.WorkOrder) ([]*ent.Activity, error) {
	return obj.QueryActivities().All(ctx)
}

func (workOrderResolver) WorkOrderType(ctx context.Context, obj *ent.WorkOrder) (*ent.WorkOrderType, error) {
	typ, err := obj.Edges.TypeOrErr()
	if ent.IsNotLoaded(err) {
		return obj.QueryType().Only(ctx)
	}
	return typ, err
}

func (workOrderResolver) WorkOrderTemplate(ctx context.Context, obj *ent.WorkOrder) (*ent.WorkOrderTemplate, error) {
	t, err := obj.QueryTemplate().Only(ctx)
	return t, ent.MaskNotFound(err)
}

func (workOrderResolver) Location(ctx context.Context, obj *ent.WorkOrder) (*ent.Location, error) {
	loc, err := obj.Edges.LocationOrErr()
	if ent.IsNotLoaded(err) {
		loc, err = obj.QueryLocation().Only(ctx)
	}
	return loc, ent.MaskNotFound(err)
}

func (workOrderResolver) Project(ctx context.Context, obj *ent.WorkOrder) (*ent.Project, error) {
	prj, err := obj.Edges.ProjectOrErr()
	if ent.IsNotLoaded(err) {
		prj, err = obj.QueryProject().Only(ctx)
	}
	return prj, ent.MaskNotFound(err)
}

func (workOrderResolver) CreationDate(_ context.Context, obj *ent.WorkOrder) (int, error) {
	secs := int(obj.CreationDate.Unix())
	return secs, nil
}

func (workOrderResolver) InstallDate(_ context.Context, obj *ent.WorkOrder) (*int, error) {
	secs := int(obj.InstallDate.Unix())
	return &secs, nil
}

func (workOrderResolver) EquipmentToAdd(ctx context.Context, obj *ent.WorkOrder) ([]*ent.Equipment, error) {
	return obj.QueryEquipment().Where(equipment.FutureState(models.FutureStateInstall.String())).All(ctx)
}

func (workOrderResolver) EquipmentToRemove(ctx context.Context, obj *ent.WorkOrder) ([]*ent.Equipment, error) {
	return obj.QueryEquipment().Where(equipment.FutureState(models.FutureStateRemove.String())).All(ctx)
}

func (workOrderResolver) LinksToAdd(ctx context.Context, obj *ent.WorkOrder) ([]*ent.Link, error) {
	return obj.QueryLinks().Where(link.FutureState(models.FutureStateInstall.String())).All(ctx)
}

func (workOrderResolver) LinksToRemove(ctx context.Context, obj *ent.WorkOrder) ([]*ent.Link, error) {
	return obj.QueryLinks().Where(link.FutureState(models.FutureStateRemove.String())).All(ctx)
}

func (workOrderResolver) Properties(ctx context.Context, obj *ent.WorkOrder) ([]*ent.Property, error) {
	props, err := obj.Edges.PropertiesOrErr()
	if ent.IsNotLoaded(err) {
		return obj.QueryProperties().All(ctx)
	}
	return props, err
}

func (workOrderResolver) CheckListCategories(ctx context.Context, obj *ent.WorkOrder) ([]*ent.CheckListCategory, error) {
	return obj.QueryCheckListCategories().All(ctx)
}

func (workOrderResolver) fileOfType(ctx context.Context, workOrder *ent.WorkOrder, typ file.Type) ([]*ent.File, error) {
	return workOrder.QueryFiles().Where(file.TypeEQ(typ)).All(ctx)
}

func (r workOrderResolver) Images(ctx context.Context, obj *ent.WorkOrder) ([]*ent.File, error) {
	return r.fileOfType(ctx, obj, file.TypeIMAGE)
}

func (r workOrderResolver) Files(ctx context.Context, obj *ent.WorkOrder) ([]*ent.File, error) {
	return r.fileOfType(ctx, obj, file.TypeFILE)
}

func (workOrderResolver) Hyperlinks(ctx context.Context, obj *ent.WorkOrder) ([]*ent.Hyperlink, error) {
	return obj.QueryHyperlinks().All(ctx)
}

func (workOrderResolver) Comments(ctx context.Context, obj *ent.WorkOrder) ([]*ent.Comment, error) {
	return obj.QueryComments().All(ctx)
}

func (workOrderResolver) Owner(ctx context.Context, obj *ent.WorkOrder) (*ent.User, error) {
	owner, err := obj.Edges.OwnerOrErr()
	if ent.IsNotLoaded(err) {
		owner, err = obj.QueryOwner().Only(ctx)
	}
	if err != nil {
		return nil, fmt.Errorf("querying owner: %w", err)
	}
	return owner, nil
}

func (workOrderResolver) AssignedTo(ctx context.Context, obj *ent.WorkOrder) (*ent.User, error) {
	assignee, err := obj.Edges.AssigneeOrErr()
	if ent.IsNotLoaded(err) {
		assignee, err = obj.QueryAssignee().Only(ctx)
	}
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("querying assignee: %w", err)
	}
	return assignee, nil
}

func (r mutationResolver) AddWorkOrder(
	ctx context.Context,
	input models.AddWorkOrderInput,
) (*ent.WorkOrder, error) {
	return r.internalAddWorkOrder(ctx, input, false)
}

func (r mutationResolver) convertToTemplatePropertyInputs(
	ctx context.Context,
	template *ent.WorkOrderTemplate,
	properties []*models.PropertyInput,
) ([]*models.PropertyInput, error) {
	client := r.ClientFrom(ctx).PropertyType
	inputs := make([]*models.PropertyInput, 0, len(properties))
	for _, p := range properties {
		name, err := client.Query().
			Where(propertytype.ID(p.PropertyTypeID)).
			Select(propertytype.FieldName).
			String(ctx)
		if err != nil {
			return nil, fmt.Errorf("cannot query property type name from id: %w", err)
		}
		id, err := template.
			QueryPropertyTypes().
			Where(propertytype.Name(name)).
			OnlyID(ctx)
		if err != nil {
			return nil, err
		}
		input := *p
		input.PropertyTypeID = id
		inputs = append(inputs, &input)
	}
	return inputs, nil
}

func (r mutationResolver) internalAddWorkOrder(
	ctx context.Context,
	input models.AddWorkOrderInput,
	skipMandatoryPropertiesCheck bool,
) (*ent.WorkOrder, error) {
	workOrderTemplate, err := r.addWorkOrderTemplate(ctx, input.WorkOrderTypeID)
	if err != nil {
		return nil, err
	}
	tPropInputs, err := r.convertToTemplatePropertyInputs(ctx, workOrderTemplate, input.Properties)
	if err != nil {
		return nil, fmt.Errorf("convert to template property inputs: %w", err)
	}
	propInput, err := r.validatedPropertyInputsFromTemplate(ctx, tPropInputs, workOrderTemplate.ID, models.PropertyEntityWorkOrder, skipMandatoryPropertiesCheck)
	if err != nil {
		return nil, fmt.Errorf("validating property for template : %w", err)
	}
	mutation := r.ClientFrom(ctx).
		WorkOrder.Create().
		SetName(input.Name).
		SetTypeID(input.WorkOrderTypeID).
		SetNillableStatus(input.Status).
		SetNillablePriority(input.Priority).
		SetTemplateID(workOrderTemplate.ID).
		SetNillableProjectID(input.ProjectID).
		SetNillableLocationID(input.LocationID).
		SetNillableDescription(input.Description).
		SetCreationDate(time.Now()).
		SetNillableIndex(input.Index).
		SetNillableAssigneeID(input.AssigneeID)
	if input.OwnerID != nil {
		mutation = mutation.SetOwnerID(*input.OwnerID)
	} else {
		v, ok := viewer.FromContext(ctx).(*viewer.UserViewer)
		if !ok {
			return nil, gqlerror.Errorf("could not be executed in automation")
		}
		mutation = mutation.SetOwner(v.User())
	}
	wo, err := mutation.Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "creating work order")
	}
	for _, clInput := range input.CheckListCategories {
		_, err := r.createOrUpdateCheckListCategory(ctx, clInput, wo.ID)
		if err != nil {
			return nil, errors.Wrap(err, "creating check list category")
		}
	}
	if _, err := r.AddProperties(propInput,
		resolverutil.AddPropertyArgs{
			Context:    ctx,
			EntSetter:  func(b *ent.PropertyCreate) { b.SetWorkOrderID(wo.ID) },
			IsTemplate: pointer.ToBool(true),
		},
	); err != nil {
		return nil, errors.Wrap(err, "creating work order properties")
	}
	return wo, nil
}

func (r mutationResolver) EditWorkOrder(
	ctx context.Context,
	input models.EditWorkOrderInput,
) (*ent.WorkOrder, error) {
	client := r.ClientFrom(ctx)
	wo, err := client.WorkOrder.Get(ctx, input.ID)
	if err != nil {
		return nil, errors.Wrap(err, "querying work order")
	}
	mutation := client.WorkOrder.
		UpdateOne(wo).
		SetName(input.Name).
		SetNillableDescription(input.Description).
		SetNillableIndex(input.Index).
		SetNillableStatus(input.Status).
		SetNillablePriority(input.Priority)

	if input.AssigneeID != nil {
		mutation.SetAssigneeID(*input.AssigneeID)
	} else {
		mutation.ClearAssignee()
	}
	if input.OwnerID != nil {
		mutation.SetOwnerID(*input.OwnerID)
	}
	if input.InstallDate != nil {
		mutation.SetInstallDate(*input.InstallDate)
	} else {
		mutation.ClearInstallDate()
	}
	if input.ProjectID != nil {
		mutation.SetProjectID(*input.ProjectID)
	} else {
		mutation.ClearProject()
	}
	if input.LocationID != nil {
		mutation.SetLocationID(*input.LocationID)
	} else {
		mutation.ClearLocation()
	}
	tPropInputs := input.Properties
	tmpl, err := wo.QueryTemplate().Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
		}
	} else {
		tPropInputs, err = r.convertToTemplatePropertyInputs(ctx, tmpl, tPropInputs)
		if err != nil {
			return nil, fmt.Errorf("convert to template property inputs: %w", err)
		}
	}
	for _, pin := range tPropInputs {
		err = r.updateProperty(ctx, wo.QueryProperties(), pin)
		if err != nil {
			return nil, fmt.Errorf("updating work order property value: %w", err)
		}
	}
	ids := make([]int, 0, len(input.CheckListCategories))
	for _, clInput := range input.CheckListCategories {
		cli, err := r.createOrUpdateCheckListCategory(ctx, clInput, wo.ID)
		if err != nil {
			return nil, err
		}
		ids = append(ids, cli.ID)
	}
	currentCL, err := wo.QueryCheckListCategories().IDs(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "querying checklist categories of work order %q", wo.ID)
	}
	addedCLIds, deletedCLIds := resolverutil.GetDifferenceBetweenSlices(currentCL, ids)
	mutation.
		RemoveCheckListCategoryIDs(deletedCLIds...).
		AddCheckListCategoryIDs(addedCLIds...)
	return mutation.Save(ctx)
}

func (r mutationResolver) updateProperty(
	ctx context.Context,
	query *ent.PropertyQuery,
	input *models.PropertyInput) error {
	propertyQuery := query.
		Where(property.HasTypeWith(propertytype.ID(input.PropertyTypeID)))
	if input.ID != nil {
		propertyQuery = propertyQuery.
			Where(property.ID(*input.ID))
	}
	existingProperty, err := propertyQuery.Only(ctx)
	if err != nil {
		if input.ID == nil {
			return errors.Wrapf(err, "querying property type %q", input.PropertyTypeID)
		}
		return errors.Wrapf(err, "querying property type %q and id %q", input.PropertyTypeID, *input.ID)
	}
	client := r.ClientFrom(ctx)
	typ, err := client.PropertyType.Get(ctx, input.PropertyTypeID)
	if err != nil {
		return errors.Wrapf(err, "querying property type %q", input.PropertyTypeID)
	}
	if typ.Editable && typ.IsInstanceProperty {
		updater := client.Property.UpdateOneID(existingProperty.ID)
		if r.updatePropValues(ctx, input, updater) != nil {
			return errors.Wrap(err, "saving work order property value update")
		}
	}
	return nil
}

func (r mutationResolver) createOrUpdateCheckListCategory(
	ctx context.Context,
	clInput *models.CheckListCategoryInput,
	workOrderID int) (*ent.CheckListCategory, error) {
	client := r.ClientFrom(ctx)
	cl := client.CheckListCategory
	var clc *ent.CheckListCategory
	var err error
	if clInput.ID == nil {
		clc, err = cl.Create().
			SetTitle(clInput.Title).
			SetNillableDescription(clInput.Description).
			SetWorkOrderID(workOrderID).
			Save(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "creating check list category")
		}
	} else {
		clc, err = cl.UpdateOneID(*clInput.ID).
			SetTitle(clInput.Title).
			SetNillableDescription(clInput.Description).
			Save(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "updating check list category")
		}
	}
	mutation := cl.UpdateOneID(clc.ID)
	addedCLIds, deletedCLIds, err := r.createOrUpdateCheckListItems(ctx, clc, clInput.CheckList)
	if err != nil {
		return nil, errors.Wrap(err, "updating check list category items")
	}
	return mutation.
		RemoveCheckListItemIDs(deletedCLIds...).
		AddCheckListItemIDs(addedCLIds...).
		Save(ctx)
}

func (r mutationResolver) createOrUpdateCheckListItems(
	ctx context.Context,
	clc *ent.CheckListCategory,
	inputs []*models.CheckListItemInput) ([]int, []int, error) {
	ids := make([]int, 0, len(inputs))
	for _, input := range inputs {
		cli, err := r.createOrUpdateCheckListItem(ctx, input, clc.ID)
		if err != nil {
			return nil, nil, err
		}
		if cli != nil {
			ids = append(ids, cli.ID)
		}
	}
	currentCLIds := clc.QueryCheckListItems().IDsX(ctx)
	addedCLIds, deletedCLIds := resolverutil.GetDifferenceBetweenSlices(currentCLIds, ids)
	return addedCLIds, deletedCLIds, nil
}

func (r mutationResolver) createOrUpdateCheckListItem(
	ctx context.Context,
	input *models.CheckListItemInput,
	checklistCategoryID int) (*ent.CheckListItem, error) {
	client := r.ClientFrom(ctx)
	cl := client.CheckListItem
	var cli *ent.CheckListItem
	var err error
	if input.ID == nil {
		cli, err = cl.Create().
			SetTitle(input.Title).
			SetType(input.Type.String()).
			SetNillableIndex(input.Index).
			SetNillableEnumValues(input.EnumValues).
			SetNillableHelpText(input.HelpText).
			SetNillableChecked(input.Checked).
			SetNillableStringVal(input.StringValue).
			SetNillableEnumSelectionModeValue(input.EnumSelectionMode).
			SetNillableSelectedEnumValues(input.SelectedEnumValues).
			SetNillableYesNoVal(convertYesNoResponseToYesNoVal(input.YesNoResponse)).
			SetCheckListCategoryID(checklistCategoryID).
			Save(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "creating check list item")
		}
	} else {
		cli, err = cl.UpdateOneID(*input.ID).
			SetTitle(input.Title).
			SetType(input.Type.String()).
			SetNillableIndex(input.Index).
			SetNillableEnumValues(input.EnumValues).
			SetNillableHelpText(input.HelpText).
			SetNillableChecked(input.Checked).
			SetNillableStringVal(input.StringValue).
			SetNillableEnumSelectionModeValue(input.EnumSelectionMode).
			SetNillableSelectedEnumValues(input.SelectedEnumValues).
			SetNillableYesNoVal(convertYesNoResponseToYesNoVal(input.YesNoResponse)).
			Save(ctx)
	}
	if err != nil {
		return nil, errors.Wrap(err, "updating check list item")
	}
	return r.createOrUpdateCheckListItemFiles(ctx, cli, input.Files)
}

func toIDSet(ids []int) map[int]bool {
	idSet := make(map[int]bool, len(ids))
	for _, id := range ids {
		idSet[id] = true
	}
	return idSet
}

func (r mutationResolver) deleteRemovedCheckListItemFiles(ctx context.Context, item *ent.CheckListItem, currentFileIDs []int, inputFileIDs []int) (*ent.CheckListItem, map[int]bool, error) {
	client := r.ClientFrom(ctx)
	_, deletedFileIDs := resolverutil.GetDifferenceBetweenSlices(currentFileIDs, inputFileIDs)
	deletedIDSet := toIDSet(deletedFileIDs)
	for _, fileID := range deletedFileIDs {
		err := client.File.DeleteOneID(fileID).Exec(ctx)
		if err != nil {
			return nil, nil, fmt.Errorf("deleting checklist file: file=%q: %w", fileID, err)
		}
	}
	item, err := client.CheckListItem.UpdateOne(item).RemoveFileIDs(deletedFileIDs...).Save(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("removing checklist files, item: %q: %w", item, err)
	}
	return item, deletedIDSet, nil
}

func (r mutationResolver) createAddedCheckListItemFiles(ctx context.Context, item *ent.CheckListItem, fileInputs []*models.FileInput) (*ent.CheckListItem, error) {
	for _, input := range fileInputs {
		if input.ID != nil {
			continue
		}
		_, err := r.createImage(
			ctx,
			&models.AddImageInput{
				ImgKey:   input.StoreKey,
				FileName: input.FileName,
				FileSize: func() int {
					if input.SizeInBytes != nil {
						return *input.SizeInBytes
					}
					return 0
				}(),
				Modified: time.Now(),
				ContentType: func() string {
					if input.MimeType != nil {
						return *input.MimeType
					}
					return "image/jpeg"
				}(),
				Annotation: input.Annotation,
			},
			func(create *ent.FileCreate) error {
				create.SetChecklistItem(item)
				return nil
			},
		)
		if err != nil {
			return nil, err
		}
	}
	return item, nil
}

func (r mutationResolver) createOrUpdateCheckListItemFiles(ctx context.Context, item *ent.CheckListItem, fileInputs []*models.FileInput) (*ent.CheckListItem, error) {
	client := r.ClientFrom(ctx)
	currentFileIDs, err := client.CheckListItem.QueryFiles(item).IDs(ctx)
	if err != nil {
		return nil, fmt.Errorf("querying checklist files, item=%q: %w", item.ID, err)
	}
	inputFileIDs := make([]int, 0, len(fileInputs))
	for _, fileInput := range fileInputs {
		if fileInput.ID == nil {
			continue
		}
		inputFileIDs = append(inputFileIDs, *fileInput.ID)
	}
	item, deletedIDSet, err := r.deleteRemovedCheckListItemFiles(ctx, item, currentFileIDs, inputFileIDs)
	if err != nil {
		return nil, err
	}
	item, err = r.createAddedCheckListItemFiles(ctx, item, fileInputs)
	if err != nil {
		return nil, err
	}
	for _, input := range fileInputs {
		if input.ID == nil {
			continue
		}
		if _, ok := deletedIDSet[*input.ID]; ok {
			continue
		}
		existingFile, err := client.File.Get(ctx, *input.ID)
		if err != nil {
			return nil, fmt.Errorf("querying file: file=%q: %w", *input.ID, err)
		}
		if existingFile.Name == input.FileName {
			continue
		}
		_, err = client.File.UpdateOne(existingFile).SetName(input.FileName).SetModifiedAt(time.Now()).Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("updating file name: file=%q: %w", existingFile.ID, err)
		}
	}
	return item, nil
}

func (r mutationResolver) addWorkOrderTypeCategoryDefinitions(ctx context.Context, input models.AddWorkOrderTypeInput, workOrderTypeID int) error {
	client := r.ClientFrom(ctx)
	for _, categoryInput := range input.CheckListCategories {
		checkListCategoryDefinition, err := client.CheckListCategoryDefinition.Create().
			SetTitle(categoryInput.Title).
			SetNillableDescription(categoryInput.Description).
			SetWorkOrderTypeID(workOrderTypeID).
			Save(ctx)
		if err != nil {
			return errors.Wrap(err, "creating check list category definition")
		}
		for _, clInput := range categoryInput.CheckList {
			if _, err = client.CheckListItemDefinition.Create().
				SetTitle(clInput.Title).
				SetType(clInput.Type.String()).
				SetNillableIndex(clInput.Index).
				SetNillableHelpText(clInput.HelpText).
				SetNillableEnumValues(clInput.EnumValues).
				SetNillableEnumSelectionModeValue((*checklistitemdefinition.EnumSelectionModeValue)(clInput.EnumSelectionMode)).
				SetCheckListCategoryDefinitionID(checkListCategoryDefinition.ID).
				Save(ctx); err != nil {
				return errors.Wrap(err, "creating check list item definition")
			}
		}
	}
	return nil
}

func (r mutationResolver) addWorkOrderTemplate(
	ctx context.Context,
	workOrderTypeID int,
) (*ent.WorkOrderTemplate, error) {
	client := r.ClientFrom(ctx)
	workOrderType, err := client.WorkOrderType.Query().
		Where(workordertype.ID(workOrderTypeID)).
		WithPropertyTypes().
		WithCheckListCategoryDefinitions().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("querying work order type: %w", err)
	}
	workOrderTemplate, err := client.WorkOrderTemplate.
		Create().
		SetName(workOrderType.Name).
		SetDescription(workOrderType.Description).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("creating work order template: %w", err)
	}
	for _, pt := range workOrderType.Edges.PropertyTypes {
		_, err = client.PropertyType.Create().
			SetName(pt.Name).
			SetType(pt.Type).
			SetNodeType(pt.NodeType).
			SetIndex(pt.Index).
			SetCategory(pt.Category).
			SetNillableStringVal(pt.StringVal).
			SetNillableIntVal(pt.IntVal).
			SetNillableBoolVal(pt.BoolVal).
			SetNillableFloatVal(pt.FloatVal).
			SetNillableLatitudeVal(pt.LatitudeVal).
			SetNillableLongitudeVal(pt.LongitudeVal).
			SetIsInstanceProperty(pt.IsInstanceProperty).
			SetNillableRangeFromVal(pt.RangeFromVal).
			SetNillableRangeToVal(pt.RangeToVal).
			SetEditable(pt.Editable).
			SetMandatory(pt.Mandatory).
			SetDeleted(pt.Deleted).
			SetWorkOrderTemplateID(workOrderTemplate.ID).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("creating property type: %w", err)
		}
	}
	for _, categoryInput := range workOrderType.Edges.CheckListCategoryDefinitions {
		cd, err := client.CheckListCategoryDefinition.Create().
			SetTitle(categoryInput.Title).
			SetDescription(categoryInput.Description).
			SetWorkOrderTemplateID(workOrderTemplate.ID).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("creating check list category definition: %w", err)
		}
		checkLists, err := categoryInput.QueryCheckListItemDefinitions().All(ctx)
		if err != nil {
			return nil, err
		}
		for _, checkList := range checkLists {
			var enumSelectionMode *checklistitemdefinition.EnumSelectionModeValue
			if checkList.EnumSelectionModeValue != "" {
				enumSelectionMode = &checkList.EnumSelectionModeValue
			}
			_, err := client.CheckListItemDefinition.Create().
				SetTitle(checkList.Title).
				SetType(checkList.Type).
				SetIndex(checkList.Index).
				SetNillableHelpText(checkList.HelpText).
				SetNillableEnumValues(checkList.EnumValues).
				SetNillableEnumSelectionModeValue(enumSelectionMode).
				SetCheckListCategoryDefinitionID(cd.ID).
				Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("creating check list item definition: %w", err)
			}
		}
	}
	return workOrderTemplate, err
}

func (r mutationResolver) deleteWorkOrderTemplate(ctx context.Context, id int) (int, error) {
	client, logger := r.ClientFrom(ctx), r.logger.For(ctx).With(zap.Int("id", id))
	pTypes, err := client.PropertyType.Query().
		Where(propertytype.HasWorkOrderTemplateWith(workordertemplate.ID(id))).
		All(ctx)
	if err != nil {
		logger.Error("cannot query properties of work order template", zap.Error(err))
		return id, fmt.Errorf("querying work order template property types: %w", err)
	}
	for _, pType := range pTypes {
		if err := client.PropertyType.DeleteOne(pType).
			Exec(ctx); err != nil {
			logger.Error("cannot delete property of work order template", zap.Error(err))
			return id, fmt.Errorf("deleting work order template property type: %w", err)
		}
	}
	switch err := client.WorkOrderTemplate.DeleteOneID(id).Exec(ctx); err.(type) {
	case nil:
		logger.Info("deleted work order template")
		return id, nil
	case *ent.NotFoundError:
		err := gqlerror.Errorf("work order template not found")
		logger.Error(err.Message)
		return id, err
	default:
		logger.Error("cannot delete work order template", zap.Error(err))
		return id, fmt.Errorf("deleting work order template: %w", err)
	}
}

func (r mutationResolver) AddWorkOrderType(
	ctx context.Context, input models.AddWorkOrderTypeInput) (*ent.WorkOrderType, error) {
	client := r.ClientFrom(ctx)
	typ, err := client.WorkOrderType.
		Create().
		SetName(input.Name).
		SetNillableDescription(input.Description).
		Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return nil, gqlerror.Errorf("A work order type with the name %v already exists", input.Name)
		}
		return nil, errors.Wrap(err, "creating work order type")
	}
	if err := r.AddPropertyTypes(ctx, func(ptc *ent.PropertyTypeCreate) {
		ptc.SetWorkOrderTypeID(typ.ID)
	}, input.Properties...); err != nil {
		return nil, err
	}
	err = r.addWorkOrderTypeCategoryDefinitions(ctx, input, typ.ID)
	if err != nil {
		return nil, errors.Wrap(err, "creating checklist category definitions")
	}
	return typ, nil
}

func (r mutationResolver) EditWorkOrderType(
	ctx context.Context, input models.EditWorkOrderTypeInput,
) (*ent.WorkOrderType, error) {
	client := r.ClientFrom(ctx)
	wot, err := client.WorkOrderType.Get(ctx, input.ID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, gqlerror.Errorf("A work order template with id=%q does not exist", input.ID)
		}
		if ent.IsConstraintError(err) {
			return nil, gqlerror.Errorf("A work order template with the name %v already exists", input.Name)
		}
		return nil, errors.Wrapf(err, "updating work order template: id=%q", input.ID)
	}
	for _, p := range input.Properties {
		if p.ID == nil {
			if err := r.AddPropertyTypes(ctx, func(b *ent.PropertyTypeCreate) { b.SetWorkOrderTypeID(input.ID) }, p); err != nil {
				return nil, err
			}
		} else if err := r.updatePropType(ctx, p); err != nil {
			return nil, err
		}
	}
	mutation := client.WorkOrderType.
		UpdateOneID(input.ID).
		SetName(input.Name).
		SetNillableDescription(input.Description)
	currentCategories, err := wot.QueryCheckListCategoryDefinitions().IDs(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "querying checklist category definitions: id=%q", wot.ID)
	}
	ids := make([]int, 0, len(input.CheckListCategories))
	for _, categoryInput := range input.CheckListCategories {
		category, err := r.createOrUpdateCheckListCategoryDefinition(ctx, categoryInput, wot.ID)
		if err != nil {
			return nil, err
		}
		ids = append(ids, category.ID)
	}
	_, deletedCategoryIds := resolverutil.GetDifferenceBetweenSlices(currentCategories, ids)
	mutation = mutation.RemoveCheckListCategoryDefinitionIDs(deletedCategoryIds...)
	return mutation.Save(ctx)
}

func (r mutationResolver) createOrUpdateCheckListCategoryDefinition(
	ctx context.Context,
	categoryInput *models.CheckListCategoryDefinitionInput,
	wotID int) (*ent.CheckListCategoryDefinition, error) {
	client := r.ClientFrom(ctx)
	cl := client.CheckListCategoryDefinition
	var category *ent.CheckListCategoryDefinition
	var err error
	if categoryInput.ID == nil {
		category, err = cl.Create().
			SetTitle(categoryInput.Title).
			SetNillableDescription(categoryInput.Description).
			SetWorkOrderTypeID(wotID).
			Save(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "creating check list category definition")
		}
	} else {
		category, err = cl.UpdateOneID(*categoryInput.ID).
			SetTitle(categoryInput.Title).
			SetNillableDescription(categoryInput.Description).
			SetWorkOrderTypeID(wotID).
			Save(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "updating check list category definition")
		}
	}
	currentCL, err := category.QueryCheckListItemDefinitions().IDs(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "querying checklist item definitions: id=%q", category.ID)
	}
	ids := make([]int, 0, len(categoryInput.CheckList))
	for _, clInput := range categoryInput.CheckList {
		cli, err := r.createOrUpdateCheckListDefinition(ctx, clInput, category.ID)
		if err != nil {
			return nil, err
		}
		ids = append(ids, cli.ID)
	}
	_, deletedCLIds := resolverutil.GetDifferenceBetweenSlices(currentCL, ids)
	return category.Update().
		RemoveCheckListItemDefinitionIDs(deletedCLIds...).
		Save(ctx)
}

func (r mutationResolver) createOrUpdateCheckListDefinition(
	ctx context.Context,
	clInput *models.CheckListDefinitionInput,
	categoryID int) (*ent.CheckListItemDefinition, error) {
	client := r.ClientFrom(ctx)
	cl := client.CheckListItemDefinition
	if clInput.ID == nil {
		cli, err := cl.Create().
			SetTitle(clInput.Title).
			SetType(clInput.Type.String()).
			SetNillableIndex(clInput.Index).
			SetNillableEnumValues(clInput.EnumValues).
			SetNillableEnumSelectionModeValue((*checklistitemdefinition.EnumSelectionModeValue)(clInput.EnumSelectionMode)).
			SetNillableHelpText(clInput.HelpText).
			SetCheckListCategoryDefinitionID(categoryID).
			Save(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "creating check list definition")
		}
		return cli, nil
	}

	cli, err := cl.UpdateOneID(*clInput.ID).
		SetTitle(clInput.Title).
		SetType(clInput.Type.String()).
		SetNillableIndex(clInput.Index).
		SetNillableEnumValues(clInput.EnumValues).
		SetNillableEnumSelectionModeValue((*checklistitemdefinition.EnumSelectionModeValue)(clInput.EnumSelectionMode)).
		SetNillableHelpText(clInput.HelpText).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "updating check list definition")
	}
	return cli, nil
}

func (r mutationResolver) RemoveWorkOrderType(ctx context.Context, id int) (int, error) {
	client, logger := r.ClientFrom(ctx), r.logger.For(ctx).With(zap.Int("id", id))
	switch count, err := client.WorkOrderType.Query().
		Where(workordertype.ID(id)).
		QueryWorkOrders().
		Count(ctx); {
	case err != nil:
		logger.Error("cannot query work order count of type", zap.Error(err))
		return id, fmt.Errorf("querying work orders for type: %w", err)
	case count > 0:
		logger.Warn("work order type has existing work orders", zap.Int("count", count))
		return id, gqlerror.Errorf("cannot delete work order type with %d existing work orders", count)
	}
	pTypes, err := client.PropertyType.Query().
		Where(propertytype.HasWorkOrderTypeWith(workordertype.ID(id))).
		All(ctx)
	if err != nil {
		logger.Error("cannot query properties of work order type", zap.Error(err))
		return id, fmt.Errorf("querying work order property types: %w", err)
	}
	for _, pType := range pTypes {
		if err := client.PropertyType.DeleteOne(pType).
			Exec(ctx); err != nil {
			logger.Error("cannot delete property of work order type", zap.Error(err))
			return id, fmt.Errorf("deleting work order property type: %w", err)
		}
	}
	switch err := client.WorkOrderType.DeleteOneID(id).Exec(ctx); err.(type) {
	case nil:
		logger.Info("deleted work order type")
		return id, nil
	case *ent.NotFoundError:
		err := gqlerror.Errorf("work order type not found")
		logger.Error(err.Message)
		return id, err
	default:
		logger.Error("cannot delete work order type", zap.Error(err))
		return id, fmt.Errorf("deleting work order type: %w", err)
	}
}

func convertYesNoResponseToYesNoVal(response *models.YesNoResponse) *checklistitem.YesNoVal {
	if response == nil {
		return nil
	}
	yesNoVal := checklistitem.YesNoVal(*response)
	err := checklistitem.YesNoValValidator(yesNoVal)
	if err != nil {
		return nil
	}
	return &yesNoVal
}

func (r mutationResolver) TechnicianWorkOrderUploadData(ctx context.Context, input models.TechnicianWorkOrderUploadInput) (*ent.WorkOrder, error) {
	client := r.ClientFrom(ctx)
	wo, err := client.WorkOrder.Query().Where(workorder.ID(input.WorkOrderID)).WithAssignee().Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("querying work order %q: err %w", input.WorkOrderID, err)
	}
	assignee, err := wo.Edges.AssigneeOrErr()
	if err != nil || assignee == nil {
		return nil, fmt.Errorf(
			"work order %q is not assigned to a technician: err %w",
			input.WorkOrderID,
			err,
		)
	}
	v, ok := viewer.FromContext(ctx).(*viewer.UserViewer)
	if !ok {
		return nil, gqlerror.Errorf("could not be executed in automation")
	}
	if assignee.Email != v.User().Email {
		return nil, fmt.Errorf(
			"mismatch between work order %q assginee %q and technician %q: err %w",
			input.WorkOrderID,
			wo.Edges.Assignee.Email,
			v.User().Email,
			err,
		)
	}
	for _, clInput := range input.Checklist {
		checkListItem, err := client.CheckListItem.
			UpdateOneID(clInput.ID).
			SetNillableChecked(clInput.Checked).
			SetNillableStringVal(clInput.StringValue).
			SetNillableSelectedEnumValues(clInput.SelectedEnumValues).
			SetNillableYesNoVal(convertYesNoResponseToYesNoVal(clInput.YesNoResponse)).
			Save(ctx)
		if clInput.WifiData != nil && len(clInput.WifiData) > 0 {
			_, err := r.CreateWiFiScans(ctx, clInput.WifiData, ScanParentIDs{checklistItemID: &clInput.ID})
			if err != nil {
				return nil, fmt.Errorf("creating wifi scans, item %q: err %w", clInput.ID, err)
			}
		}
		if clInput.CellData != nil && len(clInput.CellData) > 0 {
			_, err := r.CreateCellScans(ctx, clInput.CellData, ScanParentIDs{checklistItemID: &clInput.ID})
			if err != nil {
				return nil, fmt.Errorf("creating cell scans, item %q: err %w", clInput.ID, err)
			}
		}
		if clInput.FilesData != nil && len(clInput.FilesData) > 0 {
			_, err := r.createOrUpdateCheckListItemFiles(ctx, checkListItem, clInput.FilesData)
			if err != nil {
				return nil, fmt.Errorf("creating and saving images while uploading a work order: %q: err %w", input.WorkOrderID, err)
			}
		}
		if err != nil {
			return nil, fmt.Errorf("updating checklist item %q: err %w", clInput.ID, err)
		}
	}
	if _, err = r.AddComment(ctx, models.CommentInput{
		EntityType: models.CommentEntityWorkOrder,
		ID:         input.WorkOrderID,
		Text:       v.User().Email + " uploaded data",
	}); err != nil {
		return nil, fmt.Errorf("adding technician uploaded data comment: %w", err)
	}
	return client.WorkOrder.
		Query().
		Where(workorder.ID(input.WorkOrderID)).
		WithComments().
		WithCheckListCategories().
		Only(ctx)
}
