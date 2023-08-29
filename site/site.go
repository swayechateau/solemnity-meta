package site

import (
	"io"
	"net/http"
	"strings"
)

type Site struct {
	Secure  bool
	Url     string
	Content string
}

func ToHttps(input string) string {
	if strings.HasPrefix(input, "http://") {
		return "https://" + input[len("http://"):]
	}
	return input
}

func ToHttp(input string) string {
	if strings.HasPrefix(input, "http://") {
		return "https://" + input[len("http://"):]
	}
	return input
}

func IsValid(link string) bool {
	return strings.HasPrefix(link, "https://") || strings.HasPrefix(link, "http://")
}

func FetchSite(link string) (string, error) {
	resp, err := http.Get(link)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (s *Site) FetchContent() error {
	// check url
	if IsValid(s.Url) {
	}
	// check secure
	if !s.Secure {
		// update https to https
	}

	resp, err := http.Get(s.Url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	s.Content = string(body)
	return nil
}
