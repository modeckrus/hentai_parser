package parser

import (
	"hentai_parser/go/module/geziyor"
	"hentai_parser/go/module/geziyor/client"
	"hentai_parser/go/module/models"
	"log"
	"net/http/cookiejar"
	"strconv"
	"strings"
)

//ParseChapter with url
func (p *Parser) ParseChapter(windex int, jobs chan string, result chan models.Chapter) {
	log.Println("Worker", windex, "; started...")
	for chapterURL := range jobs {
		url := p.BaseURL + chapterURL

		g := geziyor.NewGeziyor(&geziyor.Options{
			StartURLs:      []string{url},
			AllowedDomains: p.AllowedDomains,

			StartRequestsFunc: func(g *geziyor.Geziyor) {
				g.Opt.AllowedDomains = p.AllowedDomains

				g.GetRendered(url, func(g *geziyor.Geziyor, r *client.Response) {
					if !r.IsHTML() {
						log.Println("Not Html")
						jobs <- chapterURL
						return
					}
					pageCountS := r.HTMLDoc.Find("span.pages-count").First().Text()
					pageCount, err := strconv.Atoi(pageCountS)
					if err != nil {
						log.Println(err)
						jobs <- chapterURL
						return
					}
					log.Println("pageCount: ", pageCountS)

					mangaLink := r.HTMLDoc.Find(".manga-link").First()
					text := mangaLink.Text()

					images := make([]string, pageCount)
					image, err := p.getImage(g, r)
					if err != nil {
						log.Println(err)
						jobs <- chapterURL
						return
					}
					images[0] = image
					for i := 1; i < pageCount-1; i++ {
						//http://c.aaa200.rocks/auto/03/32/44/002.png_res.jpg
						urls := strings.Split(image, "/")
						ourls := strings.Split(urls[len(urls)-1], ".")
						ourls[0] = "00" + strconv.Itoa(i+1)
						urls[len(urls)-1] = strings.Join(ourls, ".")
						nurl := strings.Join(urls, "/")
						log.Println("New Image: ", nurl)
						images[i] = nurl
					}
					chapter := models.Chapter{
						PageCount: pageCount,
						Text:      text,
						URL:       chapterURL,
						Images:    images,
					}
					result <- chapter
				})

			},
			CookiesDisabled: true,
		})
		cjar, err := cookiejar.New(&cookiejar.Options{})
		if err != nil {
			log.Println(err)
			jobs <- chapterURL
		}
		g.Client.Jar = cjar
		g.Start()
		log.Println("Worker", windex, "; Stoped...")
	}
}
