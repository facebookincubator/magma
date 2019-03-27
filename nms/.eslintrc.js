/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @format
 */

module.exports.extends = ['eslint-config-fbcnms'];
module.exports.overrides = [
  {
    files: ['.eslintrc.js'],
    rules: {
      quotes: ['warn', 'single'],
    },
  },
  {
    env: {
      jest: true,
      node: true,
    },
    files: ['**/__mocks__/**/*.js', '**/__tests__/**/*.js'],
  },
  {
    env: {
      node: true,
    },
    files: [
      '.eslintrc.js',
      'babel.config.js',
      'jest.config.js',
      'jest.*.config.js',
      'fbcnms-packages/eslint-config-fbcnms/**/*.js',
      'fbcnms-packages/fbcnms-auth/**/*.js',
      'fbcnms-packages/fbcnms-babel-register/**/*.js',
      'fbcnms-packages/fbcnms-express-middleware/**/*.js',
      'fbcnms-packages/fbcnms-logging/**/*.js',
      'fbcnms-packages/fbcnms-sequelize-models/**/*.js',
      'fbcnms-packages/fbcnms-ui/stories/**/*.js',
      'fbcnms-packages/fbcnms-util/**/*.js',
      'fbcnms-packages/fbcnms-webpack-config/**/*.js',
      'fbcnms-projects/*/config/webpack.*.js',
      'fbcnms-projects/*/scripts/**/*.js',
      'fbcnms-projects/*/server/**/*.js',
    ],
    rules: {
      'no-console': 'off',
    },
  },
  {
    files: ['**/tgnms/**/*.js'],
    rules: {
      // tgnms doesn't want this because there's too many errors
      'flowtype/require-valid-file-annotation': 'off',
    },
  },
];
module.exports.settings = {
  react: {
    version: 'detect',
  },
};
