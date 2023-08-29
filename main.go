package main

import (
	"encoding/json"
	"fmt"
	"log"
	"meta/meta"
	"meta/site"
	"net/http"
	"strings"
)

type ResponseData struct {
	Link      string            `json:"link"`
	LinkValid bool              `json:"linkValid"`
	All       bool              `json:"all"`
	Content   string            `json:"content,omitempty"`
	Meta      meta.MetaResponse `json:"meta,omitempty"`
}

func main() {
	http.HandleFunc("/api", metaHandler)
	// Serve static files from the "public" directory
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", http.StripPrefix("/", fs))

	port := "8080"
	fmt.Printf("Server is running on port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func metaHandler(w http.ResponseWriter, r *http.Request) {
	website := site.Site{
		// set site url as https (Default: true)
		Secure: isFalse(r.URL.Query().Get("secure")),
		// website url to grab meta data from
		Url: r.URL.Query().Get("link"),
	}
	// list all meta not just concise list (Default: false)
	all := isTrue(r.URL.Query().Get("all"))

	if website.Url == "" {
		http.Error(w, "Missing required 'link' parameter", http.StatusBadRequest)
		return
	}

	linkValid := site.IsValid(website.Url)
	var metaData []meta.Meta

	if linkValid {
		var err error
		err = website.FetchContent()
		if err != nil {
			http.Error(w, "Error fetching website content", http.StatusInternalServerError)
			log.Printf("Error fetching website content: %v", err)
			return
		}
		metaData = meta.GetMetaResponse(website.Content)
	}

	responseData := ResponseData{
		Link:      link,
		LinkValid: linkValid,
		All:       all,
		Meta:      metaData,
	}

	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		log.Printf("Error encoding JSON response: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)

	log.Printf("Response: %s", responseData)
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
