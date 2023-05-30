package service

import (
	"looking-for-remote-jobs/src/crawler"
	"strings"
)

func GetAllOportunities(job string) string {
	result := replaceSpaceForPlace(job)
	crawler.GetOpportunitiesFlexJobs(result)
	return ""
}

func replaceSpaceForPlace(text string) string {
	str := text
	result := strings.Replace(str, " ", "+", -1)
	return result
}
