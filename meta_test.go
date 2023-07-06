package meta

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMetaData(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(`
			<html>
				<head>
					<title>Example Website</title>
					<meta name="description" content="This is an example website">
					<meta property="og:image" content="https://example.com/image.jpg">
					<meta name="twitter:url" content="https://twitter.com/example">
				</head>
				<body></body>
			</html>
		`))
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	m := Meta{}
	metaData := m.GetMetaData(server.URL, false)

	expectedTitle := "Example Website"
	if metaData["title"] != expectedTitle {
		t.Errorf("Expected title '%s', but got '%s'", expectedTitle, metaData["title"])
	}

	expectedDescription := "This is an example website"
	if metaData["description"] != expectedDescription {
		t.Errorf("Expected description '%s', but got '%s'", expectedDescription, metaData["description"])
	}

	expectedImage := "https://example.com/image.jpg"
	if metaData["image"] != expectedImage {
		t.Errorf("Expected image '%s', but got '%s'", expectedImage, metaData["image"])
	}

	expectedURL := "https://twitter.com/example"
	if metaData["url"] != expectedURL {
		t.Errorf("Expected URL '%s', but got '%s'", expectedURL, metaData["url"])
	}
}
