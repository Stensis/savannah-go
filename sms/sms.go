package sms

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"go-service/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func SendSMS(to, message, env string) (*model.AfricasTalkingResponse, error) {
	// Africa's Talking API endpoint
	var apiURL string = os.Getenv("AT_URL")

	// Prepare the data for URL-encoded body
	data := url.Values{}
	data.Set("username", os.Getenv("AT_PROD_USERNAME"))
	data.Set("to", to)
	data.Set("message", message)
	data.Set("enqueue", "1")

	// Create the HTTP POST request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set request headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("apiKey", os.Getenv("AT_PROD_API_KEY"))

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read and print the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Decode the response body
	var smsResponse model.AfricasTalkingResponse
	if err := xml.Unmarshal(body, &smsResponse); err != nil {
		return nil, fmt.Errorf("failed to decode xml response: %v", err)
	}

	return &smsResponse, nil
}
