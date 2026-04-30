package foobar_test

import (
	"testing"

	"github.com/padok-team/github-actions-tutorial/foobar"
)

func TestSequence(t *testing.T) {
	tests := []struct {
		length    int
		expected  []string
		expectErr bool
	}{
		{
			length:    7,
			expected:  []string{"1", "2", "foo", "4", "bar"},
			expectErr: false,
		},
		{
			length:    15,
			expected:  []string{"1", "2", "foo", "4", "bar", "foo", "7", "8", "foo", "bar", "11", "foo", "13", "14", "foobar"},
			expectErr: false,
		},
		{
			length:    0,
			expected:  []string{},
			expectErr: false,
		},
		{
			length:    -3,
			expected:  nil,
			expectErr: true,
		},
	}

	for _, test := range tests {
		actual, err := foobar.Sequence(test.length)

		if err != nil && !test.expectErr {
			t.Errorf("got unexpected error: %s", err)
		}
		if err == nil && test.expectErr {
			t.Error("expected error, got nil")
		}

		if !sequencesEqual(test.expected, actual) {
			t.Errorf("expected %q, got %q", test.expected, actual)
		}
	}
}

func sequencesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
