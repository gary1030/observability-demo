package joke

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gary1030/learning-o11y/pkg/log"
	"github.com/gary1030/learning-o11y/pkg/otel"
	otelresty "github.com/gary1030/learning-o11y/pkg/otel-resty"
	"github.com/go-resty/resty/v2"
)

type Joke struct {
	ID        int    `json:"id"`
	Punchline string `json:"punchline"`
	Setup     string `json:"setup"`
	Type      string `json:"type"`
}

type Repository struct {
	httpClient *resty.Client
}

func NewRepository() (*Repository, error) {
	client := otelresty.NewClient()

	return &Repository{client}, nil
}

func (r *Repository) GetRandomJoke(ctx context.Context) (*Joke, error) {
	_, span := otel.StartNewSpan(ctx)
	defer span.End()

	resp, err := r.httpClient.R().
		SetContext(ctx).
		EnableTrace().
		Get("https://official-joke-api.appspot.com/random_joke")

	if err != nil {
		log.Warn("HTTP request failed: %v", err)
		return nil, err
	}

	if resp.StatusCode() != 200 {
		log.Warn("Non-OK HTTP status:", resp.Status())
		return nil, fmt.Errorf("non-OK HTTP status: %s", resp.Status())
	}

	var joke Joke
	if err := json.Unmarshal(resp.Body(), &joke); err != nil {
		log.Warn("Error unmarshalling response:", err)
		return nil, err
	}

	return &joke, nil
}
