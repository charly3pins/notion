package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/charly3pins/notion"
)

const (
	baseURL       = "https://api.notion.com"
	apiVersion    = "v1"
	headerVersion = "2021-05-13"
)

func main() {
	token := os.Getenv("ACCESS_TOKEN")
	if token == "" {
		log.Fatal("ACCESS_TOKEN env var missing")
	}
	cli := notion.Client{
		Config: notion.ClientConfig{
			BaseURL:       baseURL,
			APIVersion:    apiVersion,
			HeaderVersion: headerVersion,
			Token:         token,
		},
		Client: http.Client{
			Timeout: 20 * time.Second,
		},
	}
	resp, err := cli.ListDatabase(nil)
	if err != nil {
		log.Fatal(err)
	}
	enc, err := json.MarshalIndent(resp, "  ", "   ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(enc))
}
