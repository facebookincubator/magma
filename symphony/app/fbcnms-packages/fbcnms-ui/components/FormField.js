/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import type {WithStyles} from '@material-ui/core';

import * as React from 'react';
import Text from './design-system/Text';
import Typography from '@material-ui/core/Typography';

import {withStyles} from '@material-ui/core/styles';

const styles = theme => ({
  root: {
    display: 'flex',
    whiteSpace: 'nowrap',
  },
  labelName: {
    color: theme.palette.grey.A700,
    fontWeight: 500,
    marginRight: '4px',
  },
  value: {
    textOverflow: 'ellipsis',
    overflowWrap: 'break-word',
    whiteSpace: 'nowrap',
    overflow: 'hidden',
  },
});

type Props = WithStyles<typeof styles> & {
  label: string,
  value?: ?React.Node,
};

class FormField extends React.Component<Props> {
  render() {
    const {classes, label, value} = this.props;
    return (
      <div className={classes.root}>
        <Text className={classes.labelName} variant="body2">
          {label}:
        </Text>
        {typeof value === 'string' || typeof value === 'number' ? (
          <Typography
            className={classes.value}
            variant="body2"
            color="secondary"
            title={value}>
            {value}
          </Typography>
        ) : (
          value
        )}
      </div>
    );
  }
}

export default withStyles(styles)(FormField);
