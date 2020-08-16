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

type Discussion struct {
	Id                 int          // Discussion Id
	ParentId           int          // Id of the directly associated row or sheet: present only when the direct association is not clear (see List Discussions)
	ParentType         string       // SHEET or ROW: present only when the direct association is not clear (see List Discussions)
	AccessLevel        string       // User's permissions on the discussion
	CommentAttachments []Attachment // Array of Attachment objects
	CommentCount       int          // The number of comments in the discussion
	Comments           []Comment    // Array of Comment objects
	CreatedBy          User         // User object containing name and email of the creator of the discussion
	LastCommentedAt    time.Time    // Time of most recent comment
	LastCommentedUser  User         // User object containing name and email of the author of the most recent comment
	ReadOnly           bool         // Indicates whether the user can modify the discussion
	Title              string       // Read Only. Discussion title automatically created by duplicating the first 100 characters of the top-level comment
}
