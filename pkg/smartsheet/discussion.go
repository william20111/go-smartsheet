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
	id                 int          // Discussion Id
	parentId           int          // Id of the directly associated row or sheet: present only when the direct association is not clear (see List Discussions)
	parentType         string       // SHEET or ROW: present only when the direct association is not clear (see List Discussions)
	accessLevel        string       // User's permissions on the discussion
	commentAttachments []Attachment // Array of Attachment objects
	commentCount       int          // The number of comments in the discussion
	comments           []Comment    // Array of Comment objects
	createdBy          User         // User object containing name and email of the creator of the discussion
	lastCommentedAt    time.Time    // Time of most recent comment
	lastCommentedUser  User         // User object containing name and email of the author of the most recent comment
	readOnly           bool         // Indicates whether the user can modify the discussion
	title              string       // Read Only. Discussion title automatically created by duplicating the first 100 characters of the top-level comment
}
