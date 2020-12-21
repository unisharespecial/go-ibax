/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package model

// ProgressComplete status of installation progress
const ProgressComplete = "complete"

// Install is model
type Install struct {
	Progress string `gorm:"not null;size:10"`
}