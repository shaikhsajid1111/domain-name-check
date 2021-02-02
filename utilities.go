package main

import (
	"net/http"
	"strings"
)

func checkDomains(domainName string) map[string]bool {

	var URLS [23]string
	var domains = [23]string{".com",".app", ".io", ".live",
		".club", ".rocks", ".social", ".desi", ".buzz",
		".click", ".link", ".news", ".video", ".ninja", ".reviews", ".today",
		".website", ".work", ".academy", ".site", ".online", ".social", ".desi"}
	if !strings.Contains(domainName, "http") && !strings.Contains(domainName, "https") {
		domainName = "http://" + domainName
	}
	for i := range domains {
		
		URLS[i] = domainName + domains[i]
	}

	var results = make(map[string]bool,23) 

	for i := range URLS{
		results[URLS[i]] = domainExists(URLS[i])
	}

	return results
}

func domainExists(url string) bool {
	response, error := http.Get(url)

	if error != nil {
		return false
	}

	status := response.StatusCode

	if status >= 200 && status <= 400 || status >= 100 && status <= 101{
		return true
	}

	return false
}
