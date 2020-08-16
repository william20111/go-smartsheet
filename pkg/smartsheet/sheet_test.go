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

import (
	"reflect"
	"testing"
	"time"
)

func TestSheet_GetColumnById(t *testing.T) {
	type fields struct {
		Id                         int
		FromId                     int
		OwnerId                    int
		AccessLevel                string
		Attachments                []Attachment
		Columns                    []Column
		CreatedAt                  time.Time
		CrossSheetReferences       []CrossSheetReference
		DependenciesEnabled        bool
		Discussions                []Discussion
		EffectiveAttachmentOptions []string
		Favorite                   bool
		GanttEnabled               bool
		HasSummaryFields           bool
		ModifiedAt                 time.Time
		Name                       string
		Owner                      string
		Permalink                  string
		ProjectSettings            ProjectSettings
		ReadOnly                   bool
		ResourceManagementEnabled  bool
		Rows                       []Row
		ShowParentRowsForFilters   bool
		Source                     Source
		Summary                    SheetSummary
		TotalRowCount              int
		UserPermissions            SheetUserPermissions
		UserSettings               SheetUserSettings
		Version                    int
		Workspace                  Workspace
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Column
		wantErr bool
	}{
		{
			"pass",
			fields{
				Columns: []Column{
					{
						Id:    1,
						Title: "test",
					},
					{
						Id:    2,
						Title: "test2",
					},
				},
			},
			args{
				id: 1,
			},
			&Column{
				Id:    1,
				Title: "test",
			},
			false,
		},
		{
			"fail",
			fields{
				Columns: []Column{
					{
						Id:    1,
						Title: "test",
					},
					{
						Id:    2,
						Title: "test2",
					},
				},
			},
			args{
				id: 3,
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sheet{
				Id:                         tt.fields.Id,
				FromId:                     tt.fields.FromId,
				OwnerId:                    tt.fields.OwnerId,
				AccessLevel:                tt.fields.AccessLevel,
				Attachments:                tt.fields.Attachments,
				Columns:                    tt.fields.Columns,
				CreatedAt:                  tt.fields.CreatedAt,
				CrossSheetReferences:       tt.fields.CrossSheetReferences,
				DependenciesEnabled:        tt.fields.DependenciesEnabled,
				Discussions:                tt.fields.Discussions,
				EffectiveAttachmentOptions: tt.fields.EffectiveAttachmentOptions,
				Favorite:                   tt.fields.Favorite,
				GanttEnabled:               tt.fields.GanttEnabled,
				HasSummaryFields:           tt.fields.HasSummaryFields,
				ModifiedAt:                 tt.fields.ModifiedAt,
				Name:                       tt.fields.Name,
				Owner:                      tt.fields.Owner,
				Permalink:                  tt.fields.Permalink,
				ProjectSettings:            tt.fields.ProjectSettings,
				ReadOnly:                   tt.fields.ReadOnly,
				ResourceManagementEnabled:  tt.fields.ResourceManagementEnabled,
				Rows:                       tt.fields.Rows,
				ShowParentRowsForFilters:   tt.fields.ShowParentRowsForFilters,
				Source:                     tt.fields.Source,
				Summary:                    tt.fields.Summary,
				TotalRowCount:              tt.fields.TotalRowCount,
				UserPermissions:            tt.fields.UserPermissions,
				UserSettings:               tt.fields.UserSettings,
				Version:                    tt.fields.Version,
				Workspace:                  tt.fields.Workspace,
			}
			got, err := s.GetColumnById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetColumnById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetColumnById() got = %v, want %v", got, tt.want)
			}
		})
	}
}
