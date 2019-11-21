/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import type {
  AddWorkOrderMutationResponse,
  AddWorkOrderMutationVariables,
} from '../../mutations/__generated__/AddWorkOrderMutation.graphql';
import type {ContextRouter} from 'react-router-dom';
import type {MutationCallbacks} from '../../mutations/MutationCallbacks.js';
import type {WithSnackbarProps} from 'notistack';
import type {WithStyles} from '@material-ui/core';
import type {WorkOrder} from '../../common/WorkOrder';

import AddWorkOrderMutation from '../../mutations/AddWorkOrderMutation';
import Breadcrumbs from '@fbcnms/ui/components/Breadcrumbs';
import Button from '@fbcnms/ui/components/design-system/Button';
import CircularProgress from '@material-ui/core/CircularProgress';
import ExpandingPanel from '@fbcnms/ui/components/ExpandingPanel';
import FormField from '@fbcnms/ui/components/design-system/FormField/FormField';
import Grid from '@material-ui/core/Grid';
import LocationTypeahead from '../typeahead/LocationTypeahead';
import MenuItem from '@material-ui/core/MenuItem';
import NameDescriptionSection from '@fbcnms/ui/components/NameDescriptionSection';
import ProjectTypeahead from '../typeahead/ProjectTypeahead';
import PropertyValueInput from '../form/PropertyValueInput';
import React from 'react';
import RelayEnvironment from '../../common/RelayEnvironment.js';
import SnackbarItem from '@fbcnms/ui/components/SnackbarItem';
import TextField from '@material-ui/core/TextField';
import UserTypeahead from '../typeahead/UserTypeahead';
import nullthrows from '@fbcnms/util/nullthrows';
import update from 'immutability-helper';
import {LogEvents, ServerLogger} from '../../common/LoggingUtils';
import {fetchQuery, graphql} from 'relay-runtime';
import {getInitialPropertyFromType} from '../../common/PropertyType';
import {priorityValues, statusValues} from '../../common/WorkOrder';
import {sortPropertiesByIndex, toPropertyInput} from '../../common/Property';
import {withRouter} from 'react-router-dom';
import {withSnackbar} from 'notistack';
import {withStyles} from '@material-ui/core/styles';

type Props = {
  workOrderTypeId: ?string,
} & WithStyles<typeof styles> &
  ContextRouter &
  WithSnackbarProps;

const styles = theme => ({
  root: {
    height: '100%',
    display: 'flex',
    flexDirection: 'column',
    padding: '40px 32px',
  },
  contentRoot: {
    display: 'flex',
    flexDirection: 'column',
    position: 'relative',
    flexGrow: 1,
    overflow: 'auto',
  },
  cards: {
    flexGrow: 1,
    overflow: 'hidden',
    overflowY: 'auto',
  },
  card: {
    display: 'flex',
    flexDirection: 'column',
  },
  input: {
    width: '250px',
    paddingBottom: '24px',
  },
  gridInput: {
    display: 'inline-flex',
  },
  nameHeader: {
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
    marginBottom: '24px',
  },
  breadcrumbs: {
    flexGrow: 1,
  },
  separator: {
    borderBottom: `1px solid ${theme.palette.grey[100]}`,
    margin: '0 0 24px -24px',
    paddingBottom: '24px',
    width: 'calc(100% + 48px)',
  },
  separator: {
    borderBottom: `1px solid ${theme.palette.grey[100]}`,
    margin: '0 0 24px -24px',
    paddingBottom: '24px',
    width: 'calc(100% + 48px)',
  },
  dense: {
    paddingTop: '9px',
    paddingBottom: '9px',
    height: '14px',
  },
  cancelButton: {
    marginRight: '8px',
  },
});

type State = {
  workOrder: ?WorkOrder,
  locationId: ?string,
};

const addWorkOrderCard__workOrderTypeQuery = graphql`
  query AddWorkOrderCard__workOrderTypeQuery($workOrderTypeId: ID!) {
    workOrderType(id: $workOrderTypeId) {
      id
      name
      description
      propertyTypes {
        id
        name
        type
        index
        stringValue
        intValue
        booleanValue
        floatValue
        latitudeValue
        longitudeValue
        rangeFromValue
        rangeToValue
        isEditable
        isInstanceProperty
      }
    }
  }
`;

