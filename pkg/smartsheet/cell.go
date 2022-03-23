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

type Cell struct {
	CellHistory
	ColumnId           int64        `json:"columnId,omitempty"`           //The Id of the column that the cell is located in
	ColumnType         string       `json:"columnType,omitempty"`         //See type definition on the Column object. Only returned if the include query string parameter contains columnType.
	ConditionalFormat  string       `json:"conditionalFormat,omitempty"`  //The format descriptor describing this cell's conditional format (see Formatting). Only returned if the include query string parameter contains format and this cell has a conditional format applied.
	DisplayValue       string       `json:"displayValue,omitempty"`       //Visual representation of cell contents, as presented to the user in the UI. See Cell Reference.
	Format             string       `json:"format,omitempty"`             //The format descriptor (see Formatting) Only returned if the include query string parameter contains format and this cell has a non-default format applied.
	Formula            string       `json:"formula,omitempty"`            //The formula for a cell, if set, for instance =COUNTM([Assigned To]3). NOTE: calculation errors or problems with a formula do not cause the API call to return an error code. Instead, the response contains the same value as in the UI, such as cell.value = "#CIRCULAR REFERENCE".
	Hyperlink          *Hyperlink   `json:"hyperlink,omitempty"`          //A hyperlink to a URL, sheet, or report
	Image              *Image       `json:"image,omitempty"`              //The image that the cell contains. Only returned if the cell contains an image.
	LinkInFromCell     *CellLink    `json:"linkInFromCell,omitempty"`     //An inbound link from a cell in another sheet. This cell's value mirrors the linked cell's value.
	LinksOutToCells    []CellLink   `json:"linksOutToCells,omitempty"`    //An array of CellLink objects. Zero or more outbound links from this cell to cells in other sheets whose values mirror this cell's value.
	ObjectValue        *ObjectValue `json:"objectValue,omitempty"`        //Optionally included object representation of the cell's value. Used for updating predecessor cell values or for multi-contact details.
	OverrideValidation bool         `json:"overrideValidation,omitempty"` //(Admin only) Indicates whether the cell value can contain a value outside of the validation limits (value = true). When using this parameter, you must also set strict to false to bypass value type checking. This property is honored for POST or PUT actions that update rows.
	Strict             *bool        `json:"strict,omitempty"`             //Set to false to enable lenient parsing. Defaults to true. You can specify this attribute in a request, but it is never present in a response. See Cell Value Parsing for more information about using this attribute.
	Value              interface{}  `json:"value,omitempty"`              //string number, or Boolean 	A string, a number, or a Boolean value -- depending on the cell type and the data in the cell. Cell values larger than 4000 characters are silently truncated. An empty cell returns no value. See Cell Reference.
}

type CellHistory struct {
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	ModifiedBy *User      `json:"modifiedBy,omitempty"`
}

type CellLink struct {
	ColumnId  int64  `json:"columnId,omitempty"`  // Column Id of the linked cell
	RowId     int    `json:"rowId,omitempty"`     // Row Id of the linked cell
	SheetId   int    `json:"sheetId,omitempty"`   // Sheet Id of the sheet that the linked cell belongs to
	SheetName string `json:"sheetName,omitempty"` // Sheet name of the linked cell
	Status    string `json:"status,omitempty"`
}

type Hyperlink struct {
	ReportId int    `json:"reportId,omitempty"` // If non-null, this hyperlink is a link to the report with this Id.
	SheetId  int    `json:"sheetId,omitempty"`  // If non-null, this hyperlink is a link to the sheet with this Id.
	SightId  int    `json:"sightId,omitempty"`  // If non-null, this hyperlink is a link to the Sight with this Id.
	Url      string `json:"url,omitempty"`      // When the hyperlink is a URL link, this property contains the URL value. When the hyperlink is a sheet/report/Sight link (that is, sheetId, reportId, or sightId is non-null), this property contains the permalink to the sheet, report, or Sight.
}

type Image struct {
	Id      string `json:"id,omitempty"`      // Image Id
	AltText string `json:"altText,omitempty"` // Alternate text for the image
	Height  int    `json:"height,omitempty"`  // Original height (in pixels) of the uploaded image
	Width   int    `json:"width,omitempty"`   // Original width (in pixels) of the uploaded image
}

type ObjectValue struct {
	ObjectType string `json:"objectType,omitempty"`
}
