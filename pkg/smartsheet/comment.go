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

type Comment struct {
	Id           int          // Comment Id
	DiscussionId int          // (optional) 	Discussion Id
	Attachments  []Attachment // Array of Attachment objects
	CreatedAt    time.Time    // Time of creation
	CreatedBy    User         // User object containing name and email of the comment's author
	ModifiedAt   time.Time    // Time of last modification
	Text         string       // Comment body
}
