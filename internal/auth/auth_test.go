package auth

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		header http.Header
		value  string
	}{
		"Simple API Key": {header: http.Header{"Authorization": {"ApiKey some_token"}}, value: "some_token"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			val, err := GetAPIKey(tc.header)
			if err != nil {
				t.Fatalf("%v", err)
			}
			diff := cmp.Diff(tc.value, val)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
