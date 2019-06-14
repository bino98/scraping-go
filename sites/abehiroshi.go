package sites

import (
	"github.com/PuerkitoBio/goquery"
	"../site"
)

func Abehiroshi() site.PersonAppearance {
	appearanceData := site.PersonAppearanceData{ Tv: tvAppearance(), Movie: movieAppearance() }
	return site.PersonAppearance{ Name: "abehiroshi", Data: appearanceData }
}

func tvAppearance() [][]string {
	appearance := [][]string{}
	doc := site.Scraping("http://abehiroshi.la.coocan.jp/tv/tv.htm", "cp932")
	
	doc.Find("body > div > table > tbody > tr").Each(func(idx int, s *goquery.Selection) {
		date := s.Find("td:nth-child(1)")
		desc := s.Find("td:nth-child(2)")

		appearance = append(appearance, []string{site.Sanitize(date.Text()), site.Sanitize(desc.Text())})
	})

	return appearance
}

func movieAppearance() [][]string {
	appearance := [][]string{}
	doc := site.Scraping("http://abehiroshi.la.coocan.jp/movie/eiga.htm", "cp932")
	doc.Find("body > center:nth-child(3) > table > tbody > tr").Each(func(idx int, s *goquery.Selection) {
		date := s.Find("td:nth-child(1)")
		desc := s.Find("td:nth-child(2)")
		appearance = append(appearance, []string{site.Sanitize(date.Text()), site.Sanitize(desc.Text())})
	})
	
	return appearance
}
