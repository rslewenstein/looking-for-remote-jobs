package tests

import (
	"looking-for-remote-jobs/src/service"
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
		for _, v := range service.GetTelegramMessage(test.message) {
			if !strings.Contains(v.Url, test.expected) {
				t.Error("Test failed")
			}else{
				return
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
		for _, v := range service.GetTelegramMessage(test.message) {
			if !strings.Contains(v.Url, test.expected) {
				t.Error("Test failed")
			}
		}
	}
}
