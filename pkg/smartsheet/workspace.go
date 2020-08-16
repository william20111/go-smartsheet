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
	Id          int        // Workspace Id
	AccessLevel string     // User's permissions on the workspace
	Favorite    bool       // Returned only if the user has marked the workspace as a favorite in their "Home" tab (value = true)
	Folders     []Folder   // Array of Folder objects
	Name        string     // Workspace name
	Permalink   string     // URL that represents a direct link to the workspace in Smartsheet
	Reports     []Report   // Array of Report objects
	Sheets      []Sheet    // Array of Sheet objects
	Sights      []Sight    // Array of Sight objects
	Templates   []Template // Array of Template objects
}

type Folder struct {
	Id        int        // Folder Id
	Favorite  bool       // Returned only if the user has marked the folder as a favorite in their "Home" tab (value = true)
	Folders   []Folder   // Array of Folder objects
	Name      string     // Folder name
	Permalink string     // URL that represents a direct link to the folder in Smartsheet
	Reports   []Report   // Array of Report objects
	Sheets    []Sheet    // Array of Sheet objects
	Sights    []Sight    // Array of Sight objects
	Templates []Template // Array of Template objects
}

type Report struct {
	Scope        Scope   // A report's scope can be defined as the sheet ids and workspace ids that make up the report.
	SourceSheets []Sheet // Array of Sheet objects (without rows), representing the sheets that rows in the report originated from. Only included in the Get Report response if the include parameter specifies sourceSheets.
}

type Sight struct {
	Id              int       // Sight Id
	AccessLevel     string    // User's permissions on the Sight
	BackgroundColor string    // The hex color, for instance #E6F5FE
	ColumnCount     int       //	Number of columns that the Sight contains
	CreatedAt       time.Time //	Time of creation
	Favorite        bool      //	Indicates whether the user has marked the Sight as a favorite
	ModifiedAt      time.Time //	Time of last modification
	Name            string    //	Sight name
	Permalink       string    //	URL that represents a direct link to the Sight in Smartsheet
	Source          Source    //	A Source object indicating the Sight (aka dashboard) from which this Sight was created, if any
	Widgets         []Widget  //	Array of Widget objects
	Workspace       Workspace //	A Workspace object, limited to only 2 attributes:
}

type Widget struct {
	Id            int         // 	Widget Id
	Type          string      // Type of widget. See table below to see how UI widget names map to type.
	Contents      interface{} // object 	Data that specifies the contents of the widget. NOTE: The type of WidgetContent object (and attributes within) depends on the value of widget.type.
	Height        int         // 	Number of rows that the widget occupies on the Sight
	ShowTitle     bool        //	True indicates that the client should display the widget title. NOTE: This is independent of the title string which may be null or empty.
	ShowTitleIcon bool        //	True indicates that the client should display the sheet icon in the widget title
	Title         string      // Title of the widget
	TitleFormat   string      // FormatDescriptor
	Version       int         // Widget version int //
	ViewMode      int         //	1 indicates content is centered. 2 indicates content is left aligned. Must use a query parameter of level=2 to see this information.
	Width         int         // 	Number of columns that the widget occupies on the Sight
	XPosition     int         // 	X-coordinate of widget's position on the Sight
	YPosition     int         // 	Y-coordinate of widget's position on the Sight
}
