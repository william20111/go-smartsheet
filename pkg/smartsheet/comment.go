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
	id           int          // Comment Id
	discussionId int          // (optional) 	Discussion Id
	attachments  []Attachment // Array of Attachment objects
	createdAt    time.Time    // Time of creation
	createdBy    User         // User object containing name and email of the comment's author
	modifiedAt   time.Time    // Time of last modification
	text         string       // Comment body
}
