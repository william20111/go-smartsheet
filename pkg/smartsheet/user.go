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
	Id                        int        `json:"id,omitempty"`                        //   User Id
	Admin                     bool       `json:"admin,omitempty"`                     //   Indicates whether the user is a system admin (can manage user accounts and organization account)
	CustomWelcomeScreenViewed *time.Time `json:"customWelcomeScreenViewed,omitempty"` //    Timestamp of viewing an Enterprise Custom Welcome Screen by the current user
	Email                     string     `json:"email,omitempty"`                     //User's primary email address
	FirstName                 string     `json:"firstName,omitempty"`                 //User's first name
	GroupAdmin                bool       `json:"groupAdmin,omitempty"`                //Indicates whether the user is a group admin (can create and edit groups)
	LastLogin                 *time.Time `json:"lastLogin,omitempty"`                 //Last login time of the current user
	LastName                  string     `json:"lastName,omitempty"`                  //User's last name
	LicensedSheetCreator      bool       `json:"licensedSheetCreator,omitempty"`      // Indicates whether the user is a licensed user (can create and own sheets)
	Name                      string     `json:"name,omitempty"`                      //User's full name (read-only)
	ProfileImage              *Image     `json:"profile,omitempty"`                   //An Image object representing the profile image associated with the user account
	ResourceViewer            bool       `json:"resourceViewer,omitempty"`            //Indicates whether the user is a resource viewer (can access resource views)
	SheetCount                int        `json:"sheetCount,omitempty"`                //The number of sheets owned by the current user within the organization account
	Status                    string     `json:"status,omitempty"`                    //User status, set to one of the following values: ACTIVE, DECLINED, or PENDING
}
