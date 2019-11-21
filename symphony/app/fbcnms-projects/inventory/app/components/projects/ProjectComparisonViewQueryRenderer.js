/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import InventoryQueryRenderer from '../InventoryQueryRenderer';
import ListAltIcon from '@material-ui/icons/ListAlt';
import MapButtonGroup from '@fbcnms/ui/components/map/MapButtonGroup';
import MapIcon from '@material-ui/icons/Map';
import ProjectsMap from './ProjectsMap';
import ProjectsTableView from './ProjectsTableView';
import React from 'react';
import SearchIcon from '@material-ui/icons/Search';
import Text from '@fbcnms/ui/components/design-system/Text';
import {graphql} from 'relay-runtime';
import {makeStyles} from '@material-ui/styles';
import {useState} from 'react';

import type {FilterValue} from '../comparison_view/ComparisonViewTypes';

const useStyles = makeStyles(theme => ({
  noResultsRoot: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'center',
    marginTop: '100px',
  },
  noResultsLabel: {
    color: theme.palette.grey[600],
  },
  searchIcon: {
    color: theme.palette.grey[600],
    marginBottom: '6px',
    fontSize: '36px',
  },
  bar: {
    borderBottom: '2px solid #f0f0f0',
  },
  groupButtons: {
    display: 'flex',
    justifyContent: 'flex-end',
  },
  projectsTable: {
    padding: '24px',
  },
}));

type Props = {
  limit?: number,
  filters: Array<FilterValue>,
  displayMode: ?'map' | 'table',
  onProjectSelected: (projectID: string) => void,
};

const projectSearchQuery = graphql`
  query ProjectComparisonViewQueryRendererSearchQuery(
    $limit: Int
    $filters: [ProjectFilterInput!]!
  ) {
    projectSearch(limit: $limit, filters: $filters) {
      ...ProjectsTableView_projects
      ...ProjectsMap_projects
    }
  }
`;

const ProjectComparisonViewQueryRenderer = (props: Props) => {
  const classes = useStyles();
  const {filters, limit, onProjectSelected} = props;
  const [resultsDisplayMode, setResultsDisplayMode] = useState('table');

  return (
    <InventoryQueryRenderer
      query={projectSearchQuery}
      variables={{
        limit: limit,
        filters: filters.map(f => ({
          filterType: f.name.toUpperCase(),
          operator: f.operator.toUpperCase(),
          stringValue: f.stringValue,
          propertyValue: f.propertyValue,
          idSet: f.idSet,
        })),
      }}
      render={props => {
        const {projectSearch} = props;
        if (!projectSearch || projectSearch.length === 0) {
          return (
            <div className={classes.noResultsRoot}>
              <SearchIcon className={classes.searchIcon} />
              <Text variant="h6" className={classes.noResultsLabel}>
                No results found
              </Text>
            </div>
          );
        }
        return (
          <>
            <div className={classes.bar}>
              <div className={classes.groupButtons}>
                <MapButtonGroup
                  onIconClicked={id => {
                    setResultsDisplayMode(id === 'table' ? 'table' : 'map');
                  }}
                  buttons={[
                    {item: <ListAltIcon />, id: 'table'},
                    {item: <MapIcon />, id: 'map'},
                  ]}
                />
              </div>
            </div>
            {resultsDisplayMode === 'map' ? (
              <ProjectsMap projects={projectSearch} />
            ) : (
              <ProjectsTableView
                className={classes.projectsTable}
                projects={projectSearch}
                onProjectSelected={onProjectSelected}
              />
            )}
          </>
        );
      }}
    />
  );
};

export default ProjectComparisonViewQueryRenderer;
