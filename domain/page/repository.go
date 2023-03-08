package page

import (
	"errors"
	"net/http"
)

var (
	ErrFailedCreatePage = errors.New("failed to create page")
	ErrFailedGetHTML    = errors.New("failed to get HTML content")

	ErrPageNotExists        = errors.New("page does not exist")
	ErrCreatingGetRequest   = errors.New("failed to create get request")
	ErrFailedDoRequest      = errors.New("failed to do request")
	ErrResourceNotAvailable = errors.New("resource not available")
)

type (
	Repository interface {
		GetPage(url string) (Page, error)
	}

	HTMLContentGetter interface {
		GetHTMLContent(url string) (string, error)
	}

	RequestDoer interface {
		Do(req *http.Request) (*http.Response, error)
	}
)
