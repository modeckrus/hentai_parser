package models

//Chapter ...
type Chapter struct {
	PageCount int      `json:"pageCount"`
	Text      string   `json:"text"`
	URL       string   `json:"url"`
	Images    []string `json:"images"`
}
