package crawler

import (
	"fmt"
	"looking-for-remote-jobs/src/model"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

func GetOpportunitiesFlexJobs(job string) []model.Opportunity {
	var baseUrl string = "https://www.flexjobs.com"
	space := regexp.MustCompile(`\s+`)
	resultsArray := []model.Opportunity{}

	c := colly.NewCollector(
		colly.AllowedDomains("flexjobs.com", "www.flexjobs.com"),
	)

	c.OnHTML("li.m-0", func(element *colly.HTMLElement) {
		opportunities := element.DOM

		result := model.Opportunity{
			Title:       space.ReplaceAllString(strings.TrimSpace(opportunities.Find("a.job-link").Text()), " "),
			Description: space.ReplaceAllString(strings.TrimSpace(opportunities.Find("div.job-description").Text()), " "),
			Date:        space.ReplaceAllString(strings.TrimSpace(opportunities.Find("div.job-age").Text()), " "),
			Url:         baseUrl + opportunities.Find("a.job-link").AttrOr("href", " "),
		}

		if result.Date == "New! Today" || result.Date == "New! Yesterday" || result.Date == "New! 2 days ago" {
			resultsArray = append(resultsArray, result)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.Visit(baseUrl + "/search?search=" + job + "&location=&srt=date")

	//writeJSON(resultsArray)
	return resultsArray
}

// func writeJSON(data []model.Opportunity) {
// 	f, err := json.MarshalIndent(data, "", " ")
// 	util.CheckError(err)

// 	_ = ioutil.WriteFile("file-test.json", f, 0644)
// }
