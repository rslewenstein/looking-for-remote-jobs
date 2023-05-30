package main

import (
	"fmt"
	"looking-for-remote-jobs/src/crawler"
)

// import (
// 	"fmt"
// 	"log"
// 	"looking-for-remote-jobs/src/config"

// 	"net/http"
// )

// func main() {
// 	config.LoadEnv()

// 	fmt.Printf("Listening on port %d", config.Port)
// 	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil))
// }

type Opportunity struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Url         string `json:"url"`
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	crawler.GetOpportunitiesFlexJobs()
}
