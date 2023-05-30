package crawler

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func GetOpportunitiesFlexJobs(job string) string {
	url := "https://www.flexjobs.com/search?search=" + job + "&location=&srt=date"

	response, err := http.Get(url)
	defer response.Body.Close()
	check(err)

	if response.StatusCode > 400 {
		fmt.Println("Status code:", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	check(err)

	//jobCatergory, err := doc.Find("div.job-category-jobs").Html()
	// jobCatergory, err := doc.Find("div.job-category-jobs").Find("a.job-title").Html()
	jobCatergory, err := doc.Find("div.job-category-jobs").Find("ul.p-0").Html()

	fmt.Println(jobCatergory)
	return jobCatergory
}
