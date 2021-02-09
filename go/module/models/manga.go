package models

//Manga ...
type Manga struct {
	Text     string    `json:"text"`
	URL      string    `json:"url"`
	Chapters []Chapter `json:"chapters"`
	Tags     []string  `json:"tags"`
}
