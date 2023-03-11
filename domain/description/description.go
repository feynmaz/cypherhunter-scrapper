package description

import (
	"encoding/json"
	"errors"

	scrapper "github.com/feynmaz/cypherhunter-scrapper"
	"github.com/feynmaz/cypherhunter-scrapper/tools"
)

var (
	ErrMissingName        = errors.New("missing `name` field")
	ErrMissingHomepage    = errors.New("missing `homepage` field")
	ErrInvalidHomepageURL = errors.New("homepage URL is not valid")
)

// Description is an aggregate that represents crypto project description
type Description struct {
	project   *scrapper.Project
	investors []*scrapper.Investor
}

func New(name, homepage string) (Description, error) {
	if name == "" {
		return Description{}, ErrMissingName
	}
	if homepage == "" {
		return Description{}, ErrMissingHomepage
	}
	if !tools.IsValidURL(homepage) {
		return Description{}, ErrInvalidHomepageURL
	}
	project := &scrapper.Project{
		Name:     name,
		Homepage: homepage,
	}
	investors := []*scrapper.Investor{}

	return Description{
		project:   project,
		investors: investors,
	}, nil
}

func (d *Description) GetName() string {
	return d.project.Name
}

func (d *Description) GetHomepage() string {
	return d.project.Homepage
}

func (d *Description) GetInvestors() []*scrapper.Investor {
	return d.investors
}

func (d *Description) SetInvestor(name, homepage string) {
	d.investors = append(d.investors, &scrapper.Investor{
		Name:     name,
		Homepage: homepage,
	})
}

func (d Description) String() string {
	m := make(map[string]any)
	m["name"] = d.project.Name
	m["homepage"] = d.project.Homepage

	investors := make([]map[string]string, len(d.investors))
	for i, investor := range d.investors {
		investors[i] = map[string]string{
			"name":     investor.Name,
			"homepage": investor.Homepage,
		}
	}
	m["investors"] = investors

	jsonBytes, _ := json.Marshal(m)
	s := string(jsonBytes)
	return s
}
