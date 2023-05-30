package crawler

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// type Opportunity struct {
// 	Title       string `json:"title"`
// 	Description string `json:"description"`
// 	Date        string `json:"date"`
// 	Url         string `json:"url"`
// }

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func GetOpportunitiesFlexJobs() string {
	url := "https://www.flexjobs.com/search?search=.net+developer&location=&srt=date"

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

// func GetOpportunitiesFlexJobs(text string) string {
// 	opportunity := Opportunity{}
// 	opportunities := make([]Opportunity, 0)

// 	newCol := colly.NewCollector(
// 		colly.AllowedDomains("www.flexjobs.com/search?search=.net+developer&location=&srt=date", "https://www.flexjobs.com/search?search=.net+developer&location=&srt=date"),
// 	)

// 	newCol.OnHTML(".tab_item_name", func(element *colly.HTMLElement) {
// 		opportunity.Title = element.Text
// 		opportunities = append(opportunities, opportunity)
// 	})

// 	newCol.OnHTML(".discount_final_price", func(element *colly.HTMLElement) {
// 		opportunity.Description = element.Text
// 		opportunities = append(opportunities, opportunity)
// 	})

// 	newCol.OnHTML(".discount_pct", func(element *colly.HTMLElement) {
// 		opportunity.Date = element.Text
// 		opportunities = append(opportunities, opportunity)
// 	})

// 	newCol.OnHTML(".discount_original_price", func(element *colly.HTMLElement) {
// 		opportunity.Url = element.Text
// 		opportunities = append(opportunities, opportunity)
// 	})

// 	newCol.Visit("https://www.flexjobs.com/search?search=.net+developer&location=&srt=date")

// 	newCol.OnScraped(func(r *colly.Response) {
// 		opportunities = append(opportunities, opportunity)
// 		opportunity = Opportunity{}
// 	})

// 	encode := json.NewEncoder(os.Stdout)
// 	encode.SetIndent("", " ")
// 	encode.Encode(opportunities)
// }
