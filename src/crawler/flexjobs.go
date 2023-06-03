package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"looking-for-remote-jobs/src/model"
	"looking-for-remote-jobs/src/util"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

// const (
// 	baseUrl = "https://www.flexjobs.com/"
// )

func GetOpportunitiesFlexJobs(job string) string {
	space := regexp.MustCompile(`\s+`)
	vagas := []model.Opportunity{}
	c := colly.NewCollector(
		colly.AllowedDomains("flexjobs.com", "www.flexjobs.com"),
	)

	c.OnHTML("li.m-0", func(element *colly.HTMLElement) {
		opportunities := element.DOM

		vaga := model.Opportunity{
			Title:       space.ReplaceAllString(strings.TrimSpace(opportunities.Find("a.job-link").Text()), " "),
			Description: space.ReplaceAllString(strings.TrimSpace(opportunities.Find("div.job-description").Text()), " "),
			Date:        "",
			Url:         opportunities.Find("a.job-link").AttrOr("href", " "),
		}
		vagas = append(vagas, vaga)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.Visit("https://www.flexjobs.com/search?search=.net+developer&location=&srt=date")

	writeJSON(vagas)
	return ""
}

func writeJSON(data []model.Opportunity) {
	f, err := json.MarshalIndent(data, "", " ")
	util.CheckError(err)

	_ = ioutil.WriteFile("file-test.json", f, 0644)
}
