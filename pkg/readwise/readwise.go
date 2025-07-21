package readwise

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"
)

type Readwise struct {
	apiKey     string
	logger     *slog.Logger
	httpClient *http.Client
}

type Response struct {
	Results []Result `json:"results"`
}

type Result struct {
	SourceUrl string    `json:"source_url"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Summary   string    `json:"summary"`
	CreatedAt time.Time `json:"created_at"`
	Location  string    `json:"location"`
}

func (r *Readwise) GetByTag(tag string, archived *bool) (Response, error) {
	r.logger.Info("Getting items by tag", "tag", tag)

	req, err := http.NewRequest("GET", "https://readwise.io/api/v3/list/?tag="+tag, nil)
	if err != nil {
		r.logger.Error("Error creating request", "error", err)
		return Response{}, err
	}
	req.Header.Add("Authorization", "Token "+r.apiKey)
	resp, err := r.httpClient.Do(req)
	if err != nil {
		r.logger.Error("Error getting items by tag", "error", err)
		return Response{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		r.logger.Error("Error getting items by tag", "status", resp.StatusCode)
		return Response{}, errors.New("status code " + strconv.Itoa(resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		r.logger.Error("Error reading response body", "error", err)
		return Response{}, err
	}
	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		r.logger.Error("Error decoding response", "error", err)
		return Response{}, err
	}

	sort.Slice(response.Results, func(i, j int) bool {
		return response.Results[i].CreatedAt.Before(response.Results[j].CreatedAt)
	})

	for index, result := range response.Results {
		if archived != nil && *archived {
			fmt.Printf("%d. [%s](%s) - %s\n", index+1, result.Title, result.SourceUrl, result.Location)
		} else {
			fmt.Printf("%d. [%s](%s)\n", index+1, result.Title, result.SourceUrl)
		}
	}

	return response, nil
}

func NewReadwise() *Readwise {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	apiKey := os.Getenv("READWISE_API_KEY")
	if apiKey == "" {
		log.Fatal("READWISE_API_KEY is not set")
	}

	return &Readwise{
		apiKey:     apiKey,
		logger:     logger,
		httpClient: &http.Client{},
	}
}
