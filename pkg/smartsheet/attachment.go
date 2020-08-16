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

type Attachment struct {
	Id                 int       // Attachment Id
	ParentId           int       // The Id of the parent
	AttachmentType     string    // Attachment type (one of BOX_COM, DROPBOX*, EGNYTE*, EVERNOTE*, FILE, GOOGLE_DRIVE, LINK, or ONEDRIVE) *Not supported for all account types, see below.
	AttachmentSubType  string    // Attachment sub type, valid only for the following:
	MimeType           string    // Attachment MIME type (PNG, etc.)
	ParentType         string    // The type of object the attachment belongs to (one of COMMENT, ROW, or SHEET)
	CreatedAt          time.Time // A timestamp of when the attachment was originally added
	CreatedBy          User      // User object containing name and email of the creator of this attachment
	Name               string    // Attachment name
	SizeInKb           int       // The size of the file, if the attachmentType is FILE
	Url                string    // Attachment temporary URL (files only)
	UrlExpiresInMillis int       // Attachment temporary URL time to live (files only)
}
