package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		header  http.Header
		want    string
		wantErr bool
	}{
		{
			name:    "missing auth header",
			header:  http.Header{"Authorization": []string{""}},
			want:    "",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		got, err := GetAPIKey(tc.header)
		if (err != nil) != tc.wantErr {
			fmt.Sprintf("Got issue in test case %s", tc.name)
			t.Fatalf("expected error: %v, got %v", tc.wantErr, err)
		}
		if got != tc.want {
			t.Errorf("expected: %q, got %q", tc.want, got)
		}
	}
}
