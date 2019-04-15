/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

// Package storage contains common definitions to be used across service
// storage interfaces
package storage

import "fmt"

type TypeAndKey struct {
	Type string
	Key  string
}

func (tk TypeAndKey) String() string {
	return fmt.Sprintf("%s-%s", tk.Type, tk.Key)
}
