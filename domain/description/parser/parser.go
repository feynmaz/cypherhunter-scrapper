package parser

import (
	"fmt"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/feynmaz/cypherhunter-scrapper/domain/description"
)

type htmlParser struct{}

func New() *htmlParser {
	return &htmlParser{}
}

func (h *htmlParser) GetDescriptionFromHTML(html string) (description.Description, error) {
	if html == "" {
		return description.Description{}, description.ErrEmptyHTML
	}

	htmlDoc := soup.HTMLParse(html)
	main := htmlDoc.Find("body").Find("div").Find("div").Find("main")
	content := main.Find("div").Find("div")

	name := content.Find("h1").Text()
	homepage := ""
	for _, child := range content.Children() {
		if child.NodeValue == "a" {
			homepage = child.Attrs()["href"]
		}
	}
	descr, err := description.New(name, homepage)
	if err != nil {
		return description.Description{}, fmt.Errorf("%w: %w", description.ErrFailedCreateDescription, err)
	}

	investorsRecords := getInvestorsRecords(content)
	if investorsRecords == nil {
		return description.Description{}, description.ErrNoInvestors
	}
	for _, investorRecord := range investorsRecords {
		name := investorRecord.Attrs()["title"]
		// TODO: get actual link
		homepage := investorRecord.Attrs()["href"]

		descr.SetInvestor(name, homepage)
	}

	return descr, nil
}

func getInvestorsRecords(content soup.Root) []soup.Root {
	for _, section := range content.FindAll("section") {
		for _, h2 := range section.FindAll("h2") {
			if strings.Contains(h2.Text(), "Investors") {
				children := section.Children()
				divInvestors := children[len(children)-1]
				return divInvestors.FindAll("a")
			}
		}
	}

	return nil
}
