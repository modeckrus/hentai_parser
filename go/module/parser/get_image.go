package parser

import (
	"fmt"
	"hentai_parser/go/module/geziyor"
	"hentai_parser/go/module/geziyor/client"
	"strings"
)

func (p *Parser) getImage(g *geziyor.Geziyor, r *client.Response) (string, error) {
	img := r.HTMLDoc.Find(".manga-img_0").First()
	iurl := img.AttrOr("src", "")
	iurl = strings.Replace(iurl, "//", "http://", 1)
	iurl = strings.Split(iurl, "?")[0]
	if iurl == "" {
		return "", fmt.Errorf("Empty image address")
	}
	return iurl, nil
}
