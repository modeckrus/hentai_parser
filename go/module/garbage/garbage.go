// // To the extent possible under law, the Yawning Angel has waived all copyright
// // and related or neighboring rights to orhttp_example, using the creative
// // commons "cc0" public domain dedication. See LICENSE or
// // <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package garbage

// import (
// 	// Things needed by the actual interface.

// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"net/http/cookiejar"
// 	"net/url"
// 	"os"
// 	"strconv"
// 	"strings"

// 	// Things needed by the example code.
// 	"fmt"
// 	"io/ioutil"

// 	"hentai_parser/go/module/geziyor"
// 	"hentai_parser/go/module/geziyor/client"
// 	"hentai_parser/go/module/geziyor/export"

// 	"github.com/PuerkitoBio/goquery"
// 	"github.com/pkg/errors"
// 	"golang.org/x/net/proxy"
// )

// const baseURL = "http://wwv.allhen.me"

// func worker(id int, jobs chan Chapeter, result chan<- Chapeter) {
// 	for chapter := range jobs {
// 		log.Println("Worker ", id, "start job")
// 		url := baseURL + chapter.URL

// 		g := geziyor.NewGeziyor(&geziyor.Options{
// 			StartURLs: []string{url},
// 			AllowedDomains: []string{
// 				"c.aaa200.rocks",
// 				"wwv.allhen.me",
// 			},

// 			StartRequestsFunc: func(g *geziyor.Geziyor) {
// 				g.Opt.AllowedDomains = []string{
// 					"c.aaa200.rocks",
// 					"wwv.allhen.me",
// 				}

// 				g.GetRendered(url, func(g *geziyor.Geziyor, r *client.Response) {
// 					if !r.IsHTML() {
// 						log.Println("Not Html")
// 						jobs <- chapter
// 						return
// 					}
// 					pageCountS := r.HTMLDoc.Find("span.pages-count").First().Text()
// 					pageCount, err := strconv.Atoi(pageCountS)
// 					if err != nil {
// 						log.Println(err)
// 						jobs <- chapter
// 						return
// 					}
// 					log.Println("pageCount: ", pageCountS)
// 					images := make([]Image, pageCount)
// 					image, err := getImage(0, g, r)
// 					if err != nil {
// 						log.Println(err)
// 						jobs <- chapter
// 						return
// 					}
// 					images[0] = *image
// 					for i := 1; i < pageCount; i++ {
// 						url := image.URL
// 						// url = strings.Replace(url, "//", "http://", 1)
// 						// url = strings.Split(url, "?")[0]
// 						log.Println(url)
// 						//http://c.aaa200.rocks/auto/03/32/44/002.png_res.jpg
// 						urls := strings.Split(url, "/")
// 						ourls := strings.Split(urls[len(urls)-1], ".")
// 						ourls[0] = "00" + strconv.Itoa(i+1)
// 						urls[len(urls)-1] = strings.Join(ourls, ".")
// 						nurl := strings.Join(urls, "/")
// 						nimage := Image{
// 							URL: nurl,
// 						}
// 						log.Println("New Image: ", nimage)
// 						images[i] = nimage
// 					}
// 					chapter.Vols.PageCount = pageCount
// 					chapter.Vols.Text = chapter.Text
// 					chapter.Vols.Images = images
// 					result <- chapter
// 				})

// 			},
// 			ParseFunc:       chapterParse,
// 			Exporters:       []export.Exporter{&export.JSON{}},
// 			CookiesDisabled: true,
// 		})
// 		cjar, err := cookiejar.New(&cookiejar.Options{})
// 		if err != nil {
// 			log.Println(err)
// 			jobs <- chapter
// 			return
// 		}
// 		g.Client.Jar = cjar
// 		g.Start()
// 		log.Println("Worker ", id, "Finished Job")
// 	}
// }

