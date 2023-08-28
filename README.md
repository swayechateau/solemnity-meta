Solemnity Meta Api
=======

Rebuilt in Go

Meta Grabber gets meta data from a website for creating sharable cards.


Installation
---------------

Meta Grabber requires --`Go` v1.2+-- to run.

As this API is built using [go](https://go.dev), please run the following command to install.

### Required Before Install

 - [go](https://go.dev)
 - [git](https://git-scm.com/downloads/)

```sh
$ git clone https://github.com/swayechateau/solemnity-meta.git
$ cd solemnity-meta
$ go mod download
```

Development 

```sh
$ go run main.go
```

Production 
Either host on a server with go support or use one with docker support `will add instruction later`

How it works
------------------
There are two ways to get a websites meta data with meta grabber by using either POST or GET Request


The follow fileds are supported.

| Field | Type | Required | Default |
| ----- | ---- | -------- | ------- |
| link | String | Yes | null |
| all | Boolean | No | false |


For instance a GET Request

```sh
curl http://localhost:8080/?link=https://www.youtube.com/watch?v=jgbVa274m9k&all=true
```

And a POST Request

```sh
curl -X POST -d 'link=facbook.com' -d 'all=true' http://localhost:8080/
```

### Api Response Examples

Full

``` json
{
    "title":"Facebook \u2013 log in or sign up",
    "url":"http:\/\/facebook.com",
    "meta": { 
        "referrer":"default",
        "og:site_name":"Facebook",
        "og:url":"https:\/\/www.facebook.com\/",
        "og:image":"https:\/\/www.facebook.com\/images\/fb_icon_325x325.png",
        "og:locale":"en_GB",
        "description":"Create an account or log in to Facebook. Connect with friends, family and other people you know. Share photos and videos, send messages and get updates.",
        "robots":"noodp,noydir"
    }
}
```

Short

``` json
{
    "title":"Facebook \u2013 log in or sign up",
    "url":"https:\/\/www.facebook.com\/",
    "image":"https:\/\/www.facebook.com\/images\/fb_icon_325x325.png",
    "description":"Create an account or log in to Facebook. Connect with friends, family and other people you know. Share photos and videos, send messages and get updates."
}
```