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

type Template struct {
	id             int      // Template Id
	_type          string   // Type of the template. One of sheet or report. Only applicable to public templates
	accessLevel    string   // User's permissions on the template
	blank          bool     // Indicates whether the template is blank. Only applicable to public templates
	categories     []string // List of categories this template belongs to. Only applicable to public templates
	description    string   // Template description
	globalTemplate string   // Type of global template. One of: BLANK_SHEET, PROJECT_SHEET, or TASK_LIST. Only applicable to blank public templates
	image          string   // URL to the small preview image for this template. Only applicable to non-blank public templates
	largeImage     string   // URL to the large preview image for this template. Only applicable to non-blank public templates
	locale         string   // Locale of the template. Only applicable to public templates
	name           string   // Template name
	tags           []string // List of search tags for this template. Only applicable to non-blank public templates
}
