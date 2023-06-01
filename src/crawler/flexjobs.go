package crawler

import (
	"fmt"
	"net/http"

	"looking-for-remote-jobs/src/util"

	"github.com/PuerkitoBio/goquery"
)

const (
	baseUrl = "https://www.flexjobs.com/"
)

func GetOpportunitiesFlexJobs(job string) string {
	url := baseUrl + "search?search=" + job + "&location=&srt=date"

	response, err := http.Get(url)
	util.CheckError(err)
	defer response.Body.Close()

	if response.StatusCode > 400 {
		fmt.Println("Status code:", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	util.CheckError(err)

	//jobCategory, err := doc.Find("div.job-category-jobs").Html()
	// jobCategory, err := doc.Find("div.job-category-jobs").Find("a.job-title").Html()
	// jobCategory, err := doc.Find("div.job-category-jobs").Find("ul.p-0").Html()
	//jobCategory, err := doc.Find("div.job-category-jobs").Find("li.m-0").Html()

	// var s []string

	// // for _, v := range strings.Split(jobCategory, "data-title=\"") {
	// 	for _, v := range strings.Split(jobCategory, "<li class=") {
	// 	s = append(s, v)
	// 	//fmt.Println(strconv.QuoteRuneToASCII((v)))
	// }
	// fmt.Println(s)

	doc.Find("div.job-category-jobs").Find("ul.p-0").Each(func(index int, item *goquery.Selection) {
		title := item.Find("a.job-title")
		url, _ := item.Find("li").Attr("data-url")
		description := item.Find(".job-description")
		
		teste := []string{title.Text(), description.Text(), url}
		fmt.Println(teste)
	})

	return ""
}
