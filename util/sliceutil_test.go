package util

import "testing"

func TestContains(t *testing.T) {
	cases := []struct {
		elem  string
		slice []string
		expected bool
	} {
		{"a", []string{"b", "c", "a", "d"}, true},
		{"1", []string{"b", "d", "e"}, false},
		{"", []string{"1", "2"}, false},
		{"a", nil, false},
	}

	for _, tc := range cases {
		found := Contains(tc.slice, tc.elem)
		if found != tc.expected {
			t.Errorf("contains %s in %s should return %t but got %t ", tc.elem, tc.slice, tc.expected, found)
		}
	}
}
