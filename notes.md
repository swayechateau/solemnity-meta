# Notes

short json (all is false)

```json
{
    "site_name" = og:site_name
    "theme_colour" = theme_color
    "locale" = og:locale
    "url" = og:url
    "title" = og:title
    "type" = og:type
    "description" = description || og:description
    "keywords" = keywords // break into an array
    "image" = og:image
    "image_alt" = twitter:image:alt
}
```

long json (all is true)
```json
{
    "site_name" = og:site_name
    "theme_colour" = theme_color
    "locale" = og:locale
    "url" = og:url
    "title" = og:title
    "type" = og:type
    "description" = description || og:description
    "keywords" = keywords // break into an array
    "image" = og:image
    "image_alt" = twitter:image:alt
}
```

Add github project image 1280×640px for best display
Add xml support as a return type
Add a user agent 'Mozilla/5.0 (Windows NT 5.1; rv:31.0) Gecko/20100101 Firefox/31.0'

Issues
twitter.com - errors wih redirect exceeding 10 


		{
			"name": "og:video:url" "og:video:secure_url",
			"content": "https://www.youtube.com/embed/VEDA3rYkENs"
		},
		{
			"name": "og:video:type",
			"content": "text/html"
		},
		{
			"name": "og:video:width",
			"content": "960"
		},
		{
			"name": "og:video:height",
			"content": "720"
		},

		{
			"name": "og:video:tag",
			"content": "Philipp Beesen"
		},
		{
			"name": "og:video:tag",
			"content": "Memories"
		},
		{
			"name": "og:video:tag",
			"content": "Shadow Warrior"
		},
	
 	Url  string `json:"uri,omitempty"`
	Type string `json:"type,omitempty"`