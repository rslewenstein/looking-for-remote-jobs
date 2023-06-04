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
	result := replaceSpaceForPlus(job)
	return crawler.GetOpportunitiesFlexJobs(result)
}

func replaceSpaceForPlus(text string) string {
	str := text
	result := strings.Replace(str, " ", "+", -1)
	return result
}
