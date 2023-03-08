package services

import (
	"strings"
	"testing"

	"github.com/feynmaz/cypherhunter-scrapper/domain/description"
	"github.com/feynmaz/cypherhunter-scrapper/domain/page"
	"github.com/stretchr/testify/assert"
)

type mockPageRepo struct{}

func (m mockPageRepo) GetPage(url string) (page.Page, error) {
	p, err := page.New(url)
	if err != nil {
		return page.Page{}, ErrFailedGetPage
	}
	p.SetHTML("<div>investors</div>")
	if p.GetURL() == "https://www.cypherhunter.com/en/p/no-investors/" {
		p.SetHTML("<div></div>")
	}
	return p, nil
}

type mockDescriptionRepo struct{}

func (m mockDescriptionRepo) GetDescriptionFromHTML(html string) (description.Description, error) {
	if !strings.Contains(html, "investors") {
		return description.Description{}, ErrFailedGetDescription
	}
	name := "mock_name"
	homepage := "https://minaprotocol.com/"
	d, _ := description.New(name, homepage)
	return d, nil
}

func TestScrapper_GetDescription(t *testing.T) {
	scrapper, err := NewScrapper(
		WithPageRepo(mockPageRepo{}),
		WithDescriptionRepo(mockDescriptionRepo{}),
	)
	assert.NoError(t, err)

	testCases := []struct {
		test        string
		url         string
		expectedErr error
	}{
		{
			test:        "failed to get page",
			url:         "ethereum_page.com",
			expectedErr: ErrFailedGetPage,
		},
		{
			test:        "failed to get description",
			url:         "https://www.cypherhunter.com/en/p/no-investors/",
			expectedErr: ErrFailedGetDescription,
		},
		{
			test:        "failed to get description",
			url:         "https://www.cypherhunter.com/en/p/mina/",
			expectedErr: nil,
		},
	}

	for _, c := range testCases {
		t.Run(c.test, func(t *testing.T) {
			descr, err := scrapper.GetDescription(c.url)

			if c.expectedErr != nil {
				assert.ErrorIs(t, err, c.expectedErr)
			} else {
				assert.NotNil(t, descr.GetName())
				assert.NotNil(t, descr.GetHomepage())
			}
		})
	}
}
