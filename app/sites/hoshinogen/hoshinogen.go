package hoshinogen

import (
	"github.com/PuerkitoBio/goquery"
	"../../site"
)

func Apperances() site.SiteAppearance {
	appearanceData := site.SiteAppearanceData{ Tv: tvAppearanceHoshinogen() }
	return site.SiteAppearance{ Name: "hoshinogen", Data: appearanceData }
}

func tvAppearanceHoshinogen() [][]string {
	appearance := [][]string{}
	doc := site.Scraping("https://www.hoshinogen.com/news/tv/", "utf-8")
	
	doc.Find("#news > div > div > section > article").Each(func(idx int, s *goquery.Selection) {
		date := s.Find("time")
		desc := s.Find("h1")
		appearance = append(appearance, []string{site.Sanitize(date.Text()), site.Sanitize(desc.Text())})
	})

	return appearance
}
