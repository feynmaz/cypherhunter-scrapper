package cypherhunter_adapter

import (
	"testing"

	"github.com/feynmaz/cypherhunter-scrapper/domain/page"

	"github.com/stretchr/testify/assert"
)

type MockHTMLContentGetter struct{}

func (m MockHTMLContentGetter) GetHTMLContent(url string) (string, error) {
	html := "<div>content</div>"

	if url == "https://www.cypherhunter.com/en/p/not-exists/" {
		return "", page.ErrPageNotExists
	}

	return html, nil
}

func TestPageRepo_GetPage(t *testing.T) {
	mockHTMLContentGetter := MockHTMLContentGetter{}
	cypherhunterAdapter := New(mockHTMLContentGetter)

	testCases := []struct {
		test        string
		url         string
		expectedErr error
	}{
		{
			test:        "fail page URL validation",
			url:         "",
			expectedErr: page.ErrInvalidURL,
		},
		{
			test:        "failed to get HTML content",
			url:         "https://www.cypherhunter.com/en/p/not-exists/",
			expectedErr: page.ErrFailedGetHTML,
		},
		{
			test:        "valid URL",
			url:         "https://www.cypherhunter.com/en/p/ethereum/",
			expectedErr: nil,
		},
	}

	for _, c := range testCases {
		t.Run(c.test, func(t *testing.T) {
			_, err := cypherhunterAdapter.GetPage(c.url)

			if c.expectedErr != nil {
				assert.ErrorIs(t, err, c.expectedErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
