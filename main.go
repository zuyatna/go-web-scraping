package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
	"github.com/zuyatna/go-web-scraping/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	loginURL := os.Getenv("LOGIN_URL")
	targetURL := os.Getenv("TARGET_URL")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	var isDummy bool

	if loginURL == "" || targetURL == "" || username == "" || password == "" {
		log.Print("One or more required environment variables are missing.")
		log.Print("Using dummy values for testing purposes.")

		isDummy = true
	} else {
		log.Print("All required environment variables are set.")
		isDummy = false
	}

	configData := make(map[string]string)

	if isDummy {
		dummyServer := server.DummyServer()
		defer dummyServer.Close()

		url := dummyServer.URL
		log.Printf("Dummy server running at: %s", url)

		res, err := http.Get(url)
		if err != nil {
			log.Fatalf("Failed to make GET request: %v", err)
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Fatalf("Failed to close response body: %v", err)
			}
		}(res.Body)

		if res.StatusCode != 200 {
			log.Fatalf("Unexpected status code: %d", res.StatusCode)
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatalf("Failed to parse HTML: %v", err)
		}

		doc.Find(".config-row").Each(func(i int, s *goquery.Selection) {
			// Extract key and value, then store in map
			key := s.Find(".config-key").Text()
			value := s.Find(".config-value").Text()

			if key != "" {
				configData[strings.TrimSpace(key)] = strings.TrimSpace(value)
			}
		})
	}

	fmt.Println("Result:")
	fmt.Println("===============================")
	for k, v := range configData {
		fmt.Printf("%s: %s\n", k, v)
	}
	fmt.Println("===============================")
}
