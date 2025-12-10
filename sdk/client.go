package martianpay

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
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
	Code      int             `json:"code"`       // Deprecated HTTP-level code
	ErrorCode string          `json:"error_code"` // Business-level error code
	Msg       string          `json:"msg"`        // Error message
	Data      json.RawMessage `json:"data"`       // Response data
}

// getHTTPStatusText returns human-readable text for HTTP status codes
func getHTTPStatusText(statusCode int) string {
	switch statusCode {
	case 400:
		return "Bad Request"
	case 401:
		return "Unauthorized"
	case 402:
		return "Payment Required"
	case 403:
		return "Forbidden"
	case 404:
		return "Not Found"
	case 405:
		return "Method Not Allowed"
	case 409:
		return "Conflict"
	case 422:
		return "Unprocessable Entity"
	case 429:
		return "Too Many Requests"
	case 500:
		return "Internal Server Error"
	case 502:
		return "Bad Gateway"
	case 503:
		return "Service Unavailable"
	case 504:
		return "Gateway Timeout"
	default:
		if statusCode >= 400 && statusCode < 500 {
			return "Client Error"
		} else if statusCode >= 500 {
			return "Server Error"
		}
		return "Unknown Status"
	}
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

	// Check HTTP status code first
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		statusText := getHTTPStatusText(resp.StatusCode)
		if len(bodyBytes) > 0 {
			return fmt.Errorf("HTTP %d %s: %s", resp.StatusCode, statusText, string(bodyBytes))
		}
		return fmt.Errorf("HTTP %d %s", resp.StatusCode, statusText)
	}

	var commonResp CommonResponse
	if err := json.NewDecoder(resp.Body).Decode(&commonResp); err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}

	// Check for business-level errors (error_code field)
	// Skip error check if error_code is empty, "ok", or "success"
	if commonResp.ErrorCode != "" && commonResp.ErrorCode != "ok" && commonResp.ErrorCode != "success" {
		return fmt.Errorf("API error [%s]: %s", commonResp.ErrorCode, commonResp.Msg)
	}

	// Legacy: check deprecated Code field
	if commonResp.Code != 0 {
		return fmt.Errorf("API error: %s", commonResp.Msg)
	}

	if err := json.Unmarshal(commonResp.Data, response); err != nil {
		return fmt.Errorf("error unmarshaling data: %v", err)
	}

	return nil
}

// sendRequestWithQuery sends an HTTP GET request with query parameters
func (c *Client) sendRequestWithQuery(method, path string, params interface{}, response interface{}) error {
	urlStr := fmt.Sprintf("%s%s", c.BaseURL, path)

	// Build query parameters
	if params != nil {
		queryParams := url.Values{}
		v := reflect.ValueOf(params)
		t := reflect.TypeOf(params)

		// Handle pointer types
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
			t = t.Elem()
		}

		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fieldType := t.Field(i)

			// Try form tag first, then fall back to json tag
			tag := fieldType.Tag.Get("form")
			if tag == "" {
				tag = fieldType.Tag.Get("json")
			}

			// Skip if field is empty or nil
			if tag == "" || tag == "-" {
				continue
			}

			// Extract field name from tag (handle comma-separated options)
			fieldName := tag
			if idx := len(tag); idx > 0 {
				if commaIdx := 0; commaIdx < idx {
					for j, c := range tag {
						if c == ',' {
							fieldName = tag[:j]
							break
						}
					}
				}
			}

			// Add non-zero values to query params
			switch field.Kind() {
			case reflect.String:
				if field.String() != "" {
					queryParams.Add(fieldName, field.String())
				}
			case reflect.Int, reflect.Int32, reflect.Int64:
				queryParams.Add(fieldName, strconv.FormatInt(field.Int(), 10))
			case reflect.Bool:
				queryParams.Add(fieldName, strconv.FormatBool(field.Bool()))
			case reflect.Ptr:
				if !field.IsNil() {
					switch field.Elem().Kind() {
					case reflect.String:
						queryParams.Add(fieldName, field.Elem().String())
					case reflect.Bool:
						queryParams.Add(fieldName, strconv.FormatBool(field.Elem().Bool()))
					}
				}
			}
		}

		if len(queryParams) > 0 {
			urlStr = urlStr + "?" + queryParams.Encode()
		}
	}

	req, err := http.NewRequest(method, urlStr, nil)
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

	// Check HTTP status code first
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		statusText := getHTTPStatusText(resp.StatusCode)
		if len(bodyBytes) > 0 {
			return fmt.Errorf("HTTP %d %s: %s", resp.StatusCode, statusText, string(bodyBytes))
		}
		return fmt.Errorf("HTTP %d %s", resp.StatusCode, statusText)
	}

	var commonResp CommonResponse
	if err := json.NewDecoder(resp.Body).Decode(&commonResp); err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}

	// Check for business-level errors (error_code field)
	// Skip error check if error_code is empty, "ok", or "success"
	if commonResp.ErrorCode != "" && commonResp.ErrorCode != "ok" && commonResp.ErrorCode != "success" {
		return fmt.Errorf("API error [%s]: %s", commonResp.ErrorCode, commonResp.Msg)
	}

	// Legacy: check deprecated Code field
	if commonResp.Code != 0 {
		return fmt.Errorf("API error: %s", commonResp.Msg)
	}

	if err := json.Unmarshal(commonResp.Data, response); err != nil {
		return fmt.Errorf("error unmarshaling data: %v", err)
	}

	return nil
}
