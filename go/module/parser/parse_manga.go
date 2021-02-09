package parser

import (
	"hentai_parser/go/module/geziyor"
	"hentai_parser/go/module/geziyor/client"
	"hentai_parser/go/module/models"
	"log"

	"github.com/PuerkitoBio/goquery"
)

//ParseManga with url
func (p *Parser) ParseManga(mangaIndex int, mjobs chan string, mresult chan models.Manga) {
	for url := range mjobs {
		var manga models.Manga
		g := geziyor.NewGeziyor(&geziyor.Options{
			AllowedDomains:  p.AllowedDomains,
			CookiesDisabled: true,
			ErrorFunc: func(g *geziyor.Geziyor, r *client.Request, err error) {
				log.Println("Error while parseManga")
				log.Println(err)
				mjobs <- url
				return
			},
			ParseHTMLDisabled: false,
			ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
				if !r.IsHTML() {

					mjobs <- url
					return
				}
				genres := make([]string, 0)
				hrefs := make([]string, 0)
				r.HTMLDoc.Find(".elementList .elem_genre a").Each(func(i int, s *goquery.Selection) {
					genres = append(genres, s.Text())
				})
				name := ""
				name = r.HTMLDoc.Find("span.name").First().Text()
				r.HTMLDoc.Find("tbody").Each(func(i int, s *goquery.Selection) {

					s.Find("tr").Each(func(i2 int, s2 *goquery.Selection) {
						s2.Find("td").Each(func(i3 int, s3 *goquery.Selection) {
							sel := s3.Find("a")
							if sel.Text() == "" {
								return
							}
							href := sel.AttrOr("href", "")
							hrefs = append(hrefs, href)
						})
					})

				})
				manga = models.Manga{
					Text: name,
					URL:  url,
					Tags: genres,
					// Chapters: chapters,
				}
				log.Println("Manga:")
				log.Println(manga)
				chapters := make([]models.Chapter, len(hrefs))
				jobs := make(chan string, len(hrefs))
				result := make(chan models.Chapter, len(hrefs))
				for i := 1; i <= p.ChromeInstances; i++ {
					go p.ParseChapter(i, jobs, result)
				}
				for i := 0; i < len(hrefs); i++ {
					link := hrefs[i]
					jobs <- link
				}
				for i := 0; i < len(hrefs); i++ {
					chapters[i] = <-result
					log.Println("RESULT: ", chapters[i])
				}
				close(jobs)
				close(result)
				manga = models.Manga{
					Text:     name,
					URL:      url,
					Tags:     genres,
					Chapters: chapters,
				}
			},
			StartRequestsFunc: func(g *geziyor.Geziyor) {
				g.GetRendered(url, func(g *geziyor.Geziyor, r *client.Response) {
					if !r.IsHTML() {
						g.GetRendered(url, g.Opt.ParseFunc)
					}
					g.Opt.ParseFunc(g, r)
				})
			},
		})

		g.Start()
		mresult <- manga
		return
	}
}
