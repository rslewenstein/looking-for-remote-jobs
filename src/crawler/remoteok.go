package crawler

import (
	"fmt"
	"looking-for-remote-jobs/src/model"
	"regexp"
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
			Url:  opportunities.Find("a.preventLink").AttrOr("href", " "),
			Date: space.ReplaceAllString(strings.TrimSpace(opportunities.Find("td.time").Text()), " "),
		}

		resultsArray = append(resultsArray, result)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.Visit(baseUrl + "/remote-" + job + "-jobs?order_by=date")

	return resultsArray
}
