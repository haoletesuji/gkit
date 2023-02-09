package gkit

import "testing"

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		email   string
		isValid bool
	}{
		{"good@exmaple.com", true},
		{"bad-example", false},
	}

	for _, test := range tests {
		isValid := IsValidEmail(test.email)
		if isValid != test.isValid {
			t.Errorf("IsValidEmail was incorrect, got: %t, want: %t.", isValid, test.isValid)
		}
	}
}
