package service

import (
	"fmt"
	"looking-for-remote-jobs/src/crawler"
	"looking-for-remote-jobs/src/model"
	"strings"
)

func GetTelegramMessage(message string) []model.Opportunity {
	var result []model.Opportunity
	if message != "" {
		result = getAllOportunities(message)
	} else {
		fmt.Println("Error to get Telegram message")
	}
	return result
}

func getAllOportunities(job string) []model.Opportunity {
	if job == "" {
		return nil
	}

	arrayResult := []model.Opportunity{}
	replaced := replaceSpaceForPlus(job)

	//TODO: Refactor these statements - FOR
	for _, v := range crawler.GetOpportunitiesFlexJobs(replaced) {
		arrayResult = append(arrayResult, v)
	}

	// TODO: Words such as DEVELOPER or DEV don't is working.
	if !strings.Contains(job, "dev") || !(strings.Contains(job, "developer")) {
		for _, v := range crawler.GetOpportunitiesHimalayas(replaced) {
			arrayResult = append(arrayResult, v)
		}
	}

	for _, v := range crawler.GetOpportunitiesWeWorkRemotely(replaced) {
		arrayResult = append(arrayResult, v)
	}

	return arrayResult
}

func replaceSpaceForPlus(text string) string {
	str := text
	result := strings.Replace(str, " ", "+", -1)
	return result
}