// // if !r.IsHTML() {
// // 	return
// // }
// // pageCountS := r.HTMLDoc.Find("span.pages-count").First().Text()
// // pageCount, err := strconv.Atoi(pageCountS)
// // if err != nil {
// // 	return
// // }
// // log.Println("pageCount: ", pageCountS)
// // images := make([]Image, pageCount)
// // image, err := getImage(0, g, r)
// // if err != nil {
// // 	log.Println(err)
// // 	return
// // }
// // images[0] = *image
// // for i := 1; i < pageCount; i++ {
// // 	url := image.URL
// // 	// url = strings.Replace(url, "//", "http://", 1)
// // 	// url = strings.Split(url, "?")[0]
// // 	log.Println(url)
// // 	//http://c.aaa200.rocks/auto/03/32/44/002.png_res.jpg
// // 	urls := strings.Split(url, "/")
// // 	ourls := strings.Split(urls[len(urls)-1], ".")
// // 	ourls[0] = "00" + strconv.Itoa(i+1)
// // 	urls[len(urls)-1] = strings.Join(ourls, ".")
// // 	nurl := strings.Join(urls, "/")
// // 	nimage := Image{
// // 		URL: nurl,
// // 	}
// // 	log.Println("New Image: ", nimage)
// // 	images[i] = nimage
// // }
// // chapter.Vols.PageCount = pageCount
// // chapter.Vols.Text = chapter.Text
// // chapter.Vols.Images = images
// // result <- chapter

// // SetDefaultHeader sets header if not exists before
// func SetDefaultHeader(header http.Header, key string, value string) http.Header {
// 	if header.Get(key) == "" {
// 		header.Set(key, value)
// 	}
// 	return header
// }

// // ConvertHeaderToMap converts http.Header to map[string]interface{}
// func ConvertHeaderToMap(header http.Header) map[string]interface{} {
// 	m := make(map[string]interface{})
// 	for key, values := range header {
// 		for _, value := range values {
// 			m[key] = value
// 		}
// 	}
// 	return m
// }

// // ConvertMapToHeader converts map[string]interface{} to http.Header
// func ConvertMapToHeader(m map[string]interface{}) http.Header {
// 	header := http.Header{}
// 	for k, v := range m {
// 		header.Set(k, v.(string))
// 	}
// 	return header
// }

// // NewRedirectionHandler returns maximum allowed redirection function with provided maxRedirect
// func NewRedirectionHandler(maxRedirect int) func(req *http.Request, via []*http.Request) error {
// 	return func(req *http.Request, via []*http.Request) error {
// 		if len(via) >= maxRedirect {
// 			return errors.Errorf("stopped after %d redirects", maxRedirect)
// 		}
// 		return nil
// 	}
// }

// type Clinet struct {
// 	Client *http.Client
// }

// func main() {

// 	parseHentai()
// }

// func parseHentai() {
// 	file, err := os.Open("super.json")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	b, err := ioutil.ReadAll(file)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	chapters := make([]Chapeter, 0)
// 	json.Unmarshal(b, &chapters)
// 	jobs := make(chan Chapeter, len(chapters))
// 	result := make(chan Chapeter, len(chapters))
// 	for w := 1; w <= 5; w++ {
// 		go worker(w, jobs, result)
// 	}
// 	for i := 0; i < len(chapters); i++ {
// 		chapter := chapters[i]
// 		jobs <- chapter
// 	}

// 	for i := 0; i < len(chapters); i++ {
// 		chapters[i] = <-result
// 		log.Println("RESULT: ", chapters[i])
// 	}
// 	close(jobs)
// 	close(result)
// 	chpatersJs, err := json.Marshal(chapters)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	hentaiFile, err := os.OpenFile("hentai.json", os.O_RDWR|os.O_CREATE, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	written, err := hentaiFile.Write(chpatersJs)
// 	if err != nil {
// 		log.Fatal(err)

// 	}
// 	log.Println("Succes; Written: ", written)
// }
// func chapterParse(g *geziyor.Geziyor, r *client.Response) {

