package martianpay

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	DefaultAPIURL = "https://api.martianpay.com"
)

// Client represents a MartianPay API client
type Client struct {
	APIKey  string
	BaseURL string
}

// NewClient creates a new MartianPay client
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:  apiKey,
		BaseURL: DefaultAPIURL,
	}
}

// CommonResponse represents the common response structure
type CommonResponse struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}

// sendRequest sends an HTTP request to the MartianPay API
func (c *Client) sendRequest(method, path string, body interface{}, response interface{}) error {
	var bodyReader io.Reader
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("error marshaling request: %v", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	url := fmt.Sprintf("%s%s", c.BaseURL, path)
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	authStr := base64.StdEncoding.EncodeToString([]byte(c.APIKey + ":"))
	req.Header.Set("Authorization", "Basic "+authStr)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	var commonResp CommonResponse
	if err := json.NewDecoder(resp.Body).Decode(&commonResp); err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}

	if commonResp.Code != 0 {
		return fmt.Errorf("API error: %s", commonResp.Msg)
	}

	if err := json.Unmarshal(commonResp.Data, response); err != nil {
		return fmt.Errorf("error unmarshaling data: %v", err)
	}

	return nil
}
