package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	tests := []struct {
		name    string
		input   http.Header
		wantStr string
		wantErr error
	}{
		{"Header does not contain Authorization header", http.Header{}, "", ErrNoAuthHeaderIncluded},
		{"Header has malformed Authorization header with > 2 components", http.Header{"Authorization": []string{"ApiKey s d"}}, "", fmt.Errorf("malformed authorization header")},
		{"Header has malformed Authorization header with = 2 components", http.Header{"Authorization": []string{"a s"}}, "", fmt.Errorf("malformed authorization header")},
		{"Header has malformed Authorization header with < 2 components", http.Header{"Authorization": []string{"ApiKey"}}, "", fmt.Errorf("malformed authorization header")},
		{"Header is correct", http.Header{"Authorization": []string{"ApiKey s"}}, "s", nil},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			gotStr, gotErr := GetAPIKey(v.input)

			if gotStr != v.wantStr {
				t.Errorf("got: %v, want %v", gotStr, v.wantStr)
			}

			if (gotErr != nil) != (v.wantErr != nil) || (gotErr != nil && gotErr.Error() != v.wantErr.Error()) {
				t.Errorf("gotErr: %v, wantErr: %v", gotErr, v.wantErr)
			}
		})
	}
}
