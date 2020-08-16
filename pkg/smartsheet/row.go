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

import "time"

type Row struct {
	Id                int          `json:"id"`                // Row Id
	SheetId           int          `json:"sheetId"`           // Parent sheet Id
	AccessLevel       string       `json:"accessLevel"`       // User's permission level on the sheet that contains the row
	Attachments       []Attachment `json:"attachments"`       // Array of Attachment objects. Only returned if the include query string parameter contains attachments.
	Cells             []Cell       `json:"cells"`             // Array of Cell objects belonging to the row
	Columns           []Column     `json:"columns"`           // Array of Column objects. Only returned if the Get Row include query string parameter contains columns.
	ConditionalFormat string       `json:"conditionalFormat"` // Describes this row's conditional format (see Formatting). Only returned if the include query string parameter contains format and this row has a conditional format applied.
	CreatedAt         time.Time    `json:"createdAt"`         // Time of creation
	CreatedBy         User         `json:"createdBy"`         // User object containing name and email of the creator of this row
	Discussions       []Discussion `json:"discussions"`       // Array of Discussion objects. Only returned if the include query string parameter contains discussions.
	Expanded          bool         `json:"expanded"`          // Indicates whether the row is expanded or collapsed
	FilteredOut       bool         `json:"filteredOut"`       // true if this row is filtered out by a column filter (and thus is not displayed in the Smartsheet app), false if the row is not filtered out. Only returned if the include query string parameter contains filters.
	Format            string       `json:"format"`            // Format descriptor (see Formatting). Only returned if the include query string parameter contains format and this row has a non-default format applied.
	InCriticalPath    bool         `json:"inCriticalPath"`    // Only returned, with a value of true, if the sheet is a project sheet with dependencies enabled and this row is in the critical path
	Locked            bool         `json:"locked"`            // Indicates whether the row is locked. In a response, a value of true indicates that the row has been locked by the sheet owner or the admin.
	LockedForUser     bool         `json:"lockedForUser"`     // Indicates whether the row is locked for the requesting user. This attribute may be present in a response, but cannot be specified in a request.
	ModifiedAt        time.Time    `json:"modifiedAt"`        // Time of last modification
	ModifiedBy        User         `json:"modifiedBy"`        // User object containing name and email of the last person to modify this row
	Permalink         string       `json:"permalink"`         // URL that represents a direct link to the row in Smartsheet. Only returned if the include query string parameter contains rowPermalink.
	Rownumber         int          `json:"rownumber"`         // Row int within the sheet (1-based - starts at 1)
	Version           int          `json:"version"`           // Sheet version int that is incremented every time a sheet is modified
}

// Return number of cells in the row
func (r Row) CellCount() int {
	return len(r.Cells)
}
