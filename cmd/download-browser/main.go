package main

import (
	"log"

	"github.com/go-rod/rod"
)

func main() {
	browser := rod.New()
	if err := browser.Connect(); err != nil {
		log.Fatal(err)
	}
	browser.MustClose()
}



