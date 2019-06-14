package site

import (
	// "github.com/microcosm-cc/bluemonday"
	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
	"strings"
	"fmt"
	"net/http"
	"time"

)


type SiteAppearanceData struct {
	Tv [][]string
	Movie [][]string
}

type SiteAppearance struct {
	Name string
	Data SiteAppearanceData
}

func Scraping(url, charset string) (*goquery.Document) {
	doc, err := fetchPage(url, charset, 10)
	
	if err != nil {
		panic("fetchPage failed")
	}
	return doc
}

func Sanitize(origin string) string {
	str := strings.Replace(origin, "\n", "", -1)
	return strings.Replace(str, " ", "", -1)
}

func fetchPage(url string, charset string, reTry int) (*goquery.Document, error) {
	client := &http.Client{Timeout: time.Duration(10*(10-reTry)) * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("res getted failed")
		if reTry == 0 {
				return nil, err
		}
		return fetchPage(url, charset, reTry-1)
	}

	utfBody, err := iconv.NewReader(resp.Body, charset, "utf8")
	if err != nil {
		fmt.Println("converted failed")
		if reTry == 0 {
			return nil, err
		}
		return fetchPage(url, charset, reTry-1)
	}

	doc, err := goquery.NewDocumentFromReader(utfBody)
	
	if err != nil {
		fmt.Println("url scarapping failed")
		if reTry == 0 {
			return nil, err
		}
		return fetchPage(url, charset, reTry-1)
	}

	return doc, err
}