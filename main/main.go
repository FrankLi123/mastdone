package main

import (
	"context"
	"encoding/json"
	"fmt"

	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

// Status represents a status from the Mastodon public timeline
type Status struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

func main() {

	postLimit := 1
	timelineURL := "https://mastodon.social/api/v1/timelines/public"
	url := fmt.Sprintf("%s?limit=%d", timelineURL, postLimit)

	// Create a custom HTTP client with a transport that enforces IPv4
	transport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: false,
			}).DialContext(ctx, "tcp4", addr)
		},
	}
	client := &http.Client{
		Transport: transport,
	}

	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status: %v", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	var statuses []Status
	err = json.Unmarshal(body, &statuses)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	for _, status := range statuses {
		fmt.Printf("ID: %s\nContent: %s\n\n", status.ID, status.Content)
	}
}
