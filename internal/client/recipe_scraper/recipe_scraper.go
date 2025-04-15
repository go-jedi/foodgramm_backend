package recipescraper

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/go-jedi/foodgrammm-backend/config"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/parser"
	recipescraper "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_scraper"
	jsoniter "github.com/json-iterator/go"
)

const defaultURL = "https://food-gpt.codewavestudio.ru"

// Custom error messages
var (
	ErrUnexpectedStatusCode     = errors.New("unexpected status code")
	ErrFailedToCreateRequest    = errors.New("failed to create request")
	ErrFailedToReadResponseBody = errors.New("failed to read response body")
	ErrFailedToMarshalData      = errors.New("failed to marshal data")
	ErrFailedToUnMarshalData    = errors.New("failed to unmarshal data")
	ErrFailedToSendRequest      = errors.New("failed to send request")
)

//go:generate mockery --name=IClient --output=mocks --case=underscore
type IClient interface {
	Get(ctx context.Context, data recipescraper.GetBody) (parser.ParsedRecipe, error)
}

type Client struct {
	httpClient *http.Client
	headers    http.Header
	url        string
}

func New(cfg config.ClientConfig, httpClient *http.Client) (*Client, error) {
	c := &Client{
		httpClient: httpClient,
		url:        cfg.RecipeScraper.URL,
	}

	if err := c.init(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) init() error {
	if c.url == "" {
		c.url = defaultURL
	}

	return nil
}

// SetDefaultHeaders sets default headers for all requests.
func (c *Client) SetDefaultHeaders(headers http.Header) {
	for key, values := range headers {
		for i := range values {
			c.headers.Add(key, values[i])
		}
	}
}

// joinURLs concatenate base URL and path.
func (c *Client) joinURLs(api string) (string, error) {
	parsedBaseURL, err := url.Parse(c.url)
	if err != nil {
		return "", fmt.Errorf("error parsing base URL: %w", err)
	}

	parsedPath, err := url.Parse(api)
	if err != nil {
		return "", fmt.Errorf("error parsing path: %w", err)
	}

	return parsedBaseURL.ResolveReference(parsedPath).String(), nil
}

func (c *Client) Get(ctx context.Context, data recipescraper.GetBody) (parser.ParsedRecipe, error) {
	const api = "/v1/menu/get"

	fullURL, err := c.joinURLs(api)
	if err != nil {
		return parser.ParsedRecipe{}, err
	}

	jsonData, err := jsoniter.Marshal(data)
	if err != nil {
		return parser.ParsedRecipe{}, fmt.Errorf("%w: %v", ErrFailedToMarshalData, err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return parser.ParsedRecipe{}, fmt.Errorf("%w: %v", ErrFailedToCreateRequest, err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return parser.ParsedRecipe{}, fmt.Errorf("%w: %v", ErrFailedToSendRequest, err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("error closing response body: %v", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return parser.ParsedRecipe{}, fmt.Errorf("%w: %v", ErrFailedToReadResponseBody, err)
	}

	if resp.StatusCode >= http.StatusBadRequest {
		return parser.ParsedRecipe{}, fmt.Errorf("%w: %d, response: %s", ErrUnexpectedStatusCode, resp.StatusCode, body)
	}

	var gr parser.ParsedRecipe
	if err := jsoniter.Unmarshal(body, &gr); err != nil {
		return parser.ParsedRecipe{}, fmt.Errorf("%w: %v", ErrFailedToUnMarshalData, err)
	}

	return gr, nil
}
