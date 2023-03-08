package page

import (
	"errors"
	"strings"

	"github.com/feynmaz/cypherhunter-scrapper/tools"
)

var (
	ErrInvalidURL = errors.New("page URL is not valid")
)

type Page struct {
	url  string
	html string
}

func New(url string) (Page, error) {
	url = strings.Trim(url, "\t\n ")

	if !tools.IsValidCypherhunterURL(url) {
		return Page{}, ErrInvalidURL
	}

	return Page{url: url}, nil
}

func (p *Page) GetURL() string {
	return p.url
}

func (p *Page) SetHTML(html string) {
	p.html = html
}

func (p *Page) GetHTML() string {
	return p.html
}
