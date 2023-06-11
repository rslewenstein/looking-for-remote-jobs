package model

type Opportunity struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Url         string `json:"url"`
}
