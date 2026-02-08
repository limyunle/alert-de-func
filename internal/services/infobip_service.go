package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type InfobipService struct {
	apiKey string
	baseURL string
}

func NewInfobipService() *InfobipService {
	return &InfobipService{
		apiKey:  os.Getenv("INFOBIP_API_KEY"),
		baseURL: os.Getenv("INFOBIP_BASE_URL"),
	}
}

func (i *InfobipService) Call(phoneNumber, message string) error {
	payload := map[string]interface{}{
		"from": "DutyAlert",
		"to":   phoneNumber,
		"text": message,
	}

	body, _ := json.Marshal(payload)

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/tts/1/advanced", i.baseURL),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "App "+i.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf("infobip returned status %d", resp.StatusCode)
	}

	return nil
}

