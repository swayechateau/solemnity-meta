# Notes

short json (all is false)

```json
{
    "site_name" : "",
    "theme_colour":"",
    "locale": "",
    "url":"",
    "title":"",
    "type":"",
    "description": "",
    "keywords":[],
    "image" : {
		"url": "",
		"alt_text":"",
        "height":"",
		"width":""
	},
	"video" : {
		"url": "",
		"type":"",
		"height":"",
		"width":"",
		"tags":[]
	}
}
```

long json (all is true)
```json
{
    "site_name" : "",
    "theme_colour":"",
    "locale": "",
    "url":"",
    "title":"",
    "type":"",
    "description": "",
    "keywords":[],
    "image" : {
		"url": "",
		"alt_text":""
	},
	"video" : {
		"url": "",
		"type":"",
		"height":"",
		"width":"",
		"tags":[]
	},
	"all":[
		{
			"name":"",
			"content":""
		}
	]
}
```

Added user agent support - google bot was the only one working across all test cases

Add github project image 1280×640px for best display
Add xml support as a return type

Issues
twitter.com - errors wih redirect exceeding 10 - fixed with user agent
twitter.com content not being converted

Added docker-compose, changed dockerfile and moved go code to app directory