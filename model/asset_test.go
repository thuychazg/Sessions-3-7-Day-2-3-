package model

import (
	"testing"
)


func TestAssetCreateRequest(t *testing.T) {

	tests := []struct {
		name string
		req  CreateAssetRequest
		valid bool
	}{
		{
			name: "valid domain",
			req: CreateAssetRequest{
				Name: "google.com",
				Type: "domain",
			},
			valid: true,
		},

		{
			name: "valid ip",
			req: CreateAssetRequest{
				Name: "127.0.0.1",
				Type: "ip",
			},
			valid: true,
		},

		{
			name: "empty name",
			req: CreateAssetRequest{
				Name: "",
				Type: "domain",
			},
			valid: false,
		},

		{
			name: "invalid type",
			req: CreateAssetRequest{
				Name: "test.com",
				Type: "abc",
			},
			valid: false,
		},
	}


	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {


			isValid :=
				tt.req.Name != "" &&
				(tt.req.Type == "domain" ||
				 tt.req.Type == "ip" ||
				 tt.req.Type == "service")


			if isValid != tt.valid {

				t.Errorf(
					"validation = %v, want %v",
					isValid,
					tt.valid,
				)
			}


		})
	}
}
