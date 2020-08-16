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
	id          int        // Workspace Id
	accessLevel string     // User's permissions on the workspace
	favorite    bool       // Returned only if the user has marked the workspace as a favorite in their "Home" tab (value = true)
	folders     []Folder   // Array of Folder objects
	name        string     // Workspace name
	permalink   string     // URL that represents a direct link to the workspace in Smartsheet
	reports     []Report   // Array of Report objects
	sheets      []Sheet    // Array of Sheet objects
	sights      []Sight    // Array of Sight objects
	templates   []Template // Array of Template objects
}

type Folder struct {
	id        int        // Folder Id
	favorite  bool       // Returned only if the user has marked the folder as a favorite in their "Home" tab (value = true)
	folders   []Folder   // Array of Folder objects
	name      string     // Folder name
	permalink string     // URL that represents a direct link to the folder in Smartsheet
	reports   []Report   // Array of Report objects
	sheets    []Sheet    // Array of Sheet objects
	sights    []Sight    // Array of Sight objects
	templates []Template // Array of Template objects
}

type Report struct {
	scope        Scope   // A report's scope can be defined as the sheet ids and workspace ids that make up the report.
	sourceSheets []Sheet // Array of Sheet objects (without rows), representing the sheets that rows in the report originated from. Only included in the Get Report response if the include parameter specifies sourceSheets.
}

type Sight struct {
	id              int       // Sight Id
	accessLevel     string    // User's permissions on the Sight
	backgroundColor string    // The hex color, for instance #E6F5FE
	columnCount     int       //	Number of columns that the Sight contains
	createdAt       time.Time //	Time of creation
	favorite        bool      //	Indicates whether the user has marked the Sight as a favorite
	modifiedAt      time.Time //	Time of last modification
	name            string    //	Sight name
	permalink       string    //	URL that represents a direct link to the Sight in Smartsheet
	source          Source    //	A Source object indicating the Sight (aka dashboard) from which this Sight was created, if any
	widgets         []Widget  //	Array of Widget objects
	workspace       Workspace //	A Workspace object, limited to only 2 attributes:
}

type Widget struct {
	id            int         // 	Widget Id
	_type         string      // Type of widget. See table below to see how UI widget names map to type.
	contents      interface{} // object 	Data that specifies the contents of the widget. NOTE: The type of WidgetContent object (and attributes within) depends on the value of widget.type.
	height        int         // 	Number of rows that the widget occupies on the Sight
	showTitle     bool        //	True indicates that the client should display the widget title. NOTE: This is independent of the title string which may be null or empty.
	showTitleIcon bool        //	True indicates that the client should display the sheet icon in the widget title
	title         string      // Title of the widget
	titleFormat   string      // FormatDescriptor
	version       int         // Widget version int //
	viewMode      int         //	1 indicates content is centered. 2 indicates content is left aligned. Must use a query parameter of level=2 to see this information.
	width         int         // 	Number of columns that the widget occupies on the Sight
	xPosition     int         // 	X-coordinate of widget's position on the Sight
	yPosition     int         // 	Y-coordinate of widget's position on the Sight
}
