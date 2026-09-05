package validator

import (
	"net/url"
)

func ValidateURL(u string) bool {
	parsed, err := url.Parse(u)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return false
	}
	return true
}
