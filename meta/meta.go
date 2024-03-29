package meta

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Meta struct {
	Name    string `json:"name" xml:"name,attr"`
	Content string `json:"content" xml:"content,attr"`
}

type MetaImage struct {
	Url     string `json:"uri,omitempty" xml:"imageUri,attr,omitempty"`
	AltText string `json:"alt_text,omitempty" xml:"imageAltText,attr,omitempty"`
	Width   string `json:"width,omitempty" xml:"imageWidth,attr,omitempty"`
	Height  string `json:"height,omitempty" xml:"imageHeight,attr,omitempty"`
}

type MetaVideo struct {
	Url    string   `json:"uri,omitempty" xml:"videoUri,attr,omitempty"`
	Type   string   `json:"type,omitempty" xml:"videoType,attr,omitempty"`
	Width  string   `json:"width,omitempty" xml:"videoWidth,attr,omitempty"`
	Height string   `json:"height,omitempty" xml:"videoHeight,attr,omitempty"`
	Tags   []string `json:"tags,omitempty" xml:"videoTags,attr,omitempty"`
}

type MetaResponse struct {
	SiteName    string    `json:"site_name,omitempty" xml:"siteName,attr,omitempty"`
	Locale      string    `json:"locale,omitempty" xml:"locale,attr,omitempty"`
	Url         string    `json:"url,omitempty" xml:"url,attr,omitempty"`
	Title       string    `json:"title,omitempty" xml:"title,attr,omitempty"`
	Type        string    `json:"type,omitempty" xml:"type,attr,omitempty"`
	Description string    `json:"description,omitempty" xml:"description,attr,omitempty"`
	Keywords    []string  `json:"keywords,omitempty" xml:"keywords,attr,omitempty"`
	Image       MetaImage `json:"image,omitempty" xml:"image,attr,omitempty"`
	Video       MetaVideo `json:"video,omitempty" xml:"video,attr,omitempty"`

	All []Meta `json:"all,omitempty" xml:"all,attr,omitempty"`
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
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		return metaTags, err
	}

	doc.Find("meta").Each(func(index int, metaTag *goquery.Selection) {
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

func FilterByName(metaSlice []Meta, requestedName string) Meta {
	var filteredMeta Meta

	for _, meta := range metaSlice {
		if meta.Name == requestedName {
			filteredMeta = meta
		}
	}

	return filteredMeta
}

func FilterByNameSlice(metaSlice []Meta, requestedName string) []Meta {
	var filteredMeta []Meta

	for _, meta := range metaSlice {
		if meta.Name == requestedName {
			filteredMeta = append(filteredMeta, meta)
		}
	}

	return filteredMeta
}
