package tests

import (
	"looking-for-remote-jobs/src/crawler"
	"strings"
	"testing"
)

func Test_GetOpportunitiesWeWorkRemotely_ShouldReturnsResults_WhenParameterIsNotNull(t *testing.T) {
	var tests = []struct {
		message  string
		expected string
	}{
		{"developer", "weworkremotely.com"},
	}

	for _, test := range tests {
		for _, v := range crawler.GetOpportunitiesWeWorkRemotely(test.message) {
			if !strings.Contains(v.Url, test.expected) {
				t.Error("Test failed")
			}
		}
	}
}

func Test_GetOpportunitiesWeWorkRemotely_ShouldntReturnsResults_WhenParameterIsNull(t *testing.T) {
	var tests = []struct {
		message  string
		expected string
	}{
		{"", ""},
	}

	for _, test := range tests {
		for _, v := range crawler.GetOpportunitiesWeWorkRemotely(test.message) {
			if !strings.Contains(v.Url, test.expected) {
				t.Error("Test failed")
			}
		}
	}
}
