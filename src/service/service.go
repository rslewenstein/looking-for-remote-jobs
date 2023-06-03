package service

import (
	"looking-for-remote-jobs/src/crawler"
	"strings"
)

func GetAllOportunities(job string) string {
	result := replaceSpaceForPlus(job)
	crawler.GetOpportunitiesFlexJobs(result)
	return ""
}

func replaceSpaceForPlus(text string) string {
	str := text
	result := strings.Replace(str, " ", "+", -1)
	return result
}
