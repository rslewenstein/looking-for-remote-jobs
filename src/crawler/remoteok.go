package crawler

import (
	"fmt"
	"looking-for-remote-jobs/src/model"

	"github.com/gocolly/colly"
)

func GetOpportunitiesHimalayas(job string) []model.Opportunity {
	var baseUrl string = "https://remoteok.com"
	// space := regexp.MustCompile(`\s+`)
	resultsArray := []model.Opportunity{}

	c := colly.NewCollector(
		colly.AllowedDomains("remoteok.com", "www.remoteok.com"),
	)

	c.OnHTML("tbody", func(element *colly.HTMLElement) {
		opportunities := element.DOM

		result := model.Opportunity{
			Url: opportunities.Find("a.preventLink").AttrOr("href", " "),
		}

		resultsArray = append(resultsArray, result)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	// c.OnHTML("li.m-0", func(element *colly.HTMLElement) {
	// 	opportunities := element.DOM

	// 	result := model.Opportunity{
	// 		Title:       space.ReplaceAllString(strings.TrimSpace(opportunities.Find("a.job-link").Text()), " "),
	// 		Description: space.ReplaceAllString(strings.TrimSpace(opportunities.Find("div.job-description").Text()), " "),
	// 		Date:        space.ReplaceAllString(strings.TrimSpace(opportunities.Find("div.job-age").Text()), " "),
	// 		Url:         baseUrl + opportunities.Find("a.job-link").AttrOr("href", " "),
	// 	}

	// 	if result.Date == "New! Today" || result.Date == "New! Yesterday" || result.Date == "New! 2 days ago" {
	// 		resultsArray = append(resultsArray, result)
	// 	}
	// })

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	//c.Visit(baseUrl + "/jobs/" + job)
	c.Visit(baseUrl + "/remote-" + job + "-jobs?order_by=date")

	return resultsArray
}
