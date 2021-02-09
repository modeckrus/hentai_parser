package middleware

import (
	"bytes"
	"hentai_parser/go/module/geziyor/client"
	"log"

	"github.com/PuerkitoBio/goquery"
)

// ParseHTML parses response if response is HTML
type ParseHTML struct {
	ParseHTMLDisabled bool
}

func (p *ParseHTML) ProcessResponse(r *client.Response) {
	if !p.ParseHTMLDisabled && r.IsHTML() {
		doc, err := goquery.NewDocumentFromReader(bytes.NewReader(r.Body))
		if err != nil {
			log.Println(err.Error())
			return
		}
		r.HTMLDoc = doc
	}
}
