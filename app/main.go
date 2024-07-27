package main

import (
	"fmt"
	"log"
	"meta/cmd/meta"
	"meta/cmd/site"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ResponseData struct {
	Link      string            `json:"link" validate:"required,url" example:"https://www.example.com" xml:"link"`
	LinkValid bool              `json:"linkValid" xml:"linkValid"`
	All       bool              `json:"all" xml:"displayAll"`
	Content   string            `json:"content,omitempty" xml:"content,omitempty"`
	Meta      meta.MetaResponse `json:"meta,omitempty" xml:"meta,omitempty"`
}

func main() {
	e := echo.New()
	// Middleware to set Access-Control-Allow-Origin header
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			return next(c)
		}
	})

	e.Use(middleware.Static("public"))

	e.GET("/api", getMetaHandler)
	e.POST("/api", postMetaHandler)
	e.Start(":5050")
}

func getMetaHandler(c echo.Context) error {
	website := site.Site{
		// set site url as https (Default: true)
		Secure: isFalse(c.QueryParam("secure")),
		// website url to grab meta data from
		Url: c.QueryParam("link"),
	}
	all := isTrue(c.QueryParam("all"))

	responseData, err := getMeta(website, all)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, responseData)
}

func postMetaHandler(c echo.Context) error {
	website := site.Site{
		// set site url as https (Default: true)
		Secure: isFalse(c.FormValue("secure")),
		// website url to grab meta data from
		Url: c.FormValue("link"),
	}
	all := isTrue(c.FormValue("all"))

	responseData, err := getMeta(website, all)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, responseData)
}

func getMeta(website site.Site, all bool) (ResponseData, error) {
	response := ResponseData{}
	// check if website url is empty
	if website.Url == "" {
		return response, fmt.Errorf("Missing required website url")
	}
	// check if website url is valid
	linkValid := website.IsValidUrl()

	if !linkValid {
		return response, fmt.Errorf("Invalid website url")
	}
	response.LinkValid = linkValid
	response.Link = website.Url
	response.All = all
	// fetch website content
	if err := website.FetchContent(); err != nil {
		log.Printf("Error fetching website content: %v", err)
		return response, fmt.Errorf("Error fetching website content: %v", err)
	}

	metaTags, err := meta.GetMetaResponse(website.Content, all)
	if err != nil {
		log.Printf("Error getting meta tags: %v", err)
		return response, fmt.Errorf("Error getting meta tags: %v", err)
	}
	response.Meta = metaTags

	return response, nil
}

func isFalse(boolean string) bool {
	switch strings.ToLower(boolean) {
	case "0", "false", "n", "no":
		return false
	default:
		return true
	}
}

func isTrue(boolean string) bool {
	switch strings.ToLower(boolean) {
	case "1", "true", "y", "yes":
		return true
	default:
		return false
	}
}
