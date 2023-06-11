package crawler

import (
	"fmt"
	"looking-for-remote-jobs/src/model"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

func GetOpportunitiesWeWorkRemotely(job string) []model.Opportunity {
	var baseUrl string = "https://weworkremotely.com"
	space := regexp.MustCompile(`\s+`)
	resultsArray := []model.Opportunity{}

	c := colly.NewCollector(
		colly.AllowedDomains("weworkremotely.com", "www.weworkremotely.com"),
	)

	c.OnHTML("li.feature", func(element *colly.HTMLElement) {
		opportunities := element.DOM

		result := model.Opportunity{
			Title:       space.ReplaceAllString(strings.TrimSpace(opportunities.Find("span.title").Text()), " "),
			Description: space.ReplaceAllString(strings.TrimSpace(opportunities.Find("span.region").Text()), " "),
			Date:        space.ReplaceAllString(strings.TrimSpace(opportunities.Find("span.featured").Text()), " "),
			Url:         baseUrl + opportunities.Find("a").AttrOr("href", " "),
		}
		resultsArray = append(resultsArray, result)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.Visit(baseUrl)

	return resultsArray
}
