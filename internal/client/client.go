package client

import (
	"net/http"
	"time"

	"github.com/go-jedi/foodgramm_backend/config"
	"github.com/go-jedi/foodgramm_backend/internal/client/openai"
	"github.com/go-jedi/foodgramm_backend/internal/client/payment"
	recipescraper "github.com/go-jedi/foodgramm_backend/internal/client/recipe_scraper"
)

const defaultTimeoutReq = 15 // second

// Client represents an HTTP client.
type Client struct {
	OpenAI        *openai.Client
	Payment       *payment.Client
	RecipeScraper *recipescraper.Client
}

// NewClient creates a new instance of HTTP client with a specified timeout.
func NewClient(cfg config.ClientConfig) (client *Client, err error) {
	c := &Client{}

	httpClient := &http.Client{
		Timeout: time.Duration(cfg.TimeoutReq) * time.Second,
	}

	if cfg.TimeoutReq <= 0 {
		httpClient.Timeout = time.Duration(defaultTimeoutReq) * time.Second
	}

	c.OpenAI, err = openai.New(cfg, httpClient)
	if err != nil {
		return nil, err
	}

	c.Payment, err = payment.New(cfg, httpClient)
	if err != nil {
		return nil, err
	}

	c.RecipeScraper, err = recipescraper.New(cfg, httpClient)
	if err != nil {
		return nil, err
	}

	return c, nil
}
