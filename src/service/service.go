package service

import (
	"looking-for-remote-jobs/src/crawler"
	"strings"
)

func GetAllOportunities(job string) string {

	// chamar metodo para replace de espaço por +
	// .net+developer
	crawler.GetOpportunitiesFlexJobs(job)
	return ""
}

func replaceSpaceForPlace(text string) string{
	str := text
	result := strings.Replace(str, " ", "+", -1)
	return result
}