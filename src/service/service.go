package service

import "looking-for-remote-jobs/src/crawler"

func GetAllOportunities(job string) string {

	// chamar metodo para replace de espaço por +
	// .net+developer
	crawler.GetOpportunitiesFlexJobs(job)
	return ""
}
