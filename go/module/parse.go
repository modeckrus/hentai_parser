package parse

import (
	"encoding/json"
	"hentai_parser/go/module/models"
	"hentai_parser/go/module/parser"
	"log"
	"os"
)

func main() {
	p := parser.Parser{
		BaseURL: "http://wwv.allhen.me",
		AllowedDomains: []string{
			"c.aaa200.rocks",
			"wwv.allhen.me",
		},
		ChromeInstances: 10,
		UserDir:         "/home/modeck/.config/google-chrome/Default/",
	}
	hrefs := []string{"http://wwv.allhen.me/a_pervert_s_daily_life"}
	mangas := make([]models.Manga, 1)
	jobs := make(chan string, len(mangas))
	result := make(chan models.Manga, len(mangas))
	for i := 1; i <= p.ChromeInstances; i++ {
		go p.ParseManga(i, jobs, result)
	}
	for i := 0; i < len(mangas); i++ {
		href := hrefs[i]
		jobs <- href
	}
	for i := 0; i < len(hrefs); i++ {
		mangas[i] = <-result
		log.Println("RESULT: ", mangas[i])
	}
	file, err := os.Create("day.json")
	if err != nil {
		log.Fatal(err)
	}
	mangaJs, err := json.Marshal(mangas)
	if err != nil {
		log.Fatal(err)
	}
	written, err := file.Write(mangaJs)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("TOTAL SUCCESS....: ", written)
}
