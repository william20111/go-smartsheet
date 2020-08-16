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
	ColumnId           int         `json:"columnId"`           //The Id of the column that the cell is located in
	ColumnType         string      `json:"columnType"`         //See type definition on the Column object. Only returned if the include query string parameter contains columnType.
	ConditionalFormat  string      `json:"conditionalFormat"`  //The format descriptor describing this cell's conditional format (see Formatting). Only returned if the include query string parameter contains format and this cell has a conditional format applied.
	DisplayValue       string      `json:"displayValue"`       //Visual representation of cell contents, as presented to the user in the UI. See Cell Reference.
	Format             string      `json:"format"`             //The format descriptor (see Formatting) Only returned if the include query string parameter contains format and this cell has a non-default format applied.
	Formula            string      `json:"formula"`            //The formula for a cell, if set, for instance =COUNTM([Assigned To]3). NOTE: calculation errors or problems with a formula do not cause the API call to return an error code. Instead, the response contains the same value as in the UI, such as cell.value = "#CIRCULAR REFERENCE".
	Hyperlink          Hyperlink   `json:"hyperlink"`          //A hyperlink to a URL, sheet, or report
	Image              Image       `json:"image"`              //The image that the cell contains. Only returned if the cell contains an image.
	LinkInFromCell     CellLink    `json:"linkInFromCell"`     //An inbound link from a cell in another sheet. This cell's value mirrors the linked cell's value.
	LinksOutToCells    []CellLink  `json:"linksOutToCells"`    //An array of CellLink objects. Zero or more outbound links from this cell to cells in other sheets whose values mirror this cell's value.
	ObjectValue        ObjectValue `json:"objectValue"`        //Optionally included object representation of the cell's value. Used for updating predecessor cell values or for multi-contact details.
	OverrideValidation bool        `json:"overrideValidation"` //(Admin only) Indicates whether the cell value can contain a value outside of the validation limits (value = true). When using this parameter, you must also set strict to false to bypass value type checking. This property is honored for POST or PUT actions that update rows.
	Strict             bool        `json:"strict"`             //Set to false to enable lenient parsing. Defaults to true. You can specify this attribute in a request, but it is never present in a response. See Cell Value Parsing for more information about using this attribute.
	Value              string      `json:"value"`              //string number, or Boolean 	A string, a number, or a Boolean value -- depending on the cell type and the data in the cell. Cell values larger than 4000 characters are silently truncated. An empty cell returns no value. See Cell Reference.
}

type CellHistory struct {
	ModifiedAt time.Time `json:"modifiedAt"`
	ModifiedBy User      `json:"modifiedBy"`
}

type CellLink struct {
	ColumnId  int    `json:"columnId"`  // Column Id of the linked cell
	RowId     int    `json:"rowId"`     // Row Id of the linked cell
	SheetId   int    `json:"sheetId"`   // Sheet Id of the sheet that the linked cell belongs to
	SheetName string `json:"sheetName"` // Sheet name of the linked cell
	Status    string `json:"status"`
}

type Hyperlink struct {
	ReportId int    `json:"reportId"` // If non-null, this hyperlink is a link to the report with this Id.
	SheetId  int    `json:"sheetId"`  // If non-null, this hyperlink is a link to the sheet with this Id.
	SightId  int    `json:"sightId"`  // If non-null, this hyperlink is a link to the Sight with this Id.
	Url      string `json:"url"`      // When the hyperlink is a URL link, this property contains the URL value. When the hyperlink is a sheet/report/Sight link (that is, sheetId, reportId, or sightId is non-null), this property contains the permalink to the sheet, report, or Sight.
}

type Image struct {
	Id      string `json:"id"`      // Image Id
	AltText string `json:"altText"` // Alternate text for the image
	Height  int    `json:"height"`  // Original height (in pixels) of the uploaded image
	Width   int    `json:"width"`   // Original width (in pixels) of the uploaded image
}

type ObjectValue struct {
	objectType string
}
