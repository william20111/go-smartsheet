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

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"time"
)

type Row struct {
	Id                int64        `json:"id,omitempty"`                // Row Id
	SheetId           int64        `json:"sheetId,omitempty"`           // Parent sheet Id
	AccessLevel       string       `json:"accessLevel,omitempty"`       // User's permission level on the sheet that contains the row
	Attachments       []Attachment `json:"attachments,omitempty"`       // Array of Attachment objects. Only returned if the include query string parameter contains attachments.
	Cells             []Cell       `json:"cells"`                       // Array of Cell objects belonging to the row
	Columns           []Column     `json:"columns,omitempty"`           // Array of Column objects. Only returned if the Get Row include query string parameter contains columns.
	ConditionalFormat string       `json:"conditionalFormat,omitempty"` // Describes this row's conditional format (see Formatting). Only returned if the include query string parameter contains format and this row has a conditional format applied.
	CreatedAt         *time.Time   `json:"createdAt,omitempty"`         // Time of creation
	CreatedBy         *User        `json:"createdBy,omitempty"`         // User object containing name and email of the creator of this row
	Discussions       []Discussion `json:"discussions,omitempty"`       // Array of Discussion objects. Only returned if the include query string parameter contains discussions.
	Expanded          bool         `json:"expanded,omitempty"`          // Indicates whether the row is expanded or collapsed
	FilteredOut       bool         `json:"filteredOut,omitempty"`       // true if this row is filtered out by a column filter (and thus is not displayed in the Smartsheet app), false if the row is not filtered out. Only returned if the include query string parameter contains filters.
	Format            string       `json:"format,omitempty"`            // Format descriptor (see Formatting). Only returned if the include query string parameter contains format and this row has a non-default format applied.
	InCriticalPath    bool         `json:"inCriticalPath,omitempty"`    // Only returned, with a value of true, if the sheet is a project sheet with dependencies enabled and this row is in the critical path
	Locked            bool         `json:"locked,omitempty"`            // Indicates whether the row is locked. In a response, a value of true indicates that the row has been locked by the sheet owner or the admin.
	LockedForUser     bool         `json:"lockedForUser,omitempty"`     // Indicates whether the row is locked for the requesting user. This attribute may be present in a response, but cannot be specified in a request.
	ModifiedAt        *time.Time   `json:"modifiedAt,omitempty"`        // Time of last modification
	ModifiedBy        *User        `json:"modifiedBy,omitempty"`        // User object containing name and email of the last person to modify this row
	Permalink         string       `json:"permalink,omitempty"`         // URL that represents a direct link to the row in Smartsheet. Only returned if the include query string parameter contains rowPermalink.
	Rownumber         int64        `json:"rownumber,omitempty"`         // Row int within the sheet (1-based - starts at 1)
	Version           int64        `json:"version,omitempty"`           // Sheet version int that is incremented every time a sheet is modified
	ToTop             bool         `json:"toTop,omitempty"`
	ToBottom          bool         `json:"toBottom,omitempty"`
}

// Return number of cells in the row
func (r Row) CellCount() int {
	return len(r.Cells)
}

// Return ResultObject object
func (c Client) AddRow(sheetId int64, rows []Row) (*[]Row, error) {
	var res ResultObject
	resp, err := c.post(fmt.Sprintf("%s/sheets/%d/rows", apiEndpoint, sheetId), rows, nil)
	if err != nil {
		return nil, err
	}
	if dErr := c.decodeJSON(resp, &res); dErr != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
	}
	var result []Row
	err = mapstructure.Decode(res.Result, &result)
	return &result, nil
}
