package silurlshortener

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/savannahghi/serverutils"
)

// client bundles the configurations needed to interact with url shortener's server
type client struct {
	httpClient *http.Client
	BaseURL    string
	APIKey     string
}

// NewURLShortener initializes url shortener client
func NewURLShortener() (*client, error) {
	baseURL := fmt.Sprintf("https://%s/rest/v3", serverutils.MustGetEnvVar("URL_SHORTENER_DOMAIN"))

	return &client{
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
		BaseURL: baseURL,
		APIKey:  serverutils.MustGetEnvVar("URL_SHORTENER_API_KEY"),
	}, nil
}

// ShortenURL method is used to shorten a given URL
func (s *client) ShortenURL(ctx context.Context, payload *ShortenURLPayload) (*ShortenURLResponse, error) {
	response, err := s.makeRequest(ctx, http.MethodPost, "/short-urls", nil, payload)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d. wanted %d", response.StatusCode, http.StatusOK)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var shortenURLResponse *ShortenURLResponse

	err = json.Unmarshal(data, &shortenURLResponse)
	if err != nil {
		return nil, err
	}

	return shortenURLResponse, nil
}

// makeRequest is used to make a HTTP request to the server
func (c *client) makeRequest(ctx context.Context, method, path string, queryParams url.Values, body interface{}) (*http.Response, error) {
	urlPath := fmt.Sprintf("%s%s", c.BaseURL, path)

	var request *http.Request

	switch method {
	case http.MethodPost:
		encoded, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		payload := bytes.NewBuffer(encoded)

		req, err := http.NewRequestWithContext(ctx, method, urlPath, payload)
		if err != nil {
			return nil, err
		}

		request = req

	default:
		return nil, fmt.Errorf("unsupported http method: %s", method)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Api-Key", c.APIKey)

	if queryParams != nil {
		request.URL.RawQuery = queryParams.Encode()
	}

	return c.httpClient.Do(request)
}
