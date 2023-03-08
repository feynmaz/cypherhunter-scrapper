package getter

import (
	"fmt"
	"io"
	"net/http"

	"github.com/feynmaz/cypherhunter-scrapper/domain/page"
)

type getter struct {
	requester page.RequestDoer
}

func New(requester page.RequestDoer) *getter {
	return &getter{requester: requester}
}

func (g *getter) GetHTMLContent(url string) (string, error) {
	html, err := getWithRequester(url, g.requester)
	return html, err
}

// rewritten from github.com/anaskhan96/soup GetWithClient
func getWithRequester(url string, requester page.RequestDoer) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("%w: %w", page.ErrCreatingGetRequest, err)
	}

	// Perform request
	resp, err := requester.Do(req)
	if err != nil {
		return "", fmt.Errorf("%w: %w", page.ErrFailedDoRequest, err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		bytes, _ := io.ReadAll(resp.Body)
		return string(bytes), nil

	case http.StatusNotFound:
		return "", page.ErrPageNotExists

	default:
		return "", page.ErrResourceNotAvailable
	}
}
