package meta

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type Meta struct{}

func (m *Meta) GetMetaData(link string, all bool) map[string]interface{} {
	userAgent := "Mozilla/5.0 (Windows NT 5.1; rv:31.0) Gecko/20100101 Firefox/31.0"
	client := &http.Client{}
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return nil
	}
	req.Header.Set("User-Agent", userAgent)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading HTTP response body:", err)
		return nil
	}

	htmlBody := string(body)
	if !strings.HasPrefix(link, "http") {
		link = "http://" + link
	}

	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return nil
	}

	response := make(map[string]interface{})
	response["title"] = getTitle(doc)
	response["website"] = link
	response["meta"] = make(map[string]string)

	if all {
		getAllMetaTags(doc, response["meta"].(map[string]string))
	} else {
		getSpecificMetaTags(doc, response)
	}

	return response
}

func getTitle(doc *html.Node) string {
	var getTitleNode func(*html.Node) string
	getTitleNode = func(n *html.Node) string {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			return n.FirstChild.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			title := getTitleNode(c)
			if title != "" {
				return title
			}
		}
		return ""
	}
	return getTitleNode(doc)
}

func getAllMetaTags(doc *html.Node, metaTags map[string]string) {
	var getMetaTags func(*html.Node)
	getMetaTags = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "meta" {
			for _, attr := range n.Attr {
				if attr.Key == "name" || attr.Key == "rel" || attr.Key == "itemprop" || attr.Key == "property" {
					metaTags[attr.Key] = attr.Val
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			getMetaTags(c)
		}
	}
	getMetaTags(doc)
}

func getSpecificMetaTags(doc *html.Node, response map[string]interface{}) {
	var getDescriptionAndImage func(*html.Node)
	getDescriptionAndImage = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "meta" {
			var attrVal string
			for _, attr := range n.Attr {
				if attr.Key == "content" {
					attrVal = attr.Val
				}
			}
			switch {
			case hasMetaDescription(n):
				response["description"] = attrVal
			case hasMetaImage(n):
				response["image"] = attrVal
			case hasMetaURL(n):
				response["url"] = attrVal
			case n.Attr[0].Key == "name" && n.Attr[0].Val == "twitter:domain":
				response["twitter_domain"] = attrVal
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			getDescriptionAndImage(c)
		}
	}
	getDescriptionAndImage(doc)
}

func hasMetaImage(meta *html.Node) bool {
	for _, attr := range meta.Attr {
		if attr.Key == "itemprop" && attr.Val == "image" ||
			attr.Key == "property" && attr.Val == "og:image" ||
			attr.Key == "name" && (attr.Val == "twitter:image" || attr.Val == "twitter:image:src") {
			return true
		}
	}
	return false
}

func hasMetaDescription(meta *html.Node) bool {
	for _, attr := range meta.Attr {
		if attr.Key == "itemprop" && attr.Val == "description" ||
			attr.Key == "property" && attr.Val == "og:description" ||
			attr.Key == "name" && attr.Val == "twitter:description" {
			return true
		}
	}
	return false
}

func hasMetaURL(meta *html.Node) bool {
	for _, attr := range meta.Attr {
		if attr.Key == "itemprop" && attr.Val == "url" ||
			attr.Key == "property" && attr.Val == "og:url" ||
			attr.Key == "name" && attr.Val == "twitter:url" {
			return true
		}
	}
	return false
}

func main() {
	m := Meta{}
	metaData := m.GetMetaData("example.com", false)
	fmt.Println(metaData)
}