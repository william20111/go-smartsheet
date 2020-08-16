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

type User struct {
	id                        int       //   User Id
	admin                     bool      //   Indicates whether the user is a system admin (can manage user accounts and organization account)
	customWelcomeScreenViewed time.Time //    Timestamp of viewing an Enterprise Custom Welcome Screen by the current user
	email                     string    //User's primary email address
	firstName                 string    //User's first name
	groupAdmin                bool      //Indicates whether the user is a group admin (can create and edit groups)
	lastLogin                 time.Time //Last login time of the current user
	lastName                  string    //User's last name
	licensedSheetCreator      bool      // Indicates whether the user is a licensed user (can create and own sheets)
	name                      string    //User's full name (read-only)
	profileImage              Image     //An Image object representing the profile image associated with the user account
	resourceViewer            bool      //Indicates whether the user is a resource viewer (can access resource views)
	sheetCount                int       //The number of sheets owned by the current user within the organization account
	status                    string    //User status, set to one of the following values: ACTIVE, DECLINED, or PENDING
}
