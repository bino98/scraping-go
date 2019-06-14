package site

import (
	// "github.com/microcosm-cc/bluemonday"
	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
	"strings"
	"fmt"
	"net/http"
)


type PersonAppearanceData struct {
	Tv [][]string
	Movie [][]string
}

type PersonAppearance struct {
	Name string
	Data PersonAppearanceData
}

func Scraping(url, charset string) (*goquery.Document) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("res getted failed")
	}
	defer res.Body.Close()

	utfBody, err := iconv.NewReader(res.Body, charset, "utf8")
	if err != nil {
		fmt.Println("converted failed")
	}

	doc, err := goquery.NewDocumentFromReader(utfBody)
	
	if err != nil {
		fmt.Println(err)
		fmt.Println("url scarapping failed")
	}

	return doc
}

func Sanitize(origin string) string {
	str := strings.Replace(origin, "\n", "", -1)
	return strings.Replace(str, " ", "", -1)
}