class AddWorkOrderCard extends React.Component<Props, State> {
  state = {
    locationId: null,
    workOrder: null,
  };
  componentDidMount() {
    this.getWorkOrder().then(workOrder => {
      this.setState({
        workOrder,
      });
    });
  }
  render() {
    const {classes} = this.props;
    const {workOrder} = this.state;
    if (!workOrder) {
      return (
        <div className={classes.root}>
          <CircularProgress />
        </div>
      );
    }
    return (
      <div className={classes.root}>
        <div className={classes.nameHeader}>
          <Breadcrumbs
            className={classes.breadcrumbs}
            breadcrumbs={[
              {
                id: 'workOrders',
                name: 'WorkOrders',
                onClick: () => this.navigateToMainPage(),
              },
              {
                id: `new_workOrder_` + Date.now(),
                name: 'New WorkOrder',
              },
            ]}
            size="large"
          />
          <Button
            className={classes.cancelButton}
            skin="regular"
            onClick={this.navigateToMainPage}>
            Cancel
          </Button>
          <Button onClick={this._saveWorkOrder}>Save</Button>
        </div>
        <div className={classes.contentRoot}>
          <div className={classes.cards}>
            <Grid container spacing={2}>
              <Grid item xs={8} sm={8} lg={8} xl={8}>
                <ExpandingPanel title="Details">
                  <NameDescriptionSection
                    name={workOrder.name}
                    description={workOrder.description}
                    onNameChange={value =>
                      this._setWorkOrderDetail('name', value)
                    }
                    onDescriptionChange={value =>
                      this._setWorkOrderDetail('description', value)
                    }
                  />
                  <div className={classes.separator} />
                  <Grid container spacing={2}>
                    <Grid item xs={12} sm={6} lg={4} xl={4}>
                      <FormField label="Project">
                        <ProjectTypeahead
                          className={classes.gridInput}
                          margin="dense"
                          onProjectSelection={project =>
                            this._setWorkOrderDetail('projectId', project?.id)
                          }
                        />
                      </FormField>
                    </Grid>
                    {workOrder.workOrderType && (
                      <Grid item xs={12} sm={6} lg={4} xl={4}>
                        <FormField label="Type">
                          <TextField
                            disabled
                            variant="outlined"
                            margin="dense"
                            className={classes.gridInput}
                            value={workOrder.workOrderType.name}
                          />
                        </FormField>
                      </Grid>
                    )}
                    <Grid item xs={12} sm={6} lg={4} xl={4}>
                      <FormField label="Priority">
                        <TextField
                          select
                          className={classes.gridInput}
                          variant="outlined"
                          value={workOrder.priority}
                          InputProps={{
                            classes: {
                              input: classes.dense,
                            },
                          }}
                          onChange={event => {
                            this._setWorkOrderDetail(
                              'priority',
                              event.target.value,
                            );
                          }}>
                          {priorityValues.map(option => (
                            <MenuItem key={option.value} value={option.value}>
                              {option.label}
                            </MenuItem>
                          ))}
                        </TextField>
                      </FormField>
                    </Grid>
                    <Grid item xs={12} sm={6} lg={4} xl={4}>
                      <FormField label="Status">
                        <TextField
                          select
                          className={classes.gridInput}
                          variant="outlined"
                          value={workOrder.status}
                          InputProps={{
                            classes: {
                              input: classes.dense,
                            },
                          }}
                          onChange={event => {
                            this._setWorkOrderDetail(
                              'status',
                              event.target.value,
                            );
                          }}>
                          {statusValues.map(option => (
                            <MenuItem key={option.value} value={option.value}>
                              {option.label}
                            </MenuItem>
                          ))}
                        </TextField>
                      </FormField>
                    </Grid>
                    <Grid item xs={12} sm={6} lg={4} xl={4}>
                      <FormField label="Location">
                        <LocationTypeahead
                          headline={null}
                          className={classes.gridInput}
                          margin="dense"
                          onLocationSelection={location =>
                            this._setWorkOrderDetail(
                              'locationId',
                              location?.id ?? null,
                            )
                          }
                        />
                      </FormField>
                    </Grid>
                    {workOrder.properties.map((property, index) => (
                      <Grid key={property.id} item xs={12} sm={6} lg={4} xl={4}>
                        <PropertyValueInput
                          required={false}
                          disabled={false}
                          label={property.propertyType.name}
                          className={classes.gridInput}
                          margin="dense"
                          inputType="Property"
                          property={property}
                          headlineVariant="form"
                          fullWidth={true}
                          onChange={this._propertyChangedHandler(index)}
                        />
                      </Grid>
                    ))}
                  </Grid>
                </ExpandingPanel>
              </Grid>
              <Grid item xs={4} sm={4} lg={4} xl={4}>
                <ExpandingPanel title="Team">
                  <UserTypeahead
                    className={classes.input}
                    headline="Assignee"
                    onUserSelection={user =>
                      this._setWorkOrderDetail('assignee', user)
                    }
                    margin="dense"
                  />
                </ExpandingPanel>
              </Grid>
            </Grid>
          </div>
        </div>
      </div>
    );
  }

