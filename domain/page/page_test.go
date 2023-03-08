package page

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPage_New(t *testing.T) {
	testCases := []struct {
		test        string
		urlIn       string
		urlOut      string
		expectedErr error
	}{
		{
			test:        "invalid url validation",
			urlIn:       "ethereum_page.com",
			expectedErr: ErrInvalidURL,
		},
		{
			test:        "valid url with spaces",
			urlIn:       "  \t  https://www.cypherhunter.com/en/p/ethereum/  \n",
			urlOut:      "https://www.cypherhunter.com/en/p/ethereum/",
			expectedErr: nil,
		},
		{
			test:        "valid url",
			urlIn:       "https://www.cypherhunter.com/en/p/ethereum/",
			urlOut:      "https://www.cypherhunter.com/en/p/ethereum/",
			expectedErr: nil,
		},
	}

	for _, c := range testCases {
		t.Run(c.test, func(t *testing.T) {
			page, err := New(c.urlIn)

			if c.expectedErr != nil {
				assert.ErrorIs(t, err, c.expectedErr)
			} else {
				assert.Equal(t, page.GetURL(), c.urlOut)
				assert.Equal(t, page.GetHTML(), "")
			}
		})
	}
}

func TestPage_SetHTML(t *testing.T) {
	page, err := New("https://www.cypherhunter.com/en/p/ethereum/")
	assert.NoError(t, err)

	testHtml := "<div><h2>Ethereum</h2></div>"

	page.SetHTML(testHtml)
	assert.Equal(t, page.GetHTML(), testHtml)
}
