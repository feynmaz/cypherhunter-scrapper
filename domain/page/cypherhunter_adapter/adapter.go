package cypherhunter_adapter

import (
	"fmt"

	"github.com/feynmaz/cypherhunter-scrapper/domain/page"
)

type cypherhunterAdapter struct {
	htmlContentGetter page.HTMLContentGetter
}

func New(htmlContentGetter page.HTMLContentGetter) *cypherhunterAdapter {
	return &cypherhunterAdapter{
		htmlContentGetter: htmlContentGetter,
	}
}

func (a *cypherhunterAdapter) GetPage(url string) (page.Page, error) {
	p, err := page.New(url)
	if err != nil {
		return page.Page{}, fmt.Errorf("%w: %w", page.ErrFailedCreatePage, err)
	}

	content, err := a.htmlContentGetter.GetHTMLContent(p.GetURL())
	if err != nil {
		return page.Page{}, fmt.Errorf("%w: %w", page.ErrFailedGetHTML, err)
	}

	p.SetHTML(content)

	return p, nil
}
