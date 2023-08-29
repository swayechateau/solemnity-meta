package meta

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
