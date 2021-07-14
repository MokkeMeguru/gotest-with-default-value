package model_test

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/MokkeMeguru/gotest-with-default-value/internal/domain/model"
	"github.com/MokkeMeguru/gotest-with-default-value/pkg/testutil"
	"github.com/xorcare/golden"
)

func TestProfile_Validate(t *testing.T) {
	type args struct {
		profile model.Profile
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				profile: model.Profile{},
			},
			wantErr: false,
		},
		{
			name: "another man",
			args: args{
				profile: model.Profile{
					FirstName: "Potter",
					LastName:  "Harry",
				},
			},
			wantErr: true,
		},
		{
			name: "another man 2 (weak point)",
			args: args{
				profile: model.Profile{
					FirstName: " ",
					LastName:  " ",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := testutil.FillTestData(&tt.args.profile); err != nil {
				t.Errorf("test data filler failed: %v", err)
			}
			if err := tt.args.profile.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProfile_FullName(t *testing.T) {
	type fields struct {
		FirstName     string
		LastName      string
		FavoriteFoods []model.Food
		Skill         map[string]int
		Post          *model.Post
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name:    "normal",
			fields:  fields{},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := model.Profile{
				FirstName:     tt.fields.FirstName,
				LastName:      tt.fields.LastName,
				FavoriteFoods: tt.fields.FavoriteFoods,
				Skill:         tt.fields.Skill,
				Post:          tt.fields.Post,
			}
			if err := testutil.FillTestData(&m); err != nil {
				t.Errorf("test data filler failed: %v", err)
			}
			got, err := m.FullName()
			if (err != nil) != tt.wantErr {
				t.Errorf("Profile.FullName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			gotJson, err := json.Marshal(got)
			if err != nil {
				t.Errorf("marshal failed: %v %v", got, err)
			}
			golden.Assert(t, gotJson)
		})
	}
}
