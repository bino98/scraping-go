package main

import (
	"./app/sites/abehiroshi"
	"./app/sites/hoshinogen"
	"./app/db"
	"./app/model"
	"sync"
)

func main() {
	var apperances []model.Appearance

	funcs := []func() {
		func() { apperances = append(apperances, model.ConvertApperanceFromSiteApperance(abehiroshi.Apperances())...) },
		func() { apperances = append(apperances, model.ConvertApperanceFromSiteApperance(hoshinogen.Apperances())...) },
	}

	var waitGroup sync.WaitGroup

	for _, apperanceFunc := range funcs {
		waitGroup.Add(1)
		go func(function func()) {
			defer waitGroup.Done()
			function()
		}(apperanceFunc)
	}

	waitGroup.Wait()
	
	db := db.Init()
	defer db.Close()

	for _, apperance := range apperances {
		db.AutoMigrate(&model.Appearance{})
		db.Where(
			model.Appearance{
				Name: apperance.Name,
				AppearanceType: apperance.AppearanceType,
				Date: apperance.Date,
				Description: apperance.Description }).FirstOrCreate(&model.Appearance{})
	}
}
