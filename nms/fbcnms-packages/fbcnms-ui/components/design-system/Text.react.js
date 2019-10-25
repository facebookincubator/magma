/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import * as React from 'react';
import classNames from 'classnames';
import {makeStyles} from '@material-ui/styles';

const useStyles = makeStyles(({symphony}) => ({
  h1: symphony.typography.h1,
  h2: symphony.typography.h2,
  h3: symphony.typography.h3,
  h4: symphony.typography.h4,
  h5: symphony.typography.h5,
  h6: symphony.typography.h6,
  subtitle1: symphony.typography.subtitle1,
  subtitle2: symphony.typography.subtitle2,
  subtitle3: symphony.typography.subtitle3,
  body1: symphony.typography.body1,
  body2: symphony.typography.body2,
  caption: symphony.typography.caption,
  overline: symphony.typography.overline,
  lightColor: {
    color: symphony.palette.white,
  },
  regularColor: {
    color: symphony.palette.D900,
  },
  primaryColor: {
    color: symphony.palette.primary,
  },
  errorColor: {
    color: symphony.palette.R600,
  },
  lightWeight: {
    fontWeight: 300,
  },
  regularWeight: {
    fontWeight: 400,
  },
  mediumWeight: {
    fontWeight: 500,
  },
  boldWeight: {
    fontWeight: 600,
  },
}));

type Props = {
  children: React.Node,
  variant?:
    | 'h1'
    | 'h2'
    | 'h3'
    | 'h4'
    | 'h5'
    | 'h6'
    | 'subtitle1'
    | 'subtitle2'
    | 'subtitle3'
    | 'body1'
    | 'body2'
    | 'caption'
    | 'overline',
  className?: string,
  weight?: 'inherit' | 'light' | 'regular' | 'medium' | 'bold',
  color?: 'light' | 'regular' | 'primary' | 'error',
};

const Text = (props: Props) => {
  const {children, variant, className, color, weight} = props;
  const classes = useStyles();
  return (
    <span
      className={classNames(
        classes[variant],
        classes[`${color ?? 'regular'}Color`],
        classes[`${weight ? weight : 'inhreit'}Weight`],
        className,
      )}>
      {children}
    </span>
  );
};

Text.defaultProps = {
  variant: 'body1',
  color: 'regular',
  weight: 'inherit',
};

export default Text;
