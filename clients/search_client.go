package clients

import (
	"fmt"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
)

type SearchResult struct {
	Query   string       `json:"query"`
	Total   int          `json:"total"`
	Results []NoteResult `json:"results"`
}

type NoteResult struct {
	ID         uint           `json:"id"`
	Title      string         `json:"title"`
	Content    string         `json:"content"`
	CategoryID uint           `json:"category_id"`
	Category   CategoryResult `json:"category"`
}

type CategoryResult struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type SearchClient struct {
	client  *resty.Client
	baseURL string
}

func NewSearchClient(baseURL string) *SearchClient {
	client := resty.New().
		SetBaseURL(baseURL).
		SetTimeout(5*time.Second).
		SetRetryCount(2).
		SetRetryWaitTime(500*time.Millisecond).
		SetHeader("Content-Type", "application/json")

	client.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		log.Printf("[NoteApp → SearchService] %s %s", req.Method, req.URL)
		return nil
	})

	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		log.Printf("[NoteApp ← SearchService] Status: %d | Time: %v",
			resp.StatusCode(), resp.Time())
		return nil
	})

	return &SearchClient{client: client, baseURL: baseURL}
}
func (sc *SearchClient) Search(query string, categoryID string) (*SearchResult, error) {
	var result SearchResult
	var errResp map[string]string

	req := sc.client.R().
		SetQueryParam("q", query).
		SetResult(&result).
		SetError(&errResp)

	if categoryID != "" {
		req.SetQueryParam("category_id", categoryID)
	}

	resp, err := req.Get("/search")
	if err != nil {
		return nil, fmt.Errorf("search request failed: %w", err)
	}

	if resp.IsError() {
		msg := errResp["error"]
		if msg == "" {
			msg = "unknown error from search service"
		}
		return nil, fmt.Errorf("search service error (%d): %s", resp.StatusCode(), msg)
	}

	return &result, nil
}

func (sc *SearchClient) HealthCheck() bool {
	resp, err := sc.client.R().Get("/health")
	if err != nil || resp.IsError() {
		return false
	}
	return true
}
