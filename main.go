package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type ResponseData struct {
	Link      string     `json:"link"`
	LinkValid bool       `json:"linkValid"`
	All       string     `json:"all"`
	Content   string     `json:"content,omitempty"`
	Metadata  []MetaData `json:"metadata,omitempty"`
}

type MetaData struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func main() {

	http.HandleFunc("/api", metaHandler)
	// Serve static files from the "public" directory
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", http.StripPrefix("/", fs))

	http.HandleFunc("/handle", func(w http.ResponseWriter, r *http.Request) {
		link := r.URL.Query().Get("link")
		all := r.URL.Query().Get("all")

		if link == "" {
			http.Error(w, "Missing required 'link' parameter", http.StatusBadRequest)
			return
		}

		linkValid := isLinkValid(link)
		var content string
		var metaData []MetaData

		if linkValid {
			var err error
			content, err = fetchWebsiteContent(link)
			if err != nil {
				http.Error(w, "Error fetching website content", http.StatusInternalServerError)
				log.Printf("Error fetching website content: %v", err)
				return
			}

			metaData = extractMetaData(content)
		}

		responseData := ResponseData{
			Link:      link,
			LinkValid: linkValid,
			All:       all,
			// Content:   content,
			Metadata: metaData,
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

	})

	port := "8080"
	fmt.Printf("Server is running on port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func metaHandler(w http.ResponseWriter, r *http.Request) {
	link := r.URL.Query().Get("link")
	// all := r.URL.Query().Get("all")

	if link == "" {
		http.Error(w, "Missing required 'link' parameter", http.StatusBadRequest)
		return
	}

	resp, err := http.Get("https://" + link)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// return string(body)

	//allFlag := isAllFlagSet(all)

	response := string(body)
	log.Printf("Response: %s", response)
	// fmt.Sprintf("Received link: %s, all: %s", link, allFlag)

	fmt.Fprintf(w, response)
}

func isAllFlagSet(all string) bool {
	switch strings.ToLower(all) {
	case "1", "true", "y", "yes":
		return true
	default:
		return false
	}
}

func isLinkValid(link string) bool {
	return strings.HasPrefix(link, "https://") || strings.HasPrefix(link, "http://")
}

func fetchWebsiteContent(link string) (string, error) {
	resp, err := http.Get(link)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func extractMetaData(content string) []MetaData {
	var metaData []MetaData

	// Use regular expression to match meta tags
	metaTagPattern := `<meta\s+(?:name|property)="([^"]+)"\s+content="([^"]+)"[^>]*>`
	re := regexp.MustCompile(metaTagPattern)
	matches := re.FindAllStringSubmatch(content, -1)

	// Process matched meta tags
	for _, match := range matches {
		nameOrProperty := match[1]
		content := match[2]
		metaData = append(metaData, MetaData{Name: nameOrProperty, Content: content})
	}

	return metaData
}
