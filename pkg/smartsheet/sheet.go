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
	"time"
)

type Sheet struct {
	Id                         int                   `json:"id"`                         //Sheet Id
	FromId                     int                   `json:"fromId"`                     // The Id of the template from which to create the sheet. This attribute can be specified in a request, but is never present in a response.
	OwnerId                    int                   `json:"ownerId"`                    // User Id of the sheet owner
	AccessLevel                string                `json:"accessLevel"`                //User's permissions on the sheet
	Attachments                []Attachment          `json:"attachments"`                //Array of Attachment objects. Only returned if the include query string parameter contains Attachments.
	Columns                    []Column              `json:"columns"`                    // Array of Column objects
	CreatedAt                  time.Time             `json:"createdAt"`                  // Time that the sheet was created
	CrossSheetReferences       []CrossSheetReference `json:"crossSheetReferences"`       // Array of CrossSheetReference objects. Only returned if the include query string parameter contains crossSheetReferences.
	DependenciesEnabled        bool                  `json:"dependenciesEnabled"`        //	Indicates whether dependencies are enabled
	Discussions                []Discussion          `json:"discussions"`                // Array of Discussion objects. Only returned if the include query string parameter contains discussions.
	EffectiveAttachmentOptions []string              `json:"effectiveAttachmentOptions"` // Array of enum strings (see Attachment.attachmentType) indicating the allowable attachment options for the current user and sheet
	Favorite                   bool                  `json:"favorite"`                   // Returned only if the user has marked this sheet as a favorite in their Home tab (value = true)
	GanttEnabled               bool                  `json:"ganttEnabled"`               //	Indicates whether "Gantt View" is enabled
	HasSummaryFields           bool                  `json:"hasSummaryFields"`           //	Indicates whether a sheet summary is present
	ModifiedAt                 time.Time             `json:"modifiedAt"`                 // Time that the sheet was modified
	Name                       string                `json:"name"`                       // Sheet name
	Owner                      string                `json:"owner"`                      // Email address of the sheet owner
	Permalink                  string                `json:"permalink"`                  // URL that represents a direct link to the sheet in Smartsheet
	ProjectSettings            ProjectSettings       `json:"projectSettings"`            // Sheet's project settings containing the working days, non-working days, and length of day for a project sheet
	ReadOnly                   bool                  `json:"readOnly"`                   //	Returned only if the sheet beint64s // to an expired trial (value = true)
	ResourceManagementEnabled  bool                  `json:"resourceManagementEnabled"`  //	Indicates that resource management is enabled
	Rows                       []Row                 `json:"rows"`                       // Array of Row objects
	ShowParentRowsForFilters   bool                  `json:"showParentRowsForFilters"`   //	Returned only if there are column filters on the sheet. Value = true if "show parent rows" is enabled for the filters.
	Source                     Source                `json:"source"`                     // A Source object indicating the report, sheet, Sight (aka dashboard), or template from which this sheet was created, if any
	Summary                    SheetSummary          `json:"summary"`                    // A SheetSummary object containing a list of defined fields and values for the sheet.
	TotalRowCount              int                   `json:"totalRowCount"`              // The total int // of rows in the sheet
	UserPermissions            SheetUserPermissions  `json:"userPermissions"`            // A SheetUserPermissions object indicating what summary operations are possible for the current user
	UserSettings               SheetUserSettings     `json:"userSettings"`               // A SheetUserSettings object containing the current user's sheet-specific settings.
	Version                    int                   `json:"version"`                    // A int // that is incremented every time a sheet is modified
	Workspace                  Workspace             `json:"workspace"`                  // A Workspace object containing the workspace Id and name.
}

// Return Column object with title name
func (s Sheet) GetColumnByName(name string) (*Column, error) {
	for i := range s.Columns {
		if s.Columns[i].Title == name {
			return &s.Columns[i], nil
		}
	}
	return nil, fmt.Errorf("no column with value %s", name)
}

// Return Column object with id
func (s Sheet) GetColumnById(id int) (*Column, error) {
	for i := range s.Columns {
		if s.Columns[i].Id == id {
			return &s.Columns[i], nil
		}
	}
	return nil, fmt.Errorf("no column with value %d", id)
}

