package client

import (
	"net/http"
	"time"

	"github.com/go-jedi/foodgrammm-backend/config"
	"github.com/go-jedi/foodgrammm-backend/internal/client/openai"
	"github.com/go-jedi/foodgrammm-backend/internal/client/payment"
)

const defaultTimeoutReq = 15 // second

// Client represents an HTTP client.
type Client struct {
	OpenAI  *openai.Client
	Payment *payment.Client
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

	c.OpenAI, err = openai.NewOpenAI(cfg, httpClient)
	if err != nil {
		return nil, err
	}

	c.Payment, err = payment.NewClient(cfg, httpClient)
	if err != nil {
		return nil, err
	}

	return c, nil
}
