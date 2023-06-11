package crawler

import (
	"fmt"
	"looking-for-remote-jobs/src/model"
	"looking-for-remote-jobs/src/util"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func GetOpportunitiesHimalayas(job string) []model.Opportunity {
	var baseUrl string = "https://remoteok.com"
	space := regexp.MustCompile(`\s+`)
	resultsArray := []model.Opportunity{}

	c := colly.NewCollector(
		colly.AllowedDomains("remoteok.com", "www.remoteok.com"),
	)

	c.OnHTML("tr.job", func(element *colly.HTMLElement) {
		opportunities := element.DOM

		result := model.Opportunity{
			Url:  baseUrl + opportunities.Find("a.preventLink").AttrOr("href", " "),
			Date: space.ReplaceAllString(strings.TrimSpace(opportunities.Find("td.time").Text()), " "),
		}

		if strings.Contains(result.Date, "h") || checkLastThreeDays(result.Date) {
			resultsArray = append(resultsArray, result)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	//TODO: I don't know what happened. Until yesterday It was working :/
	// Words such as DEVELOPER or DEV don't is working.
	c.Visit(baseUrl + "/remote-" + job + "-jobs?order_by=date")

	return resultsArray
}

func checkLastThreeDays(text string) bool {
	var results []string
	var checkDay bool = true
	if strings.Contains(text, "d") {
		results = strings.Split(text, "d")
	}

	number, err := strconv.Atoi(results[0])
	util.CheckError(err)
	if number > 3 {
		checkDay = false
	}
	return checkDay
}
