package main

import (
	"log"
	"net/http"
	"time"

	"github.com/feynmaz/cypherhunter-scrapper/domain/description/parser"
	"github.com/feynmaz/cypherhunter-scrapper/domain/page/cypherhunter_adapter"
	"github.com/feynmaz/cypherhunter-scrapper/domain/page/getter"
	"github.com/feynmaz/cypherhunter-scrapper/services"
)

func main() {
	url := "https://www.cypherhunter.com/en/p/mina/"

	httpClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 5,
		},
		Timeout: 2 * time.Second,
	}

	htmlContentGetter := getter.New(httpClient)
	cypherhunterAdapter := cypherhunter_adapter.New(htmlContentGetter)

	htmlParser := parser.New()

	scrapper, err := services.NewScrapper(
		services.WithPageRepo(cypherhunterAdapter),
		services.WithDescriptionRepo(htmlParser),
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	description, err := scrapper.GetDescription(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Print(description.String())
}
