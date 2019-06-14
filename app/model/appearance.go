package model

import (
  "../site"
  "github.com/jinzhu/gorm"
)

type Appearance struct {
  gorm.Model
  Name  string
  AppearanceType string
  Date string
  Description string
}

func ConvertApperanceFromSiteApperance(data site.SiteAppearance) (appearances []Appearance) {
  for _, tv := range data.Data.Tv {
    appearance := Appearance{}
    appearance.Name = data.Name
    appearance.AppearanceType = "tv"
    appearance.Date = tv[0]
    appearance.Description = tv[1]
    appearances = append(appearances, appearance)
  }

  for _, movie := range data.Data.Movie {
    appearance := Appearance{}
    appearance.Name = data.Name
    appearance.AppearanceType = "movie"
    appearance.Date = movie[0]
    appearance.Description = movie[1]
    appearances = append(appearances, appearance)
  }

  return appearances
}
