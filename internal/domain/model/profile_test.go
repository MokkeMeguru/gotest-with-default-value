package model_test

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/MokkeMeguru/gotest-with-default-value/internal/domain/model"
	model_fixture "github.com/MokkeMeguru/gotest-with-default-value/internal/domain/model/fixture"
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
				profile: model_fixture.Profile_Normal,
			},
			wantErr: false,
		},
		{
			name: "another man",
			args: args{
				profile: model_fixture.Profile_AnotherMan,
			},
			wantErr: true,
		},
		{
			name: "empty man",
			args: args{
				profile: model_fixture.Profile_EmptyMan,
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
	tests := []struct {
		name    string
		m       model.Profile
		want    string
		wantErr bool
	}{
		{
			name:    "normal",
			m:       model_fixture.Profile_Normal,
			want:    "",
			wantErr: false,
		},
		{
			name:    "another man",
			m:       model_fixture.Profile_AnotherMan,
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := testutil.FillTestData(&tt.m); err != nil {
				t.Errorf("test data filler failed: %v", err)
			}
			got, err := tt.m.FullName()
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
