package getter

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/feynmaz/cypherhunter-scrapper/domain/page"
	"github.com/stretchr/testify/assert"
)

type MockRequestDoer struct{}

func (m MockRequestDoer) Do(req *http.Request) (*http.Response, error) {
	if req.URL.Host == "" {
		return nil, fmt.Errorf("mock error")
	}

	response := http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString("<div>content</div>")),
	}

	if req.URL.Path == "/en/p/not-available/" {
		response.StatusCode = http.StatusServiceUnavailable
		return &response, nil
	}

	if req.URL.Path == "/not-exists/" {
		response.StatusCode = http.StatusNotFound
		return &response, nil
	}

	return &response, nil
}

func TestHTMLContentGetter_GetHTMLContent(t *testing.T) {
	mockRequestDoer := MockRequestDoer{}
	getter := New(mockRequestDoer)

	testCases := []struct {
		test        string
		url         string
		expectedErr error
	}{
		{
			test:        "failed to create get request",
			url:         "ethereum_site.com:badport",
			expectedErr: page.ErrCreatingGetRequest,
		},
		{
			test:        "failed to do request",
			url:         "ethereum_site.com",
			expectedErr: page.ErrFailedDoRequest,
		},
		{
			test:        "resource not available",
			url:         "https://www.cypherhunter.com/en/p/not-available/",
			expectedErr: page.ErrResourceNotAvailable,
		},
		{
			test:        "page does not exist",
			url:         "https://www.cypherhunter.com/not-exists/",
			expectedErr: page.ErrPageNotExists,
		},
		{
			test:        "valid url and page",
			url:         "https://www.cypherhunter.com/en/p/ethereum/",
			expectedErr: nil,
		},
	}

	for _, c := range testCases {
		t.Run(c.test, func(t *testing.T) {
			_, err := getter.GetHTMLContent(c.url)

			if c.expectedErr != nil {
				assert.ErrorIs(t, err, c.expectedErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
