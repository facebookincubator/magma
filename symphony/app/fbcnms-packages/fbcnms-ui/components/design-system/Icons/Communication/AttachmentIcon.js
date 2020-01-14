/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import type {SvgIconStyleProps} from '../SvgIcon';

import React from 'react';
import SvgIcon from '../SvgIcon';

const AttachmentIcon = (props: SvgIconStyleProps) => (
  <SvgIcon {...props}>
    <g transform="translate(3,4)">
      <path
        d="M1.613 14.399a5.53 5.53 0 007.784 0l7.43-7.375a3.951 3.951 0 000-5.62 4.023 4.023 0 00-5.66 0L5.15 7.375a2.471 2.471 0 000 3.513c.977.969 2.562.969 3.539 0l5.307-5.268-1.416-1.405-5.37 5.331c-.39.387-1.097-.316-.708-.702l6.078-6.034a2.018 2.018 0 012.83 0 1.982 1.982 0 010 2.81l-7.43 7.375a3.525 3.525 0 01-4.952 0 3.462 3.462 0 010-4.917L9.75 1.405 8.336 0 1.613 6.672a5.432 5.432 0 000 7.727z"
        fillRule="evenodd"
      />
    </g>
  </SvgIcon>
);

export default AttachmentIcon;
