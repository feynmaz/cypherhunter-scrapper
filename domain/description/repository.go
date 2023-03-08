package description

import "errors"

var (
	ErrEmptyHTML               = errors.New("provided html is empty")
	ErrFailedCreateDescription = errors.New("failed to create description")

	ErrNoInvestors = errors.New("no information about investors")
)

type Repository interface {
	GetDescriptionFromHTML(html string) (Description, error)
}
