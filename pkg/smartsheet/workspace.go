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

type Workspace struct {
	Id          int        `json:"id"`          // Workspace Id
	AccessLevel string     `json:"accessLevel"` // User's permissions on the workspace
	Favorite    bool       `json:"favorite"`    // Returned only if the user has marked the workspace as a favorite in their "Home" tab (value = true)
	Folders     []Folder   `json:"folders"`     // Array of Folder objects
	Name        string     `json:"name"`        // Workspace name
	Permalink   string     `json:"permalink"`   // URL that represents a direct link to the workspace in Smartsheet
	Reports     []Report   `json:"reports"`     // Array of Report objects
	Sheets      []Sheet    `json:"sheets"`      // Array of Sheet objects
	Sights      []Sight    `json:"sights"`      // Array of Sight objects
	Templates   []Template `json:"templates"`   // Array of Template objects
}

type Folder struct {
	Id        int        `json:"id"`        // Folder Id
	Favorite  bool       `json:"favorite"`  // Returned only if the user has marked the folder as a favorite in their "Home" tab (value = true)
	Folders   []Folder   `json:"folders"`   // Array of Folder objects
	Name      string     `json:"name"`      // Folder name
	Permalink string     `json:"permalink"` // URL that represents a direct link to the folder in Smartsheet
	Reports   []Report   `json:"reports"`   // Array of Report objects
	Sheets    []Sheet    `json:"sheets"`    // Array of Sheet objects
	Sights    []Sight    `json:"sights"`    // Array of Sight objects
	Templates []Template `json:"templates"` // Array of Template objects
}

type Report struct {
	Scope        Scope   `json:"scope"`        // A report's scope can be defined as the sheet ids and workspace ids that make up the report.
	SourceSheets []Sheet `json:"sourceSheets"` // Array of Sheet objects (without rows), representing the sheets that rows in the report originated from. Only included in the Get Report response if the include parameter specifies sourceSheets.
}

type Sight struct {
	Id              int       `json:"id"`              // Sight Id
	AccessLevel     string    `json:"accessLevel"`     // User's permissions on the Sight
	BackgroundColor string    `json:"backgroundColor"` // The hex color, for instance #E6F5FE
	ColumnCount     int       `json:"columnCount"`     //	Number of columns that the Sight contains
	CreatedAt       time.Time `json:"createdAt"`       //	Time of creation
	Favorite        bool      `json:"favorite"`        //	Indicates whether the user has marked the Sight as a favorite
	ModifiedAt      time.Time `json:"modifiedAt"`      //	Time of last modification
	Name            string    `json:"name"`            //	Sight name
	Permalink       string    `json:"permalink"`       //	URL that represents a direct link to the Sight in Smartsheet
	Source          Source    `json:"source"`          //	A Source object indicating the Sight (aka dashboard) from which this Sight was created, if any
	Widgets         []Widget  `json:"widgets"`         //	Array of Widget objects
	Workspace       Workspace `json:"workspace"`       //	A Workspace object, limited to only 2 attributes:
}

type Widget struct {
	Id            int         `json:"id"`            // 	Widget Id
	Type          string      `json:"type"`          // Type of widget. See table below to see how UI widget names map to type.
	Contents      interface{} `json:"contents"`      // object 	Data that specifies the contents of the widget. NOTE: The type of WidgetContent object (and attributes within) depends on the value of widget.type.
	Height        int         `json:"height"`        // 	Number of rows that the widget occupies on the Sight
	ShowTitle     bool        `json:"showTitle"`     //	True indicates that the client should display the widget title. NOTE: This is independent of the title string which may be null or empty.
	ShowTitleIcon bool        `json:"showTitleIcon"` //	True indicates that the client should display the sheet icon in the widget title
	Title         string      `json:"title"`         // Title of the widget
	TitleFormat   string      `json:"titleFormat"`   // FormatDescriptor
	Version       int         `json:"version"`       // Widget version int //
	ViewMode      int         `json:"viewMode"`      //	1 indicates content is centered. 2 indicates content is left aligned. Must use a query parameter of level=2 to see this information.
	Width         int         `json:"width"`         // 	Number of columns that the widget occupies on the Sight
	XPosition     int         `json:"xPosition"`     // 	X-coordinate of widget's position on the Sight
	YPosition     int         `json:"yPosition"`     // 	Y-coordinate of widget's position on the Sight
}
