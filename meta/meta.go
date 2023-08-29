package meta

import "regexp"

type Meta struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type MetaImage struct {
	Url     string `json:"uri,omitempty"`
	AltText string `json:"alt_text,omitempty"`
}

type MetaVideo struct {
	Url    string   `json:"uri,omitempty"`
	Type   string   `json:"type,omitempty"`
	Width  string   `json:"width,omitempty"`
	Height string   `json:"height,omitempty"`
	Tags   []string `json:"tags,omitempty"`
}

type MetaResponse struct {
	SiteName    string    `json:"site_name,omitempty"`
	Locale      string    `json:"locale,omitempty"`
	Url         string    `json:"url,omitempty"`
	Tilte       string    `json:"titile,omitempty"`
	Type        string    `json:"type,omitempty"`
	Description string    `json:"description,omitempty"`
	Keywords    []string  `json:"keywords,omitempty"`
	Image       MetaImage `json:"image,omitempty"`
	Video       MetaVideo `json:"content,omitempty"`

	All []Meta `json:"all,omitempty"`
}

// func (*Meta) getImage() MetaImage{

// }

// func (*Meta) getVideo() MetaVideo{

// }

func ExtractMeta(content string) []Meta {
	var metaData []Meta

	// Use regular expression to match meta tags
	metaTagPattern := `<meta\s+(?:(?:name|property|rel|itemprop)="([^"]+)")\s+content="([^"]+)"[^>]*>`
	re := regexp.MustCompile(metaTagPattern)
	matches := re.FindAllStringSubmatch(content, -1)

	// Process matched meta tags
	for _, match := range matches {
		attrName := match[1]
		content := match[2]
		metaData = append(metaData, Meta{Name: attrName, Content: content})
	}

	return metaData
}

func GetMetaResponse(content string, all bool) MetaResponse {
	var response MetaResponse
	metaData := ExtractMeta(content)

	return response
}
