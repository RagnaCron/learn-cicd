package auth

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		key           string
		value         string
		expected      string
		expectedError string
	}{
		"Simple API Key":                   {key: "Authorization", value: "ApiKey som_token", expected: "some_token"},
		"No Auth header":                   {key: "Authorization", expectedError: "no authorization header"},
		"Malformed Auth Header":            {key: "Authorization", value: "-", expectedError: "malformed authorization header"},
		"Malformed Auth Header; No ApiKey": {key: "Authorization", value: "Bearer xxxxxx", expectedError: "malformed authorization header"},
		"Not expected":                     {key: "Authorization", value: "ApiKey xxxxxx", expected: "xxxxxx", expectedError: "not expecting an error"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			header := http.Header{}
			header.Add(tc.key, tc.value)

			output, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), tc.expectedError) {
					return
				}
				t.Errorf("Unexpected: TestGetAPIKey:%v\n", err)
				return
			}

			if output != tc.expected {
				t.Errorf("Unexpected: TestGetAPIKey:%s", output)
				return
			}
		})
	}
}