  async getWorkOrder(): Promise<WorkOrder> {
    const response = await fetchQuery(
      RelayEnvironment,
      addWorkOrderCard__workOrderTypeQuery,
      {
        workOrderTypeId: this.props.workOrderTypeId,
      },
    );
    const workOrderType = response.workOrderType;
    const initialProps = (workOrderType.propertyTypes || [])
      .map(propType => getInitialPropertyFromType(propType))
      .sort(sortPropertiesByIndex);
    return {
      id: 'workOrder@tmp',
      workOrderType: workOrderType,
      workOrderTypeId: workOrderType.id,
      name: workOrderType.name,
      description: workOrderType.description,
      locationId: null,
      location: null,
      properties: initialProps,
      workOrders: [],
      ownerName: '',
      creationDate: '',
      installDate: '',
      status: 'PENDING',
      priority: 'NONE',
      equipmentToAdd: [],
      equipmentToRemove: [],
      linksToAdd: [],
      linksToRemove: [],
      files: [],
      images: [],
      assignee: '',
      projectId: null,
    };
  }

  _saveWorkOrder = () => {
    const {
      name,
      description,
      locationId,
      projectId,
      assignee,
      status,
      priority,
      properties,
    } = nullthrows(this.state.workOrder);
    const workOrderTypeId = nullthrows(this.state.workOrder?.workOrderTypeId);
    const variables: AddWorkOrderMutationVariables = {
      input: {
        name,
        description,
        locationId,
        workOrderTypeId,
        assignee,
        projectId,
        status,
        priority,
        properties: toPropertyInput(properties),
      },
    };

    const callbacks: MutationCallbacks<AddWorkOrderMutationResponse> = {
      onCompleted: (response, errors) => {
        if (errors && errors[0]) {
          this._enqueueError(errors[0].message);
        } else {
          // navigate to main page
          this.props.history.push(this.props.match.url);
        }
      },
      onError: () => {
        this._enqueueError('Error saving work order');
      },
    };
    ServerLogger.info(LogEvents.SAVE_PROJECT_BUTTON_CLICKED, {
      source: 'workOrder_details',
    });
    AddWorkOrderMutation(variables, callbacks);
  };

  _enqueueError = (message: string) => {
    this.props.enqueueSnackbar(message, {
      children: key => (
        <SnackbarItem id={key} message={message} variant="error" />
      ),
    });
  };

  _setWorkOrderDetail = (
    key:
      | 'name'
      | 'description'
      | 'assignee'
      | 'projectId'
      | 'locationId'
      | 'priority'
      | 'status',
    value,
  ) => {
    this.setState(prevState => {
      return {
        // $FlowFixMe Set state for each field
        workOrder: update(prevState.workOrder, {[key]: {$set: value}}),
      };
    });
  };

  _propertyChangedHandler = index => property =>
    this.setState(prevState => {
      return {
        workOrder: update(prevState.workOrder, {
          properties: {[index]: {$set: property}},
        }),
      };
    });

  navigateToMainPage = () => {
    ServerLogger.info(LogEvents.WORK_ORDERS_SEARCH_NAV_CLICKED, {
      source: 'work_order_details',
    });
    const {match} = this.props;
    this.props.history.push(match.url);
  };
}

export default withSnackbar(withRouter(withStyles(styles)(AddWorkOrderCard)));
