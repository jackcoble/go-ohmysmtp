package ohmysmtp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Production endpoint for OhMySMTP
const API_URL = "https://app.ohmysmtp.com/api/v1"

type OhMySMTPClient struct {
	apiKey string
	Client *http.Client
}

// Expected payload when sending to API
type Payload struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Subject  string `json:"subject"`
	TextBody string `json:"textbody"`
}

// Return new OhMySMTP client
func NewClient(apiKey string) *OhMySMTPClient {
	return &OhMySMTPClient{
		apiKey: apiKey,
	}
}

// Send an email through the API
func (c *OhMySMTPClient) Send(payload *Payload) error {
	url := fmt.Sprintf("%s/%s", API_URL, "send")

	// Marshal the payload into JSON format
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Create and configure HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewReader(payloadJSON))
	if err != nil {
		return err
	}

	// Add appropriate headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("OhMySMTP-Server-Token", c.apiKey)

	// Execute the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
