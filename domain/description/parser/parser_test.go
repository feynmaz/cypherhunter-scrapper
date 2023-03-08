package parser

import (
	"os"
	"testing"

	"github.com/feynmaz/cypherhunter-scrapper/domain/description"
	"github.com/stretchr/testify/assert"
)

func TestDescriptionRepo_GetDescriptionFromHTML_Ok(t *testing.T) {
	data, err := os.ReadFile("test_data/mina.html")
	if err != nil {
		t.Fatalf("failed to read test data file: %v", err)
	}
	html := string(data)
	parser := New()

	descr, err := parser.GetDescriptionFromHTML(html)

	assert.NoError(t, err)
	assert.Equal(t, descr.GetName(), "Mina")
	assert.Equal(t, descr.GetHomepage(), "https://minaprotocol.com/?utm_source=cypherhunter")
	assert.True(t, len(descr.GetInvestors()) > 0)
}

func TestDescriptionRepo_GetDescriptionFromHTML_Error(t *testing.T) {
	testCases := []struct {
		test        string
		html        string
		expectedErr error
	}{
		{
			test:        "empty html validation",
			html:        "",
			expectedErr: description.ErrEmptyHTML,
		},
		{
			test: "empty homepage description init error",
			html: `
			<body><div><div><main><div><div>
				content
				<h1>Mina</h1>
				no homepage link provided
			</div></div></main></div></div></body>
			`,
			expectedErr: description.ErrFailedCreateDescription,
		},
		{
			test: "no investors error",
			html: `
			<body><div><div><main><div><div>
				content
				<h1>Mina</h1>
				<a href="https://minaprotocol.com/">homepage</a>
			</div></div></main></div></div></body>
			`,
			expectedErr: description.ErrNoInvestors,
		},
	}

	parser := New()

	for _, c := range testCases {
		t.Run(c.test, func(t *testing.T) {
			_, err := parser.GetDescriptionFromHTML(c.html)

			if c.expectedErr != nil {
				assert.ErrorIs(t, err, c.expectedErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
