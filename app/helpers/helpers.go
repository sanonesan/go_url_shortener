package helpers

import (
	"os"
	"strings"
)

func EnforceHTTPProtocol(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}

	return url
}

// RemoveDomainError ...
func RemoveDomainError(url string) bool {
	// basically this functions removes all the commonly found
	// prefixes from URL such as http, https, www
	// then checks of the remaining string is the DOMAIN itself
	if url == os.Getenv("DOMAIN") {
		return false
	}
	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]

	if newURL == os.Getenv("DOMAIN")+":"+os.Getenv("API_PORT") {
		return false
	}
	if newURL == os.Getenv("DOMAIN")+":"+os.Getenv("DB_PORT") {
		return false
	}

	return true
}
