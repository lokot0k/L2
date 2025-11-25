package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
		{"a10b15cde", "aaaaaaaaaabbbbbbbbbbbbbbbcde", false},
		{"а3б2д", "аааббд", false},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result, err := unpack(test.input)
			if test.err && err == nil {
				t.Errorf("ожидалась ошибка для %v", test.input)
			}
			if !test.err && err != nil {
				t.Errorf("неожиданная ошибка для %v: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("для строки %v, ожидалось %v, получили %v", test.input, test.expected, result)
			}
		})
	}
}
