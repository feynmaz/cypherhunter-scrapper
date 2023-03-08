package tools

import (
	netUrl "net/url"
	"regexp"
)

// AI
// IsValidCypherhunterURL checks if url matches pattern ^https://www\.cypherhunter\.com/(\w+)/p/(\w+(-\w+)*)(\/)?$
func IsValidCypherhunterURL(url string) bool {
	// Compile a regular expression that matches the pattern
	re := regexp.MustCompile(`^https://www\.cypherhunter\.com/(\w+)/p/(\w+(-\w+)*)(\/)?$`)

	// Test the url against the regular expression
	return re.MatchString(url)
}

// IsValidURL checks if url is valid
func IsValidURL(url string) bool {
	// Parse the url and check for errors
	u, err := netUrl.Parse(url)
	if err != nil {
		return false
	}
	// Check if the url has a scheme and a host
	return u.Scheme != "" && u.Host != ""
}
