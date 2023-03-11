package services

import (
	"errors"
	"fmt"

	"github.com/feynmaz/cypherhunter-scrapper/domain/description"
	"github.com/feynmaz/cypherhunter-scrapper/domain/page"
)

var (
	ErrFailedGetPage = errors.New("failed to get page")

	ErrFailedGetDescription = errors.New("failed to get description")
)

type ScrapperConfiguration func(s *scrapper) error

type scrapper struct {
	pageRepo        page.Repository
	descriptionRepo description.Repository
}

// NewScrapper is a factory creating a new scraper
func NewScrapper(cfgs ...ScrapperConfiguration) (*scrapper, error) {
	s := &scrapper{}

	for _, cfg := range cfgs {
		if err := cfg(s); err != nil {
			return nil, err
		}
	}

	return s, nil
}

func WithPageRepo(pr page.Repository) ScrapperConfiguration {
	return func(s *scrapper) error {
		s.pageRepo = pr
		return nil
	}
}

func WithDescriptionRepo(dr description.Repository) ScrapperConfiguration {
	return func(s *scrapper) error {
		s.descriptionRepo = dr
		return nil
	}
}

func (s *scrapper) GetDescription(url string) (description.Description, error) {
	p, err := s.pageRepo.GetPage(url)
	if err != nil {
		return description.Description{}, fmt.Errorf("%w: %w", ErrFailedGetPage, err)
	}

	descr, err := s.descriptionRepo.GetDescriptionFromHTML(p.GetHTML())
	if err != nil {
		return description.Description{}, fmt.Errorf("%w: %w", ErrFailedGetDescription, err)
	}

	return descr, nil
}
