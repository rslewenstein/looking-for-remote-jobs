package service

import (
	"strings"
	"testing"
)

func Test_GetTelegramMessage_ShouldReturnResults_WhenParameterIsNotNull(t *testing.T) {
	var tests = []struct {
		message  string
		expected string
	}{
		{".net developer", "flexjobs.com"},
	}

	for _, test := range tests {
		for _, v := range GetTelegramMessage(test.message) {
			if !strings.Contains(v.Url, test.expected) {
				t.Error("Test failed")
			}
		}
	}
}

func Test_GetTelegramMessage_ShouldntReturnResults_WhenParameterIsNull(t *testing.T) {
	var tests = []struct {
		message  string
		expected string
	}{
		{"", ""},
	}

	for _, test := range tests {
		for _, v := range GetTelegramMessage(test.message) {
			if !strings.Contains(v.Url, test.expected) {
				t.Error("Test failed")
			}
		}
	}
}
