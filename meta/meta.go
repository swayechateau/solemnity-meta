package meta

import (
	"log"
	"strings"
)

type Meta struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type MetaImage struct {
	Url     string `json:"uri,omitempty"`
	AltText string `json:"alt_text,omitempty"`
	Width   string `json:"width,omitempty"`
	Height  string `json:"height,omitempty"`
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
	Title       string    `json:"title,omitempty"`
	Type        string    `json:"type,omitempty"`
	Description string    `json:"description,omitempty"`
	Keywords    []string  `json:"keywords,omitempty"`
	Image       MetaImage `json:"image,omitempty"`
	Video       MetaVideo `json:"video,omitempty"`

	All []Meta `json:"all,omitempty"`
}

func GetMetaResponse(content string, all bool) (MetaResponse, error) {
	var response MetaResponse
	metaTags, err := ExtractMeta(content)

	if err != nil {
		return response, err
	}

	if all {
		response.All = metaTags
	}

	for _, meta := range metaTags {
		switch meta.Name {
		case "keywords":
			response.Keywords = SplitKeywords(meta.Content)
		case "description":
			fallthrough
		case "og:description":
			response.Description = meta.Content
		case "title":
			fallthrough
		case "og:title":
			response.Title = meta.Content
		case "og:site_name":
			response.SiteName = meta.Content
		case "og:url":
			response.Url = meta.Content
		case "og:locale":
			response.Locale = meta.Content
		case "Type":
			response.Type = meta.Content
		case "og:image":
			response.Image.Url = meta.Content
		case "twitter:image:alt":
			response.Image.AltText = meta.Content
		case "og:image:width":
			response.Image.Width = meta.Content
		case "og:image:height":
			response.Image.Height = meta.Content
		case "og:video:url":
			fallthrough
		case "og:video:secure_url":
			response.Video.Url = meta.Content
		case "og:video:type":
			response.Video.Type = meta.Content
		case "og:video:width":
			response.Video.Width = meta.Content
		case "og:video:height":
			response.Video.Height = meta.Content
		case "og:video:tag":
			response.Video.Tags = append(response.Video.Tags, meta.Content)
		}
	}

	return response, nil
}

func ExtractMeta(content string) ([]Meta, error) {
	var metaTags []Meta
	log.Printf("Extracting website content from: %v", content)
	doc, err := NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		return metaTags, err
	}

	doc.Find("meta").Each(func(index int, metaTag *Selection) {
		name, _ := metaTag.Attr("name")
		property, _ := metaTag.Attr("property")
		rel, _ := metaTag.Attr("rel")
		itemprop, _ := metaTag.Attr("itemprop")
		content, _ := metaTag.Attr("content")

		if name != "" || property != "" {
			metaTags = append(metaTags, Meta{Name: name + property + rel + itemprop, Content: content})
		}
	})

	return metaTags, nil
}

func SplitKeywords(keywords string) []string {
	return strings.Split(keywords, ", ")
}
