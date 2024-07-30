package auth

import (
    "net/http"
    "testing"
)

func TestGetAPIKey(t *testing.T) {
    tests := []struct {
        name string
        headers http.Header
        expectedKey string
        expectedError error
    }{
        {
            name: "No Authorization Header",
            headers: http.Header{},
            expectedKey: "",
            expectedError: ErrNoAuthHeaderIncluded,
        },
        {
            name: "Malformed Authorization Header",
            headers: http.Header{
                "Authorization": []string{"Bearer 1234567890"},
            },
            expectedKey: "",
            expectedError: errors.New("malformed authorization header"),
        },
        {
            name: "Valid Authorization Header",
            headers: http.Header{
                "Authorization": []string{"ApiKey 1234567890"},
            },
            expectedKey: "1234567890",
            expectedError: nil,
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            key, err := GetAPIKey(test.headers)
            if err != nil {
                if test.expectedError == nil {
                    t.Errorf("Expected no error, got %v", err)
                } else if !strings.Contains(err.Error(), test.expectedError.Error()) {
                    t.Errorf("Expected error %v, got %v", test.expectedError, err)
                }
            } else {
                if test.expectedError != nil {
                    t.Errorf("Expected error %v, got %v", test.expectedError, err)
                } else {
                    if test.expectedKey != key {
                        t.Errorf("Expected key %v, got %v", test.expectedKey, key)
                    }
                }
            }
        })
    }
}