package main

import (
	"log"
	"meta/meta"
	"meta/site"
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
	e.GET("/api", getMeta)
	e.Start(":8080")
}

func getMeta(c echo.Context) error {
	website := site.Site{
		// set site url as https (Default: true)
		Secure: isFalse(c.QueryParam("secure")),
		// website url to grab meta data from
		Url: c.QueryParam("link"),
	}
	all := isTrue(c.QueryParam("all"))

	if website.Url == "" {
		return c.String(http.StatusBadRequest, "Missing required 'link' parameter")
	}
	linkValid := website.IsValidUrl()

	if linkValid {
		err := website.FetchContent()
		if err != nil {
			log.Printf("Error fetching website content: %v", err)
			return c.String(http.StatusInternalServerError, "Error fetching website content")
		}
	}

	metaTags, _ := meta.GetMetaResponse(website.Content, all)

	responseData := ResponseData{
		Link:      website.Url,
		LinkValid: linkValid,
		All:       all,
		Meta:      metaTags,
	}
	return c.JSON(http.StatusOK, responseData)
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
