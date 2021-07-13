package model_test

import (
	"testing"

	"github.com/MokkeMeguru/gotest-with-default-value/internal/domain/model"
	"github.com/MokkeMeguru/gotest-with-default-value/pkg/testutil"
)

func TestValidate(t *testing.T) {
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
