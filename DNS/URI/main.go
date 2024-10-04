package main

import (
	"net/url"
)

type ParsedURL struct {
	protocol string
	username string
	password string
	hostname string
	port     string
	pathname string
	search   string
	hash     string
}

func newParsedURL(urlString string) ParsedURL {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return ParsedURL{}
	}
	pw, _ := parsedURL.User.Password()

	return ParsedURL{
		protocol: parsedURL.Scheme,
		username: parsedURL.User.Username(),
		password: pw,
		hostname: parsedURL.Hostname(),
		port:     parsedURL.Port(),
		pathname: parsedURL.Path,
		search:   parsedURL.RawQuery,
		hash:     parsedURL.Fragment,
	}
}