// 	pageCountS := r.HTMLDoc.Find("span.pages-count").First().Text()
// 	pageCount, err := strconv.Atoi(pageCountS)
// 	if err != nil {
// 		return
// 	}
// 	log.Println("pageCount: ", pageCountS)
// 	images := make([]Image, pageCount)
// 	image, err := getImage(0, g, r)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	images[0] = *image
// 	for i := 1; i < pageCount; i++ {
// 		url := image.URL
// 		// url = strings.Replace(url, "//", "http://", 1)
// 		// url = strings.Split(url, "?")[0]
// 		log.Println(url)
// 		//http://c.aaa200.rocks/auto/03/32/44/002.png_res.jpg
// 		urls := strings.Split(url, "/")
// 		ourls := strings.Split(urls[len(urls)-1], ".")
// 		ourls[0] = "00" + strconv.Itoa(i+1)
// 		urls[len(urls)-1] = strings.Join(ourls, ".")
// 		nurl := strings.Join(urls, "/")
// 		nimage := Image{
// 			URL: nurl,
// 		}
// 		log.Println("New Image: ", nimage)
// 		images[i] = nimage
// 	}
// }

// func getImage(index int, g *geziyor.Geziyor, r *client.Response) (*Image, error) {
// 	img := r.HTMLDoc.Find(".manga-img_0").First()
// 	iurl := img.AttrOr("src", "")
// 	iurl = strings.Replace(iurl, "//", "http://", 1)
// 	iurl = strings.Split(iurl, "?")[0]
// 	page := Image{
// 		URL: iurl,
// 	}
// 	pageJs, err := json.Marshal(page)
// 	if err != nil {
// 		return nil, err
// 	}
// 	log.Println("Page:  ", string(pageJs))
// 	return &page, nil
// }

// func getManga() {
// 	jsonS = "["

// 	getPage(baseURL + "/a_pervert_s_daily_life")
// 	file, err := os.OpenFile("super.json", os.O_RDWR, 0777)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	jsonS += "]"
// 	written, err := file.WriteString(jsonS)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println("All success: ", written)
// }

// var jsonS string

// func getPage(murl string) (*os.File, error) {
// 	g := geziyor.NewGeziyor(&geziyor.Options{
// 		StartURLs: []string{murl},
// 		StartRequestsFunc: func(g *geziyor.Geziyor) {
// 			g.GetRendered(murl, g.Opt.ParseFunc)

// 		},
// 		ParseFunc:       quotesParse,
// 		Exporters:       []export.Exporter{&export.JSON{}},
// 		CookiesDisabled: false,
// 	})
// 	// tbProxyURL, err := url.Parse("socks5://69.163.163.241:30765")
// 	tbProxyURL, err := url.Parse("socks5://127.0.0.1:9050")
// 	if err != nil {
// 		log.Printf("Failed to parse proxy URL: %v\n", err)
// 		// wg.Done()
// 		return nil, err
// 	}

// 	// Get a proxy Dialer that will create the connection on our
// 	// behalf via the SOCKS5 proxy.  Specify the authentication
// 	// and re-create the dialer/transport/client if tor's
// 	// IsolateSOCKSAuth is needed.
// 	tbDialer, err := proxy.FromURL(tbProxyURL, proxy.Direct)
// 	if err != nil {
// 		log.Printf("Failed to obtain proxy dialer: %v\n", err)
// 		// wg.Done()
// 		return nil, err
// 	}

// 	// Make a http.Transport that uses the proxy dialer, and a
// 	// http.Client that uses the transport.
// 	tbTransport := &http.Transport{Dial: tbDialer.Dial}
// 	g.Client.Transport = tbTransport
// 	// g.Client.SetCookies("nhentai.net", []*http.Cookie{
// 	// 	{
// 	// 		Name:     "csrftoken",
// 	// 		Value:    "yqoqw2yJztlWkSUVtO4SDEwwzVKoAiHwr4VWqvWMDB2OBmxoqMyDDr1snNCrzECj",
// 	// 		Domain:   "nhentai.net",
// 	// 		HttpOnly: false,
// 	// 	},
// 	// })
// 	// g.Client.SetCookies(".nhentai.net", []*http.Cookie{
// 	// 	{
// 	// 		Name:     "__cfduid",
// 	// 		Value:    "d801b7eb8003ea4356e533d23a61285da1612528519",
// 	// 		Domain:   ".nhentai.net",
// 	// 		HttpOnly: true,
// 	// 	},
// 	// })
// 	g.Start()
// 	return nil, nil
// }

