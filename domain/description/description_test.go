package description

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDescription_New(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		homepage    string
		expectedErr error
	}

	testsCases := []testCase{
		{
			test:        "Empty name validation",
			name:        "",
			homepage:    "https://ethereum.org/en/",
			expectedErr: ErrMissingName,
		},
		{
			test:        "Empty homepage validation",
			name:        "Ethereum",
			homepage:    "",
			expectedErr: ErrMissingHomepage,
		},
		{
			test:        "Homepage bad URL",
			name:        "Ethereum",
			homepage:    "ethereum_site.com",
			expectedErr: ErrInvalidHomepageURL,
		},
		{
			test:        "Valid name and homepage",
			name:        "Ethereum",
			homepage:    "https://ethereum.org/en/",
			expectedErr: nil,
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.test, func(t *testing.T) {
			descr, err := New(tc.name, tc.homepage)

			if tc.expectedErr != nil {
				assert.ErrorIs(t, err, tc.expectedErr)
			} else {
				assert.Equal(t, descr.GetName(), tc.name)
				assert.Equal(t, descr.GetHomepage(), tc.homepage)
				assert.Equal(t, len(descr.GetInvestors()), 0)
			}
		})
	}
}

func TestDescription_String(t *testing.T) {
	m := map[string]any{
		"name":     "Mina",
		"homepage": "https://minaprotocol.com/?utm_source=cypherhunter",
		"investors": []map[string]string{
			{
				"name":     "FTX Ventures",
				"homepage": "",
			},
			{
				"name":     "Finality Capital Partners",
				"homepage": "",
			},
			{
				"name":     "PANTERA Capital",
				"homepage": "",
			},
		},
	}
	jsonBytes, _ := json.Marshal(m)
	expectedString := string(jsonBytes)

	d, err := New("Mina", "https://minaprotocol.com/?utm_source=cypherhunter")
	if err != nil {
		t.Fatalf("failed to init Description: %v", err)
	}

	d.SetInvestor("FTX Ventures", "")
	d.SetInvestor("Finality Capital Partners", "")
	d.SetInvestor("PANTERA Capital", "")

	s := d.String()

	assert.Equal(t, expectedString, s)
}