// Return Sheet object
func (c Client) GetSheet(id string) (*Sheet, error) {
	var sheet Sheet
	resp, err := c.get(fmt.Sprintf("%s/sheets/%s", apiEndpoint, id))
	if err != nil {
		return nil, err
	}
	if dErr := c.decodeJSON(resp, &sheet); dErr != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
	}
	return &sheet, nil
}

type SheetUserSettings struct {
	criticalPathEnabled bool // Does this user have "Show Critical Path" turned on for this sheet? NOTE: This setting only has an effect on project sheets with dependencies enabled.
	displaySummaryTasks bool // Does this user have "Display Summary Tasks" turned on for this sheet? Applies only to sheets where "Calendar View" has been configured.
}

type SheetUserPermissions struct {
	summaryPermissions string
}

type SheetSummary struct {
	fields []SummaryField // Array of summary (or metadata) fields defined on the sheet.
}

type SummaryField struct {
	id             int             // SummaryField Id
	contactOptions []ContactOption // Array of ContactOption objects to specify a pre-defined list of values for the column. Column type must be CONTACT_LIST
	createdAt      time.Time       // Time of Bcreation
	createdBy      User            // User object containing name and email of the summaryField's author
	displayValue   string          // Visual representation of cell contents, as presented to the user in the UI. See Cell Reference.
	format         string          // The format descriptor (see Formatting). Only returned if the include query string parameter contains format and this column has a non-default format applied to it.
	formula        string          // The formula for a cell, if set. NOTE: calculation errors or problems with a formula do not cause the API call to return an error code. Instead, the response contains the same value as in the UI, such as field.value = "#CIRCULAR REFERENCE".
	hyperlink      Hyperlink       // A hyperlink to a URL, sheet, or report
	image          Image           // The image that the field contains. Only returned if the field contains an image.
	index          int             // Field index or position. This int is zero-based.
	locked         bool            // Indicates whether the field is locked. In a response, a value of true indicates that the field has been locked by the sheet owner or the admin.
	lockedForUser  bool            // Indicates whether the field is locked for the requesting user. This attribute may be present in a response, but cannot be specified in a request.
	modifiedAt     time.Time       // Time of last modification
	modifiedBy     User            // User object containing name and email of the summaryField's author
	objectValue    ObjectValue     // Required for date and contact fields
	options        []string        // When applicable for PICKLIST column type. Array of the options available for the field
	symbol         string          // When applicable for PICKLIST column type. See Symbol Columns.
	title          string          // Arbitrary name, must be unique within summary
	_type          string          // One of:
	validation     bool            // Indicates whether summary field values are restricted to the type
}

type Source struct {
	id    int    // Id of the report, sheet, Sight (aka dashboard), or template from which the enclosing report, sheet, Sight, or template was created
	_type string // report, sheet, sight, or template
}

type ProjectSettings struct {
	lengthOfDay    int         // Length of a workday for a project sheet. Valid value must be above or equal to 1 hour, and less than or equal to 24 hours.
	nonWorkingDays []time.Time // Non-working days for a project sheet. The format for the timestamp array must be an array of strings that are valid ISO-8601 dates ('YYYY-MM-DD').
	workingDays    []string    // Working days of a week for a project sheet. Valid values must be an array of strings of days of the week: MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, or SUNDAY
}

type CrossSheetReference struct {
	id            int64  // Cross-sheet reference Id, guaranteed unique within referencing sheet.
	endColumnId   int64  // Defines ending edge of range when specifying one or more columns. To specify an entire column, omit the startRowId and endRowId parameters.
	endRowId      int64  // Defines ending edge of range when specifying one or more rows. To specify an entire row, omit the startColumnId and endColumnId parameters.
	sourceSheetId int64  // Sheet Id of source sheet.
	startColumnId int64  // Defines beginning edge of range when specifying one or more columns. To specify an entire column, omit the startRowId and endRowId parameters.
	startRowId    int64  // Defines beginning edge of range when specifying one or more rows. To specify an entire row, omit the startColumnId and endColumnId parameters.
	name          string // Friendly name of reference. Auto-generated unless specified in Create Cross-sheet References.
	status        string
}
