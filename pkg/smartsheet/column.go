/*
 * Copyright 2020 wfleming@grumpysysadm.com
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package smartsheet

type Column struct {
	Id               int              `json:"id"`               // Column Id
	SystemColumnType string           `json:"systemColumnType"` // When applicable, one of: AUTO_NUMBER, CREATED_BY, CREATED_DATE, MODIFIED_BY, MODIFIED_DATE. See System Columns.
	Type             string           `json:"type"`             // Column type
	AutoNumberFormat AutoNumberFormat `json:"autoNumberFormat"` // Present when systemColumnType == AUTO_NUMBER
	ContactOptions   []ContactOption  `json:"contactOptions"`   // Array of ContactOption objects to specify a pre-defined list of values for the column. Column type must be CONTACT_LIST
	Description      string           `json:"description"`      // Column description.
	Format           string           `json:"format"`           // The format descriptor (see Formatting). Only returned if the include query string parameter contains format and this column has a non-default format applied to it.
	Hidden           bool             `json:"hidden"`           // Indicates whether the column is hidden
	Index            int              `json:"index"`            // Column index or position. This number is zero-based.
	Locked           bool             `json:"locked"`           // Indicates whether the column is locked. In a response, a value of true indicates that the column has been locked by the sheet owner or the admin.
	LockedForUser    bool             `json:"lockedForUser"`    // Indicates whether the column is locked for the requesting user. This attribute may be present in a response, but cannot be specified in a request.
	Options          []string         `json:"options"`          // Array of the options available for the column
	Primary          bool             `json:"primary"`          // Returned only if the column is the Primary Column (value = true)
	Symbol           string           `json:"symbol"`           // When applicable for CHECKBOX or PICKLIST column types. See Symbol Columns.
	Tags             []string         `json:"tags"`             // Set of tags to indicate special columns. Each element in the array is set to one of the following values:
	Title            string           `json:"title"`            // Column title
	Validation       bool             `json:"validation"`       // Indicates whether validation has been enabled for the column (value = true)
	Version          int              `json:"version"`          // Read only. The level of the column type. Each element in the array is set to one of the following values:
	Width            int              `json:"width"`            // Display width of the column in pixels
}

type AutoNumberFormat struct {
	Fill           string `json:"fill"`           // Indicates zero-padding. Must be between 0 and 10 "0" (zero) characters.
	Prefix         string `json:"prefix"`         // The prefix. Can include the date tokens:
	StartingNumber int    `json:"startingNumber"` // The starting number for the auto-id
	Suffix         string `json:"suffix"`         // The suffix. Can include the date tokens:
}
