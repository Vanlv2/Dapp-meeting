package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// CloudflareStreamResponse định nghĩa cấu trúc phản hồi từ Cloudflare Stream API
type CloudflareStreamResponse struct {
	Success  bool                   `json:"success"`
	Result   CloudflareStreamResult `json:"result"`
	Errors   []interface{}          `json:"errors"`
	Messages []interface{}          `json:"messages"`
}

type CloudflareStreamResult struct {
	UID   string `json:"uid"`
	RTMPS struct {
		UrlNew string `json:"urlNew"`
	} `json:"rtmps"`
	Meta struct {
		MeetingUrl string `json:"meetingUrl"`
	} `json:"meta"`
}

// CloudflareService cung cấp các phương thức tương tác với Cloudflare Stream API.
type CloudflareService struct {
	AccountID string
	APIToken  string
}

// NewCloudflareService khởi tạo một instance mới của CloudflareService.
func NewCloudflareService(accountID, apiToken string) *CloudflareService {
	return &CloudflareService{
		AccountID: accountID,
		APIToken:  apiToken,
	}
}

// CreateLiveInput gửi request tới Cloudflare Stream API để tạo live input và trả về URL tối ưu (RTMPS URL).
func (s *CloudflareService) CreateLiveInput(meetingUrl string) (string, error) {
	// Endpoint của Cloudflare Stream API để tạo live input.
	apiURL := fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/stream/live_inputs", s.AccountID)

	// Chuẩn bị payload request theo định dạng của Cloudflare.
	payload := map[string]interface{}{
		"meta": map[string]string{
			"meeting_url": meetingUrl,
		},
		"recording": map[string]string{
			"mode": "automatic",
		},
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %w", err)
	}

	// Tạo request HTTP POST.
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create http request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.APIToken))

	// Gửi request với timeout.
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send http request: %w", err)
	}
	defer resp.Body.Close()

	// Kiểm tra status code của phản hồi.
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("cloudflare stream api error: status %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	// Giải mã phản hồi JSON.
	var cfResp CloudflareStreamResponse
	if err := json.NewDecoder(resp.Body).Decode(&cfResp); err != nil {
		return "", fmt.Errorf("failed to decode cloudflare response: %w", err)
	}

	if !cfResp.Success {
		return "", fmt.Errorf("cloudflare stream api returned errors: %v", cfResp.Errors)
	}

	// Trích xuất RTMPS URL từ phản hồi.
	if cfResp.Result.RTMPS.UrlNew == "" {
		return "", errors.New("failed to extract rtmps url from cloudflare response")
	}

	return cfResp.Result.RTMPS.UrlNew, nil
}