// func quotesParse(g *geziyor.Geziyor, r *client.Response) {
// 	r.HTMLDoc.Find("tbody").Each(func(i int, s *goquery.Selection) {
// 		g.Exports <- map[string]interface{}{
// 			"text": s.Find("tr td a").Text(),
// 			"url":  s.Find("tr td a").AttrOr("href", ""),
// 		}
// 		s.Find("tr").Each(func(i2 int, s2 *goquery.Selection) {
// 			s2.Find("td").Each(func(i3 int, s3 *goquery.Selection) {
// 				sel := s3.Find("a")
// 				if sel.Text() == "" {
// 					return
// 				}
// 				text := sel.Text()
// 				text = strings.ReplaceAll(text, "\n", "")
// 				text = strings.ReplaceAll(text, " ", "")
// 				u := map[string]interface{}{
// 					"text": text,
// 					"url":  sel.AttrOr("href", ""),
// 				}
// 				js, err := json.Marshal(u)
// 				if err != nil {
// 					return
// 				}
// 				jss := string(js)
// 				jsonS += jss + ","
// 			})
// 		})

// 	})
// 	// r.HTMLDoc.Find("div.quote").Each(func(i int, s *goquery.Selection) {
// 	// 	g.Exports <- map[string]interface{}{
// 	// 		"text":   s.Find("span.text").Text(),
// 	// 		"author": s.Find("small.author").Text(),
// 	// 	}
// 	// })
// 	// if href, ok := r.HTMLDoc.Find("li.next > a").Attr("href"); ok {
// 	// 	g.Get(r.JoinURL(href), quotesParse)
// 	// }
// 	file, err := os.OpenFile("index.html", os.O_RDWR, 0777)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	written, err := file.Write(r.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println("Succes: ", written)

// }

// // func main() {
// // 	// wg := sync.WaitGroup{}
// // 	// for i := 1; i <= 24; i++ {
// // 	// 	go func(index int) {
// // 	// 		wg.Add(1)
// // 	// 		url := fmt.Sprintf("https://i.nhentai.net/galleries/1840561/%v.jpg", index)
// // 	// 		file, err := DownloadImage(url, "hentai")
// // 	// 		if err != nil {
// // 	// 			log.Fatal(err)
// // 	// 		}
// // 	// 		log.Println("Filename: ", file.Name())
// // 	// 		wg.Done()
// // 	// 	}(i)
// // 	// }
// // 	// wg.Wait()
// // 	file, err := DownloadPage("https://nhentai.net/g/346564/", "hentai")
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// // 	log.Println(file.Name())
// // }

// //DownloadPage ...
// func DownloadPage(addr string, prefix string) (*os.File, error) {
// 	// tbProxyURL, err := url.Parse("socks5://69.163.163.241:30765")
// 	tbProxyURL, err := url.Parse("socks5://127.0.0.1:9050")
// 	if err != nil {
// 		log.Printf("Failed to parse proxy URL: %v\n", err)
// 		// wg.Done()
// 		return nil, err
// 	}

// 	// Get a proxy Dialer that will create the connection on our
// 	// behalf via the SOCKS5 proxy.  Specify the authentication
// 	// and re-create the dialer/transport/client if tor's
// 	// IsolateSOCKSAuth is needed.
// 	tbDialer, err := proxy.FromURL(tbProxyURL, proxy.Direct)
// 	if err != nil {
// 		log.Printf("Failed to obtain proxy dialer: %v\n", err)
// 		// wg.Done()
// 		return nil, err
// 	}

// 	// Make a http.Transport that uses the proxy dialer, and a
// 	// http.Client that uses the transport.
// 	tbTransport := &http.Transport{Dial: tbDialer.Dial}
// 	cjar, err := cookiejar.New(&cookiejar.Options{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	client := &http.Client{
// 		Transport: tbTransport,
// 		Jar:       cjar,
// 	}
// 	baseUrl, err := url.Parse("nhentai.net")
// 	if err != nil {
// 		return nil, err
// 	}
// 	client.Jar.SetCookies(baseUrl, []*http.Cookie{
// 		{
// 			Name:  "csrftoken",
// 			Value: "yqoqw2yJztlWkSUVtO4SDEwwzVKoAiHwr4VWqvWMDB2OBmxoqMyDDr1snNCrzECj",
// 		},
// 	})
// 	dotbaseUrl, err := url.Parse(".nhentai.net")
// 	if err != nil {
// 		return nil, err
// 	}
// 	client.Jar.SetCookies(dotbaseUrl, []*http.Cookie{
// 		{
// 			Name:  "__cfduid",
// 			Value: "d801b7eb8003ea4356e533d23a61285da1612528519",
// 		},
// 	})

