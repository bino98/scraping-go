package main

import (
	"./sites"
	"./db"
	"fmt"
)

func main() {
	abehiroshiData := sites.Abehiroshi()
	dbConnection := db.Init()
	fmt.Println(abehiroshiData)
	fmt.Println(dbConnection)
}
