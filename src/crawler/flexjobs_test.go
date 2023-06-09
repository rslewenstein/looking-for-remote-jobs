package crawler

import (
	"strings"
	"testing"
)

func Test_GetOpportunitiesFlexJobs_ShouldReturnsResults_WhenParameterIsNotNull(t *testing.T) {
	var tests = []struct {
		message  string
		expected string
	}{
		{".net developer", "flexjobs.com"},
	}

	for _, test := range tests {
		for _, v := range GetOpportunitiesFlexJobs(test.message) {
			if !strings.Contains(v.Url, test.expected) {
				t.Error("Test failed")
			}
		}
	}
}

func Test_GetOpportunitiesFlexJobs_ShouldntReturnsResults_WhenParameterIsNull(t *testing.T) {
	var tests = []struct {
		message  string
		expected string
	}{
		{"", ""},
	}

	for _, test := range tests {
		for _, v := range GetOpportunitiesFlexJobs(test.message) {
			if !strings.Contains(v.Url, test.expected) {
				t.Error("Test failed")
			}
		}
	}
}