// 	// Example: Fetch something.  Real code will probably want to use
// 	// client.Do() so they can change the User-Agent.
// 	resp, err := client.Get(addr)
// 	if err != nil {
// 		log.Printf("Failed to issue GET request: %v\n", err)
// 		// wg.Done()
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	fmt.Printf("GET returned: %v\n", resp.Status)
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Printf("Failed to read the body: %v\n", err)
// 		// wg.Done()
// 		return nil, err
// 	}
// 	//"https://i.nhentai.net/galleries/1840561/1.jpg"
// 	filename := "index.html"
// 	_, err = os.Stat(prefix)
// 	if os.IsNotExist(err) {
// 		log.Println("Creating dir")
// 		err = os.Mkdir(prefix, 0777)
// 		if err != nil {
// 			log.Println("Error while creating folder")
// 			return nil, err
// 		}
// 	} else {
// 		log.Println("Folder exist")
// 	}

// 	file, err := os.Create(prefix + "/" + filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	written, err := file.Write(body)
// 	if err != nil {
// 		log.Println("Error while copying into file")
// 		return nil, err
// 	}
// 	log.Printf("Success; Written: %v", written)
// 	return file, nil
// }

// //DownloadImage used proxy
// func DownloadImage(addr string, prefix string) (*os.File, error) {
// 	// tbProxyURL, err := url.Parse("socks5://69.163.163.241:30765")
// 	tbProxyURL, err := url.Parse("socks5://127.0.0.1:9050")
// 	if err != nil {
// 		log.Printf("Failed to parse proxy URL: %v\n", err)
// 		// wg.Done()
// 		return nil, err
// 	}

// 	// Get a proxy Dialer that will create the connection on our
// 	// behalf via the SOCKS5 proxy.  Specify the authentication
// 	// and re-create the dialer/transport/client if tor's
// 	// IsolateSOCKSAuth is needed.
// 	tbDialer, err := proxy.FromURL(tbProxyURL, proxy.Direct)
// 	if err != nil {
// 		log.Printf("Failed to obtain proxy dialer: %v\n", err)
// 		// wg.Done()
// 		return nil, err
// 	}

// 	// Make a http.Transport that uses the proxy dialer, and a
// 	// http.Client that uses the transport.
// 	tbTransport := &http.Transport{Dial: tbDialer.Dial}
// 	client := &http.Client{Transport: tbTransport}

// 	// Example: Fetch something.  Real code will probably want to use
// 	// client.Do() so they can change the User-Agent.
// 	resp, err := client.Get(addr)
// 	if err != nil {
// 		log.Printf("Failed to issue GET request: %v\n", err)
// 		// wg.Done()
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	fmt.Printf("GET returned: %v\n", resp.Status)
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Printf("Failed to read the body: %v\n", err)
// 		// wg.Done()
// 		return nil, err
// 	}
// 	//"https://i.nhentai.net/galleries/1840561/1.jpg"
// 	sAddr := strings.Split(addr, "/")
// 	filename := sAddr[len(sAddr)-1]
// 	_, err = os.Stat(prefix)
// 	if os.IsNotExist(err) {
// 		log.Println("Creating dir")
// 		err = os.Mkdir(prefix, 0777)
// 		if err != nil {
// 			log.Println("Error while creating folder")
// 			return nil, err
// 		}
// 	} else {
// 		log.Println("Folder exist")
// 	}

// 	file, err := os.Create(prefix + "/" + filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	written, err := file.Write(body)
// 	if err != nil {
// 		log.Println("Error while copying into file")
// 		return nil, err
// 	}
// 	log.Printf("Success; Written: %v", written)
// 	return file, nil
// }
