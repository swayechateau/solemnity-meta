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

func (s *Site) IsValidUrl() bool {
	prefix := prefix(s.Secure)
	if s.Url == "" {
		return false
	}

	if !strings.Contains(s.Url, "://") {
		s.Url = prefix + s.Url
		return true
	}

	if !IsUrlSupported(s.Url) {
		return false
	}

	if strings.HasPrefix(s.Url, "http://") && s.Secure {
		s.Url = ToHttps(s.Url)
		return true
	}

	if strings.HasPrefix(s.Url, "https://") && !s.Secure {
		s.Url = ToHttp(s.Url)
		return true
	}

	if strings.HasPrefix(s.Url, "https://") && s.Secure || strings.HasPrefix(s.Url, "http://") && !s.Secure {
		return true
	}

	return false
}

func prefix(secure bool) string {
	if secure {
		return "https://"
	}
	return "http://"
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

func IsUrlSupported(link string) bool {
	return strings.HasPrefix(link, "https://") || strings.HasPrefix(link, "http://")
}

func (s *Site) FetchContent() error {
	client := &http.Client{}

	req, err := http.NewRequest("GET", s.Url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; Googlebot/2.1; +http://www.google.com/bot.html) Chrome/W.X.Y.Z Safari/537.36")

	resp, err := client.Do(req)
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
